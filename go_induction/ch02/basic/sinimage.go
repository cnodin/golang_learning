package main

import (
	"image/color"
	"image"
	"math"
	"os"
	"github.com/gpmgo/gopm/modules/log"
	"image/png"
)

//图片大小
const size = 300

func main() {
	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(0, 0, color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s) * size/2

		//用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}

	//创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err.Error())
	}

	png.Encode(file, pic)
	file.Close()
}
