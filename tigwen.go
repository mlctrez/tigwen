package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/google/go-github/v25/github"
	"github.com/kevinburke/ssh_config"
	"github.com/mitchellh/go-homedir"
	"github.com/mlctrez/tigwen/internal/files"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func githubToken() (token string, err error) {
	var u *user.User
	u, err = user.Current()
	if err != nil {
		return "", err
	}

	tokenFile := filepath.Join(u.HomeDir, ".github_token")
	tokenBytes, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}

	token = strings.TrimSpace(string(tokenBytes))

	if !strings.HasPrefix(token, "ghp_") {
		return "", fmt.Errorf(".github_token must be new form and start with ghp_")
	}

	return token, nil
}

func githubClient() (client *github.Client, err error) {

	token, err := githubToken()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.TODO(), ts)

	client = github.NewClient(tc)

	return client, err
}

func getUserAndRepo() (user, repo string, err error) {
	if _, ferr := os.Stat("go.mod"); os.IsNotExist(ferr) {

		var dir string
		if dir, err = os.Getwd(); err != nil {
			return
		}

		repo = filepath.Base(dir)
		dir = filepath.Dir(dir)
		user = filepath.Base(dir)
		dir = filepath.Dir(dir)
		if filepath.Base(dir) != "github.com" {
			err = errors.New("grandparent of current directory is not github.com")
		}
		return
	}
	var file *os.File
	file, err = os.Open("go.mod")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`module github.com/(\w*)/(\w*)`)
	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())
		if len(match) == 3 {
			user = match[1]
			repo = match[2]
			break
		}
	}
	err = scanner.Err()

	if user == "" || repo == "" {
		err = errors.New("unable to determine github user or repo: check go.mod file")
	}

	return
}

func main() {

	repoPath, err := os.Getwd()
	checkErr(err)

	if _, err = os.Stat(filepath.Join(repoPath, ".git")); err == nil {
		log.Fatal("existing .git directory found")
	}

	userName, repoName, err := getUserAndRepo()
	checkErr(err)

	client, err := githubClient()
	checkErr(err)

	repo := &github.Repository{Name: &repoName}
	repo.Name = &repoName

	repository, response, err := client.Repositories.Create(context.Background(), "", repo)
	checkErr(err)

	_ = repository
	_ = response

	goModule := strings.Join([]string{"github.com", userName, repoName}, "/")

	checkErr(files.WriteFile("README.md", map[string]string{"RepoName": repoName, "GoModule": goModule}))

	checkErr(files.WriteFile("LICENSE", map[string]string{
		"LicenseCopyright": fmt.Sprintf("Copyright %d Matt Crawford", time.Now().Year()),
	}))

	checkErr(files.WriteFile(".gitignore", nil))

	gitRepo, err := git.PlainInit(repoPath, false)
	checkErr(err)

	worktree, err := gitRepo.Worktree()
	checkErr(err)

	for _, name := range []string{"README.md", ".gitignore", "LICENSE"} {
		_, err = worktree.Add(name)
		checkErr(err)
	}
	_, err = worktree.Commit("initial commit", &git.CommitOptions{})
	checkErr(err)

	err = gitRepo.CreateBranch(&config.Branch{Name: "master"})
	checkErr(err)

	_, err = gitRepo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{fmt.Sprintf("git@github.com:%v/%v.git", userName, repoName)},
	})
	checkErr(err)

	identity := ssh_config.Get("github.com", "IdentityFile")
	if identity == "" {
		identity = "~/.git/id_rsa"
	}

	expand, err := homedir.Expand(identity)
	checkErr(err)

	fmt.Println("using identity file", identity)

	callback, err := ssh.NewKnownHostsCallback()
	checkErr(err)

	auth, err := ssh.NewPublicKeysFromFile("git", expand, "")
	checkErr(err)
	clientConfig, err := auth.ClientConfig()
	checkErr(err)
	clientConfig.HostKeyCallback = callback

	fmt.Println(auth)

	err = gitRepo.Push(&git.PushOptions{
		Auth:       auth,
		RemoteName: "origin",
	})
	checkErr(err)

}
