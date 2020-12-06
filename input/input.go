package input

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

const callSkip = 2

func path() string {
	_, main, _, _ := runtime.Caller(callSkip)

	return filepath.Join(filepath.Dir(main), "input")
}

func Open() (*os.File, error) {
	return os.Open(path())
}

func Read() ([]byte, error) {
	return ioutil.ReadFile(path())
}
