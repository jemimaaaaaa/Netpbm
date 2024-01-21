package Netpbm 

// PBM represents a PBM image
type PBM struct{
    data [][]bool
    width, height int
    magicNumber string
}

// Function ReadPBM reads a PBM  file and creates a data structure to represent the image
func ReadPBM(filename string) (*PBM, error) {
	file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	scanner.Scan()
	magicNumber := scanner.Text()

    if magicNumber != "P1" && magicNumber != "P4" {
        return nil, errors.New("unsupported file type")
    }

// Read image dimensions
    scanner.Scan()
	dimensions := strings.Fields(scanner.Text())
    if len(dimensions) != 2 {
        return nil, errors.New("invalid image dimensions")
    }
    width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

    var data [][]bool
	for scanner.Scan() {
		line := scanner.Text()
		if magicNumber == "P1" {row := make([]bool, width)
			for i, char := range strings.Fields(line) {
				pixel, _ := strconv.Atoi(char)
				row[i] = pixel == 1
			}
            data = append(data, row)
            } else if magicNumber == "P4" {// Allows binary data to be read
                reader := bufio.NewReader(file)
                // Ignore white space that might exist after dimensions
                reader.Discard(width % 8)
                for y := 0; y < height; y++ {
                    row := make([]bool, width)
                    for x := 0; x < width; x += 8 {
                // Read one byte (8 bits) at a time              
                        b, err := reader.ReadByte()
                        if err != nil {
                            return nil, err
                        }
                        // Convert Byte to Booleans
                        for i := 0; i < 8; i++ {
                        //  Check if the bit at position i is set                           
                         row[x+i] = b&(1<<(7-i)) != 0
                        }
                    }
                    data = append(data, row)
                }
            }
        }
    
        if err := scanner.Err(); err != nil {
            return nil, err
        }
    
        return &PBM{
            data:        data,
            width:       width,
            height:      height,
            magicNumber: magicNumber,
        }, nil
            
    // Read data image
    for i := 0; i < height && scanner.Scan(); i++ {
		line := scanner.Text()
		pbm.data[i] = parseRow(line)
	}

	return pbm, nil
}

// The size function provides information on the image's width and height.
func (pbm *PBM) Size() (int,int){
	return pbm.width, pbm.height
}

// This method returns the pixel value at (x, y) coordinates.
func (pbm *PBM) At(x, y int) bool {
    if x < 0 || y  <  0 || x >= pbm.width || y >= pbm.height {
        return false
    }
    return pbm.data[x][y]
}

// This method changes the pixel value at (x, y) coordinates.
func (pbm *PBM) Set(x, y int, value bool){
    if x < 0 || y  <  0 || x >= pbm.width || y >= pbm.height {
		return
	}
	pbm.data[y][x] = value
}

// Save saves the PBM image to a file and returns an error if there is a problem.
func (pbm *PBM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return error
	}
	defer file.Close()

	// Write magic number
	fmt.Fprintln(file, pbm.magicNumber)

	// Write image dimensions
	fmt.Fprintf(file, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)

    // 	Write the pixels
    if pbm.magicNumber == "P1" {
    for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Fprint(file, "1 ")
			} else {
				fmt.Fprint(file, "0 ")
			}
		}
		fmt.Fprintln(file)
	}

	return nil
}

// Invert inverts the colors of the PBM image.
func (pbm *PBM) Invert() {
	for i := 0; i < pbm.height; i++ {
		for j := 0; j < pbm.width; j++ {
			pbm.data[i][j] = !pbm.data[i][j]
		}
	}
}

// Flip flips the PBM image horizontally.
func (pbm *PBM) Flop() {
	for i := 0; i < pbm.height/2; i++ {
		pbm.data[i], pbm.data[pbm.height-i-1] = pbm.data[pbm.height-i-1], pbm.data[i]
	}
}

// SetMagicNumber sets the magic number of the PBM image.
func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
}