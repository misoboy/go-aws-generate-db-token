package theme

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

var (
	background = &color.RGBA{0xf5, 0xf5, 0xf5, 0xff}
	button = &color.RGBA{0xd9, 0xd9, 0xd9, 0xff}
	disabledButton = &color.RGBA{0xe7, 0xe7, 0xe7, 0xff}
	text = &color.RGBA{0x21, 0x21, 0x21, 0xff}
	disabledText = &color.RGBA{0x80, 0x80, 0x80, 0xff}
	icon = &color.RGBA{0x21, 0x21, 0x21, 0xff}
	disabledIcon = &color.RGBA{0x80, 0x80, 0x80, 0xff}
	hyperlink = &color.RGBA{0x0, 0x0, 0xd9, 0xff}
	placeholder = &color.RGBA{0x88, 0x88, 0x88, 0xff}
	primary = &color.RGBA{0x9f, 0xa8, 0xda, 0xff}
	hover = &color.RGBA{0xe7, 0xe7, 0xe7, 0xff}
	scrollBar = &color.RGBA{0x0, 0x0, 0x0, 0x99}
	shadow = &color.RGBA{0x0, 0x0, 0x0, 0x33}
)

// customTheme is a simple demonstration of a bespoke theme loaded by a Fyne app.
type customTheme struct {
}

func (customTheme) BackgroundColor() color.Color {
	return background
}

func (customTheme) ButtonColor() color.Color {
	return button
}

func (customTheme) DisabledButtonColor() color.Color {
	return disabledButton
}

func (customTheme) HyperlinkColor() color.Color {
	return hyperlink
}

func (customTheme) TextColor() color.Color {
	return text
}

func (customTheme) DisabledTextColor() color.Color {
	return disabledText
}

func (customTheme) IconColor() color.Color {
	return icon
}

func (customTheme) DisabledIconColor() color.Color {
	return disabledIcon
}

func (customTheme) PlaceHolderColor() color.Color {
	return placeholder
}

func (customTheme) PrimaryColor() color.Color {
	return primary
}

func (customTheme) HoverColor() color.Color {
	return hover
}

func (customTheme) FocusColor() color.Color {
	return primary
}

func (customTheme) ScrollBarColor() color.Color {
	return primary
}

func (customTheme) ShadowColor() color.Color {
	return shadow
}

func (customTheme) TextSize() int {
	return 15
}

func (customTheme) TextFont() fyne.Resource {
	return theme.DefaultTextBoldFont()
}

func (customTheme) TextBoldFont() fyne.Resource {
	return theme.DefaultTextBoldFont()
}

func (customTheme) TextItalicFont() fyne.Resource {
	return theme.DefaultTextBoldItalicFont()
}

func (customTheme) TextBoldItalicFont() fyne.Resource {
	return theme.DefaultTextBoldItalicFont()
}

func (customTheme) TextMonospaceFont() fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func (customTheme) Padding() int {
	return 10
}

func (customTheme) IconInlineSize() int {
	return 20
}

func (customTheme) ScrollBarSize() int {
	return 10
}

func (customTheme) ScrollBarSmallSize() int {
	return 5
}

func NewCustomTheme() fyne.Theme {
	return &customTheme{}
}
