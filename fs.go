package main

import "io"

type FileSystem interface {
	ListDirs(relPath string) []string
	Write(relPath string, w io.Writer) error
	DetectContentType(relPath string) string
}
