package components

import (
	"fmt"
	"strconv"
	"strings"

	"sodium/types"
	"sodium/utils"

	"github.com/veandco/go-sdl2/sdl"
)

func HandleButton(child types.Element, window *types.Window) {
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
	onClick := child.Attributes["OnClick"]
	onHover := child.Attributes["OnHover"]
	offFocus := child.Attributes["OffFocus"]
	onPress := child.Attributes["OnPress"]

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

	if onHover != "" {
		if handler, exists := window.Handlers[onHover]; exists {
			HandleButtonHover(window, xInt, yInt, widthInt, heightInt, handler, child)
		} else {
			fmt.Printf("No handler found for OnHover: %s\n", onHover)
		}
	}

	if offFocus != "" {
		if handler, exists := window.Handlers[offFocus]; exists {
			HandleButtonUnFocus(window, xInt, yInt, widthInt, heightInt, handler, child)
		} else {
			fmt.Printf("No handler found for OffFocus: %s\n", offFocus)
		}
	}

	if onPress != "" {
		if handler, exists := window.Handlers[onPress]; exists {
			HandleButtonPress(window, xInt, yInt, widthInt, heightInt, handler, child)
		} else {
			fmt.Printf("No handler found for OnPress: %s\n", onPress)
		}
	}

	if onClick != "" {
		if handler, exists := window.Handlers[onClick]; exists {
			HandleButtonClick(window, xInt, yInt, widthInt, heightInt, handler, child)
		} else {
			fmt.Printf("No handler found for OnClick: %s\n", onClick)
		}
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
