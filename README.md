# tigwen

A utility for creating new golang github repositories.

This utility assumes
 * `$GOPATH` is set correctly and that it is part of the current path for the go package being added
 * `~/.github_token` contains a github application token that has the repository create permission
 * The current working directory is a golang package directory

 Example Usage

 ```
 mkdir -p $GOPATH/src/github.com/username/pkgname
 cd $GOPATH/src/github.com/username/pkgname
 tigwen
 ```

 * The github token is used to create a new repository at github.com/username/pkgname
 * A new git repo is initializes in the directory `$GOPATH/src/github.com/username/pkgname`
 * The files `README.md`, `LICENSE`, and `.gitignore` are added and committed
 * The remote origin is set and the initial commit is pushed


