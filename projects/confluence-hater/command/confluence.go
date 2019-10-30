package command

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	unknown             = "unknown"
	version             = "version"
	initalVersion       = "1"
	confluenceScheme    = "https"
	requestTemplateName = "request"

	formKeyfile    = "file"
	formKeyComment = "comment"

	logFileFlagLong  = "filelog"
	logFileFlagShort = "f"
	logFileName      = "confluence-hater.log"

	viperKeyConfluenceUrl      = "confluence.hater.url"
	viperKeyConfluenceUsername = "confluence.hater.username"
	viperKeyConfluencePassword = "confluence.hater.password"

	ContentTypeDrawio   ContentType = "drawio"
	ContentTypePlantuml ContentType = "plantuml"
	ContentTypeMarkdown ContentType = "markdown"
)

const (
	usage     = `confluence-hater [sync request in json]`
	shortHelp = `confluence hater help engineeeeeeeeeeeeeeers to sync/upload/update markdown/plantuml/drawio contents to confluence pages`
	longHelp  = `  ______             ___ _                                 _     _                       
 / _____)           / __) |                               | |   | |      _               
| /      ___  ____ | |__| |_   _  ____ ____   ____ ____   | |__ | | ____| |_  ____  ____ 
| |     / _ \|  _ \|  __) | | | |/ _  )  _ \ / ___) _  )  |  __)| |/ _  |  _)/ _  )/ ___)
| \____| |_| | | | | |  | | |_| ( (/ /| | | ( (__( (/ /   | |   | ( ( | | |_( (/ /| |    
 \______)___/|_| |_|_|  |_|\____|\____)_| |_|\____)____)  |_|   |_|\_||_|\___)____)_|    
																							
confluence hater help engineeeeeeeeeeeeeeers to sync/upload/update markdown/plantuml/drawio contents to confluence pages. will just override current contents.

prerequisites
- confluences page ids. retrieve page id from "Page Information": https://<confluence_url>/pages/viewinfo.action?pageId=<page_id>
- confluecne server must support "Markdown/PlangUML/Draw.io" marcos
- supported confluence version. tested against:
	- version: 6.2.3
	- buildNumber: 7615
	- applinksVersion: 5.2.6

current support content type:
- drawio   (default width: 480)
- markdown
- plantuml

environment variables:
	CONFLUENCE_HATER_URL: confluence server url including port number. default use "https" scheme.
	CONFLUENCE_HATER_USERNAME: confluence username for authentication
	CONFLUENCE_HATER_PASSWORD: confluence password for authentication

request template:
{
    "pages": [
        {
            "id": "<page_id>",
            "contents": [
                {
                    "type": "drawio",
                    "source": "path/to/drawio/file"
                },
                {
                    "type": "markdown",
                    "source": "path/to/drawio/file"
                },
                {
                    "type": "plantuml",
                    "source": "path/to/plantuml/file"
				}
				... // content will display in this order
            ]
		}
		... // process pages sequently
    ]
}`
)

var (
	logger     *log.Logger
	httpClient http.Client

	confluenceUrl      string
	confluenceUsername string
	confluencePassword string
)

var (
	requestTemplate = `{
	"id":"{{.Id}}", 
	"type":"page",
	"title":"{{.Title}}",
	"body":{
		"storage":{
			"value": "{{.Content}}",
			"representation":"storage"
		}
	},
	"version":{
		"number": {{.Version}}
	}
}`
)

var (
	pictureTemplateHeader = `<ac:structured-macro ac:name=\"view-file\" ac:schema-version=\"1\" ac:macro-id=\"88888888-8888-8888-8888-888888888888\"><ac:parameter ac:name=\"name\"><ri:attachment ri:filename=\"`
	// attachment file name
	pictureTemplateFooter = `\" /></ac:parameter><ac:parameter ac:name=\"height\">250</ac:parameter></ac:structured-macro>`
)

var (
	drawioTemplateHeader = `<ac:structured-macro ac:name=\"drawio\" ac:schema-version=\"1\" ac:macro-id=\"88888888-8888-8888-8888-888888888888\"><ac:parameter ac:name=\"border\">true</ac:parameter><ac:parameter ac:name=\"viewerToolbar\">true</ac:parameter><ac:parameter ac:name=\"fitWindow\">false</ac:parameter><ac:parameter ac:name=\"diagramName\">`
	// attachment file name
	drawioTemplateInner = `</ac:parameter><ac:parameter ac:name=\"simpleViewer\">false</ac:parameter><ac:parameter ac:name=\"width\" /><ac:parameter ac:name=\"diagramWidth\">480</ac:parameter><ac:parameter ac:name=\"revision\">`
	// attachment revision
	drawioTemplateFooter = `</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro>`
)

var (
	plantumlTemplateHeader = `<ac:structured-macro ac:name=\"plantuml\" ac:schema-version=\"1\" ac:macro-id=\"88888888-8888-8888-8888-888888888888\"><ac:parameter ac:name=\"atlassian-macro-output-type\">INLINE</ac:parameter><ac:plain-text-body><![CDATA[`
	// plantuml file content
	plantumlTemplateFooter = `]]></ac:plain-text-body></ac:structured-macro>`
)

var (
	markdownTemplateHeader = `<ac:structured-macro ac:name=\"markdown\" ac:schema-version=\"1\" ac:macro-id=\"88888888-8888-8888-8888-888888888888\"><ac:plain-text-body><![CDATA[`
	// markdown file content
	markdownTemplateFooter = `]]></ac:plain-text-body></ac:structured-macro>`
)

type ContentType string

type getPageAttachmentResultLink struct {
	Download string `json:"download"`
}

type getPageAttachmentsResult struct {
	ID    string                      `json:"id"`
	Title string                      `json:"title"`
	Links getPageAttachmentResultLink `json:"_links"`
}

type getPageAttachmentsResponse struct {
	Results []*getPageAttachmentsResult `json:"results"`
}

type getPageResponseVersion struct {
	Number int `json:"number"`
}

type getPageResponse struct {
	Title   string                 `json:"title"`
	Version getPageResponseVersion `json:"version"`
}

type pageDetailContent struct {
	Type   string `json:"type"`
	Source string `json:"source"`
}

type pageDetail struct {
	ID       string               `json:"id"`
	Contents []*pageDetailContent `json:"contents"`
}

type parsePageRequest struct {
	Pages []*pageDetail `json:"pages"`
}

type pageProcessAttachmentContext struct {
	ID      string
	Version string
}

type pageProcessContext struct {
	ID          string
	Title       string
	Version     int
	Attachments map[string]*pageProcessAttachmentContext
}

var (
	confluenceCmd = &cobra.Command{
		Use:     usage,
		Short:   shortHelp,
		Long:    longHelp,
		Args:    cobra.ExactArgs(1),
		Version: fmt.Sprintf("info: \n\ttag: %s\n\tcommit: %s\n\tbuild date: %s", gitCommit, appVersion, buildDate),
		Run: func(cmd *cobra.Command, args []string) {
			// init parse request
			handlerName := "confluenceCmd"
			requestFileName := args[0]

			// viper
			viper.AutomaticEnv()
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

			// init file logger
			logger = log.New(os.Stdout, "", log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lshortfile)
			if viper.GetBool(logFileFlagLong) {
				logFile, errCreate := os.Create(logFileName)
				defer logFile.Close()
				if errCreate != nil {
					log.Panicf("[%s] failed to create log file: %s", handlerName, logFileName)
				}
				logger = log.New(logFile, "", log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lshortfile)
			}

			// init http client
			httpClient = http.Client{}

			// confluence info
			confluenceUrl = viper.GetString(viperKeyConfluenceUrl)
			confluenceUsername = viper.GetString(viperKeyConfluenceUsername)
			confluencePassword = viper.GetString(viperKeyConfluencePassword)
			var passwordBuilder strings.Builder
			for i := 0; i < len(confluencePassword); i++ {
				passwordBuilder.WriteString("*")
			}
			if confluenceUrl == "" || confluenceUsername == "" || confluencePassword == "" {
				logger.Panicf("[%s] not enough confluence info", handlerName)
			}
			logger.Printf("[%s] confluence url: %s", handlerName, confluenceUrl)
			logger.Printf("[%s] confluence username: %s", handlerName, confluenceUsername)
			logger.Printf("[%s] confluence password: %s", handlerName, passwordBuilder.String())

			// parse request file
			ppr := parsePageRequest{}
			requestFile, errRead := ioutil.ReadFile(requestFileName)
			if errRead != nil {
				logger.Panicf("[%s] failed to read request file, err: %+v", handlerName, errRead)
			}
			errJson := json.Unmarshal([]byte(requestFile), &ppr)
			if errJson != nil {
				logger.Panicf("[%s] failed to unmarshal request file, err: %+v", handlerName, errJson)
			}

			// validate then process
			validateParseRequestFile(ppr)
			processPages(ppr)
		},
	}
)

func Execute() error {
	confluenceCmd.Flags().BoolP(logFileFlagLong, logFileFlagShort, false, fmt.Sprintf("log to file nor not (default: %s)", logFileName))
	_ = viper.BindPFlag(logFileFlagLong, confluenceCmd.Flags().Lookup(logFileFlagLong))
	return confluenceCmd.Execute()
}

func (ct ContentType) String() string {
	switch ct {
	case ContentTypeDrawio,
		ContentTypePlantuml,
		ContentTypeMarkdown:
		return string(ct)
	default:
		return unknown
	}
}

func validateParseRequestFile(ppr parsePageRequest) {
	handlerName := "validateParseRequestFile"
	if ppr.Pages == nil || len(ppr.Pages) == 0 {
		logger.Panicf("[%s] invalid pages", handlerName)
	}
	for _, page := range ppr.Pages {
		if page.ID == "" {
			logger.Panicf("[%s] invalid page id", handlerName)
		}
		if page.Contents == nil || len(page.Contents) == 0 {
			logger.Panicf("[%s] invalid page content", handlerName)
		}
		for _, content := range page.Contents {
			if content.Type == "" || ContentType(content.Type) == unknown {
				logger.Panicf("[%s] invalid page content type", handlerName)
			}
			if content.Source == "" {
				logger.Panicf("[%s] invalid page content source", handlerName)
			}
		}
	}
	logger.Printf("[%s] validateion passed", handlerName)
}

func processPages(ppr parsePageRequest) {
	handlerName := "processPages"
	logger.Printf("[%s] start to process [%d] pages", handlerName, len(ppr.Pages))
	succeeded := []string{}
	failed := map[string]error{}
	for _, pd := range ppr.Pages {
		if err := processPage(pd); err != nil {
			failed[pd.ID] = err
		} else {
			succeeded = append(succeeded, pd.ID)
		}
	}
	logger.Printf("[%s] succeeded to process [%d] pages:\n", handlerName, len(succeeded))
	for _, id := range succeeded {
		logger.Printf("\n\tpage: %s, with link: %s://%s/pages/viewpage.action?pageId=%s\n", id, confluenceScheme, confluenceUrl, id)
	}
	logger.Printf("[%s] failed to process [%d] pages:\n", handlerName, len(failed))
	for id, err := range failed {
		logger.Printf("\n\tpage: %s, with error: %+v\n", id, err)
	}
}

func processPage(pd *pageDetail) (err error) {
	handlerName := "processPage"
	logger.Printf("[%s] start to process page: %s", handlerName, pd.ID)

	// prepare context
	ppc, err := prepareProcessPageContext(pd)
	if err != nil {
		return
	}

	// build update payload
	var paylodBuilder strings.Builder
	for _, content := range pd.Contents {
		switch content.Type {
		case ContentTypeDrawio.String():
			fileName, attaVersion, errHandle := handleDrawioAttachment(content.Source, ppc)
			if errHandle != nil {
				err = errHandle
				return
			}
			body := getDrawioContent(fileName, attaVersion)
			paylodBuilder.WriteString(body)
			break
		case ContentTypePlantuml.String():
			body, errGet := getPlantumlContent(content.Source)
			if errGet != nil {
				err = errGet
				return
			}
			paylodBuilder.WriteString(body)
			break
		case ContentTypeMarkdown.String():
			body, errGet := getMarkdownContent(content.Source)
			if errGet != nil {
				err = errGet
				return
			}
			paylodBuilder.WriteString(body)
			break
		default:
			break
		}
	}

	// compose request payload
	updateContent, err := getRequestContent(struct {
		Id      string
		Title   string
		Content string
		Version string
	}{
		pd.ID,
		ppc.Title,
		paylodBuilder.String(),
		strconv.Itoa(ppc.Version + 1),
	})
	if err != nil {
		return
	}

	// update page
	err = apiUpdatePage(pd.ID, updateContent)
	return
}

func prepareProcessPageContext(pd *pageDetail) (ppc *pageProcessContext, err error) {
	handlerName := "prepareProcessPageContext"
	logger.Printf("[%s] start to prepare page context: %s", handlerName, pd.ID)
	version, title, err := apiGetPageMetadata(pd.ID)
	if err != nil {
		return
	}
	attachments, err := apiGetPageAttachments(pd.ID)
	if err != nil {
		return
	}
	ppc = &pageProcessContext{
		ID:          pd.ID,
		Title:       title,
		Version:     version,
		Attachments: attachments,
	}
	return
}

func handleDrawioAttachment(filePath string, ppc *pageProcessContext) (fileName, version string, err error) {
	handlerName := "handleDrawioAttachment"
	fileName = filepath.Base(filePath)
	ppac, ok := ppc.Attachments[fileName]
	if !ok {
		logger.Printf("[%s] going to upload, filePath: %s, fileName: %s", handlerName, filePath, fileName)
		version = initalVersion
		err = apiUploadOrUpdatePageAttachment(ppc.ID, "", filePath, fileName, true)
	} else {
		logger.Printf("[%s] going to update, filePath: %s, fileName: %s, attachment id: %s, current version: %s", handlerName, filePath, fileName, ppac.ID, ppac.Version)
		i, _ := strconv.Atoi(ppac.Version)
		version = strconv.Itoa(i + 1)
		err = apiUploadOrUpdatePageAttachment(ppc.ID, ppac.ID, filePath, fileName, false)
	}
	return
}

func getDrawioContent(fileName, attaVersion string) (body string) {
	return fmt.Sprintf("%s%s%s%s%s", drawioTemplateHeader, fileName, drawioTemplateInner, attaVersion, drawioTemplateFooter)
}

func getPlantumlContent(filePath string) (body string, err error) {
	handlerName := "getPlantumlContent"
	var plantumlBuilder strings.Builder
	plantumlFile, errOpen := os.Open(filePath)
	if errOpen != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to open file, err: %+v", handlerName, errOpen))
		return
	}
	defer plantumlFile.Close()
	fileScanner := bufio.NewScanner(plantumlFile)
	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = strings.Replace(input, "\t", "    ", -1)
		input = strings.Replace(input, "\"", "\\\"", -1)
		plantumlBuilder.WriteString(input)
		plantumlBuilder.WriteString(`\n`)
	}
	if errScan := fileScanner.Err(); errScan != nil {
		err = errors.New(fmt.Sprintf("[%s] scanner err: %+v", handlerName, errScan))
		return
	}
	return fmt.Sprintf("%s%s%s", plantumlTemplateHeader, plantumlBuilder.String(), plantumlTemplateFooter), nil
}

func getMarkdownContent(filePath string) (body string, err error) {
	handlerName := "getMarkdownContent"
	var markdownBuilder strings.Builder
	markdownFile, errOpen := os.Open(filePath)
	if errOpen != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to open file, err: %+v", handlerName, errOpen))
		return
	}
	defer markdownFile.Close()
	fileScanner := bufio.NewScanner(markdownFile)
	for fileScanner.Scan() {
		input := fileScanner.Text()
		input = strings.Replace(input, "\t", "    ", -1)
		input = strings.Replace(input, "\"", "\\\"", -1)
		markdownBuilder.WriteString(input)
		markdownBuilder.WriteString(`\n`)
	}
	if errScan := fileScanner.Err(); errScan != nil {
		err = errors.New(fmt.Sprintf("[%s] scanner err: %+v", handlerName, errScan))
		return
	}
	return fmt.Sprintf("%s%s%s", markdownTemplateHeader, markdownBuilder.String(), markdownTemplateFooter), nil
}

func getRequestContent(parameter interface{}) (body string, err error) {
	handlerName := "getRequestContent"
	t, errParse := template.New(requestTemplateName).Parse(requestTemplate)
	if errParse != nil {
		err = errors.New(fmt.Sprintf("[%s] failed parse template, err: %+v", handlerName, errParse))
		return
	}
	buffer := new(bytes.Buffer)
	if errExecute := t.Execute(buffer, parameter); errExecute != nil {
		err = errors.New(fmt.Sprintf("[%s] failed execute template, err: %+v", handlerName, errExecute))
		return
	}
	beforeUnescape := string(buffer.Bytes())
	body = html.UnescapeString(beforeUnescape)
	return
}

func apiGetPageMetadata(id string) (version int, title string, err error) {
	handlerName := "apiGetPageMetadata"

	// prepare request
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s://%s/rest/api/content/%s?expand=version", confluenceScheme, confluenceUrl, id), nil)
	req.SetBasicAuth(confluenceUsername, confluencePassword)

	// send request
	response, errHttp := httpClient.Do(req)
	if errHttp != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to send request: %+v", handlerName, errHttp))
		return
	}

	// unmarshal response
	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to read response body: %+v", handlerName, err))
		return
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("[%s] code not 200, res body: %s", handlerName, string(resBody)))
		return
	}
	contentMetadata := getPageResponse{}
	if err = json.Unmarshal(resBody, &contentMetadata); err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to unmarshal response body: %+v", handlerName, err))
		return
	}

	// get metadata
	version = contentMetadata.Version.Number
	title = contentMetadata.Title
	return
}

func apiGetPageAttachments(id string) (titleMap map[string]*pageProcessAttachmentContext, err error) {
	handlerName := "apiGetPageAttachments"

	// prepare request
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s://%s/rest/api/content/%s/child/attachment", confluenceScheme, confluenceUrl, id), nil)
	req.SetBasicAuth(confluenceUsername, confluencePassword)

	// send request
	response, errHttp := httpClient.Do(req)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to send request: %+v", handlerName, errHttp))
		return
	}

	// unmarshal response
	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to read response body: %+v", handlerName, err))
		return
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("[%s] code not 200, res body: %s", handlerName, string(resBody)))
		return
	}
	contentAttachments := getPageAttachmentsResponse{}
	if err = json.Unmarshal(resBody, &contentAttachments); err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to unmarshal response body: %+v", handlerName, err))
		return
	}

	// get attachments
	titleMap = map[string]*pageProcessAttachmentContext{}
	if contentAttachments.Results == nil {
		return
	}
	for _, atta := range contentAttachments.Results {
		ppac := &pageProcessAttachmentContext{}
		ppac.ID = atta.ID
		downloadLink := atta.Links.Download
		downloadUrl, errParseUrl := url.Parse(downloadLink)
		if errParseUrl != nil {
			err = errors.New(fmt.Sprintf("[%s] failed to parse download link to url, err: %+v", handlerName, errParseUrl))
			return
		}
		queryMap, errParseQuery := url.ParseQuery(downloadUrl.RawQuery)
		if errParseQuery != nil {
			err = errors.New(fmt.Sprintf("[%s] failed to parse query in download url, err: %+v", handlerName, errParseQuery))
			return
		}
		attachmentVersion := queryMap[version][0]
		ppac.Version = attachmentVersion
		titleMap[atta.Title] = ppac
	}
	return
}

func apiUpdatePage(id, body string) (err error) {
	handlerName := "apiUpdatePage"

	// prepare request
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("%s://%s/rest/api/content/%s", confluenceScheme, confluenceUrl, id), bytes.NewBuffer([]byte(body)))
	req.SetBasicAuth(confluenceUsername, confluencePassword)
	req.Header.Add("Content-Type", "application/json")

	// send request
	response, errHttp := httpClient.Do(req)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to send request: %+v", handlerName, errHttp))
		return
	}

	// unmarshal response
	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to read response body: %+v", handlerName, err))
		return
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("[%s] code not 200, res body: %s", handlerName, string(resBody)))
	}
	return
}

func apiUploadOrUpdatePageAttachment(contentId, attachmentId, filePath, fileName string, upload bool) (err error) {
	handlerName := "apiUploadOrUpdatePageAttachment"

	// process file
	fp, errOpen := os.Open(filePath)
	if errOpen != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to open file, err: %+v", handlerName, errOpen))
		return
	}
	defer fp.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, errCreate := writer.CreateFormFile(formKeyfile, fileName)
	if errCreate != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to create form-data header, err: %+v", handlerName, errCreate))
		return
	}
	_, errCopy := io.Copy(part, fp)
	if errCopy != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to copy file to writer, err: %+v", handlerName, errCopy))
		return
	}
	errWrite := writer.WriteField(formKeyComment, fmt.Sprintf("file updated at: %s", time.Now().UTC().String()))
	if errWrite != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to write form data, err: %+v", handlerName, errWrite))
		return
	}
	errClose := writer.Close()
	if errClose != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to close writer, err: %+v", handlerName, errClose))
		return
	}

	// prepare request
	var req *http.Request
	if upload {
		req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("%s://%s/rest/api/content/%s/child/attachment", confluenceScheme, confluenceUrl, contentId), body)
	} else {
		req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("%s://%s/rest/api/content/%s/child/attachment/%s/data", confluenceScheme, confluenceUrl, contentId, attachmentId), body)
	}
	req.SetBasicAuth(confluenceUsername, confluencePassword)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("X-Atlassian-Token", "no-check")

	// send request
	response, errHttp := httpClient.Do(req)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to send request: %+v", handlerName, errHttp))
		return
	}

	// unmarshal response
	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%s] failed to read response body: %+v", handlerName, err))
		return
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("[%s] code not 200, res body: %s", handlerName, string(resBody)))
	}
	return
}
