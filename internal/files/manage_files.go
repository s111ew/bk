package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Alias struct {
	Name string
	Path string
}

func (a Alias) String() string {
	return fmt.Sprintf("%s=%s\n", a.Name, a.Path)
}

func MakeAliasFileIfNotExists(alias_file_path string) error {
	_, err := os.Stat(alias_file_path)

	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return nil
	}

	return err
}

func EnsureZshrcConfigured() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	zshrcPath := filepath.Join(home, ".zshrc")

	data, _ := os.ReadFile(zshrcPath)
	contents := string(data)

	if strings.Contains(contents, "# >>> bk init >>>") {
		return nil
	}

	block := `
		# >>> bk init >>>
		source ~/.bk_aliases
		bk_cd() {
			local resolved
			resolved=$(bk resolve "$1" 2>/dev/null)
			if [ -n "$resolved" ]; then
				builtin cd "$resolved"
			else
				builtin cd "$1"
			fi
		}
		alias cd=bk_cd
		# <<< bk init <<<
		`

	f, err := os.OpenFile(zshrcPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(block)
	return err
}

func LoadAliases(alias_file_path string) ([]Alias, error) {
	res, err := os.ReadFile(alias_file_path)
	if err != nil {
		return nil, err
	}

	aliases := bytesToAliases(res)

	return aliases, nil
}

func WriteAliases(alias_file_path string, aliases []Alias) error {
	f, err := os.OpenFile(alias_file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(aliasesToBytes(aliases))
	if err != nil {
		return err
	}
}

func aliasesToBytes(aliases []Alias) []byte {
	var byteStream []byte

	for _, a := range aliases {
		byteStream = append(byteStream, []byte(a.String())...)
	}

	return byteStream
}

func bytesToAliases(bytes []byte) []Alias {
	aliasStrings := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	var aliases []Alias

	for _, a := range aliasStrings {
		al := strings.Split(strings.TrimSpace(a), "=")
		alias := Alias{
			Name: al[0],
			Path: al[1],
		}
		aliases = append(aliases, alias)
	}

	return aliases
}
