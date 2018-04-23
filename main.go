package main

import (
	"flag"
	"log"
)

func main() {
	var w int
	flag.IntVar(&w, "w", 500, "Largura da foto")
	flag.Parse()
	err := Dir("./in", "./out", w)
	if err != nil {
		log.Fatal(err)
	}
}
