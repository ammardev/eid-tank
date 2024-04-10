package engine

import (
	"image"
	"log"
	"os"

	"github.com/faiface/pixel"
	_ "image/png"
)

var sprites = map[string]*pixel.Sprite{}

func loadPicture(pic string) pixel.Picture {
	file, err := os.Open(pic)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Failed to load the picture: ", err)
	}

	return pixel.PictureDataFromImage(img)
}

func CreateSprite(path string) *pixel.Sprite {
	if sprites[path] != nil {
		return sprites[path]
	}

	picture := loadPicture(path)
	sprites[path] = pixel.NewSprite(picture, picture.Bounds())

	return sprites[path]
}
