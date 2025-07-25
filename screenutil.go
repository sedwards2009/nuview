package nuview

import "github.com/gdamore/tcell/v2"

type ScreenWriter interface {
	GetContent(x, y int) (primary rune, combining []rune, style tcell.Style, width int)
	SetContent(x int, y int, primary rune, combining []rune, style tcell.Style)
	Size() (width, height int)
	Fill(rune, tcell.Style)
}

type TranslateScreenWriter interface {
	ScreenWriter
	AbsolutePosition(x int, y int) (absX int, absY int)
	NewClipXY(x int, y int) TranslateScreenWriter
}

// -------------------------------------------------------------------------
type TranslateScreenWriterAdapter struct {
	screen tcell.Screen
}

func NewTranslateScreenWriterAdapter(screen tcell.Screen) *TranslateScreenWriterAdapter {
	return &TranslateScreenWriterAdapter{screen: screen}
}

func (a *TranslateScreenWriterAdapter) GetContent(x, y int) (primary rune, combining []rune, style tcell.Style, width int) {
	return a.screen.GetContent(x, y)
}

func (a *TranslateScreenWriterAdapter) SetContent(x int, y int, primary rune, combining []rune, style tcell.Style) {
	a.screen.SetContent(x, y, primary, combining, style)
}

func (a *TranslateScreenWriterAdapter) Size() (width, height int) {
	return a.screen.Size()
}

func (a *TranslateScreenWriterAdapter) Fill(r rune, style tcell.Style) {
	a.screen.Fill(r, style)
}

func (a *TranslateScreenWriterAdapter) AbsolutePosition(x int, y int) (absX int, absY int) {
	return x, y
}

func (a *TranslateScreenWriterAdapter) NewClipXY(x int, y int) TranslateScreenWriter {
	width, height := a.screen.Size()
	c := NewClippingScreenWriter(a, x, y, width, height)
	return c
}

//-------------------------------------------------------------------------

type ClippingScreenWriter struct {
	writer ScreenWriter
	x      int
	y      int
	width  int
	height int
}

// NewClippingScreenWriter creates a new ScreenWriter that clips the output to the specified rectangle.
// The x and y parameters specify the top-left corner of the clipping area, and width and height specify
// the size of the clipping area.
//
// Functions on this type use a relative coordinate system based on the clipping area.
// For example, if the clipping area is 10x10 starting at (5,5), then SetContent(0, 0, ...) will write to
// (5, 5) in the original screen coordinates.
func NewClippingScreenWriter(w ScreenWriter, x, y, width, height int) *ClippingScreenWriter {
	return &ClippingScreenWriter{
		writer: w,
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func (c *ClippingScreenWriter) GetContent(x, y int) (primary rune, combining []rune, style tcell.Style, width int) {
	return c.writer.GetContent(c.x+x, c.y+y)
}

func (c *ClippingScreenWriter) SetContent(x int, y int, primary rune, combining []rune, style tcell.Style) {
	if x < 0 || y < 0 || x >= c.width || y >= c.height {
		return
	}
	c.writer.SetContent(x+c.x, y+c.y, primary, combining, style)
}

func (c *ClippingScreenWriter) Size() (width int, height int) {
	return c.width, c.height
}

func (c *ClippingScreenWriter) Fill(r rune, style tcell.Style) {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			c.writer.SetContent(c.x+x, c.y+y, r, nil, style)
		}
	}
}

func (c *ClippingScreenWriter) AbsolutePosition(x int, y int) (absX int, absY int) {
	return x + c.x, y + c.y
}

func (c *ClippingScreenWriter) NewClipXY(x int, y int) TranslateScreenWriter {
	r := &ClippingScreenWriter{
		writer: c.writer,
		x:      c.x + x,
		y:      c.y + y,
		width:  c.width - x,
		height: c.height - y,
	}
	return r
}
