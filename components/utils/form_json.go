package utils

import (
	"github.com/astaxie/beego/validation"
	"encoding/json"
	"fmt"
)

type Err struct {
	Name    string `json:"name"`
	Message string `json:"name"`
}

type Field struct {
	Value  string `json:"value"`
	Errors []Err  `json:"errors"`
}

func NewInstant(Errors []*validation.Error, f map[string]string) map[string]Field {
	var fields = make(map[string]Field)
	var F Field
	var ok bool
	for _, err := range Errors {
		if F, ok = fields[err.Key]; !ok {
			//not exists, add
			F = Field{Value: f[err.Key]}
		}
		F.Errors = append(F.Errors, Err{err.Key, err.Message})
		fields[err.Key] = F
	}
	return fields
}

/**
 errors with the same key will appear only once.
 */
func NewSingleErrorInstant(Errors []*validation.Error) map[string]string {
	var fields = make(map[string]string)
	for _, err := range Errors {
		if _, ok := fields[err.Key]; !ok {
			//not exists, add
			fields[err.Key] = err.Message
		}
	}
	return fields
}

func NewInstantToByte(Errors []*validation.Error, f map[string]string) []byte {
	b, err := json.Marshal(NewInstant(Errors, f))
	if err != nil {
		fmt.Println("json err:", err) //todo err return
	}
	return b
}
