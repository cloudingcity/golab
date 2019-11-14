package utils

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

// RenderTable returns initialized table instance.
func RenderTable(w io.Writer, headers []string, rows [][]string) {
	table := tablewriter.NewWriter(w)
	table.SetColWidth(1000)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)

	table.SetHeader(headers)
	table.AppendBulk(rows)
	table.Render()
}
