package sodium

import (
	"sodium/types"
)

func Query(query string, uiRoot []types.Element, window *types.Window) types.Element {
	// Iterate through the root UI elements
	for _, child := range uiRoot {
		// Check if the current element matches the query
		if child.Attributes["Id"] == query {
			return child
		}
		// Recursively search in children
		result := Query(query, child.Children, window)
		if result.Attributes != nil {
			return result // Return the element if found
		}
	}
	// Return an empty element if not found
	return types.Element{}
}
