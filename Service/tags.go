package Service

import (
	"awesomeProject/DAO"
	"awesomeProject/structs"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type T struct {
	Tags []string `json:"tags"`
}

var t T

type F string

func CreateDefaultTags() {
	for _, e := range F("Assets/tags.json").convertJson() {
		DAO.InsertDefaultTag(structs.Tag{TagName: e, TagUser: "SystemGenerated"})
	}

}

func CreateTags(tag *structs.Tag) {
	DAO.InsertTag(tag)
}

func GetTags() []structs.Tag {
	return DAO.GetTags()
}

func (f F) convertJson() []string {
	_file, err := os.Open(f.toString())
	if err != nil {
		log.Fatal("Error until reading Json file: error code: " + err.Error())
	}
	defer _file.Close()
	decoder := json.NewDecoder(_file)
	if err = decoder.Decode(&t); err != nil {
		log.Fatal("error ocured until decodig file to struct  with error code: " + err.Error())
	}
	return t.Tags
}

func (f F) toString() string {
	return String(f)
}
