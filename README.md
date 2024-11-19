# Sodium Framework Documentation

## Overview

The **Sodium Framework** is a UI framework built in Go, leveraging SDL2 for rendering graphical elements. It allows developers to define and manage user interfaces declaratively and programmatically. The framework includes event handling for UI interactions such as clicks, hovers, presses, and focus changes.

This documentation will walk you through the components and usage of the Sodium Framework.

---

## Table of Contents
- [Sodium Framework Documentation](#sodium-framework-documentation)
	- [Overview](#overview)
	- [Table of Contents](#table-of-contents)
	- [Project Structure](#project-structure)
	- [Key Features](#key-features)
	- [How It Works](#how-it-works)
	- [Getting Started](#getting-started)
		- [Prerequisites](#prerequisites)
		- [Installation](#installation)
	- [Usage Guide](#usage-guide)
		- [Defining the UI](#defining-the-ui)
		- [Handling Events](#handling-events)
	- [Code Reference](#code-reference)
		- [Main Functions](#main-functions)
		- [Components](#components)
	- [Example](#example)
		- [UI Definition (`main.ui`)](#ui-definition-mainui)
		- [Main Code (`main.go`)](#main-code-maingo)
	- [Future Improvements](#future-improvements)

---

## Project Structure

The project is divided into the following modules:

1. **`main`**: Entry point for the application.
2. **`components`**: Handles rendering and logic for UI components like buttons, rectangles, etc.
3. **`sodium`**: Core of the framework, including initialization, event handling, and rendering logic.
4. **`sodium/types`**: Defines the data structures used in the framework, such as `Window`, `Element`, and event maps.
5. **`sodium/utils`**: Utility functions, such as color conversion and dimension parsing.

---

## Key Features

- **Declarative UI**: Define the UI in XML-like syntax (e.g., `<sodium>`).
- **Custom Event Handlers**: Bind handlers for user interactions like clicks, hovers, presses, and focus changes.
- **Dynamic Rendering**: Update UI elements dynamically during runtime.
- **SDL2 Integration**: Leverages SDL2 for cross-platform rendering.

---

## How It Works

1. **Initialization**: The `Start` function initializes SDL2 and the main application window.
2. **UI Loading**: The UI layout is defined in an XML-like format and loaded using `sodium.LoadUI()`.
3. **Event Handling**: Events are handled via custom functions defined in the `Handlers` map.
4. **Rendering Loop**: The `Loop` function updates the UI and processes events in a continuous loop.

---

## Getting Started

### Prerequisites

1. **Go**: Make sure Go is installed on your machine.
2. **SDL2**: Install the SDL2 library. Use package managers like `apt`, `brew`, `pacman`, or `choco` depending on your OS.

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/E3nviction/sodium
   cd sodium
   ```
2. Install dependencies:
   ```bash
   go get ./...
   ```

3. Run the project:
   ```bash
   go run .
   ```

---

## Usage Guide

### Defining the UI

The UI is defined in an XML-like format. Here's an example:

```xml
<sodium>
  <window Width="800" Height="600">
    <rect
      X="0"
      Y="0"
      Width="100%"
      Height="100%"
      BackgroundColor="#1b1b1b"
    />
    <button
      Id="buttonhello-world"
      X="10px"
      Y="10px"
      Width="150px"
      Height="50px"
      BackgroundColor="#333333"
      OnClick="Click"
    />
  </window>
</sodium>
```

- **Elements**:
  - `<rect>`: Renders a rectangle.
  - `<button>`: Renders a clickable button with event handlers.

### Handling Events

Events are handled by binding functions in the `Handlers` map:

```go
w.Handlers = map[string]func(element types.Element){
    "Click": func(self types.Element) {
        fmt.Println("Button clicked!")
    },
    "Hover": func(self types.Element) {
        self.Attributes["BackgroundColor"] = "#552222"
    },
}
```

- **Supported Events**:
  - `Click`: Triggered when the element is clicked.
  - `Hover`: Triggered when the mouse hovers over the element.
  - `Press`: Triggered when the element is pressed.
  - `OffFocus`: Triggered when the element loses focus.

---

## Code Reference

### Main Functions

- **`Init`**: Initializes the UI and sets up event handlers.
- **`Loop`**: Main rendering loop for the application.
- **`Start`**: Entry point for the framework. Initializes SDL2 and runs the application.

### Components

- **`HandleButton`**: Manages button rendering and event handling.
- **`Query`**: Searches for UI elements by ID.

---

## Example

Here’s a minimal example of how to create a basic application using Sodium:

### UI Definition (`main.ui`)
```xml
<sodium>
  <window Width="800" Height="600">
    <button
      Id="greetButton"
      X="100px"
      Y="100px"
      Width="200px"
      Height="50px"
      BackgroundColor="#333333"
      OnClick="Greet"
    />
  </window>
</sodium>
```

### Main Code (`main.go`)
```go
package main

import (
	"fmt"
	"sodium"
	"sodium/types"
)

func Init(w *types.Window) {
	UI, _ := sodium.LoadUI("main.ui")
	w.AddtoDictionary("UI", UI)

	w.Handlers = map[string]func(element types.Element){
		"Greet": func(self types.Element) {
			fmt.Println("Hello from Sodium!")
		},
	}
	sodium.DrawWindow(UI, w)
}

func Loop(w *types.Window) {
	UI := w.GetFromDictionary("UI")
	sodium.DrawUI(UI.([]types.Element), w)
}

func main() {
	sodium.Start(Init, Loop)
}
```

---

## Future Improvements

- **Extended Component Library**: Add more UI components (e.g., sliders, text inputs).
- **Labels and Text**: Support rendering text and labels.
- **Rounded Corners, Borders, Shadows and Gradients**: Introduce rounded corners, borders, shadows, and gradients.
- **Enhanced Event Handling**: Support more complex interactions and custom events.
- **Styling**: Introduce CSS-like styling for components.

---

This documentation provides an overview of the Sodium Framework, helping you get started with building interactive UI applications in Go. For any additional help, feel free to contribute or open an issue!