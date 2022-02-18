package main

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM, meta.Meta),
	goldmark.WithParserOptions(parser.WithAutoHeadingID()),
)

type note struct {
	Content template.HTML
	Metadata map[string]interface{}
}

func renderMD(src []byte) (*note, error) {
	ctx := parser.NewContext()

	var buf bytes.Buffer
	if err := md.Convert(src, &buf); err != nil {
		return nil, err
	}

	return &note{
		Content: template.HTML(buf.Bytes()),
		Metadata: meta.Get(ctx),
	}, nil
}
