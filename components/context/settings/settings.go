package settings

import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
)

const (
	SETTINGS_PATH = "conf/settings.json"
)

type Setting struct {
	SiteInfo  SiteInfo        `json:"site_info"`
	Profile   Profile         `json:"profile"`
	AuthSites map[string]Auth `json:"auth_sites"`
}

type SiteInfo struct {
	Name       string `json:"name"`
	SourceCode string `json:"source_code"`
}

type Profile struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Bio    string `json:"bio"`
}

type Auth struct {
	ClientId string `json:"client_id"`
	Url      string `json:"url"`
}

var S Setting

func init() {
	//json_path := os.Getenv("GOPATH")
	f, err := os.Open( /*json_path + "/src/" + */ SETTINGS_PATH) //todo path
	if err != nil {
		log.Println(err)
		return
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &S); err != nil {
		log.Println("JSON unmarshaling failed: %s", err)
		return
	}
}
