package main

import (
	"embed"
	"flag"
	"fmt"
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
	srcAbs string
	out = flag.String("out", ".", "The output directory to put files in")
	outAbs string
	latex = flag.Bool("latex", false, "Whether to render LaTeX equations")
	//go:embed template.html
	noteTemplates embed.FS
)

func init() {
	flag.Parse()
	var err error
	srcAbs, err = filepath.Abs(*src)
	if err != nil {
		fmt.Printf("Error resolving absolute source path: %v", err)
		return
	}

	outAbs, err = filepath.Abs(*out)
	if err != nil {
		fmt.Printf("Error resolving absolute output path: %v", err)
		return
	}
}

func main() {
	noteTmp := template.Must(template.ParseFS(noteTemplates, "*"))
	filepath.WalkDir(srcAbs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error walking directory: %v", err)
		}

		dest := destPath(path)

		// Make directory if it doesn't exist
		if d.IsDir() {
			if err := os.MkdirAll(dest, os.ModePerm); err != nil {
				fmt.Printf("Error making directory: %v", err)
				return err
			}
			return nil
		}

		if filepath.Base(path)[0:1] == "." {
			return nil
		}

		srcFile, err := os.Open(path)
		if err != nil {
			fmt.Printf("Error opening src path: %v", err)
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Error opening output path: %v", err)
			return err
		}
		defer destFile.Close()

		// Copy any non-markdown assets as-is
		if filepath.Ext(path) != ".md" {
			if _, err = io.Copy(destFile, srcFile); err != nil {
				fmt.Printf("Error copying assets: %v", err)
				return err
			}
			return nil
		}

		// Read source file
		mdSource, err := ioutil.ReadAll(srcFile)
		if err != nil {
			fmt.Printf("Error reading source file: %v", err)
			return err
		}

		// Render markdown
		note, err := renderMD(mdSource)
		if err != nil {
			fmt.Printf("Error rendering markdown: %v", err)
			return err
		}

		if err = noteTmp.Execute(destFile, note); err != nil {
			fmt.Printf("Error executing template: %v", err)
			return err
		}

		return nil
	})
}

func destPath(path string) string {
	res := strings.Replace(path, srcAbs, outAbs, 1)
	if filepath.Ext(path) == ".md" {
		return res[0:len(res)-2]+"html"
	} else {
		return res
	}
}
