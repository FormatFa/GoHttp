package util

import "strconv"

const KB int = 1024
const MB int = 1024 * 1024

// TODO: use float display
func FormatSize(size int) string {
	if size < 1024 {
		return strconv.Itoa(size) + "b"
	} else if size >= 1024 && size < MB {
		return strconv.Itoa((size / 1024)) + "k"
	} else {
		return strconv.Itoa((size / 1024 / 1024)) + "m"
	}

}
