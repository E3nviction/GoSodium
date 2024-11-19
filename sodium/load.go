package sodium

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"sodium/types"
	"strings"
)

// Tag and attribute regex patterns
var tagRegex = regexp.MustCompile(`<(/?)(\w+)([^/>]*)\s*(/?)>`)
var attrRegex = regexp.MustCompile(`(\w+)=("[^"]*"|\S+)`)

// LoadUI loads and parses a .ui file into a nested structure of Elements
func LoadUI(filename string) ([]types.Element, error) {
	if !strings.HasSuffix(filename, ".ui") {
		return nil, errors.New("file must have a .ui extension")
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	stack := []types.Element{{Tag: "root", Attributes: make(map[string]string), Children: []types.Element{}}}

	for _, match := range tagRegex.FindAllStringSubmatch(string(content), -1) {
		closingTag, tagName, attrs, selfClosing := match[1], match[2], match[3], match[4]

		// If closing tag, pop the last element
		if closingTag == "/" {
			if len(stack) > 1 {
				child := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, child)
			}
			continue
		}

		// Parse attributes into a map
		attributes := make(map[string]string)
		for _, attrMatch := range attrRegex.FindAllStringSubmatch(attrs, -1) {
			key := attrMatch[1]
			value := strings.Trim(attrMatch[2], `"`)
			attributes[key] = value
		}

		// Add new element to the stack
		newElement := types.Element{Tag: tagName, Attributes: attributes, Children: []types.Element{}}
		stack = append(stack, newElement)

		// If self-closing, immediately pop
		if selfClosing == "/" {
			if len(stack) > 1 {
				child := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, child)
			}
		}
	}

	// Return the children of the root element
	return stack[0].Children, nil
}
