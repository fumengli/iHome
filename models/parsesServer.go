package models

import (
	"encoding/json"
	"io/ioutil"
)

func GetUrl() (url string, err error) {
	buffer, err := ioutil.ReadFile("./conf/server.conf")
	if err != nil {
		return "", err
	}
	infoes := make(map[string]interface{})
	err = json.Unmarshal(buffer, &infoes)
	if err != nil {
		return "", err
	}
	if infoes["hostIP"] == nil || infoes["port"] == nil {
		return "", err
	}
	url = infoes["hostIP"].(string) + ":" + infoes["port"].(string) + "/"
	return
}
