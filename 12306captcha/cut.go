/*
	计算公式 292 * 190
	1个292*30  6个68*68
	(292-(68*4))/5 = 4
	headerH = 30
	headerfill = 10
	(190-30-2*68-2*10) = 4

	gap = 4
*/

package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var pics []image.Image

const (
	headerH    = 30
	headerfill = 10
	cellSize   = 68
	gap        = 4
)

// func main() {
// 	fmt.Println("cat img")
// 	f, _ := os.Open("12306.jpg")
// 	srcimg, ext, _ := image.Decode(f)
// 	fmt.Println(ext)
// 	srcpoint := srcimg.Bounds().Size()
// 	fmt.Println(srcpoint)

// 	var zerox, zeroy = 0, 0

// 	header := image.NewRGBA(image.Rect(0, 0, srcpoint.X, headerH))

// 	draw.Draw(header, srcimg.Bounds().Add(image.Pt(zerox, zeroy)), srcimg, srcimg.Bounds().Min, draw.Src)

// 	for x := gap; x < srcimg.Bounds().Dx()-cellSize; x += cellSize + gap {
// 		for y := headerH + headerfill; y < srcimg.Bounds().Dy()-cellSize; y += cellSize + gap {
// 			fmt.Println(x, y)
// 			new_img, _ := Subimg(x, y, srcimg)
// 			pics = append(pics, new_img)
// 		}
// 	}

// 	h, _ := os.Create("header.jpg")
// 	jpeg.Encode(h, header, nil)

// 	for i, m := range pics {
// 		name := strconv.Itoa(i+1) + "_" + f.Name()
// 		f, _ := os.Create(name)
// 		jpeg.Encode(f, m, nil)
// 	}

// }

func cut12306(srcimg image.Image) {
	srcpoint := srcimg.Bounds().Size()

	var zerox, zeroy = 0, 0

	header := image.NewRGBA(image.Rect(0, 0, srcpoint.X, headerH))

	draw.Draw(header, srcimg.Bounds().Add(image.Pt(zerox, zeroy)), srcimg, srcimg.Bounds().Min, draw.Src)
	pics = append(pics, header)

	for x := gap; x < srcimg.Bounds().Dx()-cellSize; x += cellSize + gap {
		for y := headerH + headerfill; y < srcimg.Bounds().Dy()-cellSize; y += cellSize + gap {
			// fmt.Println(x, y)
			new_img, _ := Subimg(x, y, srcimg)
			pics = append(pics, new_img)
		}
	}

}

//OpenImageFromPath 从文件读图片
func OpenImageFromPath(name string) (image.Image, string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, "", err
	}
	return image.Decode(f)
}

// ReadImageFromString 读取图片
func ReadImageFromString(str string) (image.Image, string, error) {
	unbase64Str, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, "", err
	}
	b := bytes.NewBuffer(unbase64Str)
	return image.Decode(b)
}

//Subimg 剪切图片
func Subimg(x, y int, src image.Image) (image.Image, error) {
	var subImg image.Image
	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+cellSize, y+cellSize)).(*image.YCbCr)
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+cellSize, y+cellSize)).(*image.RGBA)
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+cellSize, y+cellSize)).(*image.NRGBA)
	} else {
		return nil, errors.New("图片解码失败")
	}
	return subImg, nil
}

// SaveImage 保存图片 path路径
func SaveImage(path string, src image.Image) error {
	f, err := os.OpenFile(path, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	ext := filepath.Ext(path)

	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})
	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(f, src)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	}
	return err
}

// func test() {
// 	const (
// 		barheight  = 30
// 		celllength = 75
// 	)
// 	fmt.Println("cat img")
// 	f, _ := os.Open("12306.jpeg")

// 	srcimg, _ := jpeg.Decode(f)
// 	srcpoint := srcimg.Bounds().Size()
// 	fmt.Println(srcpoint)

// 	img2 := image.NewRGBA(image.Rect(0, 0, 75, 75))

// 	for x := 0; x < srcimg.Bounds().Dx(); x++ {
// 		for y := barheight; y < srcimg.Bounds().Dy(); y++ {
// 			if y == 75+barheight {
// 				break
// 			}
// 			img2.Set(x, y-barheight, srcimg.At(x, y))
// 		}
// 		if x == 75 {
// 			break
// 		}
// 	}
// 	f2, _ := os.Create("cut_" + f.Name())
// 	jpeg.Encode(f2, img2, nil)
// }
