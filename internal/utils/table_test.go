package utils

import (
	"reflect"
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	table := NewTable(nil)
	elem := reflect.ValueOf(table).Elem()

	assert.Equal(t, int64(1000), elem.FieldByName("mW").Int())
	assert.Equal(t, int64(tablewriter.ALIGN_LEFT), elem.FieldByName("hAlign").Int())
	assert.Equal(t, " ", elem.FieldByName("pCenter").String())
	assert.Equal(t, " ", elem.FieldByName("pColumn").String())
	assert.False(t, elem.FieldByName("hdrLine").Bool())
}
