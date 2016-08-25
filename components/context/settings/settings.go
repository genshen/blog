package settings

import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
)

const (
	settings_path = "gensh.me/blog/conf/settings.json"
)

type Setting struct {
	Menus     []Menu          `json:"menus"`
	Profile   Profile         `json:"profile"`
	AuthSites map[string]Auth `json:"auth_sites"`
}

type Profile struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Bio    string `json:"bio"`
}

type Menu struct {
	SubMenus []SubMenu  `json:"submenus"`
	Name     string     `json:"name"`
}

type SubMenu struct {
	Name string  `json:"name"`
	Url  string  `json:"url"`
}

type Auth struct {
	ClientId string `json:"client_id"`
	Url      string `json:"url"`
}

var S Setting

func init() {
	json_path := os.Getenv("GOPATH")
	f, err := os.Open(json_path + "/src/" + settings_path)
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
	}
}