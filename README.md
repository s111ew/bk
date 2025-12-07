bk — directory alias manager for zsh
=========================================================================


bk is a dependency free command line tool written in Go that provides an easy way 
for zsh users to manage their directory aliases by adding/removing/updating/listing 
aliases without having to manually edit .zshrc.

bk also provides a modification to the existing 'cd' command, allowing users
to 'cd' into aliased directories without the use of the '~' prefix.


Example Workflow:
-------------------------------------------------------------------------

  Navigate to a frequently used directory that you'd like to bookmark, then run:

  bk --add my-fave-dir

  Now you can navigate to that directory from anywhere in your file system by
  simply running:

  cd my-fave-dir


Installation:
-------------------------------------------------------------------------

  If you have Go 1.20+ installed:
  -----------------------------------------------------------------------
  From your terminal run:

  go install github.com/s111ew/bk@latest

  This installs the bk binary into your $GOBIN (or $HOME/go/bin if GOBIN is not set).
  Ensure $GOBIN or $HOME/go/bin is in your PATH:

  export PATH=$PATH:$(go env GOPATH)/bin

  Now you can run:

  bk --help
  bk --add my-dir
  etc.


  If you have Go 1.20+ installed but prefer to build from source:
  -----------------------------------------------------------------------
  Clone the repository:

  git clone github.com/s111ew/bk.git
  cd bk

  Build the binary:

  go build -o bk

  Move the binary to your PATH:

  sudo mv bk /usr/local/bin/
  

Usage:
-------------------------------------------------------------------------
  bk <command> [arguments]


Commands:
-------------------------------------------------------------------------
  -a  --add     <alias> [path]     Add a new alias/path pair
  -r  --remove  <alias>            Remove an existing alias
  -u  --update  <alias> [path]     Update the path associated with an alias
  -g  --get     <alias>            Print the path for a given alias
  -l  --list                       List all stored aliases
  -h  --help                       Print full help text


Descriptions:
-------------------------------------------------------------------------
  add     Adds an alias and path pair.
        - <alias> is required and must be unique.
        - [path] is optional; if omitted, the current working directory is used.

  remove  Removes an alias (and its stored path).
        - <alias> must already exist.

  update  Updates the path associated with an existing alias.
        - <alias> must already exist.
        - [path] is optional; if omitted, the current working directory is used.

  resolve Prints the path associated with the given alias.
        - <alias> must already exist.

  list    Prints all aliases and their paths in a table format.


Notes:
-------------------------------------------------------------------------
  To enable 'cd <alias>' functionality, bk automatically injects a small
  integration block into your ~/.zshrc on first run.