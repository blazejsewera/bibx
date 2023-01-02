package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/blazejsewera/bibx"
)

var bibs []string

func main() {
	var dir string
	var err error

	dir, err = parseDirArg()
	if err != nil {
		printErr(err)
		return
	}

	err = filepath.WalkDir(dir, extractFromMd)
	if err != nil {
		printErr(err)
		return
	}

	merged := bibx.Merge(bibs)
	fmt.Fprint(os.Stdout, merged)
}

func parseDirArg() (string, error) {
	var dir string
	var err error
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no directory specified, using current working dir")
		dir, err = os.Getwd()
		if err != nil {
			return "", err
		}
	} else {
		dir = os.Args[1]
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			return "", err
		}
	}
	return dir, nil
}

func printErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}

func extractFromMd(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	ext := filepath.Ext(path)
	if ext != ".md" {
		return nil
	}

	err = extractBib(path)
	if err != err {
		return err
	}

	return nil
}

func extractBib(path string) error {
	file, err := os.Open(path)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}

	bibs = append(bibs, bibx.Extract(file)...)

	return nil
}
