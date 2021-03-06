package utils

import (
	"os"
	"strconv"
)

// TODO: test this on real network
func FormatMilliSat(milliSat int64) string {
	return strconv.FormatFloat(float64(milliSat)/1000, 'f', 3, 64)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
