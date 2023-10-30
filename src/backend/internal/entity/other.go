package entity

import "strings"

type LocalFile struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewLocalFile(path *string) *LocalFile {
	if path == nil || *path == "" {
		return nil
	}
	paths := strings.Split(*path, "/")
	return &LocalFile{paths[len(paths)-1], *path}
}
