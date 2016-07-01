// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"fmt"
	"os"
	"regexp"
	"path/filepath"

	"github.com/spf13/cobra"
)

func runFind(starts []string, pattern *regexp.Regexp) error {
	for _, start := range starts {
		err := filepath.Walk(start, func(path string, info os.FileInfo, err error) error {
			if pattern.MatchString(filepath.Base(path)) {
				fmt.Println(path)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

var findCmd = &cobra.Command{
	Use:     "find <pattern> <files and directories...>",
	Short:   "find files and directories with names matching pattern.",
	Aliases: []string{"ls", "list"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Usage()
		}
		pattern, err := regexp.Compile(args[0])
		if err != nil {
			return err
		}
		starts := args[1:]
		if len(starts) == 0 {
			starts = []string{"."}
		}
		return runFind(starts, pattern)
	},
}
