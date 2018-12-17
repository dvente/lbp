package lbp

import (
	"reflect"
	"testing"
)

func TestLocalBinaryPatterns(t *testing.T) {
	var test_img = [][]uint8{
		{255,   6, 255,   0,  141,   0},
		{ 48, 250, 204, 166,  223,  63},
		{  8,   0, 159,  50,  255,  30},
		{167, 255,  63,  40,  128, 255},
		{  0, 255,  30,  34,  255,  24},
		{146, 241, 255,   0,  189, 126}}

	var ref = [][]int{
		{  0, 251,   0, 255,  96, 255},
		{143,   0,  20, 153,  64,  56},
		{238, 255,  12, 191,   0, 252},
		{129,  64,  62, 159, 199,   0},
		{255,   4, 255, 175,   0, 254},
		{  3,   5,   0, 255,   4,  24}}
	lbps := LocalBinaryPatterns(test_img,8,1)
	if !reflect.DeepEqual(lbps,ref){
			t.Errorf("Expected %v, got %v",ref,lbps )

	}

}
