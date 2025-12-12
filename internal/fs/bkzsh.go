package fs

const (
	bkZshContents = `
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
`

	zshInsert = `
# >>> bk init >>>
source ~/.bk.zsh
# <<< bk init <<<
`
)
