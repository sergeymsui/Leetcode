package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

    if image[sr][sc] == color {
        return image
    }

	currentColor := image[sr][sc]

	image[sr][sc] = color

	if (sr-1) >= 0 && image[sr-1][sc] == currentColor {
		floodFill(image, sr-1, sc, color)
	}

	if (sc-1) >= 0 && image[sr][sc-1] == currentColor {
		floodFill(image, sr, sc-1, color)
	}

	if sr+1 < len(image) && image[sr+1][sc] == currentColor {
		floodFill(image, sr+1, sc, color)
	}

	if sc+1 < len(image[sr]) && image[sr][sc+1] == currentColor {
		floodFill(image, sr, sc+1, color)
	}

	return image
}

func main() {
	image := [][]int{{0, 0, 0}, {0, 0, 0}}
	sr := 0
	sc := 0
	color := 0

	pix := floodFill(image, sr, sc, color)

	for _, s := range pix {
		for _, v := range s {
			fmt.Printf("%d ", v)
		}
		fmt.Println("")
	}
}
