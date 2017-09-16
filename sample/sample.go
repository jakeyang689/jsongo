package main

import (
	"jsongo"
	"fmt"
)

const JStr = "{\"a\":{\"v\":\"value from json string\"},\"b\":[1,2,3], \"c\":true}"

func HandleObject(k string, v jsongo.JValue){
	if v.Type() == jsongo.Array{
		jsongo.RetrieveArray(v.ToArray(), HandleArray)
	}else if v.Type() == jsongo.Object{
		jsongo.RetrieveObject(v.ToObject(), HandleObject)
	}else{
		str := string(v.ToString())
		fmt.Printf("%s, %s\n",k,str)
	}
}

func HandleArray(v jsongo.JValue){
	fmt.Printf("%s\n",v.ToString())
}

func DecodeJson(obj jsongo.JObject){
	//the common operation for a json object
	val := obj.GetValue("a")
	if val.Type() == jsongo.Object{
		subobj := val.ToObject()
		subval := subobj.GetValue("v")
		fmt.Println(subval.ToString())
	}

	val = obj.GetValue("b")
	if val.Type() == jsongo.Array{
		subobj := val.ToArray()
		for i := 0; i < len(subobj); i++{
			fmt.Printf("%s\n",subobj.At(i).ToString())
		}
	}

	val = obj.GetValue("c")
	fmt.Printf("%d %s\n",val.Type(),val.ToString())

	//uncomment the code below to retrieve a json object
	//jsongo.RetrieveObject(obj,HandleObject)
}

func main() {
	//parse json
	obj, err := jsongo.ParseString(JStr)
	if err == nil{
		DecodeJson(obj)
	}

	obj, err = jsongo.ParseFile("./sample.json")
	if err == nil{
		DecodeJson(obj)
	}

	//encode json
	var eobj jsongo.JObject
	eobj.New() //it must be called here to make a new map[string]interface{}
	eobj["object"] = obj
	eobj["number"] = 1
	eobj["string"] = "a sample string"
	var arr jsongo.JArray
	arr.Append(1)
	arr.Append(2)
	eobj["array"] = arr
	eobj["bool"] = true
	eobj["null"] = nil
	
	jstr := jsongo.DumpString(&eobj)
	fmt.Printf("%s\n",jstr)
}
