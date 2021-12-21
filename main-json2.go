package main

import(
	"fmt"
	"encoding/json"
)

func main(){
	data := []byte(`{
			"id": "23",
			"name": "vivi",
			"info":{
				"age": 29,
				"gender": "female"
			},
			"jobs":
			[{"id": 1, "title": "manager"},
			 {"id": 2, "title": "engineer"}
			]
		}`)

	var jsonObj map[string]interface{}
	json.Unmarshal([]byte(data), &jsonObj)
	id := jsonObj["id"].(string)
	name := jsonObj["name"].(string)
	info := jsonObj["info"].(map[string]interface{})
	jobs := jsonObj["jobs"].([]interface{})
	for _, v := range jobs{
		p := v.(map[string]interface{})
		id := p["id"]
		title := p["title"]
		fmt.Println(id, title)

	}
	fmt.Println(id, name)
	fmt.Println(info["age"])
	
}