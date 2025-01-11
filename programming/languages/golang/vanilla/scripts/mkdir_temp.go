package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// go playground: https://go.dev/play/p/BbdPww_wKP7
func main() {
	const rel = "tmp"
	var paths []string

	path, err := os.MkdirTemp("", "")
	fmt.Println(`os.MkdirTemp("", ""):`, path, err)
	paths = append(paths, path)

	println()

	path, err = os.MkdirTemp("/tmp", "")
	fmt.Println(`os.MkdirTemp("/tmp", ""):`, path, err)
	paths = append(paths, path)

	path, err = os.MkdirTemp("/tmp/", "")
	fmt.Println(`os.MkdirTemp("/tmp/", ""):`, path, err)
	paths = append(paths, path)

	println()

	if _, err := os.Stat(rel); os.IsNotExist(err) {
		err := os.MkdirAll(rel, os.ModePerm)
		fmt.Println(`os.MkdirAll(rel, os.ModePerm):`, err)
	}

	path, err = os.MkdirTemp("tmp", "")
	fmt.Println(`os.MkdirTemp("tmp", ""):`, path, err)
	abs, err := filepath.Abs(path)
	fmt.Println("filepath.Abs(path):", abs, err)
	paths = append(paths, path)

	path, err = os.MkdirTemp("tmp/", "")
	fmt.Println(`os.MkdirTemp("tmp/", ""):`, path, err)
	abs, err = filepath.Abs(path)
	fmt.Println("filepath.Abs(path):", abs, err)
	paths = append(paths, path)

	path, err = os.MkdirTemp("./tmp", "")
	fmt.Println(`os.MkdirTemp("./tmp", ""):`, path, err)
	abs, err = filepath.Abs(path)
	fmt.Println("filepath.Abs(path):", abs, err)
	paths = append(paths, path)

	path, err = os.MkdirTemp("./tmp/", "")
	fmt.Println(`os.MkdirTemp("./tmp/", ""):`, path, err)
	abs, err = filepath.Abs(path)
	fmt.Println("filepath.Abs(path):", abs, err)
	paths = append(paths, path)

	println()

	path, err = os.MkdirTemp("", "xpto")
	fmt.Println(`os.MkdirTemp("", "xpto"):`, path, err)
	paths = append(paths, path)

	path, err = os.MkdirTemp("", "xpto_")
	fmt.Println(`os.MkdirTemp("", "xpto_"):`, path, err)
	paths = append(paths, path)

	println()

	path, err = os.MkdirTemp("", "*foobar")
	fmt.Println(`os.MkdirTemp("", "*foobar"):`, path, err)
	paths = append(paths, path)

	path, err = os.MkdirTemp("", "foo*bar")
	fmt.Println(`os.MkdirTemp("", "foo*bar"):`, path, err)
	paths = append(paths, path)

	path, err = os.MkdirTemp("", "foobar*")
	fmt.Println(`os.MkdirTemp("", "foobar*"):`, path, err)
	paths = append(paths, path)

	path, err = os.MkdirTemp("", "*foo*bar*")
	fmt.Println(`os.MkdirTemp("", "*foo*bar*"):`, path, err)
	paths = append(paths, path)

	for _, path := range paths {
		if err := os.RemoveAll(path); err != nil {
			fmt.Println("os.RemoveAll(path):", err)
		}
	}
}
