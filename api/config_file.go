package main

import (
	"bufio"
	"os"
	"strings"
)

func configFileLoad(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		i := strings.Index(line, "=")
		if i == -1 {
			continue
		}

		key := line[:i]
		value := line[i+1:]

		if !strings.HasPrefix(key, "COMMENTO_") {
			continue
		}

		os.Setenv(key[9:], value)
	}

	return nil
}
