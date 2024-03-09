package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type colour struct {
	R, G, B, A int
}

func main() {
	file, err := os.Open("x.png")
	if err != nil {
		log.Fatal("error open file: ", err)
	}
	defer file.Close()

	images, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("error decoding image: ", err)
	}

	pix := make(map[colour]int64)
	x, y := 300, 300
	pixel := images.At(x, y)
	r, g, b, a := pixel.RGBA()

	var c colour
	if a > 0 {
		c = colour{R: int(r >> 8), G: int(g >> 8), B: int(b >> 8)}
		pix[c]++
	}

	fmt.Printf("RGBA values at (%d, %d): R:%d, G:%d, B:%d \n", x, y, c.R, c.G, c.B)
}
