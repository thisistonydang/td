package td

import (
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not check for existence of file '%s': %v\n", filename, err)
		return false
	}
	return true
}
