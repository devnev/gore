// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

func runSearch(paths []string, pattern *regexp.Regexp) error {
	return walkRead(func(path string, data []byte) {
		var line bytes.Buffer
		for i := 0; i < len(data)+1; i++ {
			if i < len(data) && data[i] != '\n' {
				line.WriteByte(data[i])
				continue
			}
			if pattern.Match(line.Bytes()) {
				fmt.Println(path)
				return
			}
			line.Reset()
		}
	}, paths...)
}

var searchCmd = &cobra.Command{
	Use:   "search <pattern> <files and directories...>",
	Short: "list files with lines matching the pattern.",
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
		return runSearch(starts, pattern)
	},
}
