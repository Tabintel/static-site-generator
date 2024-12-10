package parser

import (
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
)

type MarkdownContent struct {
    Content    string
    Title      string
    Date       string
    Tags       []string
    HTMLOutput string
}

func ParseMarkdown(content []byte) *MarkdownContent {
    extensions := parser.CommonExtensions | parser.AutoHeadingIDs
    parser := parser.NewWithExtensions(extensions)
    html := markdown.ToHTML(content, parser, nil)
    
    return &MarkdownContent{
        Content:    string(content),
        HTMLOutput: string(html),
    }
}
