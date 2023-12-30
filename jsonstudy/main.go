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

func EncodeJson() {

	ashuCourses := []course{
		{"ReactJS", 250, "Udemy", "ashu123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// packing this data as json file
	finalJson, err := json.MarshalIndent(ashuCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var AshuCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &AshuCourse)
		fmt.Printf("%#v\n", AshuCourse)
	} else {
		fmt.Println("Json is not valid")
	}

	// some cases where we just wnt to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	
	for k, v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is %T\n ", k, v, v)
	}

}

func main() {
	fmt.Println("Json Study in Golang")
	// EncodeJson()
	DecodeJson()
}
