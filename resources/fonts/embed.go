package fonts

import (
	"embed"
	"fmt"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *.ttf
var fs embed.FS

func ArcadeClassic() (font.Face, error) {
	return readFont("ArcadeClassic-ov2x.ttf")
}

func AtariFont() (font.Face, error) {
	return readFont("AtariFontFullVersion-ZJ23.ttf")
}

func readFont(fontFile string) (font.Face, error) {
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
