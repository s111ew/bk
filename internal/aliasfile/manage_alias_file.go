package aliasfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
	return fmt.Sprintf("alias %s=%s", alias, path)
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
