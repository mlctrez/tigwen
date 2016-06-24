# tigwen

A utility for creating new golang github repositories.

This utility assumes
* `$GOPATH` is part of the current directory
* `~/.github_token` contains a github app token that has repo create permissions

Example Usage

```
mkdir -p $GOPATH/src/github.com/username/pkgname
cd $GOPATH/src/github.com/username/pkgname
tigwen
```

What it does

* A new repository is created at github.com/username/pkgname
* A git repository is initialized in the directory `$GOPATH/src/github.com/username/pkgname`
* The files `README.md`, `LICENSE`, and `.gitignore` are added and committed
* The remote origin is set and the initial commit is pushed

TODO:

* Make the LICENSE and .gitignore content configurable and not hardcoded in the go source


