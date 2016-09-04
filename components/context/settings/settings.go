package settings

import (
	"os"
	"log"
	"sync"
	"io/ioutil"
	"encoding/json"
	"gensh.me/blog/models"
)

const (
	settings_path = "gensh.me/blog/conf/settings.json"
)

type Setting struct {
	Categories []models.Category  `json:"categories"`
	Profile    Profile            `json:"profile"`
	AuthSites  map[string]Auth    `json:"auth_sites"`
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

var s Setting
var mu sync.RWMutex

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
	if err := json.Unmarshal(data, &s); err != nil {
		log.Println("JSON unmarshaling failed: %s", err)
		return
	}
	LoadCategories()
}

func GetSettings() Setting {
	mu.RLock()
	defer mu.RUnlock()
	return s
}