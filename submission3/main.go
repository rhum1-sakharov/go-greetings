package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	textFile := readFileArg()
	fmt.Println("Reading from text file", textFile, "=>")
	readTextFile(textFile)
}

func readFileArg() string {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("No text file provided")
	}
	if len(args) > 2 {
		log.Print("Additional args provided, don't know what to do with them")
	}

	return args[1]
}

func readTextFile(textFile string) {
	f, err := os.Open(textFile)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, f); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
