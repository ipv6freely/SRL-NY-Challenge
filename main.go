package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"strings"

	"github.com/BourgeoisBear/rasterm"
	_ "golang.org/x/term"
)

//go:embed assets/7220_IXR-D3L.jpg
var d3lJPG []byte

//go:embed assets/7220_IXR-D5.jpg
var d5JPG []byte

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(2)
	}

	model := strings.ToLower(strings.TrimSpace(os.Args[1]))
	data, err := selectModelBytes(model)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		usage()
		os.Exit(2)
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error decoding image:", err)
		os.Exit(1)
	}

	if err := render(img); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	fmt.Println()

}

func render(img image.Image) error {
	// Kitty
	if rasterm.IsKittyCapable() {
		if err := rasterm.KittyWriteImage(os.Stdout, img, rasterm.KittyImgOpts{}); err == nil {
			return nil
		}
	}

	// iTerm2 / WezTerm
	if rasterm.IsItermCapable() {
		opts := rasterm.ItermImgOpts{
			Name: "panel.jpg",
			DisplayInline: true,
		}
		if err := rasterm.ItermWriteImageWithOptions(os.Stdout, img, opts); err == nil {
			return nil
		}
	}

	return fmt.Errorf(
		"No supported terminal image protocol available",
	)
}

func selectModelBytes(model string) ([]byte, error) {
	switch model {
	case "d3l":
		return d3lJPG, nil
	case "d5":
		return d5JPG, nil
	default:
		return nil, fmt.Errorf(`unknown model %q (use "d3l" or "d5")`, model)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [d3l|d5]\n", os.Args[0])
}
