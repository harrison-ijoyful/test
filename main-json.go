package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`
		{"cId" : "A" ,
         "cType" : "English" , "saddr" : { "hsnId" : "C" , "addr" : "中正路12號" } , "persons" : [{"id" : 1 , "name" : "Daniel"},{"id" : 2 , "name" : "Allen"},{"id" : 3 , "name" : "Sam"}]}`)
	var jsonObj map[string]interface{}
	json.Unmarshal([]byte(data), &jsonObj)
	classID := jsonObj["cId"].(string)
	classType := jsonObj["cType"].(string)

	fmt.Println(classID)
	fmt.Println(classType)

	studentsAddr := jsonObj["saddr"].(map[string]interface{})
	hsnID := studentsAddr["hsnId"].(string)
	addr := studentsAddr["addr"].(string)

	fmt.Println(hsnID)
	fmt.Println(addr)

	persons := jsonObj["persons"].([]interface{})
	for _, p := range persons {
		person := p.(map[string]interface{})
		id := int(person["id"].(float64))
		name := person["name"].(string)
		fmt.Printf("%d , %v \n", id, name)
	}
}