// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"log"
	"os"
	// "io"
	// "os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}

}

func main() {
	files := []string{"/etc/hosts", "/etc/resolv.conf", "/etc/nsswitch.conf"}

	for _, cfile := range files {
		fmt.Printf("Name: ", cfile, "\n-----------------------\n")
		readFilesLines(cfile)
	}

}

func readFilesLines(cfile string) {
	file, err := os.Open(cfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println("#####")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
