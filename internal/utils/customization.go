package utils

import (
	"encoding/json"
	"os"
)

type CustomizationOptions struct {
	BotName string `json:"BotName"`
	DocsURL string `json:"DocsURL"`
	CDN     struct {
		BaseURL string
		Files   map[string]string
	}
}

var BotCustoms CustomizationOptions

func LoadCustoms() {
	customsFile, err := os.Open("customizations/botproperties.json")
	if err != nil {
		Log.Error("Couldn't load customizations file: " + err.Error())
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(customsFile)
	if err = jsonParser.Decode(&BotCustoms); err != nil {
		Log.Error("Couldn't parse customizations file: " + err.Error())
		os.Exit(1)
	}
}

func GetCdnUri(fileName string) string {
	if val, ok := BotCustoms.CDN.Files[fileName]; ok {
		return BotCustoms.CDN.BaseURL + val
	} else {
		Log.Error("Couldn't file an customization entry for the file " + fileName)
	}
	return ""
}
