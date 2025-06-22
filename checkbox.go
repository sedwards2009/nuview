package nuview

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Checkbox implements a simple box for boolean values which can be checked and
// unchecked.
//
// See https://github.com/rivo/tview/wiki/Checkbox for an example.
type Checkbox struct {
	*Box

	// Whether or not this checkbox is disabled/read-only.
	disabled bool

	// Whether or not this box is checked.
	checked bool

	// The text to be displayed before the input area.
	label string

	// The screen width of the label area. A value of 0 means use the width of
	// the label text.
	labelWidth int

	labelRight      string
	labelRightWidth int

	// The label style.
	labelStyle tcell.Style

	// The style of the unchecked checkbox.
	uncheckedStyle tcell.Style

	// The style of the checked checkbox.
	checkedStyle tcell.Style

	// The style of the checkbox when it is currently focused.
	focusStyle tcell.Style

	uncheckedString       string // String shown when unchecked
	checkedString         string // String shown when checked
	cursorCheckedString   string // String shown when checked and the cursor is on it
	cursorUncheckedString string // String shown when unchecked and the cursor is on it

	// An optional function which is called when the user changes the checked
	// state of this checkbox.
	changed func(checked bool)

	// An optional function which is called when the user indicated that they
	// are done entering text. The key which was pressed is provided (tab,
	// shift-tab, or escape).
	done func(tcell.Key)

	// A callback function set by the Form class and called when the user leaves
	// this form item.
	finished func(tcell.Key)

	sync.RWMutex
}

// NewCheckbox returns a new input field.
func NewCheckbox() *Checkbox {
	return &Checkbox{
		Box:            NewBox(),
		labelStyle:     Styles.CheckboxLabelStyle,
		uncheckedStyle: Styles.CheckboxUncheckedStyle,
		checkedStyle:   Styles.CheckboxCheckedStyle,
		focusStyle:     Styles.CheckboxFocusStyle,

		uncheckedString:       Styles.CheckboxUncheckedString,
		checkedString:         Styles.CheckboxCheckedString,
		cursorCheckedString:   Styles.CheckboxCursorCheckedString,
		cursorUncheckedString: Styles.CheckboxCursorUncheckedString,
	}
}

// SetChecked sets the state of the checkbox. This also triggers the "changed"
// callback if the state changes with this call.
func (c *Checkbox) SetChecked(checked bool) {
	c.Lock()
	defer c.Unlock()

	if c.checked != checked {
		if c.changed != nil {
			c.changed(checked)
		}
		c.checked = checked
	}
}

// IsChecked returns whether or not the box is checked.
func (c *Checkbox) IsChecked() bool {
	c.RLock()
	defer c.RUnlock()
	return c.checked
}

// SetLabel sets the text to be displayed before the input area.
func (c *Checkbox) SetLabel(label string) {
	c.Lock()
	defer c.Unlock()
	c.label = label
}

// GetLabel returns the text to be displayed before the input area.
func (c *Checkbox) GetLabel() string {
	c.RLock()
	defer c.RUnlock()
	return c.label
}

// SetLabel sets the text to be displayed before the input area.
func (c *Checkbox) SetLabelRight(labelRight string) {
	c.Lock()
	defer c.Unlock()
	c.labelRight = labelRight
}

// GetLabel returns the text to be displayed before the input area.
func (c *Checkbox) GetLabelRight() string {
	c.RLock()
	defer c.RUnlock()
	return c.labelRight
}

// SetLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func (c *Checkbox) SetLabelWidth(width int) {
	c.Lock()
	defer c.Unlock()
	c.labelWidth = width
}

// SetLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func (c *Checkbox) SetLabelRightWidth(width int) {
	c.Lock()
	defer c.Unlock()
	c.labelRightWidth = width
}

// SetLabelColor sets the color of the label.
func (c *Checkbox) SetLabelColor(color tcell.Color) {
	c.Lock()
	defer c.Unlock()
	c.labelStyle = c.labelStyle.Foreground(color)
}

// SetLabelStyle sets the style of the label.
func (c *Checkbox) SetLabelStyle(style tcell.Style) {
	c.Lock()
	defer c.Unlock()
	c.labelStyle = style
}

func (c *Checkbox) SetLabelFocusedColor(color tcell.Color) {
	c.Lock()
	defer c.Unlock()
	c.labelStyle = c.labelStyle.Foreground(color)
}

func (c *Checkbox) SetFieldTextFocusedColor(color tcell.Color) {
	c.Lock()
	defer c.Unlock()
	c.focusStyle = c.focusStyle.Foreground(color)
}

func (c *Checkbox) SetFieldBackgroundFocusedColor(color tcell.Color) {
	c.Lock()
	defer c.Unlock()
	c.focusStyle = c.focusStyle.Background(color)
}

// SetFieldBackgroundColor sets the background color of the input area.
func (c *Checkbox) SetFieldBackgroundColor(color tcell.Color) {
	c.Lock()
	defer c.Unlock()
	c.uncheckedStyle = c.uncheckedStyle.Background(color)
	c.checkedStyle = c.checkedStyle.Background(color)
	c.focusStyle = c.focusStyle.Foreground(color)
}

// SetFieldTextColor sets the text color of the input area.
func (c *Checkbox) SetFieldTextColor(color tcell.Color) {
	c.Lock()
	defer c.Unlock()
	c.uncheckedStyle = c.uncheckedStyle.Foreground(color)
	c.checkedStyle = c.checkedStyle.Foreground(color)
	c.focusStyle = c.focusStyle.Background(color)
}

// SetUncheckedStyle sets the style of the unchecked checkbox.
func (c *Checkbox) SetUncheckedStyle(style tcell.Style) {
	c.Lock()
	defer c.Unlock()
	c.uncheckedStyle = style
}

// SetCheckedStyle sets the style of the checked checkbox.
func (c *Checkbox) SetCheckedStyle(style tcell.Style) {
	c.Lock()
	defer c.Unlock()
	c.checkedStyle = style
}

// SetActivatedStyle sets the style of the checkbox when it is currently
// focused.
func (c *Checkbox) SetActivatedStyle(style tcell.Style) {
	c.Lock()
	defer c.Unlock()
	c.focusStyle = style
}

// SetCheckedString sets the string to be displayed when the checkbox is
// checked (defaults to "X"). The string may contain color tags (consider
// adapting the checkbox's various styles accordingly). See [Escape] in
// case you want to display square brackets.
func (c *Checkbox) SetCheckedString(checked string) {
	c.Lock()
	defer c.Unlock()
	c.checkedString = checked
}

// SetUncheckedString sets the string to be displayed when the checkbox is
// not checked (defaults to the empty space " "). The string may contain color
// tags (consider adapting the checkbox's various styles accordingly). See
// [Escape] in case you want to display square brackets.
func (c *Checkbox) SetUncheckedString(unchecked string) {
	c.Lock()
	defer c.Unlock()
	c.uncheckedString = unchecked
}

// SetFormAttributes sets attributes shared by all form items.
func (c *Checkbox) SetFormAttributes(labelWidth int, labelColor, bgColor, fieldTextColor, fieldBgColor tcell.Color) {
	c.Lock()
	defer c.Unlock()
	c.labelWidth = labelWidth
	c.SetLabelColor(labelColor)
	c.backgroundColor = bgColor
	c.SetFieldTextColor(fieldTextColor)
	c.SetFieldBackgroundColor(fieldBgColor)
}

// GetFieldWidth returns this primitive's field width.
func (c *Checkbox) GetFieldWidth() int {
	c.RLock()
	defer c.RUnlock()
	return 1
}

// GetFieldHeight returns this primitive's field height.
func (c *Checkbox) GetFieldHeight() int {
	c.RLock()
	defer c.RUnlock()
	return 1
}

// SetDisabled sets whether or not the item is disabled / read-only.
func (c *Checkbox) SetDisabled(disabled bool) {
	c.Lock()
	defer c.Unlock()
	c.disabled = disabled
	if c.finished != nil {
		c.finished(-1)
	}
}

// SetChangedFunc sets a handler which is called when the checked state of this
// checkbox was changed. The handler function receives the new state.
func (c *Checkbox) SetChangedFunc(handler func(checked bool)) {
	c.Lock()
	defer c.Unlock()
	c.changed = handler
}

// SetDoneFunc sets a handler which is called when the user is done using the
// checkbox. The callback function is provided with the key that was pressed,
// which is one of the following:
//
//   - KeyEscape: Abort text input.
//   - KeyTab: Move to the next field.
//   - KeyBacktab: Move to the previous field.
func (c *Checkbox) SetDoneFunc(handler func(key tcell.Key)) {
	c.Lock()
	defer c.Unlock()
	c.done = handler
}

// SetFinishedFunc sets a callback invoked when the user leaves this form item.
func (c *Checkbox) SetFinishedFunc(handler func(key tcell.Key)) {
	c.Lock()
	defer c.Unlock()
	c.finished = handler
}

// Focus is called when this primitive receives focus.
func (c *Checkbox) Focus(delegate func(p Primitive)) {
	c.Lock()
	defer c.Unlock()
	// If we're part of a form and this item is disabled, there's nothing the
	// user can do here so we're finished.
	if c.finished != nil && c.disabled {
		c.finished(-1)
		return
	}

	c.Box.Focus(delegate)
}

// Draw draws this primitive onto the screen.
func (c *Checkbox) Draw(screen tcell.Screen) {
	c.RLock()
	defer c.RUnlock()
	c.Box.Draw(screen)

	// Prepare
	x, y, width, height := c.GetInnerRect()
	rightLimit := x + width
	if height < 1 || rightLimit <= x {
		return
	}

	// Draw label.
	_, labelBg, _ := c.labelStyle.Decompose()
	if c.labelWidth > 0 {
		labelWidth := c.labelWidth
		if labelWidth > width {
			labelWidth = width
		}
		printWithStyle(screen, c.label, x, y, 0, labelWidth, AlignLeft, c.labelStyle, labelBg == tcell.ColorDefault)
		x += labelWidth
		width -= labelWidth
	} else {
		_, _, drawnWidth := printWithStyle(screen, c.label, x, y, 0, width, AlignLeft, c.labelStyle, labelBg == tcell.ColorDefault)
		x += drawnWidth
		width -= drawnWidth
	}

	// Draw checkbox.
	str := c.uncheckedString
	style := c.uncheckedStyle
	if c.disabled {
		style = style.Background(c.backgroundColor)
	}
	if c.checked {
		str = c.checkedString
		style = c.checkedStyle
	}
	if c.HasFocus() {
		style = c.focusStyle
		if c.checked {
			str = c.cursorCheckedString
		} else {
			str = c.cursorUncheckedString
		}
	}

	_, _, drawnWidth := printWithStyle(screen, str, x, y, 0, width, AlignLeft, style, c.disabled)
	x += drawnWidth
	width -= drawnWidth

	if c.labelRight != "" {
		// Draw label right.
		_, labelRightBg, _ := c.labelStyle.Decompose()
		if c.labelRightWidth > 0 {
			labelRightWidth := c.labelRightWidth
			if labelRightWidth > width {
				labelRightWidth = width
			}
			printWithStyle(screen, c.labelRight, x, y, 0, labelRightWidth, AlignLeft, c.labelStyle, labelRightBg == tcell.ColorDefault)
		} else {
			printWithStyle(screen, c.labelRight, x, y, 0, width, AlignLeft, c.labelStyle, labelRightBg == tcell.ColorDefault)
		}
	}
}

// InputHandler returns the handler for this primitive.
func (c *Checkbox) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return c.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Primitive)) {
		if c.disabled {
			return
		}

		// Process key event.
		switch key := event.Key(); key {
		case tcell.KeyRune, tcell.KeyEnter: // Check.
			if key == tcell.KeyRune && event.Rune() != ' ' {
				break
			}
			c.checked = !c.checked
			if c.changed != nil {
				c.changed(c.checked)
			}
		case tcell.KeyTab, tcell.KeyBacktab, tcell.KeyEscape: // We're done.
			if c.done != nil {
				c.done(key)
			}
			if c.finished != nil {
				c.finished(key)
			}
		}
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (c *Checkbox) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return c.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if c.disabled {
			return false, nil
		}

		x, y := event.Position()
		_, rectY, _, _ := c.GetInnerRect()
		if !c.InRect(x, y) {
			return false, nil
		}

		// Process mouse event.
		if y == rectY {
			if action == MouseLeftDown {
				setFocus(c)
				consumed = true
			} else if action == MouseLeftClick {
				c.checked = !c.checked
				if c.changed != nil {
					c.changed(c.checked)
				}
				consumed = true
			}
		}

		return
	})
}
