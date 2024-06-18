package main

import (
	"fmt"
)

func EncodeRESPSimpleString(str string) string {
	return fmt.Sprintf("+%s\r\n", str)
}
