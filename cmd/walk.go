// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkReplace(replacer func(in []byte) []byte, paths ...string) error {
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			data = replacer(data)
			return ioutil.WriteFile(path, data, info.Mode())
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func walkRead(reader func(path string, in []byte), paths ...string) error {
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			reader(path, data)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
