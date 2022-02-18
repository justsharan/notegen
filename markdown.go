package main

import (
	"bytes"
	"html/template"

	toc "github.com/abhinav/goldmark-toc"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM, meta.Meta),
	goldmark.WithParserOptions(parser.WithAutoHeadingID()),
)

type note struct {
	Content template.HTML
	LaTeX bool
	Metadata map[string]interface{}
	TableOfContents template.HTML
}

func renderMD(src []byte) (*note, error) {
	ctx := parser.NewContext()
	doc := md.Parser().Parse(text.NewReader(src), parser.WithContext(ctx))

	tree, err := toc.Inspect(doc, src)
	if err != nil {
		return nil, err
	}

	list := toc.RenderList(tree)

	var tocBuf bytes.Buffer
	if err = md.Renderer().Render(&tocBuf, src, list); err != nil {
		return nil, err
	}

	var contentBuf bytes.Buffer
	if err = md.Renderer().Render(&contentBuf, src, doc); err != nil {
		return nil, err
	}

	return &note{
		Content: template.HTML(contentBuf.Bytes()),
		LaTeX: *latex,
		Metadata: meta.Get(ctx),
		TableOfContents: template.HTML(tocBuf.Bytes()),
	}, nil
}
