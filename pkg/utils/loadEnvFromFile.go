package utils

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnvFromFile(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]

			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
