package main

import (
	"github.com/GregoryKogan/mephi-av-processing/pkg/imgproc"
)

func main() {
	img, _ := imgproc.OpenPNG("assets/salt-pepper-noise.png")

	halfTone := imgproc.GetLightness(img)
	imgproc.SavePNG(halfTone, "output/lab3/original.png")

	// 1. отфильтрованное монохромное (полутоновое) изображение;
	filtered := imgproc.MedianFilter3x3(halfTone,
		[3][3]int{
			{1, 0, 1},
			{0, 1, 0},
			{1, 0, 1},
		}, 3)

	imgproc.SavePNG(filtered, "output/lab3/filtered.png")

	// 2. разностное изображение (монохромный xor или модуль разности для полутона).
	diff := imgproc.GetDifference(halfTone, filtered)
	imgproc.SavePNG(diff, "output/lab3/difference.png")
}
