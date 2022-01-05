package io

import "os"

func WriteFile(filepath string, fileContent string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fileContent)
	return err
}
