package main

import (
	"golang-design-pattern/adapter-peg/adapter"
	"log"
)

func main() {

	var hole = adapter.NewRoundHole(5)
	var rpeg = adapter.NewRoundPeg(5)
	log.Println("rpeg fits in hole:", hole.Fits(rpeg))

	var smallSqpeg = adapter.NewSquarePeg(5)
	var largeSgpeg = adapter.NewSquarePeg(10)
	//hole.Fits(smallSqpeg)  : compile error

	var smallSqpegAdapter = adapter.NewSquarePegAdapter(smallSqpeg)
	var largeSqpegAdapter = adapter.NewSquarePegAdapter(largeSgpeg)
	log.Println("smallSqpeg fits in hole :", hole.Fits(smallSqpegAdapter))
	log.Println("largeSgpeg fits in hole :", hole.Fits(largeSqpegAdapter))

}
