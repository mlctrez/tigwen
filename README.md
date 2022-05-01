# tigwen

A utility for creating new golang github repositories. It should be executed in the top level of the source tree.

If `go.mod` exists it will be used to derive the github user and repository name.

```
# go module example - $SOMEPATH is outside of GOPATH
mkdir -p $SOMEPATH/github.com/username/module
cd $SOMEPATH/github.com/username/module
go mod init github.com/username/module
tigwen
```

If `go.mod` does *not* exist in the current directory, then the user and repository are derived
from the parent directories.

```
# gopath example
mkdir -p $GOPATH/src/github.com/username/module
cd $GOPATH/src/github.com/username/module
tigwen
```

The path `~/.github_token` must be a one line file containing a [github token](https://github.com/settings/tokens) that
has repo create permissions.

* The token must have the prefix `ghp_` which is the default for any recently created tokens.

### Steps that are executed

* A new remote repository is created at github.com/username/module
* A git repository is initialized in the current directory
* The files `README.md`, `LICENSE`, and `.gitignore` are added and committed
* The remote origin is set and the initial commit is pushed

### Assumptions

* Your ssh configuration either: 
  * Has a `Host github.com` entry and IdentityFile to use for github.com 
  * Does not have this entry, `~/.ssh/id_rsa` will be used for authentication.
* Your known_hosts file has host keys for github.com 
    * See https://github.com/go-git/go-git/issues/411 for information on how to configure known_hosts to work correctly
      with the go-git / go ssh implementation



