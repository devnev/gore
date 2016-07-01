// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"github.com/spf13/cobra"
)

func runRename(dirPath string, pattern *regexp.Regexp, replacement string) error {
	items, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, info := range items {
		if info.IsDir() {
			err = runRename(path.Join(dirPath, info.Name()), pattern, replacement)
			if err != nil {
				return err
			}
		}
		newName := pattern.ReplaceAllString(info.Name(), replacement)
		if newName != info.Name() {
			err = os.Rename(path.Join(dirPath, info.Name()), path.Join(dirPath, newName))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

var renameCmd = &cobra.Command{
	Use:     "rename <pattern> <replacement> <file or directory>",
	Short:   "replace occurrences of pattern in file and directory names with replacement.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return cmd.Usage()
		}
		pattern, err := regexp.Compile(args[0])
		if err != nil {
			return err
		}
		return runRename(args[2], pattern, args[1])
	},
}
