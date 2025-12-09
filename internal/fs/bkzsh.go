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

source ~/.bk

bk_cd() {
    local resolved
    resolved=$(bk --resolve "$1" 2>/dev/null)
    if [[ -n "$resolved" ]]; then
        builtin cd "$resolved"
    else
        builtin cd "$1"
    fi
}
alias cd=bk_cd

_bk_cd_complete() {
    local -a dirs bk all
    local cur

    cur="${words[CURRENT]}"

    dirs=(${(f)"$(print -rC1 -- *(N-/))"})
    if [[ -n $cur ]]; then
        dirs=(${(M)dirs:#$cur*})
    fi

    if [[ -f "$HOME/.bk" ]]; then
        bk=("${(f)$(awk -F '=' 'NF>=2 && $1 !~ /^[[:space:]]*#/ {print $1}' "$HOME/.bk")}")
        if [[ -n $cur ]]; then
            bk=(${(M)bk:#$cur*})
        fi
    fi

    all=("${dirs[@]}" "${bk[@]}")

    compadd -Q -- $all
}

compdef _bk_cd_complete cd bk_cd

`

	zshInsert = `
# >>> bk init >>>
source ~/.bk.zsh
# <<< bk init <<<
`
)
