package main

import (
	"github.com/DanielLiu1123/xcli/pkg/cmd"
	"github.com/DanielLiu1123/xcli/pkg/model"
	"log"
)

var version = "0.0.2"

func main() {

	log.SetFlags(0)

	b := &model.BuildInfo{
		Version: version,
	}

	err := cmd.NewCmdRoot(b).Execute()
	if err != nil {
		log.Fatal(err)
	}
}
