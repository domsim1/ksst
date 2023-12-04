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

// TODO: Support INI format once mature
/**** INI HEX DATA Scheme ****
--- Entry ---
* 8 bytes: not sure..
* 4 bytes: 1, mark end of value
* 4 bytes: length of first key
--- Data ---
* variable length: key for value, size provide from previous 4 bytes
* 4 bytes: value flag, if 1 string else 0 float64
* if string:
*   4 bytes: size of string
*   variable length: value
* else float64:
*   8 bytes (big-endian): value
* 4 bytes: 1, mark end of value
* 4 bytes: length of next key
*/
