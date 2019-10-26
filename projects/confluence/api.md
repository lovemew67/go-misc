### get server version
```
https://<url>/rest/applinks/1.0/manifest
```

### get content
```
https://<url>/rest/api/content/<content_id>?expand=version,body.storage
```

### update content
```
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
```

### get attachments
```
https://<url>/rest/api/content/<content_id>/child/attachment
```

### upload attachment
```
curl -X POST \
    https://<url>/rest/api/content/<content_id>/child/attachment \
    -H "X-Atlassian-Token: no-check" \
    -F "file=@/path/to/file" \
    -F "comment=file comment" \
    -u <username>:<password>
```

### update attachment
```
curl -X POST \
    https://<url>/rest/api/content/<content>/child/attachment/<attachment_id>/data \
    -H "X-Atlassian-Token: no-check" \
    -F "file=@/path/to/file" \
    -F "comment=file comment" \
    -u <username>:<password>
```