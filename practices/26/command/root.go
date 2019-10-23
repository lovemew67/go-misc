package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	// "io"
	"io/ioutil"
	"log"
	// "mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	drawTemplate =`<ac:structured-macro ac:name=\"drawio\" ac:schema-version=\"1\" ac:macro-id=\"88f85e9d-dde0-46ed-8bb7-c80013d425d2\"><ac:parameter ac:name=\"border\">true</ac:parameter><ac:parameter ac:name=\"viewerToolbar\">true</ac:parameter><ac:parameter ac:name=\"fitWindow\">false</ac:parameter><ac:parameter ac:name=\"diagramName\">{{.Name}}</ac:parameter><ac:parameter ac:name=\"simpleViewer\">false</ac:parameter><ac:parameter ac:name=\"width\" /><ac:parameter ac:name=\"diagramWidth\">121</ac:parameter><ac:parameter ac:name=\"revision\">1</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro>`
	plantTemplate = `<ac:structured-macro ac:name=\"plantuml\" ac:schema-version=\"1\" ac:macro-id=\"c1f4c096-062b-490e-877d-24a1de8d1abc\"><ac:parameter ac:name=\"atlassian-macro-output-type\">INLINE</ac:parameter><ac:plain-text-body><![CDATA[{{.Content}}]]></ac:plain-text-body></ac:structured-macro>`
	markdownTemplate = `<ac:structured-macro ac:name=\"markdown\" ac:schema-version=\"1\" ac:macro-id=\"d3e26c03-1e7f-48ba-983b-57c88e034056\"><ac:plain-text-body><![CDATA[{{.Content}}]]></ac:plain-text-body></ac:structured-macro>`

	drawContent = `<ac:structured-macro ac:name=\"drawio\" ac:schema-version=\"1\" ac:macro-id=\"88f85e9d-dde0-46ed-8bb7-c80013d425d2\"><ac:parameter ac:name=\"border\">true</ac:parameter><ac:parameter ac:name=\"viewerToolbar\">true</ac:parameter><ac:parameter ac:name=\"fitWindow\">false</ac:parameter><ac:parameter ac:name=\"diagramName\">dia-template</ac:parameter><ac:parameter ac:name=\"simpleViewer\">false</ac:parameter><ac:parameter ac:name=\"width\" /><ac:parameter ac:name=\"diagramWidth\">121</ac:parameter><ac:parameter ac:name=\"revision\">1</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro>`
	plantContent = `<ac:structured-macro ac:name=\"plantuml\" ac:schema-version=\"1\" ac:macro-id=\"c1f4c096-062b-490e-877d-24a1de8d1abc\"><ac:parameter ac:name=\"atlassian-macro-output-type\">INLINE</ac:parameter><ac:plain-text-body><![CDATA[Entity A\nEntity B\n\n\nA->B: Hello\n\n\nB->A: World]]></ac:plain-text-body></ac:structured-macro>`
	markdownContent = `<ac:structured-macro ac:name=\"markdown\" ac:schema-version=\"1\" ac:macro-id=\"d3e26c03-1e7f-48ba-983b-57c88e034056\"><ac:plain-text-body><![CDATA[## Inquiry\n\n### Create inquiry - IM\n![](http://www.plantuml.com/plantuml/img/pPNDJkCm4CVFv2dc20TuW1mGzQ95Y8zTWSHrDOsdhOLh5ti2ojipTjDjIDj0sHMYKfFcx_dsp-bPP0nSporKWP3fsBYdseIcm8fzW0SJZTmBScouWJllvwlVrzMlKCqe-_lmPxH3LolzrhEPc0hPUBVIpf5nvD0sqSH2oyCO8y65zMEqDzvD_LgyzbMhIQfjgYgAjjhZqrne2JzoIn1IrTAFUKvFkOTq0R5hFKgu8ww3owPmgGjDR-tNaVjj_yrokHxSqd5ZG0Chbk5vcGBGQNWnUaeCnW7F22BEfGcQfi61VFhkgP7Ep4gu45_6QGubXjKZSBEWdRwRBOqM2zla4A6s584zQAgC9hb5DcwvYIbUgXmPB65ayEwD0SoHKznik2jv5uiIxabATXq93biWJ_naGk5eC5QNaOrz6NAymiaZbl2x7biJPt0-6hJkGzDMcBXPkNNusTZij4vMu7L-xsdrKrC8WaoWkF0kqiVJErta4eGmBuY5z_3g7-eIVlhmWa6BurCj-XA67KdLEwi7VUKVY5RRo_qQejL_95hrY5OV9LhjHRJQ_WRu3Dhg6vwQuVRsFZxgoqwdV8ipkSfVyoi0)]]></ac:plain-text-body></ac:structured-macro>`

	fileLogger *log.Logger
)

type contentVersion struct {
	Number int `json:"number"`
}

type getContent struct {
	Version contentVersion `json:"version"`
}

var (
	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			// say hello
			log.Printf("Inside rootCmd Run with args: %v\n", args)

			// viper
			viper.AutomaticEnv()
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

			// init
			httpClient := http.Client{}

			// file logger
			logFile, errCreate := os.Create("26.log")
			defer logFile.Close()
			if errCreate != nil {
				log.Panicf("failed to create 26.log")
			}
			fileLogger = log.New(logFile, "", log.LstdFlags|log.LUTC|log.Lmicroseconds|log.Lshortfile)

			// get content
			fileLogger.Println("==========================")
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/rest/api/content/107316117?expand=version,body.storage", viper.GetString("go.misc.url")), nil)
			req.SetBasicAuth(viper.GetString("go.misc.user"), viper.GetString("go.misc.pass"))
			response, err := httpClient.Do(req)
			if err != nil {
				fileLogger.Panicf("err get: %+v", err)
			}
			defer response.Body.Close()
			resBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fileLogger.Panicf("err read: %+v", err)
			}
			contentVersion := getContent{}
			err = json.Unmarshal(resBody, &contentVersion)
			if err != nil {
				fileLogger.Panicf("err unmarshal: %+v", err)
			}
			fileLogger.Println("current version: " + strconv.Itoa(contentVersion.Version.Number))

			// update content
			newContent := getTemplateContent("request-template", struct {
				Id      string
				Content string
				Version string
			}{
				"107316117",
				markdownContent + plantContent + drawContent + markdownContent + plantContent,
				strconv.Itoa(contentVersion.Version.Number + 1),
			})
			fileLogger.Println("==========================")
			req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("https://%s/rest/api/content/107316117", viper.GetString("go.misc.url")), bytes.NewBuffer([]byte(newContent)))
			req.SetBasicAuth(viper.GetString("go.misc.user"), viper.GetString("go.misc.pass"))
			req.Header.Add("Content-Type", "application/json")
			response, err = httpClient.Do(req)
			if err != nil {
				fileLogger.Panicf("err put: %+v", err)
			}
			defer response.Body.Close()
			resBody, err = ioutil.ReadAll(response.Body)
			if err != nil {
				fileLogger.Panicf("err read: %+v", err)
			}
			fileLogger.Println(string(resBody))

			// get attachments
			fileLogger.Println("==========================")
			req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/rest/api/content/107316117/child/attachment", viper.GetString("go.misc.url")), nil)
			req.SetBasicAuth(viper.GetString("go.misc.user"), viper.GetString("go.misc.pass"))
			response, err = httpClient.Do(req)
			if err != nil {
				fileLogger.Panicf("err get: %+v", err)
			}
			defer response.Body.Close()
			_, err = ioutil.ReadAll(response.Body)
			if err != nil {
				fileLogger.Panicf("err read: %+v", err)
			}

			// // upload attachment
			// fileLogger.Println("==========================")
			// fp, err := os.Open("dia-template")
			// if err != nil {
			// 	fileLogger.Panicf("err: %+v", err)
			// }
			// defer fp.Close()
			// body := &bytes.Buffer{}
			// writer := multipart.NewWriter(body)
			// part, err := writer.CreateFormFile("file", "dia-template")
			// if err != nil {
			// 	fileLogger.Panicf("err: %+v", err)
			// }
			// _, _ = io.Copy(part, fp)
			// _ = writer.WriteField("comment", "file comment")
			// // for key, val := range params {
			// // 	_ = writer.WriteField(key, val)
			// // }
			// err = writer.Close()
			// if err != nil {
			// 	fileLogger.Panicf("err: %+v", err)
			// }
			// req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("https://%s/rest/api/content/107316117/child/attachment", viper.GetString("go.misc.url")), body)
			// if err != nil {
			// 	fileLogger.Panicf("err: %+v", err)
			// }
			// req.SetBasicAuth(viper.GetString("go.misc.user"), viper.GetString("go.misc.pass"))
			// req.Header.Set("Content-Type", writer.FormDataContentType())
			// req.Header.Add("X-Atlassian-Token", "no-check")
			// response, err = httpClient.Do(req)
			// if err != nil {
			// 	fileLogger.Panicf("err get: %+v", err)
			// }
			// defer response.Body.Close()
			// resBody, err = ioutil.ReadAll(response.Body)
			// if err != nil {
			// 	fileLogger.Panicf("err read: %+v", err)
			// }
			// fileLogger.Println(string(resBody))
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func getTemplateContent(filePath string, parameter interface{}) (body string) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fileLogger.Panicf("err read file: %+v", err)
	}
	t, _ := template.New(filePath).Parse(string(b))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, parameter)
	if err != nil {
		fileLogger.Panicf("err template: %+v", err)
	}

	beforeEscape := string(buffer.Bytes())
	afterEscape := html.UnescapeString(beforeEscape)
	fileLogger.Println(beforeEscape)
	fileLogger.Println(afterEscape)
	return afterEscape
}
