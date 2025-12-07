package files

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Alias struct {
	Name string
	Path string
}

func (a Alias) String() string {
	return fmt.Sprintf("%s=%s\n", a.Name, a.Path)
}

func LoadAliases(alias_file_path string) ([]Alias, error) {
	res, err := os.ReadFile(alias_file_path)
	if err != nil {
		return nil, err
	}

	aliases := bytesToAliases(res)

	return aliases, nil
}

func WriteAliases(aliases []Alias, alias_file_path string) error {
	f, err := os.OpenFile(alias_file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(aliasesToBytes(aliases))
	if err != nil {
		return err
	}

	return nil
}

func aliasesToBytes(aliases []Alias) []byte {
	var buf bytes.Buffer

	for _, a := range aliases {
		fmt.Fprint(&buf, a.String())
	}

	return buf.Bytes()
}

func bytesToAliases(b []byte) []Alias {
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")

	var aliases []Alias

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		aliases = append(aliases, Alias{
			Name: parts[0],
			Path: parts[1],
		})
	}

	return aliases
}
