/*Выводит дерево каталогов и файлов (если указана опция -f).
Необходимо реализовать функцию `dirTree` внутри `main.go`
)*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func dirFileTree(out io.Writer, path string, prefix string) error {
	symbol := "├───"
	tabSymbol := "│\t"
	lastSymbol := "└───"
	tab := "\t"
	//ReadDir reads the directory named by dirname and returns a list of directory entries sorted by filename
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	count := len(files)

	for i, file := range files {

		if !(strings.HasPrefix(file.Name(), ".")) {
			if file.IsDir() {

				newPrefix := ""
				if count-1 == i {
					newPrefix = prefix + tab
					fmt.Fprintf(out, "%s%s\n", prefix+lastSymbol, file.Name())
				} else {
					newPrefix = prefix + tabSymbol
					fmt.Fprintf(out, "%s%s\n", prefix+symbol, file.Name())
				}

				newDir := filepath.Join(path, file.Name())

				dirFileTree(out, newDir, newPrefix)

			} else {
				//if printFiles {

				if file.Size() > 0 {
					if count-1 == i {
						fmt.Fprintf(out, "%s%s (%vb)\n", prefix+lastSymbol, file.Name(), file.Size())
					} else {
						fmt.Fprintf(out, "%s%s (%vb)\n", prefix+symbol, file.Name(), file.Size())
					}

				} else {
					if count-1 == i {
						fmt.Fprintf(out, "%s%s (empty)\n", prefix+lastSymbol, file.Name())
					} else {
						fmt.Fprintf(out, "%s%s (empty)\n", prefix+symbol, file.Name())
					}
				}
				//}
			}
		}

	}
	return nil
}
func dirTree(out io.Writer, path string, printFiles bool) error {

	prefix := ""
	if printFiles {
		dirFileTree(out, path, prefix)
	} else {
		subDirTree(out, path, prefix)
	}
	return nil
}

func subDirTree(out io.Writer, path string, prefix string) error {
	symbol := "├───"
	tabSymbol := "│\t"
	lastSymbol := "└───"
	tab := "\t"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	countDir := 0
	for _, fi := range files {

		if fi.IsDir() {
			countDir++
		}
	}
	i := 0
	for _, file := range files {
		if file.IsDir() {
			newPrefix := ""
			if countDir-1 == i {
				newPrefix = prefix + tab
				fmt.Fprintf(out, "%s%s\n", prefix+lastSymbol, file.Name())
			} else {
				newPrefix = prefix + tabSymbol
				fmt.Fprintf(out, "%s%s\n", prefix+symbol, file.Name())
			}

			newDir := filepath.Join(path, file.Name())

			subDirTree(out, newDir, newPrefix)
			i++
		}
	}
	return nil

}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
