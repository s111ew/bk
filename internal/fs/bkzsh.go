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

    # Standard directory completion
    dirs=(${(f)"$(compgen -d -- "$PREFIX")"})

    # bk aliases (first column)
    if [[ -f "$HOME/.bk" ]]; then
        bk=("${(f)$(awk '{print $1}' "$HOME/.bk")}")
    fi

    # Merge
    all=("${dirs[@]}" "${bk[@]}")

    compadd -- $all
}

_bk_cd_complete "$@"
`

	zshInsert = `
# >>> bk init >>>
source ~/.bk.zsh
# <<< bk init <<<
`
)
