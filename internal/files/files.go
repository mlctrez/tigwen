package files

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"os"
)

//go:embed .gitignore LICENSE README.md
var emb embed.FS

func maybeCaptureCloseError(err *error, out io.WriteCloser) func() {
	return func() {
		if ce := out.Close(); err == nil {
			*err = ce
		}
	}
}

func Contents(name string) ([]byte, error) {
	return fs.ReadFile(emb, name)
}

func WriteFile(name string, data map[string]string) (err error) {

	var out io.WriteCloser

	out, err = os.Create(name)

	defer maybeCaptureCloseError(&err, out)

	var tmp *template.Template
	if tmp, err = template.ParseFS(emb, name); err != nil {
		return
	}
	err = tmp.ExecuteTemplate(out, name, data)
	return
}
