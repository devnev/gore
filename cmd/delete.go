// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"bytes"
	"regexp"

	"github.com/spf13/cobra"
)

func runDelete(paths []string, pattern *regexp.Regexp) error {
	return walkReplace(func(data []byte) []byte {
		var line bytes.Buffer
		var out bytes.Buffer
		for i := 0; i < len(data)+1; i++ {
			if i < len(data) && data[i] != '\n' {
				line.WriteByte(data[i])
				continue
			}
			if pattern.Match(line.Bytes()) {
				line.Reset()
				continue
			}
			out.Write(line.Bytes())
			if i < len(data) {
				out.WriteByte('\n')
			}
			line.Reset()
		}
		return out.Bytes()
	}, paths...)
}

var deleteCmd = &cobra.Command{
	Use:     "delete <pattern> <files and directories...>",
	Short:   "delete lines matching pattern.",
	Aliases: []string{"d"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return cmd.Usage()
		}
		pattern, err := regexp.Compile(args[0])
		if err != nil {
			return err
		}
		return runDelete(args[1:], pattern)
	},
}
