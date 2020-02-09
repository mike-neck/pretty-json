package main

import (
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("usage: pretty-json <json string>")
	}
	jsonString := fromArg(os.Args)
	out, err := prettyjson.Format([]byte(jsonString))
	if err != nil {
		log.Fatalln("failed to format", err)
	}
	fmt.Println(string(out))
}

func fromArg(args []string) string {
	return args[1]
}
