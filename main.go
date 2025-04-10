package main

import (
	"debug/buildinfo"
	"encoding/json"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gobin-dump <binary path>")
		os.Exit(1)
	}

	bi, err := buildinfo.ReadFile(os.Args[1])
	check(err)

	data, err := json.MarshalIndent(bi, "", "  ")
	check(err)

	fmt.Printf("%s", data)
}
