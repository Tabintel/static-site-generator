package generator

import (
    "encoding/xml"
    "os"
    "path/filepath"
    //"time"
)

type RSS struct {
    XMLName xml.Name `xml:"rss"`
    Version string   `xml:"version,attr"`
    Channel Channel  `xml:"channel"`
}

type Channel struct {
    Title       string    `xml:"title"`
    Link        string    `xml:"link"`
    Description string    `xml:"description"`
    Items       []Item    `xml:"item"`
}

type Item struct {
    Title       string    `xml:"title"`
    Link        string    `xml:"link"`
    Description string    `xml:"description"`
    PubDate     string    `xml:"pubDate"`
}

func (g *Generator) GenerateRSS(posts []Post, siteURL string) error {
    items := make([]Item, len(posts))
    for i, post := range posts {
        items[i] = Item{
            Title:       post.Title,
            Link:        siteURL + "/" + post.URL,
            Description: post.Description,
            PubDate:     post.FormattedDate,
        }
    }

    rss := RSS{
        Version: "2.0",
        Channel: Channel{
            Title:       "My Static Site",
            Link:        siteURL,
            Description: "Latest blog posts",
            Items:       items,
        },
    }

    output, err := xml.MarshalIndent(rss, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(filepath.Join(g.OutputDir, "feed.xml"), output, 0644)
}
