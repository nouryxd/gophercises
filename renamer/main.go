package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// 	fileName := "birthday_001.txt"
	// 	// => Birthday - 1 of 4.txt
	// 	newName, err := match(fileName, 4)
	// 	if err != nil {
	// 		fmt.Println("no match")
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println(newName)
	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	count := 0
	type rename struct {
		filename string
		path     string
	}
	var toRename []struct {
		filename string
		path     string
	}
	for _, file := range files {
		if file.IsDir() {
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, rename{
					filename: file.Name(),
					path:     fmt.Sprintf("%s/%s", dir, file.Name()),
				})
			}
		}
	}
	for _, orig := range toRename {
		newFilename, err := match(orig.filename, count)
		if err != nil {
			panic(err)
		}
		newPath := fmt.Sprintf("%s/%s", dir, newFilename)
		fmt.Println("mv %s => %s", orig.path, newPath)
	}
	// origPath := fmt.Sprintf("%s/%s")
	// newPath := fmt.Sprintf("%s/%s")
}

func match(filename string, total int) (string, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our patter", filename)
	}
	// Birthday - 1 .txt
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil

	return "", nil
}
