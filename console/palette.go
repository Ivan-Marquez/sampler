package console

import (
	"fmt"
	ui "github.com/sqshq/termui"
)

type Theme string

const (
	ThemeDark  Theme = "dark"
	ThemeLight Theme = "light"
)

const (
	ColorOlive       ui.Color = 178
	ColorDeepSkyBlue ui.Color = 39
	ColorDeepPink    ui.Color = 162
	ColorDarkGrey    ui.Color = 240
	ColorWhite       ui.Color = 7
	ColorBlack       ui.Color = 0
	ColorClear       ui.Color = -1
)

type Palette struct {
	Colors []ui.Color
	// TODO Menu colors, like Dark, Medium, Light etc
}

func GetPalette(theme Theme) Palette {
	switch theme {
	case ThemeDark:
		return Palette{Colors: []ui.Color{ColorOlive, ColorDeepSkyBlue, ColorDeepPink, ColorOlive, ColorDeepSkyBlue, ColorDeepPink}}
	case ThemeLight:
		return Palette{Colors: []ui.Color{ColorOlive, ColorDeepSkyBlue, ColorDeepPink, ColorOlive, ColorDeepSkyBlue, ColorDeepPink}}
	default:
		panic(fmt.Sprintf("Following theme is not supported: %v", theme))
	}
}