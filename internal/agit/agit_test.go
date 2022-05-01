package agit

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/mlctrez/tigwen/internal/files"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"testing"
)

func TestInit(t *testing.T) {

	req := require.New(t)

	temp, err := os.MkdirTemp(os.TempDir(), t.Name())
	req.Nil(err)
	defer func() { req.Nil(os.RemoveAll(temp)) }()

	repo, err := git.PlainInit(temp, false)
	req.NotNil(repo)
	req.Nil(err)

	getwd, err := os.Getwd()
	defer func() { req.Nil(os.Chdir(getwd)) }()
	req.Nil(os.Chdir(temp))

	err = files.WriteFile(".gitignore", nil)
	req.Nil(err)

	command := exec.Command("git", "add", ".gitignore")
	_, err = command.CombinedOutput()
	req.Nil(err)

	command = exec.Command("git", "commit", "-m", "initial commit")
	_, err = command.CombinedOutput()
	req.Nil(err)

	err = files.WriteFile("LICENSE", nil)
	req.Nil(err)

	worktree, err := repo.Worktree()
	req.Nil(err)
	_, err = worktree.Add("LICENSE")
	req.Nil(err)

	_, err = worktree.Commit("committing", &git.CommitOptions{})
	req.Nil(err)

	command = exec.Command("git", "log")
	status, err := command.CombinedOutput()
	req.Nil(err)
	fmt.Println(string(status))

}

func Test_Git(t *testing.T) {
	req := require.New(t)
	req.True(true)

	temp, err := os.MkdirTemp(os.TempDir(), t.Name())
	req.Nil(err)
	defer func() { req.Nil(os.RemoveAll(temp)) }()

	repo, err := git.PlainInit(temp, false)
	req.Nil(err)
	err = repo.CreateBranch(&config.Branch{
		Name:   "master",
		Remote: "origin",
		Merge:  "refs/heads/master",
	})
	req.Nil(err)

}
