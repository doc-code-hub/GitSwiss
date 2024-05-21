package utilities

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// FindFolders finds all the folders in the given directory
func FindFolders(ctx context.Context, currDir string) (*[]fs.DirEntry, error) {
	entries, err := os.ReadDir(currDir)
	if err != nil {
		fmt.Printf("Failed to read directory: %s\n", err.Error())
		return nil, err
	}

	folders := []fs.DirEntry{}
	curr, err := dirEntryFromPath(currDir)
	if err != nil {
		fmt.Printf("Failed to get directory info: %s\n", err.Error())
		return nil, err
	}
	folders = append(folders, curr)
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry)
		}
	}

	return &folders, nil
}

type dirEntry struct {
	info os.FileInfo
	name string
}

func (d *dirEntry) Name() string {
	return d.name
}

func (d *dirEntry) IsDir() bool {
	return d.info.IsDir()
}

func (d *dirEntry) Type() fs.FileMode {
	return d.info.Mode().Type()
}

func (d *dirEntry) Info() (fs.FileInfo, error) {
	return d.info, nil
}

// Funzione per convertire os.FileInfo in fs.DirEntry.
func fileInfoToDirEntry(info os.FileInfo, name string) fs.DirEntry {
	return &dirEntry{
		info: info,
		name: name,
	}
}

// Funzione per creare un fs.DirEntry dato un percorso.
func dirEntryFromPath(path string) (fs.DirEntry, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return fileInfoToDirEntry(info, filepath.Base(path)), nil
}
