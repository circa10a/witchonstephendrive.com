package utils

import (
	"io/fs"
	"net/http"
)

// ConvertEmbedFsDirToHTTPSFS returns sub directory of fs
func ConvertEmbedFsDirToHTTPFS(e fs.FS, d string) (http.FileSystem, error) {
	fsys, err := fs.Sub(e, d)
	if err != nil {
		return nil, err
	}
	return http.FS(fsys), nil
}

// StrInSlice returns true if string is in slice
func StrInSlice(str string, list []string) bool {
	for _, item := range list {
		if str == item {
			return true
		}
	}
	return false
}
