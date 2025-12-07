package files

import (
	"bytes"
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
