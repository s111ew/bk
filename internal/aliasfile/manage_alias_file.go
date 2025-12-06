package aliasfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Alias struct {
	alias string
	path  string
}

func MakeAliasFileIfNotExists(path string) error {
	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return nil
	}

	return err
}

func ConstructAliasString(alias, path string) string {
	return fmt.Sprintf("%s=%s", alias, path)
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

func LoadAliases() ([]Alias, error) {
	res, err := os.ReadFile("~/.bk")
	if err != nil {
		return nil, err
	}

	aliasStrings := strings.Split(strings.TrimSpace(string(res)), "\n")

	var aliases []Alias

	for _, a := range aliasStrings {
		al := strings.Split(strings.TrimSpace(a), "=")
		alias := Alias{al[0], al[1]}
		aliases = append(aliases, alias)
	}

	return aliases, nil
}
