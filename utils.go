package main

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

func strToIntSlice(s []string) []int {
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
