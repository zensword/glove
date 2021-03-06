package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dbriemann/glove"
)

func main() {
	corpus := flag.String("corpus", "", "The path to the corpus text file.")
	output := flag.String("output", "vocab.txt", "The file to which the vocabulary shall be written.")
	minCount := flag.Uint("min-count", 5, "A threshold that defines the minimum times a word must occur to be kept in the vocabulary.")

	flag.Parse()

	if *corpus == "" {
		fmt.Println("You must specifiy a corpus file.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	wf, err := glove.LoadWordFrequenciesFromFile(*corpus)
	if err != nil {
		panic(err)
	}

	vocab := wf.ToVocabulary()
	if err = vocab.Write(*output, uint32(*minCount)); err != nil {
		panic(err)
	}
}
