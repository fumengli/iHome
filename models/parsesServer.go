package models

import (
	"encoding/json"
	"io/ioutil"
)

func GetUrl() (url string, err error) {
	ip, port, err := GetIpAndPort()
	url = ip + ":" + port + "/"
	return
}
func GetIPort() (iport string, err error) {
	ip, port, err := GetIpAndPort()
	iport = ip + ":" + port
	return
}
func GetIpAndPort() (IP, Port string, err error) {
	buffer, err := ioutil.ReadFile("./conf/server.conf")
	if err != nil {
		return "", "", err
	}
	infoes := make(map[string]interface{})
	err = json.Unmarshal(buffer, &infoes)
	if err != nil {
		return "", "", err
	}
	if infoes["hostIP"] == nil || infoes["port"] == nil {
		return "", "", err
	}
	IP = infoes["hostIP"].(string)
	Port = infoes["port"].(string)
	return

}
