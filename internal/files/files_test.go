package files

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile_Plain(t *testing.T) {
	req := require.New(t)

	temp, err := os.MkdirTemp(os.TempDir(), t.Name())
	req.Nil(err)
	defer func() { req.Nil(os.RemoveAll(temp)) }()
	getwd, err := os.Getwd()
	defer func() { req.Nil(os.Chdir(getwd)) }()

	req.Nil(os.Chdir(temp))

	name := ".gitignore"

	err = WriteFile(name, nil)
	req.Nil(err)

	filedata, err := os.ReadFile(filepath.Join(temp, name))
	req.Nil(err)

	expected, err := Contents(name)
	req.Nil(err)

	req.Equal(expected, filedata)

}

func TestWriteFile_Template(t *testing.T) {
	req := require.New(t)

	temp, err := os.MkdirTemp(os.TempDir(), t.Name())
	req.Nil(err)
	defer func() { req.Nil(os.RemoveAll(temp)) }()
	getwd, err := os.Getwd()
	defer func() { req.Nil(os.Chdir(getwd)) }()

	req.Nil(os.Chdir(temp))

	name := "LICENSE"

	err = WriteFile(name, map[string]string{
		"LicenseCopyright": "@@ COPY TEXT @@",
	})
	req.Nil(err)

	filedata, err := os.ReadFile(filepath.Join(temp, name))
	req.Nil(err)

	req.Contains(string(filedata), "@@ COPY TEXT @@")

}
