// Print the names and contents of the files in a tar or zip archive.
// Only makes sense for archives containing only text files.
//
// The archive/tar API doesn't presume random access, and I can't figure out
// how I could present a facade for tars and zips that gives random access to
// individual files in the archive without using syscall.Dup or opening a bunch
// of file descriptors that could possibly be to different files if the
// filesystem changed. I started out with a program that just detected and
// printed archive types, which seemed to satisfy most of the spirit of the
// exercise, but I decided to go a bit further and print file contents.
// Wrapping an io.Reader like in zip/zip.go and tar/tar.go is awkward and not
// that useful, since it basically assumes none of the files are binary.
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	arprint "github.com/torbiak/gopl/ex10.2"
	_ "github.com/torbiak/gopl/ex10.2/tar"
	_ "github.com/torbiak/gopl/ex10.2/zip"
)

func printArchive(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := arprint.Open(f)
	if err != nil {
		return fmt.Errorf("open archive reader: %s", err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		return fmt.Errorf("printing: %s", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: arprint FILE ...")
	}
	exitCode := 0
	for _, filename := range os.Args[1:] {
		err := printArchive(filename)
		if err != nil {
			log.Print(err)
			exitCode = 2
		}
	}
	os.Exit(exitCode)
}
