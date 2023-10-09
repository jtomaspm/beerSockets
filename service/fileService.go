package service

import (
	"bufio"
	"os"
	"strings"
)

type File struct {
	name    string
	content string
}

func stripFileExtention(fileName string) string {
	return strings.Split(fileName, ".beer")[0]
}

func SaveFile(path string, fileName string, content string) error {
	file, err := os.Create(path + fileName + ".beer")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}

func ReadFile(path string, fileName string) (File, error) {
	result := File{}
	resultBytes, err := os.ReadFile(path + fileName + ".beer")
	if err != nil {
		return result, err
	}
	result.name = fileName
	result.content = string(resultBytes)
	return result, nil
}

func ReadAllFile(path string) ([]File, error) {
	res, e := os.ReadDir(path)
	if e != nil {
		return []File{}, e
	}
	base := [100]File{}
	result := base[:0]
	for _, l := range res {
		if !l.IsDir() {
			file, e := ReadFile(path, stripFileExtention(l.Name()))
			if e != nil {
				return []File{}, e
			}
			result = append(result, file)
		}
	}
	return result, nil
}
