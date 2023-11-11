package utils

import (
	"log"
	"strings"
)

func StringToBool(s string) bool {
	if s == "1" {
		return true
	}

	if strings.ToLower(s) == "true" {
		return true
	}

	if s == "0" {
		return false
	}

	if strings.ToLower(s) == "false" {
		return false
	}

	log.Printf("The string is %s will be turned to false", s)
	return false
}
