package utils

import (
	"embed"
	"io/fs"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ConvertEmbedFsDirToHttpFs(e embed.FS, d string) (http.FileSystem, error) {
	fsys, err := fs.Sub(e, d)
	if err != nil {
		return nil, err
	}
	return http.FS(fsys), nil
}

// StrToIntSlice converts a slice of strings to ints
func StrToIntSlice(s []string) []int {
	intSlice := []int{}
	for _, s := range s {
		intValue, err := strconv.Atoi(s)
		if err != nil {
			log.Errorf("Invalid light value: %s", s)
		}
		intSlice = append(intSlice, intValue)
	}
	return intSlice
}
