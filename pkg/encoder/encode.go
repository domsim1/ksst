package encoder

import (
	"encoding/base64"

	"github.com/domsim1/ksst/pkg/util"
)

func DecodeData(data []byte) string {
	dat, err := base64.StdEncoding.DecodeString(string(data))
	util.Check(err)
	return string(dat)
}

func EncodeData(data []byte) string {
	dat := base64.StdEncoding.EncodeToString(data)
	return dat
}
