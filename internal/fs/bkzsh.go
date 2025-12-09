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

#compdef cd

_bk_cd_complete() {
    local -a dirs bk all

    dirs=(${(f)"$(printf '%s\n' -- *(N-/) )"})
    dirs=(${dirs[@]:#$PREFIX*})

    if [[ -f "$HOME/.bk" ]]; then
        bk=("${(f)$(awk -F '=' '
            NF >= 2 && $1 !~ /^[[:space:]]*#/ {print $1}
        ' "$HOME/.bk")}")
    fi

    all=("${dirs[@]}" "${bk[@]}")

    compadd -- $all
}
`

	zshInsert = `
# >>> bk init >>>
source ~/.bk.zsh
# <<< bk init <<<
`
)
