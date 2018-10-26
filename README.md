# tigwen

A utility for creating new golang github repositories. It should be executed in the top level of the source tree only.

If go.mod exists in the current directory it will be used to derive the github user and repository name.
```
# go module example - $SOMEPATH is outside of GOPATH
mkdir -p $SOMEPATH/github.com/username/pkgname
cd $SOMEPATH/github.com/username/pkgname
go mod init github.com/username/pkgname
tigwen
```

If go.mod does *not* exist in the current directory, then <somdir>/github.com/<user>/<repo>
```
# gopath example
mkdir -p $GOPATH/src/github.com/username/pkgname
cd $GOPATH/src/github.com/username/pkgname
tigwen
```
 
The path `~/.github_token` must contain a github app token that has repo create permissions.

Steps that are executed

* A new repository is created at github.com/username/pkgname
* A git repository is initialized in the current directory
* The files `README.md`, `LICENSE`, and `.gitignore` are added and committed
* The remote origin is set and the initial commit is pushed

TODO:

* Make the LICENSE and .gitignore content configurable and not hardcoded in the go source


