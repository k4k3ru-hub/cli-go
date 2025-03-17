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
	separatorWriter := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', 0)

	// Headers
	for _, header := range headers {
		fmt.Fprintf(dataWriter, "%s\t", header)
		fmt.Fprintf(separatorWriter, "%s\t", strings.Repeat("-", len(header)))
	}
	fmt.Fprintln(dataWriter)
	fmt.Fprintln(separatorWriter)
	dataWriter.Flush()
	separatorWriter.Flush()

	// Rows
	for _, row := range rows {
		for _, col := range row {
			fmt.Fprintf(dataWriter, "%v\t", col)
		}
		fmt.Fprintln(dataWriter)
	}
	dataWriter.Flush()
}
