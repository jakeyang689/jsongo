// Copyright 2017 All rights reserved.
// Author Jakeyang
// Use of this source code is governed by by a BSD-style
// license that can be found in the LICENSE file.

// Package jsongo implement a simple json operation package.
// It defines some types for json data format,
// and implement methods for the menagment of json objects.

package jsongo

import(
	"os"
	"fmt"
	"encoding/json"
)

const (
    Object = iota
    Numberic
    String
    Array
    Boolean
    Null
    Unknow
)

func catch(){
    recover()
}

//define the json value struct
type JValue struct{
	value interface{}
}

//get json value type
func (val JValue)Type() int{
    switch val.value.(type) {
       case string:
             return String
       case float64:
             return Numberic
       case bool:
             return Boolean
       case JObject:
             return Object
       case JArray:
             return Array
       default:
             return Null
    }
    return Unknow
}

//transform json value
func (val JValue) ToArray() JArray{
	defer catch()
	return val.value.(JArray)
}

func (val JValue) ToObject() JObject{
	defer catch()
	return val.value.(JObject)
}

func (val JValue) ToString() string{
	defer catch()
	switch val.value.(type) {
     case float64:
           return fmt.Sprintf("%f",val.value.(float64))
     case bool:
           return fmt.Sprintf("%t",val.value.(bool))
     case JObject:
           return fmt.Sprintf("%v",val.value.(JObject))
     case JArray:
           return fmt.Sprintf("%v",val.value.(JArray))
     case nil:
           return fmt.Sprintf("%s","nil")
     default:
           return val.value.(string)
  }
	return val.value.(string)
}

func (val JValue) ToFloat() float64{
	defer catch()
	return val.value.(float64)
}

func (val JValue) ToInt() int64{
	defer catch()
	return int64(val.value.(float64))
}

func (val JValue) ToBool() bool{
	defer catch()
	return val.value.(bool)
}

//define json array value and it's constructor
type JArray []interface{}
func NewJArray() JArray{
	var lst JArray
	return lst
}

//append a json value to json array
func (arr *JArray)Append(elem interface{}){
	*arr = append(*arr,elem)
}

//retrieve a json value from json array
func (arr *JArray)At(id int) JValue{
	var val JValue
	if len(*arr) > id{
  		v := (*arr)[id]
  		switch v.(type) {
          case string:
              val.value = string(v.(string))
          case float64:
              val.value = float64(v.(float64))
          case bool:
              val.value = bool(v.(bool))
          case map[string]interface{}:
              val.value = JObject(v.(map[string]interface{}))
          case []interface {}:
              val.value = JArray(v.([]interface {}))
          default:
              val.value = nil
      }
	}
	return val
}

//define json object value and it's constructor
type JObject map[string]interface{}
func NewJObject() JObject{
	var pnd = make(JObject)
	return pnd
}
func (obj *JObject) New (){
  (*obj) = make(JObject)
}

//get a json value from json object
func (pjnd *JObject) GetValue(key string) JValue{
  	var val JValue
  	if v, ok := (*pjnd)[key]; ok {
  		switch v.(type) {
  			case string:
                val.value = string(v.(string))
        case float64:
                val.value = float64(v.(float64))
        case bool:
                val.value = bool(v.(bool))
        case map[string]interface{}:
                val.value = JObject(v.(map[string]interface{}))
        case []interface {}:
                val.value = JArray(v.([]interface {}))
  			default:
  				val.value = nil
  		}
  	}
  	return val
}

//get value type from a json object
func (pjnd *JObject) GetValueType(key string) int{
  	if v, ok := (*pjnd)[key]; ok {
        switch v.(type) {
            case string:
                    return String
            case float64:
                    return Numberic
            case bool:
                    return Boolean
            case map[string]interface{}:
                    return Object
            case []interface {}:
                    return Array
            default:
                    return Null
        }
    }
    return Unknow
}

//parse json data from a string
func ParseString(jstr string) (obj JObject, err error){
    err = json.Unmarshal([]byte(jstr), &obj)
    return
}

//parse json data from a json file
func ParseFile(path string) (obj JObject, err error){
    file, err := os.Open(path)
    if err != nil {
            return
    }

    fi, _ := file.Stat()

    if fi.Size() == 0 {
            return
    }

    buffer := make([]byte, fi.Size())
    _, err = file.Read(buffer)
  	file.Close()
  	if err == nil{
  		obj, err := ParseString(string(buffer))
  		return obj,err
  	}
  	
  	return obj,err
}

//get a json string from a json object
func DumpString(jnd *JObject) string{
	b,err := json.Marshal(*jnd)
	if err != nil{
		return "err"
	}
	return string(b)
}

//retrieve the json object
func RetrieveObject(obj JObject, f func(string, JValue)) error{
    var err error = nil
    for k,_ := range obj{
        f(k,obj.GetValue(k))
    }
    return err
}

//retrieve the json array
func RetrieveArray(arr JArray, f func(JValue)) error{
    var err error = nil
    for id := 0; id < len(arr); id++{
        f(arr.At(id))
    }
    return err
}








