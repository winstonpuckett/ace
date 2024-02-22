package main

import (
	"os"
	"path/filepath"
)

func Open(path string) {
	stats, err := os.Stat(path)

	if err != nil {
		panic(err)
	}

	if stats.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}

		for _, entry := range entries {
			Open(path + string('/') + entry.Name())
		}
	} else if filepath.Ext(path) == ".ace" {
		data, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		scanner = &StringScanner{
			source:   data,
			position: 0,
		}

		ParseAndExecute()
	}
}
