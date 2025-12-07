package fs

import (
	"os"
	"strings"
)

func Setup(aliasFilePath, configFilePath string) error {
	if err := MakeAliasFileIfNotExists(aliasFilePath); err != nil {
		return err
	}

	if err := EnsureZshrcConfigured(configFilePath); err != nil {
		return err
	}

	return nil
}

func MakeAliasFileIfNotExists(aliasFilePath string) error {
	_, err := os.Stat(aliasFilePath)

	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		f, err := os.Create(aliasFilePath)
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
source ~/.bk
bk_cd() {
	local resolved
	resolved=$(bk --resolve "$1" 2>/dev/null)
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
