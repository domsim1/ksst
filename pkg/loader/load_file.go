package loader

import (
	"os"

	"github.com/domsim1/ksst/pkg/util"
)

func LoadRawFileData(path string) []byte {
	data, err := os.ReadFile(path)
	util.Check(err)
	return data
}
