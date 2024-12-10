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
