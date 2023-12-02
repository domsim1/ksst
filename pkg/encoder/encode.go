package encoder

import (
	"encoding/base64"
)

func DecodeData(data []byte) (string, error) {
	dat, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func EncodeData(data []byte) string {
	dat := base64.StdEncoding.EncodeToString(data)
	return dat
}
