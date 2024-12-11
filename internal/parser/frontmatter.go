package parser

import (
    "bytes"
    "gopkg.in/yaml.v2"
)

type Frontmatter struct {
    Title       string   `yaml:"title"`
    Date        string   `yaml:"date"`
    Tags        []string `yaml:"tags"`
    Description string   `yaml:"description"`
}

func ParseFrontmatter(content []byte) (*Frontmatter, []byte, error) {
    // Split content by frontmatter delimiters (---)
    parts := bytes.Split(content, []byte("---"))
    if len(parts) < 3 {
        return nil, content, nil
    }

    var meta Frontmatter
    err := yaml.Unmarshal(parts[1], &meta)
    if err != nil {
        return nil, content, err
    }

    // Return remaining content without frontmatter
    return &meta, bytes.Join(parts[2:], []byte("---")), nil
}
