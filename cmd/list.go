/*
Copyright © 2025 Satnam Sandhu <satnamsandhu70002@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/satnamSandhu2001/stackjet/database"
	"github.com/satnamSandhu2001/stackjet/internal/dto"
	"github.com/satnamSandhu2001/stackjet/internal/models"
	"github.com/satnamSandhu2001/stackjet/internal/services"
	"github.com/satnamSandhu2001/stackjet/pkg/colors"
	"github.com/spf13/cobra"
)

var (
	stackTypeFilter string
	portFilter      int

	printJsonFormat bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List applications with optional filters",
	Long: `List all applications managed by Stackjet with optional filtering capabilities.

	You can filter applications by:
	- Technology stack using --tech or -t flag
	- Port number using --port or -p flag

	Example:
	stackjet list
	stackjet list --tech nodejs
	stackjet list --port 3000
	stackjet list -t node -p 8081`,
	Run: func(cmd *cobra.Command, args []string) {

		dbConn := database.Connect()
		defer dbConn.Close()
		stackService := services.NewStackService(dbConn)

		var logBuf strings.Builder
		multiWriter := io.MultiWriter(os.Stdout, &logBuf)

		// service logic
		stacks, err := stackService.GetStackList(context.Background(), &dto.Stack_List_Request{
			Type: stackTypeFilter,
			Port: portFilter,
		})
		if err != nil {
			multiWriter.Write([]byte("__ERROR__: " + err.Error()))
			fmt.Printf(colors.Red("Failed to list stacks: "), err)
			return
		}
		if printJsonFormat {
			jsonOutput, err := json.Marshal(stacks)
			if err != nil {
				multiWriter.Write([]byte("__ERROR__: " + err.Error()))
				fmt.Printf(colors.Red("Failed to marshal stacks: "), err)
				return
			}
			multiWriter.Write(jsonOutput)
			return
		} else {

			printStacksTableList(stacks)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&stackTypeFilter, "tech", "t", "", "Filter by stack technology type")
	listCmd.Flags().IntVarP(&portFilter, "port", "p", 0, "Filter by application port")

	listCmd.Flags().BoolVar(&printJsonFormat, "json-output", false, "Print output in JSON format")
}

func printStacksTableList(stacks []models.Stack) {
	if len(stacks) == 0 {
		fmt.Println(colors.SecondaryBoldItalicBG(" No stacks found. "))
		return
	}

	headers := []string{
		"id", "name", "type", "port", "branch", "build", "start", "post start",
	}

	var rows [][]string
	for _, s := range stacks {
		rows = append(rows, []string{
			colors.Bold(fmt.Sprintf("%d", s.ID)),
			s.Name,
			s.Type,
			fmt.Sprintf("%d", s.Port),
			s.Branch,
			s.Commands.Build,
			s.Commands.Start,
			s.Commands.Post,
		})
	}

	coloredHeaders := make([]string, len(headers))
	for i, h := range headers {
		coloredHeaders[i] = colors.PrimaryBold(h)
	}

	allRowsPlain := append([][]string{headers}, rows...)
	allRowsColored := append([][]string{coloredHeaders}, rows...)

	widths := make([]int, len(headers))
	for _, row := range allRowsPlain {
		for i, col := range row {
			l := len(stripANSI(col))
			if l > widths[i] {
				widths[i] = l
			}
		}
	}

	printLine := func(left, mid, right, fill string) {
		fmt.Print(left)
		for i, w := range widths {
			fmt.Print(strings.Repeat(fill, w+2))
			if i < len(widths)-1 {
				fmt.Print(mid)
			}
		}
		fmt.Println(right)
	}

	printRow := func(row []string) {
		fmt.Print("│")
		for i, col := range row {
			padding := widths[i] - len(stripANSI(col))
			fmt.Printf(" %-*s │", widths[i], col+strings.Repeat(" ", padding))
		}
		fmt.Println()
	}

	printLine("┌", "┬", "┐", "─")
	for i, row := range allRowsColored {
		printRow(row)
		if i == 0 {
			printLine("├", "┼", "┤", "─")
		}
	}
	printLine("└", "┴", "┘", "─")
}
func stripANSI(s string) string {
	return regexp.MustCompile(`\x1b\[[0-9;]*m`).ReplaceAllString(s, "")
}
