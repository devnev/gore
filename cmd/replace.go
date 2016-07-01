// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"bytes"
	"regexp"

	"github.com/spf13/cobra"
)

func runReplace(start string, pattern *regexp.Regexp, replacement []byte) error {
	return walkReplace(func(data []byte) []byte {
		var line bytes.Buffer
		var out bytes.Buffer
		for i := 0; i < len(data)+1; i++ {
			if i < len(data) && data[i] != '\n' {
				line.WriteByte(data[i])
				continue
			}
			out.Write(pattern.ReplaceAll(line.Bytes(), replacement))
			if i < len(data) {
				out.WriteByte('\n')
			}
			line.Reset()
		}
		return out.Bytes()
	}, start)
}

var replaceCmd = &cobra.Command{
	Use:     "replace <pattern> <replacement> <file or directory>",
	Short:   "replace occurrences of pattern with replacement.",
	Aliases: []string{"re", "sub", "s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return cmd.Usage()
		}
		pattern, err := regexp.Compile(args[0])
		if err != nil {
			return err
		}
		return runReplace(args[2], pattern, []byte(args[1]))
	},
}
