package main

import (
	"log"
	mylzw "lzwgo/lzw"
	"os"
)

func main() {
	var lzwEncoder mylzw.Encoder
	var lzwDecoder mylzw.Decoder
	lzwEncoder.Init("tlp.txt", "out.lzw")

	if err := lzwEncoder.ReadFile(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	lzwEncoder.Circle()

	if err := lzwEncoder.WriteFile(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	lzwDecoder.Init("out.lzw", "reveal.txt")
	if err := lzwDecoder.ReadFile(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	lzwDecoder.Circle()
	if err := lzwDecoder.WriteFile(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
