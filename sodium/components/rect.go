package components

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"sodium/types"
	"sodium/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func HandleRect(child types.Element, window *types.Window) {
	// Retrieve attributes
	x := child.Attributes["X"]
	y := child.Attributes["Y"]
	width := child.Attributes["Width"]
	height := child.Attributes["Height"]
	backgroundcolorStr := child.Attributes["BackgroundColor"]
	alpha := child.Attributes["Alpha"]
	rounding := child.Attributes["Rounding"]
	margin := child.Attributes["Margin"]
	padding := child.Attributes["Padding"]
	paddingTop := child.Attributes["PaddingTop"]
	paddingBottom := child.Attributes["PaddingBottom"]
	paddingLeft := child.Attributes["PaddingLeft"]
	paddingRight := child.Attributes["PaddingRight"]

	// Validate required attributes
	if x == "" || y == "" || width == "" || height == "" || backgroundcolorStr == "" {
		fmt.Println("X, Y, Width, Height, BackgroundColor attributes are required for the rect.")
		return
	}

	// Convert color
	color, err := utils.ConvertColor(backgroundcolorStr, alpha)
	if err != nil {
		fmt.Println("Invalid Background color:", color)
		return
	}

	winwidth, winheight := window.Window.GetSize()

	// Parse height value first as it may be referenced in other calculations
	heightInt, err := parseDimension(height, int(winheight), child.Attributes)
	if err != nil {
		fmt.Println("Invalid height:", height)
		return
	}

	// Parse x, y, and width as absolute values
	xInt, err := parseDimension(x, int(winwidth), child.Attributes)
	if err != nil {
		fmt.Println("Invalid x:", x)
		return
	}
	yInt, err := parseDimension(y, int(winheight), child.Attributes)
	if err != nil {
		fmt.Println("Invalid y:", y)
		return
	}

	// Calculate width, supporting percentages
	widthInt, err := parseDimension(width, int(winwidth), child.Attributes)
	if err != nil {
		fmt.Println("Invalid width:", width)
		return
	}

	if rounding == "" {
		rounding = "0px"
	}

	// Parse rounding (optional)
	if strings.HasSuffix(rounding, "px") {
		rounding = strings.TrimSuffix(rounding, "px")
	}
	_, err = strconv.Atoi(rounding)
	if err != nil {
		fmt.Println("Invalid rounding:", rounding)
		return
	}

	// Handle margin (optional)
	marginInt := 0
	if margin != "" {
		marginInt, err = parseDimension(margin, int(winheight), child.Attributes)
		if err != nil {
			fmt.Println("Invalid margin:", margin)
			return
		}
		// Apply margin: shift the rect and reduce the width/height by margin value
		xInt += marginInt
		yInt += marginInt
		widthInt -= marginInt * 2
		heightInt -= marginInt * 2
	}

	// Handle padding (optional)
	if padding != "" {
		paddingInt, err := parseDimension(padding, int(winheight), child.Attributes)
		if err != nil {
			fmt.Println("Invalid padding:", padding)
			return
		}
		// Apply padding equally on all sides
		xInt -= paddingInt
		yInt -= paddingInt
		widthInt += paddingInt * 2
		heightInt += paddingInt * 2
	}

	// Handle individual padding (optional)
	if paddingTop != "" {
		paddingTopInt, err := parseDimension(paddingTop, int(winheight), child.Attributes)
		if err != nil {
			fmt.Println("Invalid paddingTop:", paddingTop)
			return
		}
		yInt += paddingTopInt
		heightInt -= paddingTopInt
	}
	if paddingBottom != "" {
		paddingBottomInt, err := parseDimension(paddingBottom, int(winheight), child.Attributes)
		if err != nil {
			fmt.Println("Invalid paddingBottom:", paddingBottom)
			return
		}
		heightInt -= paddingBottomInt
	}
	if paddingLeft != "" {
		paddingLeftInt, err := parseDimension(paddingLeft, int(winwidth), child.Attributes)
		if err != nil {
			fmt.Println("Invalid paddingLeft:", paddingLeft)
			return
		}
		xInt += paddingLeftInt
		widthInt -= paddingLeftInt
	}
	if paddingRight != "" {
		paddingRightInt, err := parseDimension(paddingRight, int(winwidth), child.Attributes)
		if err != nil {
			fmt.Println("Invalid paddingRight:", paddingRight)
			return
		}
		widthInt -= paddingRightInt
	}

	// Validate alpha (optional)
	if alpha == "" {
		alpha = "255"
	}
	if rounding == "" {
		rounding = "0"
	}

	// Draw the rectangle
	window.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	window.Renderer.FillRect(&sdl.Rect{X: int32(xInt), Y: int32(yInt), W: int32(widthInt), H: int32(heightInt)})
}

// parseDimension parses a dimension string and handles percentages or complex arithmetic expressions (e.g., "100% - height - 10px").
func parseDimension(dim string, total int, attributes map[string]string) (int, error) {
	// Handle expressions with operators (+, -, *, /)
	re := regexp.MustCompile(`(\d+%|\d+px|[+\-*/()a-zA-Z]+)`)
	tokens := re.FindAllString(dim, -1)

	var result int
	var currentOp string
	var totalPercent int
	var totalPixels int

	// Parse each token to evaluate the expression
	for _, token := range tokens {
		token = strings.TrimSpace(token)

		// Handle percentage
		if strings.HasSuffix(token, "%") {
			percentStr := strings.TrimSuffix(token, "%")
			percent, err := strconv.Atoi(percentStr)
			if err != nil {
				return 0, fmt.Errorf("invalid percentage value: %s", token)
			}
			percentValue := total * percent / 100
			totalPercent = applyOperator(totalPercent, percentValue, currentOp)
			currentOp = ""

		} else if strings.HasSuffix(token, "px") {
			// Handle pixel values (absolute values)
			pixelStr := strings.TrimSuffix(token, "px")
			pixelValue, err := strconv.Atoi(pixelStr)
			if err != nil {
				return 0, fmt.Errorf("invalid pixel value: %s", token)
			}
			totalPixels = applyOperator(totalPixels, pixelValue, currentOp)
			currentOp = ""

		} else if isAttribute(token, attributes) {
			// Handle reference to other attributes (e.g., "width", "height", "x", "y")
			attrValue, err := strconv.Atoi(attributes[token])
			if err != nil {
				// If the attribute has a suffix, we need to handle it correctly
				// Check if the attribute is in "px" or "%" format and adjust accordingly
				attrValue, err = parseDimension(attributes[token], total, attributes)
				if err != nil {
					return 0, fmt.Errorf("invalid attribute reference: %s", token)
				}
			}
			totalPixels = applyOperator(totalPixels, attrValue, currentOp)
			currentOp = ""

		} else if token == "+" || token == "-" || token == "*" || token == "/" {
			// Handle operators
			currentOp = token
		}
	}

	// Combine the results
	result = totalPercent + totalPixels
	return result, nil
}

// isAttribute checks if a token is one of the known attributes (x, y, width, height)
func isAttribute(token string, attributes map[string]string) bool {
	_, exists := attributes[token]
	return exists
}

// applyOperator applies an arithmetic operation to the accumulated result.
func applyOperator(accumulated int, value int, operator string) int {
	switch operator {
	case "+":
		return accumulated + value
	case "-":
		return accumulated - value
	case "*":
		return accumulated * value
	case "/":
		return accumulated / value
	default:
		return value
	}
}
