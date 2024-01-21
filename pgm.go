package Netpbm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PGM struct {
	data          [][]uint8
	width, height int
	magicNumber   string
	max           uint8
}

// Function to read a PGM file and create a PGM instance
func ReadPGM(filename string) (*PGM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	

	var width, height, max int
	var data [][]uint8

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	magicNumber := scanner.Text()
	if magicNumber != "P2" && magicNumber != "P5" {
		return nil, errors.New("type de fichier non pris en charge")
	}

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			_, err := fmt.Sscanf(line, "%d %d", &width, &height)
			if err == nil {
				break
			} else {
				fmt.Println("Largeur ou hauteur invalide :", err)
			}
		}
	}

	scanner.Scan()
	max, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, errors.New("valeur maximale de pixel invalide")
	}

	for scanner.Scan() {
		line := scanner.Text()
		if magicNumber == "P2" {
			row := make([]uint8, 0)
			for _, char := range strings.Fields(line) {
				pixel, err := strconv.Atoi(char)
				if err != nil {
					fmt.Println("Erreur de conversion en entier :", err)
				}
				if pixel >= 0 && pixel <= max {
					row = append(row, uint8(pixel))
				} else {
					fmt.Println("Valeur de pixel invalide :", pixel)
				}
			}
			data = append(data, row)
		}
	}

	return &PGM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
		max:         uint8(max),
	}, nil
}

// function to obtain the size of the PGM image
func (pgm *PGM) Size() (int, int) {
	return pgm.width, pgm.height
}

// At returns the value of the pixel at (x, y).
func (pgm *PGM) At(x, y int) uint8 {
	return pgm.data[x][y]
}

// Set sets the value of the pixel at (x, y).
func (pgm *PGM) Set(x, y int, value uint8) {
	pgm.data[x][y] = value
}

func (pgm *PGM) Save(filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the magic number, width, height, max value
	fmt.Fprintf(file, "%s\n%d %d\n%d\n", pgm.magicNumber, pgm.width, pgm.height, pgm.max)

	// Write the pixel data
	for _, row := range pgm.data {
		for _, pixel := range row {
			fmt.Fprintf(file, "%d ", pixel)
		}
		fmt.Fprintln(file)
	}

	return nil
}

// Invert inverts the colors of the PGM image.
func (pgm *PGM) Invert() {
	for i := 0; i < len(pgm.height); i++ {
		for j := 0; j < len(pgm.width[i]); j++ {
			pgm.data[i][j] = uint8(pgm.max) - pgm.data[i][j]
		}
	}
}

// Flip flips the PGM image horizontally.
func (pgm *PGM) Flip() {
	NumRows := pgm.width
	Numcolums := pgm.height
	for i := 0; i < NumRows; i++ {
		for j := 0; j < Numcolums/2; j++ {
			pgm.data[i][j], pgm.data[i][Numcolums-j-1] = pgm.data[i][Numcolums-j-1], pgm.data[i][j]
		}
	}
}

// Flop flops the PGM image vertically.
func (pgm *PGM) Flop() {
	numRows := len(pgm.data)
	if numRows == 0 {
		return
	}
	for i := 0; i < numRows/2; i++ {
		pgm.data[i], pgm.data[numRows-i-1] = pgm.data[numRows-i-1], pgm.data[i]
	}
}

// Allows you to update the "magic number" associated with a PGM image in an instance of the PGM structure.
func (pgm *PGM) SetMagicNumber(magicNumber string) {
	pgm.magicNumber = magicNumber
}

//  Allows you to change the pixel values of a PGM image by adjusting their scale to a new specified maximum.
func (pgm *PGM) SetMaxValue(maxValue uint8) {
	oldmax := pgm.max
	pgm.max = maxValue
	for i := 0; i < pgm.height; i++ {
		for j := 0; j < pgm.width; j++ {

			pgm.data[i][j] = pgm.data[i][j] * uint8(5) / oldmax
		}
	}

}

// Rotate90CW rotates the PGM image 90° clockwise.
func (pgm *PGM) Rotate90CW() {
	NumRows := pgm.width
	NumColums := pgm.height
	for i := 0; i < len(pgm.data); i++ {
		var temp uint8
		for j := i + 1; j < len(pgm.data[0]); j++ {
			temp = pgm.data[i][j]
			pgm.data[i][j] = pgm.data[j][i]
			pgm.data[j][i] = temp
		}
	}
	for i := 0; i < NumColums; i++ {
		for j := 0; j < NumRows/2; j++ {
			temp := pgm.data[i][j]
			pgm.data[i][j] = pgm.data[i][NumRows-j-1]
			pgm.data[i][NumRows-j-1] = temp
		}
	}
}

//  Displays the values of the Data matrix on the console, row by row, with each value separated by a space
func display(data [][]uint8) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}
}
