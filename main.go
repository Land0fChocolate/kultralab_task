package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"
	d "golang.org/x/image/draw"
)

const LoadImageFilePath = "images/Otterpng.png"
const SaveImageFilePath = "images/newOtterpng.png"

func main() {
	var input string
	var imageNRGBA *image.NRGBA

	for {
		fmt.Println("\nType a number for one of the following options:")
		fmt.Println("1 - Load image")
		fmt.Println("2 - Save image")
		fmt.Println("3 - Resize image")
		fmt.Println("4 - Rotate image (90 degrees anti-clockwise)")
		fmt.Println("0 - Quit")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Input error:", err, "\nPlease type in a number from 0 to 4.")
		}

		switch input {
		case "1":
			imageNRGBA = load(LoadImageFilePath)
		case "2":
			save(SaveImageFilePath, imageNRGBA)
		case "3":
			imageNRGBA = resize(imageNRGBA)
		case "4":
			imageNRGBA = rotate(imageNRGBA)
		case "0":
			return
		default:
			fmt.Println("Invalid input. Please type in a number from 0 to 4.")
		}
	}
}

func load(filePath string) *image.NRGBA {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot read file:", err)
		return nil
	}

	img, err := png.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
		return nil
	}

	b := img.Bounds()
	imageNRGBA := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(imageNRGBA, imageNRGBA.Bounds(), img, b.Min, draw.Src)

	fmt.Println("Load successful")
	return imageNRGBA
}

func save(filePath string, img *image.NRGBA) {
	if img == nil {
		fmt.Println("No image loaded")
		return
	}

	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
		return
	}
	png.Encode(imgFile, img)

	fmt.Println("Save successful")
	return
}

func resize(img *image.NRGBA) *image.NRGBA {
	if img == nil {
		fmt.Println("No image loaded")
		return nil
	}
	newImg := image.NewNRGBA(image.Rect(0, 0, img.Bounds().Max.X/2, img.Bounds().Max.Y/2))
	d.NearestNeighbor.Scale(newImg, newImg.Rect, img, img.Bounds(), draw.Over, nil)

	fmt.Println("Resize successful")
	return newImg
}

//rotating 90 degrees anti-clockwise
func rotate(img *image.NRGBA) *image.NRGBA {
	// newImg := image.NewNRGBA(image.Rect(0, 0, img.Bounds().Max.Y, img.Bounds().Max.X))

	// for y := 0; y < img.Bounds().Max.Y; y++ {
	// 	for x := 0; x < img.Bounds().Max.X; x++ {
	// 		newImg.Pix[(img.Bounds().Max.Y*x)+(img.Bounds().Max.Y-y-1)] = img.Pix[(img.Bounds().Max.X*y)+x]
	// 	}
	// }

	// return newImg

	if img == nil {
		fmt.Println("No image loaded")
		return nil
	}

	fmt.Println("Rotate successful")
	return imaging.Rotate90(img)
}
