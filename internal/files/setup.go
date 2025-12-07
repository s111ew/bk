package files

import (
	"os"
	"strings"
)

func MakeAliasFileIfNotExists(alias_file_path string) error {
	_, err := os.Stat(alias_file_path)

	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		f, err := os.Create(alias_file_path)
		if err != nil {
			return err
		}
		return f.Close()
	}

	return err
}

func EnsureZshrcConfigured(configFilePath string) error {
	data, _ := os.ReadFile(configFilePath)
	contents := string(data)

	if strings.Contains(contents, "# >>> bk init >>>") {
		return nil
	}

	block := `
# >>> bk init >>>
source ~/.bk_aliases
bk_cd() {
	local resolved
	resolved=$(bk res "$1" 2>/dev/null)
	if [ -n "$resolved" ]; then
		builtin cd "$resolved"
	else
		builtin cd "$1"
	fi
}
alias cd=bk_cd
# <<< bk init <<<
		`

	f, err := os.OpenFile(configFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(block)
	return err
}
