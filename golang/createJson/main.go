package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("how to create json in go lang")

	// EnCodeJson()
	DecodeJson()
}

// encoding the json
// func EnCodeJson() {
// 	lcoCourse := []course{
// 		{"ReactJS Bootcamp", 299, "LCO", "abc123", []string{"web-dev", "js"}},
// 		{"MERN Bootcamp", 299, "LCO", "abcd123", []string{"fullstack", "js"}},
// 		{"Angular Bootcamp", 299, "LCO", "abdc123", nil},
// 	}

// 	// package this data as json date
// 	finlaJson, err := json.MarshalIndent(lcoCourse, " ", "/t")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", finlaJson)

// }

// consume the json data
func DecodeJson() {
	jsonDateFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "LCO",
			"tags": ["web-dev","js"]
		}
`)

	var lcoCourse course

	checkValid := json.Valid(jsonDateFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDateFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// some cases where you just want to add data key value
	var myOnllineData map[string]interface{}
	json.Unmarshal(jsonDateFromWeb, &myOnllineData)
	fmt.Printf("%#v\n", myOnllineData)

	for k, v := range myOnllineData {
		fmt.Printf("Key is  %v and value is %v and type of data is: %T\n", k, v, v)
	}
}
