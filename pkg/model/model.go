package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

var FixJson = strings.NewReplacer("{", "{ ", ":", ": ", ",", ", ", "}", " }").Replace

func ConvertStringData(data string) (*SaveData, string, error) {
	sd := SaveData{}

	i := strings.Index(data, "{ \"")
	prefix := data[:i]
	err := json.Unmarshal([]byte(data[i:len(data)-1]), &sd)
	if err != nil {
		return nil, "", err
	}
	return &sd, prefix, nil
}

func ConvertModelToStringData(data *SaveData, prefix string) (string, error) {
	rawData, err := json.Marshal(data)
	strData := string(rawData)
	strData = FixJson(strData)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s\x00", prefix, strData), nil
}
