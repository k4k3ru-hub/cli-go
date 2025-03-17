//
// print.go
//
package cli

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)


//
// Print in table format.
//
func PrintTable(headers []string, rows [][]interface{}) {
	const padding = 2
	dataWriter := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)

	// Calculate maximum column widths.
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)+padding
	}
	for _, row := range rows {
		for i, col := range row {
			colStr := fmt.Sprintf("%v", col)
			if len(colStr)+padding > colWidths[i] {
				colWidths[i] = len(colStr)+padding
			}
		}
	}
	totalWidth := 0
	for _, width := range colWidths {
		totalWidth += width
	}

	// Headers
	for _, header := range headers {
		fmt.Fprintf(dataWriter, "%s\t", header)
	}
	fmt.Fprintln(dataWriter)
	fmt.Fprintf(dataWriter, "%s\n", strings.Repeat("-", totalWidth))

	// Rows
	for _, row := range rows {
		for _, col := range row {
			fmt.Fprintf(dataWriter, "%v\t", col)
		}
		fmt.Fprintln(dataWriter)
	}
	dataWriter.Flush()
}
