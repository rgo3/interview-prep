package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func copyFiles(files ...string) {
	copies := make(chan string, len(files))

	for _, f := range files {
		go copyFile(copies, f, f+"_new")
	}

	for range files {
		fmt.Println(<-copies)
	}

}

func copyFile(ch chan<- string, old string, new string) error {
	content, err := ioutil.ReadFile(old)
	if err != nil {
		return fmt.Errorf("Error")
	}

	info, err := os.Stat(old)
	if err != nil {
		return fmt.Errorf("Error")
	}
	mode := info.Mode()
	perm := mode.Perm()

	err = ioutil.WriteFile(new, content, perm)
	if err != nil {
		return fmt.Errorf("Error")
	}
	ch <- "done"
	return nil
}
