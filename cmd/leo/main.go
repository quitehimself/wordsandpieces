package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/quitehimself/wordsandpieces"
)

var wordFlag bool
var pieceFlag bool
var wordsFlag int
var piecesFlag int

func init() {
	flag.BoolVar(&wordFlag, "word", false, "print one random word from Leo")
	flag.BoolVar(&pieceFlag, "piece", false, "print one random sentence from Leo")
	flag.IntVar(&wordsFlag, "words", 0, "print `x` random words from Leo")
	flag.IntVar(&piecesFlag, "pieces", 0, "print `x` random sentences from Leo")
}

func main() {
	flag.Parse()

	numFlags := 0
	for _, b := range []bool{wordFlag, pieceFlag, wordsFlag != 0, piecesFlag != 0} {
		if b {
			numFlags++
		}
	}
	if numFlags != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-word|-piece|-words=x|-pieces=x]\n  (only one flag may be used)\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	var seed int64
	binary.Read(crand.Reader, binary.LittleEndian, &seed)
	rand.Seed(seed)

	if wordFlag {
		wordsFlag = 1
	} else if pieceFlag {
		piecesFlag = 1
	}

	if wordsFlag > 0 {
		for i := 0; i < wordsFlag; i++ {
			fmt.Println(wordsandpieces.Words[rand.Intn(len(wordsandpieces.Words))])
		}
	} else {
		for i := 0; i < piecesFlag; i++ {
			fmt.Println(wordsandpieces.Pieces[rand.Intn(len(wordsandpieces.Pieces))])
		}
	}
}
