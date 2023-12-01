package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/domsim1/ksst/pkg/encoder"
	"github.com/domsim1/ksst/pkg/loader"
	"github.com/domsim1/ksst/pkg/util"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		exitWithUsage()
	}
	cmd := args[0]
	sourcePath := args[1]
	destinationPath := args[2]

	fmt.Printf("loading file: %s\n", sourcePath)
	data := loader.LoadRawFileData(sourcePath)
	var strData string
	switch strings.ToLower(cmd) {
	case "encode":
		fmt.Println("encoding data")
		strData = encoder.EncodeData(data)
	case "decode":
		fmt.Println("decoding data")
		strData = encoder.DecodeData(data)
	default:
		exitWithUsage()
	}

	if len(strData) < 1 {
		exitWithUsage()
	}

	fmt.Printf("saving file: %s\n", destinationPath)
	err := os.WriteFile(destinationPath, []byte(strData), 0644)
	util.Check(err)
	fmt.Println("done!")
}

func exitWithUsage() {
	fmt.Println("Usage: ksst-encoder [encode|decode] <source_file_path> <destination_file_path>")
	os.Exit(1)
}
