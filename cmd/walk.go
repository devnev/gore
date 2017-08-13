// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/spf13/pflag"
)

var fVerbose = pflag.BoolP("verbose", "v", false, "print additional information")

func walkReplace(replacer func(in []byte) []byte, paths ...string) error {
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if *fVerbose {
					log.Printf("Error stating %v: %v", path, err.Error())
				}
				return nil
			}
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				if *fVerbose {
					log.Printf("Error reading %v: %v", path, err.Error())
				}
				return nil
			}
			replaced := replacer(data)
			if reflect.DeepEqual(data, replaced) {
				return nil
			}
			err = ioutil.WriteFile(path, replaced, info.Mode())
			if err != nil {
				if *fVerbose {
					log.Printf("Error writing %v: %v", path, err.Error())
				}
			}
			return nil
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
			if err != nil {
				if *fVerbose {
					log.Printf("Error stating %v: %v", path, err.Error())
				}
				return nil
			}
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				if *fVerbose {
					log.Printf("Error reading %v: %v", path, err.Error())
				}
				return nil
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
