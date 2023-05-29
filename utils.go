package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ParseBool(value string) bool {
	lower := strings.ToLower(value)
	return lower == "true" || lower == "1" || lower == "yes" || lower == "y"
}

func GetFilesFromDirectory(directory string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info == nil || info.IsDir() {
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == extension {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("ディレクトリの走査中にエラーが発生しました: %v", err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("指定されたディレクトリには%vファイルが見つかりませんでした", extension)
	}

	return files, nil
}

func ProcessCSVFile(file string) (string, error) {
	tables, err := ParseCSV(file)
	if err != nil {
		return "", err
	}
	ddl, err := GenerateDDL(tables)
	if err != nil {
		return "", err
	}
	return ddl, nil
}

func WriteToFile(filename string, content string) error {
	if err := ioutil.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("ファイルへの書き込み中にエラーが発生しました: %v", err)
	}
	return nil
}
