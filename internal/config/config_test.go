package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("get correct default value", func(t *testing.T) {
		c := New()

		assert.Equal(t, c.Get("host"), "https://gitlab.com")
		assert.Equal(t, c.Get("token"), "None")
	})
}

func TestLoad(t *testing.T) {
	t.Run("load success", func(t *testing.T) {
		c := New()
		err := c.Load(getTestPath())

		assert.NoError(t, err)
		assert.Equal(t, c.Get("host"), "https://abc.com")
		assert.Equal(t, c.Get("token"), "NjEaWdDcARhzYKdx4fA4")
	})

	t.Run("load fail", func(t *testing.T) {
		c := New()

		fakePath := "/foo/bar/baz"
		err := c.Load(fakePath)

		assert.Error(t, err)
	})
}

func TestEdit(t *testing.T) {
	t.Run("skip all input and get default value", func(t *testing.T) {
		dir, _ := ioutil.TempDir("", "golab")
		defer os.RemoveAll(dir)

		c := New()

		in := bytes.NewBufferString("\n\n")
		out := &bytes.Buffer{}
		c.Edit(dir, in, out)

		want := "Gitlab Host [https://gitlab.com]: Gitlab Token (scope: api) [None]: \nConfig saved to " + c.viper.ConfigFileUsed() + "\n"
		got := out.String()

		assert.Equal(t, want, got)
		assert.FileExists(t, c.viper.ConfigFileUsed())
		assert.Equal(t, c.Get("host"), "https://gitlab.com")
		assert.Equal(t, c.Get("token"), "None")
	})

	t.Run("enter something and get enter value", func(t *testing.T) {
		dir, _ := ioutil.TempDir("", "golab")
		defer os.RemoveAll(dir)

		c := New()

		in := bytes.NewBufferString("https://foo.com\nfaketoken\n")
		out := &bytes.Buffer{}
		c.Edit(dir, in, out)

		want := "Gitlab Host [https://gitlab.com]: Gitlab Token (scope: api) [None]: \nConfig saved to " + c.viper.ConfigFileUsed() + "\n"
		got := out.String()

		assert.Equal(t, want, got)
		assert.FileExists(t, c.viper.ConfigFileUsed())
		assert.Equal(t, c.Get("host"), "https://foo.com")
		assert.Equal(t, c.Get("token"), "faketoken")
	})
}

func TestList(t *testing.T) {
	c := New()

	buf := &bytes.Buffer{}
	c.List(buf)

	got := buf.String()

	assert.Contains(t, got, "NAME")
	assert.Contains(t, got, "VALUE")
	assert.Contains(t, got, "host")
	assert.Contains(t, got, "https://gitlab.com")
	assert.Contains(t, got, "host")
	assert.Contains(t, got, "None")
}

func getTestPath() string {
	dir, _ := os.Getwd()

	return dir + "/../../test"
}
