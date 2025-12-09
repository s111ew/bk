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

_bk_cd_complete() {
    local -a bk_keys bk_display

    if [[ -f "$HOME/.bk" ]]; then
        bk_keys=("${(f)$(awk -F '=' '
            NF >= 2 && $1 !~ /^[[:space:]]*#/ {print $1}
        ' "$HOME/.bk")}")
    fi

    bk_display=("${(@)bk_keys/#/~}")

    _arguments '1:directory:_files -/'

    if (( ${#bk_keys[@]} )); then
        _describe -t bookmarks bookmarks bk_display bk_keys
    fi
}

compdef _bk_cd_complete cd bk_cd
`

	zshInsert = `
# >>> bk init >>>
source ~/.bk.zsh
# <<< bk init <<<
`
)
