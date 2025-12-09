package fs

import (
	"bufio"
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

func LoadSingleAlias(aliasToFind, aliasFilePath string) (Alias, error) {
	file, err := os.Open(aliasFilePath)
	if err != nil {
		return Alias{Name: "", Path: ""}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		alias := line[0]
		path := line[1]

		if alias == aliasToFind {
			found := Alias{
				Name: alias,
				Path: path,
			}

			return found, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return Alias{Name: "", Path: ""}, err
	}

	return Alias{Name: "", Path: ""}, nil
}

func LoadAliases(aliasFilePath string) ([]Alias, error) {
	res, err := os.ReadFile(aliasFilePath)
	if err != nil {
		return nil, err
	}

	aliases := bytesToAliases(res)

	return aliases, nil
}

func WriteAliases(aliases []Alias, aliasFilePath string) error {
	f, err := os.OpenFile(aliasFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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
