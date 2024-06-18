package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const lineBreak = "\r\n"

func ParseRequest(str string) (int64, [][]string) {
	lines := strings.Split(str, lineBreak)
	argsRe := regexp.MustCompile(`\d+`)
	argsParam := lines[0]
	argsNum, err := strconv.ParseInt(argsRe.FindString(argsParam), 10, 32)

	if err != nil {
		fmt.Println("Error: Invalid number of arguments: ", err.Error())
		os.Exit(1)
	}

	var requestArgs [][]string = [][]string{}
	for i := int64(1); i <= argsNum; i++ {
		requestArgs = append(requestArgs, []string{lines[2*i-1], lines[2*i]})
	}
	return argsNum, requestArgs
}
