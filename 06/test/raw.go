package test

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID int
    Name string
    Money float64 `json:",string"`
    Skills []string
    Relationship map[string]string
    Identification Identification
    Career string
    Responsibility json.RawMessage
}

type Identification struct {
    Phone bool `json:"phone"`
    Email bool `json:"email"`
}

type Engineer struct {
    Skill string `json:"skill"`
    Description string `json:"description"`
}

type Manager struct {
    Experienced bool `json:"experienced"`
}

func Raw1() {
    var jsonBlob = []byte(`[
        {
            "ID":1,
            "Name":"Tony",
            "Career":"Engineer",
            "Responsibility":{
                "skill":"PHP&Golang&Network",
                "description":"coding"
            }
        },
        {
            "ID":2,
            "Name":"Jim",
            "Career":"Manager",
            "Responsibility":{
                "experienced":true
            }
        }
    ]`)
    
    var users []User
    if err := json.Unmarshal(jsonBlob, &users); err != nil {
        fmt.Println("error:", err)
    }

    for _, user := range users {
        switch user.Career {
        case "Engineer":
            var responsibility Engineer
            if err := json.Unmarshal(user.Responsibility, &responsibility); err != nil {
                fmt.Println("error:", err)
            }
            fmt.Println(responsibility.Description)
        case "Manager":
            var responsibility Manager
            if err := json.Unmarshal(user.Responsibility, &responsibility); err != nil {
                fmt.Println("error:", err)
            }
            fmt.Println(responsibility.Experienced)
        default:
            fmt.Println("warning:", "don't exist")
        }
    }
}