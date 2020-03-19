package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"royal-rumble/pkg/file"
	"royal-rumble/pkg/roman"
)

var (
	ErrInvalidCommand = errors.New("please input path to the file")
)

func main() {
	defer exception()
	lines := getLines()
	roman.Sort(lines)
	for _, line := range lines {
		fmt.Println(line)
	}
}

// check argument and process the argument
func getLines() []string {
	lines := []string{}
	var err error
	if len(os.Args) > 1 {
		lines, err = file.GetEachLine(os.Args[1])
		if err != nil {
			panic(err)
		}
	} else {
		panic(ErrInvalidCommand)
	}

	return lines
}

func exception() {
	if r := recover(); r != nil {
		log.Printf("%+v\n", r)
	}
}
