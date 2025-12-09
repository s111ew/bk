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
    local -a bk_keys
    local expl

    # Load alias keys from ~/.bk
    if [[ -f "$HOME/.bk" ]]; then
        bk_keys=("${(f)$(awk -F '=' 'NF>=2 && $1 !~ /^[[:space:]]*#/ {print $1}' "$HOME/.bk")}")
    fi

    # Use zsh native _arguments
    # - First argument: complete either a directory or a bk alias
    _arguments \
        '1:directory:_files -/' \
        && compadd -a bk_keys
}

# Register completion for both cd and bk_cd
compdef _bk_cd_complete cd bk_cd

`

	zshInsert = `
# >>> bk init >>>
source ~/.bk.zsh
# <<< bk init <<<
`
)
