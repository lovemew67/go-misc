package test

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id string `json:"id"`
}

type Org struct {
	Name string `json:"name"`
}

type AppWithOrg struct {
	App
	Org
}

func Raw2() {
	data := []byte(`
        {
            "id": "k34rAT4",
            "name": "My Awesome Org"
        }
    `)
	var b AppWithOrg
	_ = json.Unmarshal(data, &b)
	fmt.Printf("%#v\n", b)

	a := AppWithOrg{
		App: App{
			Id: "k34rAT4",
		},
		Org: Org{
			Name: "My Awesome Org",
		},
	}
	data, _ = json.Marshal(a)
	fmt.Println(string(data))
}