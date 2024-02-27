package main

import "fmt"

type Mobile struct {
	Brand    string
	Model    string
	OS       string
	Year     int
	Features []string
}

func (m Mobile) DisplayInfo() {
	fmt.Printf("Brand: %s\n", m.Brand)
	fmt.Printf("Model: %s\n", m.Model)
	fmt.Printf("Operating System: %s\n", m.OS)
	fmt.Printf("Year: %d\n", m.Year)

	fmt.Println("Features:")
	for _, feature := range m.Features {
		fmt.Printf("- %s\n", feature)
	}
}

func main() {
	iPhone := Mobile{
		Brand:    "Apple",
		Model:    "iPhone 12",
		OS:       "iOS",
		Year:     2020,
		Features: []string{"Face ID", "A14 Bionic chip", "Dual-camera system"},
	}

	iPhone.DisplayInfo()
}
