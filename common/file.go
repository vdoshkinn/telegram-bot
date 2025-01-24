package common

import (
	"bufio"
	"log"
	"os"
)

func ReadStringFromFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	result := ""
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}
	return result
}

func SaveStringToFile(path string, stringToSave string) error {
	err := os.WriteFile(path, []byte(stringToSave), 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func AppendToFile(path string, stringToSave string) {
	readString := ReadStringFromFile(path)
	readString += "\n" + stringToSave
	SaveStringToFile(path, readString)
}

func CreateEmptyFile(path string) error {
	return SaveStringToFile(path, "")
}
