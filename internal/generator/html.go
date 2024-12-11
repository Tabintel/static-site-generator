package generator

import (
    "html/template"
    "os"
    "path/filepath"
)

type Generator struct {
    TemplateDir string
    OutputDir   string
}

func NewGenerator(templateDir, outputDir string) *Generator {
    return &Generator{
        TemplateDir: templateDir,
        OutputDir:   outputDir,
    }
}

func (g *Generator) Generate(data interface{}, outputFile string) error {
    // Ensure output directory exists
    if err := os.MkdirAll(g.OutputDir, 0755); err != nil {
        return err
    }

    // Parse template
    tmpl, err := template.ParseFiles(filepath.Join(g.TemplateDir, "default.html"))
    if err != nil {
        return err
    }

    // Create output file
    out, err := os.Create(filepath.Join(g.OutputDir, outputFile))
    if err != nil {
        return err
    }
    defer out.Close()

    // Execute template
    return tmpl.Execute(out, data)
}

//  Generator struct
type Post struct {
    Title       string
    Date        string
    FormattedDate string
    Description string
    URL         string
    Tags        []string
    Content     template.HTML
}

type IndexData struct {
    Posts []Post
}

func (g *Generator) GenerateIndex(posts []Post) error {
    tmpl, err := template.ParseFiles(filepath.Join(g.TemplateDir, "index.html"))
    if err != nil {
        return err
    }

    out, err := os.Create(filepath.Join(g.OutputDir, "index.html"))
    if err != nil {
        return err
    }
    defer out.Close()

    return tmpl.Execute(out, IndexData{Posts: posts})
}

type TagData struct {
    Tag   string
    Posts []Post
}

func (g *Generator) GenerateTagPages(posts []Post) error {
    // Group posts by tag
    tagMap := make(map[string][]Post)
    for _, post := range posts {
        for _, tag := range post.Tags {
            tagMap[tag] = append(tagMap[tag], post)
        }
    }

    // Generate a page for each tag
    for tag, tagPosts := range tagMap {
        tagDir := filepath.Join(g.OutputDir, "tags")
        if err := os.MkdirAll(tagDir, 0755); err != nil {
            return err
        }

        tmpl, err := template.ParseFiles(filepath.Join(g.TemplateDir, "tag.html"))
        if err != nil {
            return err
        }

        out, err := os.Create(filepath.Join(tagDir, tag+".html"))
        if err != nil {
            return err
        }

        if err := tmpl.Execute(out, TagData{
            Tag:   tag,
            Posts: tagPosts,
        }); err != nil {
            out.Close()
            return err
        }
        out.Close()
    }
    return nil
}

func (g *Generator) CopyStaticAssets(themePath string) error {
    staticDir := filepath.Join(themePath, "static")
    outputStaticDir := filepath.Join(g.OutputDir, "static")

    // Create output static directory
    if err := os.MkdirAll(outputStaticDir, 0755); err != nil {
        return err
    }

    // Copy static files
    return filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories
        if info.IsDir() {
            return nil
        }

        // Read source file
        content, err := os.ReadFile(path)
        if err != nil {
            return err
        }

        // Calculate relative path and create destination path
        relPath, err := filepath.Rel(staticDir, path)
        if err != nil {
            return err
        }
        destPath := filepath.Join(outputStaticDir, relPath)

        // Ensure destination directory exists
        if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
            return err
        }

        // Write file to destination
        return os.WriteFile(destPath, content, 0644)
    })
}
