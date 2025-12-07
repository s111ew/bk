package bk

const USAGE_TEXT = `bk — directory alias manager for zsh

Usage:
  bk <command> [arguments]

Commands:
  add   <alias> [path]     Add a new alias/path pair
  rm    <alias>            Remove an existing alias
  fix   <alias> [path]     Update the path associated with an alias
  res   <alias>            Print the path for a given alias
  list                     List all stored aliases
  --help				   Print full help text
`

const HELP_TEXT = `bk — directory alias manager for zsh

Usage:
  bk <command> [arguments]

Commands:
  add   <alias> [path]     Add a new alias/path pair
  rm    <alias>            Remove an existing alias
  fix   <alias> [path]     Update the path associated with an alias
  res   <alias>            Print the path for a given alias
  list                     List all stored aliases

Descriptions:
  add   Adds an alias and path pair to ~/.bk_aliases.
        - <alias> is required and must be unique.
        - [path] is optional; if omitted, the current working directory is used.

  rm    Removes an alias (and its stored path) from ~/.bk_aliases.
        - <alias> must already exist.

  fix   Updates the path associated with an existing alias.
        - <alias> must already exist.
        - [path] is optional; if omitted, the current working directory is used.

  res   Prints the path associated with the given alias.
        - <alias> must already exist.
        - Intended for use by the shell integration (bk_cd).

  list  Prints all aliases and their paths in a table format.

Examples:
  bk add proj                   # add alias "proj" for current directory
  bk add docs ~/Documents       # add alias "docs" for ~/Documents
  bk fix proj ~/new/location    # update "proj" path
  bk rm proj                    # remove alias "proj"
  bk res docs                   # print path for "docs"
  bk list                       # show all aliases

Notes:
  To enable 'cd <alias>' functionality, bk automatically injects a small
  integration block into your ~/.zshrc on first run.
`
