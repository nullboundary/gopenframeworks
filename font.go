package gopenframeworks

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/gltext"
	"log"
	"os"
)

var fonts [16]*gltext.Font

const sampleString = "0 1 2 3 4 5 6 7 8 9 A B C D E F"

func LoadFont(filename string) {
	// Load the same font at different scale factors and directions.

	var err error
	for i := range fonts {
		fonts[i], err = loadFontFile(filename, int32(12+i))
		if err != nil {
			log.Printf("LoadFont: %v", err)
			return
		}

		//defer fonts[i].Release()
	}
}

// loadFont loads the specified font at the given scale.
func loadFontFile(filename string, fontsize int32) (*gltext.Font, error) {
	log.Printf("Loading font from: %v", filename)

	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fd.Close()
	log.Println(fontsize)

	f, _ := gltext.LoadTruetype(fd, fontsize, 32, 127, gltext.LeftToRight)
	//if err != nil {
	//	log.Printf("error:%v", err)
	//	return nil, err
	//}
	return f, nil
}

// drawString draws the same string for each loaded font.
func DrawString(str string, x float32, y float32) error {
	for i := range fonts {
		if fonts[i] == nil {
			continue
		}
		log.Println(str)
		// We need to offset each string by the height of the
		// font. To ensure they don't overlap each other.
		_, h := fonts[i].GlyphBounds()
		y := y + float32(i*h)

		// Draw a rectangular backdrop using the string's metrics.
		sw, sh := fonts[i].Metrics(sampleString)
		//gl.Color4f(0.1, 0.1, 0.1, 0.7)
		gl.Rectf(x, y, x+float32(sw), y+float32(sh))
		//Mesh(x, y, 0, x+float32(sw), y+float32(sh), 0)
		// Render the string.
		//gl.Color4f(1, 1, 1, 1)
		err := fonts[i].Printf(x, y, str)
		if err != nil {
			return err
		}
	}

	return nil
}
