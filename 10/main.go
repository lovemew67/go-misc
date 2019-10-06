package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 10")
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "no routing rule matched\n")
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
	// https://gobyexample.com/http-servers
	go func() {
		http.HandleFunc("/", rootHandler)
		http.HandleFunc("/hello", helloHandler)
		_ = http.ListenAndServe(":8080", nil)
	}()
	
	// send get request
	// https://www.kancloud.cn/digest/batu-go/153529
	response, errGet := http.Get("http://localhost:8080")
	if errGet != nil {
		log.Panicf("err get: %+v", errGet)
	}
	defer response.Body.Close()
	resBody, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Panicf("err read: %+v", errRead)
	}
	log.Println(string(resBody))

	// send post request
	// https://www.kancloud.cn/digest/batu-go/153529
	postBody := "{\"id\":2}"
	response, errPost := http.Post(
		"https://my-json-server.typicode.com/lovemew67/go-misc/posts", 
		"application/json",
		bytes.NewBuffer([]byte(postBody)),
	)
	if errPost != nil {
		log.Panicf("err post: %+v", errPost)
	}
	defer response.Body.Close()
	resBody, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Panicf("err read: %+v", errRead)
	}
	log.Println(string(resBody))

	// send post form request
	// https://www.kancloud.cn/digest/batu-go/153529
	v := url.Values{}
    v.Set("username", "go")
	v.Set("password", "misc")
	reqBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	httpClient := http.Client{}
	request, _ := http.NewRequest(
		http.MethodPost,
		"https://postman-echo.com/post",
		reqBody,
	)
	request.Header.Set(
		"Content-Type", 
		"application/x-www-form-urlencoded;param=value",
	)
	response, _ = httpClient.Do(request)
	defer response.Body.Close()
	responseBody, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Panicf("err read: %+v", errRead)
	}
	log.Println(string(responseBody))

	// send post from request
	// https://studygolang.com/articles/9467
	postParam := url.Values{
		"mobile":      {
			"xxxxxx",
		},
		"isRemberPwd": {
			"1",
		},
	}
	response, _ = http.PostForm("https://postman-echo.com/post", postParam)
	defer response.Body.Close()
	responseBody, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Panicf("err read: %+v", errRead)
	}
	log.Println(string(responseBody))

	// send patch requet with http client
	patchBody := "{\"title\": \"fake title: 2\"}"
	request, _ = http.NewRequest(
		http.MethodPatch,
		"https://my-json-server.typicode.com/lovemew67/go-misc/posts/1", 
		bytes.NewBuffer([]byte(patchBody)),
	)
	response, _ = httpClient.Do(request)
	defer response.Body.Close()
	responseBody, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Panicf("err read: %+v", errRead)
	}
	log.Println(string(responseBody))

	// send get requet with http client
	request, _ = http.NewRequest(
		http.MethodGet,
		"https://my-json-server.typicode.com/lovemew67/go-misc/posts/1",
		nil,
	)
	response, _ = httpClient.Do(request)
	defer response.Body.Close()
	responseBody, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Panicf("err read: %+v", errRead)
	}
	log.Println(string(responseBody))

	for {}
}