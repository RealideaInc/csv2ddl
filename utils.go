package main

import (
	"io/ioutil"
	"strings"
)

func ParseBool(value string) bool {
	lower := strings.ToLower(value)
	return lower == "true" || lower == "1" || lower == "yes" || lower == "y"
}

func WriteToFile(filename string, content string) error {
	bytes := []byte(content)
	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		return err
	}
	return nil
}
