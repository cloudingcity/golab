package utils

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

// NewTable returns initialized table instance.
func NewTable(w io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(w)
	table.SetColWidth(1000)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetHeaderLine(false)

	return table
}
