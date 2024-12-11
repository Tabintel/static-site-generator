package watcher

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "log"
    "os"
    "path/filepath"
)

type ProcessFn func() error

func Watch(dir string, process ProcessFn) error {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }
    defer watcher.Close()

    done := make(chan bool)
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Op&fsnotify.Write == fsnotify.Write {
                    fmt.Printf("Modified file: %s\n", event.Name)
                    if err := process(); err != nil {
                        log.Printf("Error processing: %v\n", err)
                    }
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Printf("Error: %v\n", err)
            }
        }
    }()

    err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return watcher.Add(path)
        }
        return nil
    })
    if err != nil {
        return err
    }

    <-done
    return nil
}
