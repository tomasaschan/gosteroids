package fonts

import (
	"embed"
	"fmt"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *.ttf
var fs embed.FS

func ArcadeClassic() (font.Face, error) {
	return readTTF("ArcadeClassic-ov2x.ttf")
}

func AtariFont() (font.Face, error) {
	return readOpenType("AtariFontFullVersion-ZJ23.ttf")
}

func readTTF(fontFile string) (font.Face, error) {
	data, err := fs.ReadFile(fontFile)
	if err != nil {
		return nil, fmt.Errorf("reading font file %s: %w", fontFile, err)
	}

	tt, err := truetype.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("parsing font file %s: %w", fontFile, err)
	}

	fontFace := truetype.NewFace(tt, &truetype.Options{
		Size: 24,
		DPI:  72,
	})

	return fontFace, nil
}

func readOpenType(fontFile string) (font.Face, error) {
	data, err := fs.ReadFile(fontFile)
	if err != nil {
		return nil, fmt.Errorf("reading font file %s: %w", fontFile, err)
	}

	tt, err := opentype.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("parsing font file %s: %w", fontFile, err)
	}

	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 24,
		DPI:  72,
	})
	if err != nil {
		return nil, fmt.Errorf("create new font face for %s: %w", fontFile, err)
	}

	return fontFace, nil
}
