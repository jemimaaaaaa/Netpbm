package Netpbm

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PPM struct {
	data          [][]Pixel
	width, height int
	magicNumber   string
	max           uint8
}

type Pixel struct {
	R, G, B uint8
}

func ReadPPM(filename string) (*PPM, error) {
	var err error
	var magicNumber string = ""
	var width int
	var height int
	var maxval int
	var counter int
	var headersize int
	file, err := os.ReadFile("./testImages/ppm/testP3.ppm")
	if err != nil {
	}
	splitfile := strings.SplitN(string(file), "\r\n", -1)
	for i, _ := range splitfile {
		if strings.Contains(splitfile[i], "P3") {
			magicNumber = "P3"
		} else if strings.Contains(splitfile[i], "P6") {
			magicNumber = "P6"
		}
		if strings.HasPrefix(splitfile[i], "#") && maxval != 0 {
			headersize = counter
		}
		splitl := strings.SplitN(splitfile[i], " ", -1)
		if width == 0 && height == 0 && len(splitl) >= 2 {
			width, err = strconv.Atoi(splitl[0])
			height, err = strconv.Atoi(splitl[1])
			headersize = counter
		}
		if maxval == 0 && width != 0 {
			maxval, err = strconv.Atoi(splitfile[i])
			headersize = counter
		}
		counter++

	}

	data := make([][]Pixel, height)

	for j := 0; j < height; j++ {
		data[j] = make([]Pixel, width)
	}
	var splitdata []string

	if counter > headersize {
		for i := 0; i < height; i++ {
			splitdata = strings.SplitN(splitfile[headersize+1+i], " ", -1)
			fmt.Println(splitdata)
			fmt.Print()
			//for k := 0; k < width; k++ {
			for j := 0; j < width*3; j += 3 {
				r, _ := strconv.Atoi(splitdata[j])
				fmt.Println("R : ", j, " | ", splitdata[j])
				g, _ := strconv.Atoi(splitdata[j+1])
				fmt.Println("G : ", j+1, " | ", splitdata[j])
				b, _ := strconv.Atoi(splitdata[j+2])
				fmt.Println("B : ", j+2, " | ", splitdata[j])
				data[i][j/3] = Pixel{R: uint8(r), G: uint8(g), B: uint8(b)}
				//	}
			}
		}
	}
	return &PPM{data: data, width: width, height: height, magicNumber: magicNumber, max: uint8(maxval)}, err
}

// Size returns the width and height of the image.
func (ppm *PPM) Size() (int, int) {
	return ppm.width, ppm.height
}

// At returns the value of the pixel at (x, y).
func (ppm *PPM) At(x, y int) Pixel {
	return ppm.data[x][y]
}

// Set sets the value of the pixel at (x, y).
func (ppm *PPM) Set(x, y int, value Pixel) {
	ppm.data[x][y] = Pixel{}
}

// Save saves the PPM image to a file and returns an error if there was a problem.
/*func (ppm *PPM) Save(filename string) error{
	var
}*/

// Invert inverts the colors of the PPM image.
func (ppm *PPM) Invert() {
	for i := 0; i < len(ppm.data); i++ {
		for j := 0; j < len(ppm.data[0]); j++ {
			ppm.data[i][j].R = ppm.max - ppm.data[i][j].R
			ppm.data[i][j].G = ppm.max - ppm.data[i][j].G
			ppm.data[i][j].B = ppm.max - ppm.data[i][j].B
		}
	}
}

// Flip flips the PPM image horizontally.
func (ppm *PPM) Flip() {
	if len(ppm.data[0]) > 0 {
		for i := 0; i < len(ppm.data); i++ {
			for j := 0; j < len(ppm.data[i])/2; j++ {
				startdata := ppm.data[i][j]
				ppm.data[i][j] = ppm.data[i][len(ppm.data[i])-1-j]
				ppm.data[i][len(ppm.data[i])-1-j] = startdata
			}
		}
	}
}

// Flop flops the PPM image vertically.
func (ppm *PPM) Flop() {
	if len(ppm.data) > 0 {
		for i := 0; i < len(ppm.data)/2; i++ {
			startdata := ppm.data[i]
			ppm.data[i] = ppm.data[len(ppm.data)-1-i]
			ppm.data[len(ppm.data)-1-i] = startdata
		}
	}
}

// SetMagicNumber sets the magic number of the PPM image.
func (ppm *PPM) SetMagicNumber(magicNumber string) {
	ppm.magicNumber = magicNumber
}

func (ppm *PPM) SetMaxValue(maxValue uint8) {
	oldMax := ppm.max
	ppm.max = maxValue
	for i := 0; i < len(ppm.data); i++ {
		for j := 0; j < len(ppm.data[0]); j++ {
			ppm.data[i][j].R = uint8(float64(ppm.data[i][j].R) * float64(ppm.max) / float64(oldMax))
			ppm.data[i][j].G = uint8(float64(ppm.data[i][j].G) * float64(ppm.max) / float64(oldMax))
			ppm.data[i][j].B = uint8(float64(ppm.data[i][j].B) * float64(ppm.max) / float64(oldMax))
		}
	}
}

// Rotate90CW rotates the PPM image 90Â° clockwise.
func (ppm *PPM) Rotate90CW() {
	// Height = Colums = Colonne vers le bas
	// Width = Rows = Ligne vers la droite
	NumRows := ppm.width
	NumColums := ppm.height
	for i := 0; i < NumColums; i++ {
		for j := i + 1; j < NumRows; j++ {
			vartemp := ppm.data[i][j]
			ppm.data[i][j] = ppm.data[j][i]
			ppm.data[j][i] = vartemp
		}
	}
	for i := 0; i < NumColums; i++ {
		for j := 0; j < NumRows/2; j++ {
			vartemp := ppm.data[i][j]
			ppm.data[i][j] = ppm.data[i][NumRows-j-1]
			ppm.data[i][NumRows-j-1] = vartemp
		}
	}
}

// ToPGM converts the PPM image to PGM.
func (ppm *PPM) ToPGM() *PGM {
	// Height = Colums = Colonne vers le bas
	// Width = Rows = Ligne vers la droite
	var newmagicnumber string

	if ppm.magicNumber == "P3" {
		newmagicnumber = "P2"
	} else if ppm.magicNumber == "P6" {
		newmagicnumber = "P5"
	}
	Numrows := ppm.width
	NumColumns := ppm.height
	var newdata = make([][]uint8, NumColumns)
	for i := 0; i < NumColumns; i++ {
		newdata[i] = make([]uint8, Numrows)
		for j := 0; j < Numrows; j++ {
			{
				newdata[i][j] = uint8((int(ppm.data[i][j].R) + int(ppm.data[i][j].G) + int(ppm.data[i][j].B)) / 3)
			}
		}
	}
	return &PGM{data: newdata, width: Numrows, height: NumColumns, max: ppm.max, magicNumber: newmagicnumber}
}

// ToPBM converts the PPM image to PBM.
func (ppm *PPM) ToPBM() *PBM {
	var newmagicnumber string

	if ppm.magicNumber == "P3" {
		newmagicnumber = "P1"
	} else if ppm.magicNumber == "P6" {
		newmagicnumber = "P4"
	}
	Numrows := ppm.width
	NumColumns := ppm.height
	var newdata = make([][]bool, NumColumns)
	for i := 0; i < NumColumns; i++ {
		newdata[i] = make([]bool, Numrows)
		for j := 0; j < Numrows; j++ {
			newdata[i][j] = uint8((int(ppm.data[i][j].R)+int(ppm.data[i][j].G)+int(ppm.data[i][j].B))/3) < ppm.max/2
		}
	}
	return &PBM{data: newdata, width: Numrows, height: NumColumns, magicNumber: newmagicnumber}
}