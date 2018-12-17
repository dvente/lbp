package lbp

import (
	"image"
	"image/color"
	"math"
)

type ArrayPoint struct {
	I float64
	J float64
}

func Luminance(c color.Color) (uint8) {
	//https://www.w3.org/TR/AERT/#color-contrast
	red, green, blue, _ := c.RGBA()
	return uint8(((int(red) * 299) + (int(green) * 587) + (int(blue) * 114)) / 1000)

}

func lumAtPoint(lumArr [][]uint8,i float64, j float64) (float64){
		//use nearest neighbour interpolation. can get more fancy later
		// floor to avoid problems at the boundaries
		return float64(lumArr[int(math.Floor(i))][int(math.Floor(j))])
}

func getLBPCoordinates(centreI int, centreJ int, p int, r float64) (coordArr []ArrayPoint) {
	var ans = make([]ArrayPoint, p)
	var theta float64 = 2*math.Pi / float64(p)
	for i := 0; i < p; i++ {
		var pointI float64 = -r*math.Sin(theta*float64(i)) + float64(centreI)
		var pointJ float64= r*math.Cos(theta*float64(i)) + float64(centreJ)
		ans[i] = ArrayPoint{pointI, pointJ}

	}

	return ans
}

func lbpExp(centreLum uint8, pointLum uint8) (int) {

	if centreLum >= pointLum {
		return 1
	} else {
		return 0
	}

}

func oobLbp(i int, j int) (int) {
	//here we could decide to do fancy stuff like wrapping if we want later
	return 0
}

func getLBPNumber(lumArr [][]uint8, centreI int, centreJ int, p int, r float64) (int) {

	var lbp_numb int
	height := len(lumArr)
	width := len(lumArr[0])
	centreLum := lumArr[centreI][centreJ]

	for index, point := range getLBPCoordinates(centreI,centreJ, p, r) {
		if point.I >= 0 && point.J >=0 && int(point.I) <= height && int(point.J) <= width{
			lbp_numb += oobLbp(centreI,centreJ)
		} else {
			lbp_numb += int(math.Exp2(float64(index))) * lbpExp(centreLum, uint8(lumAtPoint(lumArr,point.I,point.J)))
		}
	}

	return lbp_numb
}

func LumArrToLbpArr(lumArr [][]uint8, p int, r float64) ([][]int) {
	var lbpArr = make([][]int, len(lumArr))
	for i := range lbpArr {
		lbpArr[i] = make([]int, len(lumArr[i]))
		for j := range lbpArr[i] {
			lbpArr[i][j] = getLBPNumber(lumArr, i, j, p, r)
		}
	}

	return lbpArr

}

func imgToLumArr(img image.RGBA) [][]uint8 {
	width, height := img.Bounds().Size().X, img.Bounds().Size().Y

	var arr = make([][]uint8, height)
	for i := range arr {
		arr[i] = make([]uint8, width)
		for j := range arr[i] {
			arr[i][j] = Luminance(img.At(j, i))
		}
	}

	return arr
}

func LocalBinaryPatterns(gray [][]uint8, p int, r float64)([][]int){
	return LumArrToLbpArr(gray,p,r)
}



