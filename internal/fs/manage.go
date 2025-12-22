package fs

import (
	"bufio"
	"bytes"
	"errors"
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

func LoadOne(aliasToFind, aliasFilePath string) (Alias, error) {
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

func WriteOne(aliasNameToAdd, aliasPathToAdd, aliasFilePath string) error {
	aliases, err := LoadAll(aliasFilePath)
	if err != nil {
		return err
	}

	newAlias := Alias{Name: aliasNameToAdd, Path: aliasPathToAdd}

	for _, a := range aliases {
		if a.Name == newAlias.Name {
			return errors.New(fmt.Sprintf("bk: alias '%s' exists", aliasNameToAdd))
		}

		if a.Path == newAlias.Path {
			return errors.New(fmt.Sprintf("bk: alias for path '%s' exists", aliasPathToAdd))
		}
	}

	aliases = append(aliases, newAlias)

	if err := WriteAll(aliases, aliasFilePath); err != nil {
		return err
	}

	return nil
}

func RemoveOne(aliasName, aliasFilePath string) error {
	aliases, err := LoadAll(aliasFilePath)
	if err != nil {
		return err
	}

	var newAliases []Alias

	for _, a := range aliases {
		if a.Name != aliasName {
			newAliases = append(newAliases, a)
		}
	}

	if err := WriteAll(newAliases, aliasFilePath); err != nil {
		return err
	}

	return nil
}

func LoadAll(aliasFilePath string) ([]Alias, error) {
	res, err := os.ReadFile(aliasFilePath)
	if err != nil {
		return nil, err
	}

	aliases := bytesToAlias(res)

	return aliases, nil
}

func WriteAll(aliases []Alias, aliasFilePath string) error {
	f, err := os.OpenFile(aliasFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(aliasToBytes(aliases))
	if err != nil {
		return err
	}

	return nil
}

func aliasToBytes(aliases []Alias) []byte {
	var buf bytes.Buffer

	for _, a := range aliases {
		fmt.Fprint(&buf, a.String())
	}

	return buf.Bytes()
}

func bytesToAlias(b []byte) []Alias {
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
