package png2j

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// resizeImage Resize image
func ReSizeImage(srcFile string, width uint, height uint, destFile string) error {
	file, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Decode:", err)
		return err
	}
	file.Close()

	m := resizeImage(width, height, img)
	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()
	// write new image to file
	err = jpeg.Encode(out, m, nil)
	if err != nil {
		return err
	}
	return nil
}

// resizeImage Resize image
func resizeImage(width, height uint, img image.Image) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

// readPng Read PNG file
func readPng(src string) image.Image {
	// 读取png图片
	file, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
	}
	return img
}

// Con2jpg Convert PNG to JPG
func Con2jpg(src string, dst string) (err error) {
	var bg color.Color
	bg = color.White
	var quality int = 100
	err = ConBase(src, dst, bg, quality)
	if err != nil {
		return err
	}
	return nil
}

// ConBase Transforming the base method
func ConBase(src string, dst string, bgColor color.Color, quality int) (err error) {
	// 读取png图片
	var imgRes = readPng(src)
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(out)

	imgRect := image.Rect(0, 0, imgRes.Bounds().Max.X, imgRes.Bounds().Max.Y)
	jpg := image.NewRGBA(imgRect)
	// 使用无背景色
	//draw.Draw(jpg, jpg.Bounds(), imgRes, imgRes.Bounds().Min, draw.Src)
	draw.Draw(jpg, jpg.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	// paste image OVER to newImage
	draw.Draw(jpg, jpg.Bounds(), imgRes, imgRes.Bounds().Min, draw.Over)

	var opt jpeg.Options
	opt.Quality = quality
	err = jpeg.Encode(out, jpg, &opt)
	if err != nil {
		return err
	}
	return nil
}
