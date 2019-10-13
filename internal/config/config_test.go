package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPath = "../../test"
	fakePath = "/foo/bar/baz"
)

func TestGet(t *testing.T) {
	t.Run("load success", func(t *testing.T) {
		err := Load(testPath)

		assert.NoError(t, err)
		assert.Equal(t, Get("host"), "https://abc.com")
		assert.Equal(t, Get("token"), "NjEaWdDcARhzYKdx4fA4")
	})

	t.Run("load fail", func(t *testing.T) {
		err := Load(fakePath)

		assert.Error(t, err)
		assert.Equal(t, Get("host"), "https://gitlab.com")
		assert.Equal(t, Get("token"), "None")
	})
}

func TestConfigure(t *testing.T) {
	t.Run("skip all", func(t *testing.T) {
		Load(testPath)

		stdin := bytes.NewBufferString("\n\n")
		stdout := &bytes.Buffer{}
		Configure(stdin, stdout)

		filePath := path.Join(testPath, "golab.yaml")
		want := "Gitlab Host [https://abc.com]: Gitlab Token (scope: api) [NjEaWdDcARhzYKdx4fA4]: \nConfig saved to " + filePath + "\n"
		got := stdout.String()

		assert.Equal(t, want, got)
	})

	t.Run("enter something", func(t *testing.T) {
		dir, _ := ioutil.TempDir("", "golab")
		defer os.RemoveAll(dir)

		Load(dir)
		stdin := bytes.NewBufferString("https://foo.com\nfaketoken\n")
		stdout := &bytes.Buffer{}
		Configure(stdin, stdout)

		filePath := path.Join(dir, "golab.yaml")
		want := "Gitlab Host [https://gitlab.com]: Gitlab Token (scope: api) [None]: \nConfig saved to " + filePath + "\n"
		got := stdout.String()

		assert.Equal(t, want, got)
		assert.FileExists(t, filePath)

		Load(dir)

		assert.Equal(t, Get("host"), "https://foo.com")
		assert.Equal(t, Get("token"), "faketoken")
	})
}
