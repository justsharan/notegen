package main

import (
	"flag"
	"html/template"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	src = flag.String("src", ".", "The source directory to read files from")
	out = flag.String("out", ".", "The output directory to put files in")
	latex = flag.Bool("latex", false, "Whether to render LaTeX equations")
	noteTemplate, _ = template.ParseGlob("template.html")
)

func init() {
	flag.Parse()
}

func main() {
	filepath.WalkDir(*src, func(path string, d fs.DirEntry, err error) error {
		dest := strings.Replace(path, *src, *out, 1)

		// Make directory if it doesn't exist
		if d.IsDir() {
			err := os.MkdirAll(dest, os.ModePerm)
			return err
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(dest[0:len(dest)-2]+"html")
		if err != nil {
			return err
		}
		defer destFile.Close()

		// Copy any non-markdown assets as-is
		if filepath.Ext(path) != ".md" {
			_, err = io.Copy(destFile, srcFile)
			return err
		}

		// Read source file
		mdSource, err := ioutil.ReadAll(srcFile)
		if err != nil {
			return err
		}

		// Render markdown
		note, err := renderMD(mdSource)
		if err != nil {
			return err
		}

		return noteTemplate.Execute(destFile, note)
	})
}