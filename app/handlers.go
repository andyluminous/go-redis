package main

import (
	"strings"
)

var Handlers map[string]func(int64, [][]string) string = map[string]func(int64, [][]string) string{
	"COMMAND": command,
	"ECHO":    echo,
	"PING":    ping,
}

func ping(numberOfArgs int64, args [][]string) string {
	return EncodeRESPSimpleString("PONG")
}

func echo(numberOfArgs int64, args [][]string) string {
	var echoParams []string = []string{}
	for _, arg := range args[1:] {
		echoParams = append(echoParams, arg[1])
	}
	return EncodeRESPSimpleString(strings.Join(echoParams, " "))
}

func command(numberOfArgs int64, args [][]string) string {
	return EncodeRESPSimpleString("Welcome!")
}
