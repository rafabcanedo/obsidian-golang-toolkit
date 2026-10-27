package config

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadEnv(path string) (map[string]string, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := map[string]string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")

		if !found {
			continue
		}

		result[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}

	return result, scanner.Err()
}

func LoadJSON(path string, v any) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()

    return json.NewDecoder(f).Decode(v)
}

func LoadYAML(path string, v any) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()

    return yaml.NewDecoder(f).Decode(v)
}
