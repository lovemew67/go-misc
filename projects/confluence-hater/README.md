# confluence hater

help engineeeeeeeeeeeeeeers to sync/upload/update markdown/plantuml/drawio contents to confluence pages.

## prerequisites

- confluences page ids. retrieve page id from `Page Information`: https://<confluence_url>/pages/viewinfo.action?pageId=<page_id>
- confluecne server must support `Markdown/PlangUML/Draw.io` marcos
- supported confluence version. tested against:
```
<version>6.2.3</version>
<buildNumber>7615</buildNumber>
<applinksVersion>5.2.6</applinksVersion>
```

## usage

```go
go run main.go
```

## help

```
  ______             ___ _                                 _     _
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
}

Usage:
  confluence-hater [sync request in json] [flags]

Flags:
  -f, --filelog   log to file nor not (default: confluence-hater.log)
  -h, --help      help for confluence-hater
```

## sample log

```
2019/10/27 13:11:18.502682 confluence.go:249: [confluenceCmd] confluence url: <confluence_url>
2019/10/27 13:11:18.502682 confluence.go:250: [confluenceCmd] confluence username: <confluence_username>
2019/10/27 13:11:18.502682 confluence.go:251: [confluenceCmd] confluence password: *************
2019/10/27 13:11:18.507180 confluence.go:309: [validateParseRequestFile] validateion passed
2019/10/27 13:11:18.507180 confluence.go:314: [processPages] start to process [1] pages
2019/10/27 13:11:18.507180 confluence.go:336: [processPage] start to process page: <content_id>
2019/10/27 13:11:18.507180 confluence.go:401: [prepareProcessPageContext] start to prepare page context: <content_id>
2019/10/27 13:11:20.126257 confluence.go:424: [handleDrawioAttachment] going to upload, filePath: templates/drawio, fileName: drawio
2019/10/27 13:11:22.002290 confluence.go:324: [processPages] succeeded to process [1] pages:
2019/10/27 13:11:22.002290 confluence.go:326: 
	page: 107316117, with link: https://<confluence_url>/pages/viewpage.action?pageId=<content_id>
2019/10/27 13:11:22.002290 confluence.go:328: [processPages] failed to process [0] pages:
```

## references

```
- get server version
https://<url>/rest/applinks/1.0/manifest

- get content
https://<url>/rest/api/content/<content_id>?expand=version,body.storage

- update content
curl -X PUT \
    https://<url>/rest/api/content/<content_id> \
    -H 'Content-Type: application/json' \
    -u <username>:<password> \
    -d '{ 
    "id":"<content_id>", 
    "type":"page",
    "title":"<title>",
    "body":{
        "storage":{
            "value": "<content>",
            "representation":"storage"
        }
    },
    "version":{
        "number": <page_version>
    }
}'

- get attachments
https://<url>/rest/api/content/<content_id>/child/attachment

- upload attachment
curl -X POST \
    https://<url>/rest/api/content/<content_id>/child/attachment \
    -H "X-Atlassian-Token: no-check" \
    -F "file=@/path/to/file" \
    -F "comment=file comment" \
    -u <username>:<password>

- update attachment
curl -X POST \
    https://<url>/rest/api/content/<content>/child/attachment/<attachment_id>/data \
    -H "X-Atlassian-Token: no-check" \
    -F "file=@/path/to/file" \
    -F "comment=file comment" \
    -u <username>:<password>
```