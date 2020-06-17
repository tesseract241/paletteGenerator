package paletteGenerator

import (
    "errors"
    "image/color"
    "math"
    "math/rand"
    "time"
)

const (
    nonPositiveNumber       = "Asked for a non-positive number of colors"
)
//This random palette generator is based on a perceptual model of color vision, which says that, while 
//the cones in our eyes are sensitive to specific frequency bands, our visual cortex codifies colors
//based on where they lie among three different axes, at the ends of which are certain color components.
//Specifically this program uses a Green-Magenta, Red-Cyan and Blue-Yellow choice of coordinates.
//Thus we generate colors that are as distant as possible in this choice of coordinates, specifically generating a 
//cubical crystal, the nodes of which are our colors. We skip the nodes on the diagonal, 
//as they are less visibile according to our model, and eventually convert back to an RGB color model for ease of use.

//Converts notation for a 3-D array of size edge*edge*edge from single to triple index
func coordinatesFromIndex(index int, edge int) (x int, y int, z int) {
    x = index/(edge*edge)
    dummy := (index - x*edge*edge)
    y = dummy/edge
    z = dummy - y*edge
    return
}

//Given a color represented in a perceptive model, converts it to RGB
func rotateColorSpace(green_magenta, red_cyan, blue_yellow int) (red, green, blue int) {
    red     = (green_magenta    + blue_yellow   )/2
    green   = (red_cyan         + blue_yellow   )/2
    blue    = (green_magenta    + red_cyan      )/2
    return
}

//Finds how many nodes a cube must have for it to contain enough of them to cover numberOfColors, excluding the diagonal and an area around it.
func findEdge(numberOfColors int) int {
    //Indicating the thickness of the area around the diagonal by L, the number of nodes is found by solving the equation
    //n^3 -6Ln^2+(6L-1)n-numberOfColors = 0
    //Making the choice L=1/6 leads to a much simpler equation, and it's also a good choice for it as it leads to skipping around a third of the colors.
    delta := math.Cbrt((math.Sqrt(float64(27*27*numberOfColors*numberOfColors + 27*4*numberOfColors)) + float64(27*numberOfColors + 2))/2)
    return int(math.Ceil((delta + 1/delta + 1.)/3))
}

//Takes the number of colors to generate, returns the array containing them and an error in case the number was not positive
func GeneratePalette(numberOfColors int) ([]color.RGBA, error) {
    if numberOfColors <= 0 {
        return nil, errors.New(nonPositiveNumber)
    }
    rand.Seed(time.Now().UnixNano())
    nodesPerSide := findEdge(numberOfColors)
    skipThreshold := math.Ceil(float64(nodesPerSide)/6)
    nodeSize := 255/(nodesPerSide - 1)
    randomColors := rand.Perm(nodesPerSide*nodesPerSide*nodesPerSide)
    colors := make([]color.RGBA, numberOfColors)
    j := 0
    for i:=0;i<numberOfColors;i++ {
        var x, y, z int
        for ; ; j++ {
            x, y, z = coordinatesFromIndex(randomColors[j], nodesPerSide)
            //Avoid colors among the diagonal of the cube, as they are low contrast
            if math.Abs(float64(x - y)) >= skipThreshold || math.Abs(float64(y - z)) >= skipThreshold || math.Abs(float64(x - z)) >= skipThreshold {
                j++
                break
            }
        }
        x, y, z = x*nodeSize, y*nodeSize, z*nodeSize
        x, y, z = rotateColorSpace(x, y, z)
     colors[i] = color.RGBA{uint8(x), uint8(y), uint8(z), 255}
    }
    return colors, nil
}

func init() {
    rand.Seed(time.Now().UnixNano())
}
