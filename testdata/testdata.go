// https://clouding.city/go/refactor-2/

package testdata

import (
	"path/filepath"
	"runtime"
)

var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

func Path(rel string) string {
	return filepath.Join(basepath, rel)
}
