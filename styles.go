package nuview

import "github.com/gdamore/tcell/v2"

// Theme defines the colors used when primitives are initialized.
type Theme struct {
	// Title, border and other lines
	TitleColor    tcell.Color // Box titles.
	BorderColor   tcell.Color // Box borders.
	GraphicsColor tcell.Color // Graphics.

	// Text
	PrimaryTextColor           tcell.Color // Primary text.
	SecondaryTextColor         tcell.Color // Secondary text (e.g. labels).
	TertiaryTextColor          tcell.Color // Tertiary text (e.g. subtitles, notes).
	InverseTextColor           tcell.Color // Text on primary-colored backgrounds.
	ContrastPrimaryTextColor   tcell.Color // Primary text for contrasting elements.
	ContrastSecondaryTextColor tcell.Color // Secondary text on ContrastBackgroundColor-colored backgrounds.

	// Background
	PrimitiveBackgroundColor    tcell.Color // Main background color for primitives.
	ContrastBackgroundColor     tcell.Color // Background color for contrasting elements.
	MoreContrastBackgroundColor tcell.Color // Background color for even more contrasting elements.

	// Button
	ButtonCursorRune              rune // The symbol to draw at the end of button labels when focused.
	ButtonLabelColor              tcell.Color
	ButtonLabelFocusedColor       tcell.Color
	ButtonBackgroundColor         tcell.Color
	ButtonBackgroundFocusedColor  tcell.Color
	ButtonBackgroundDisabledColor tcell.Color
	ButtonLabelDisabledColor      tcell.Color

	// Check box
	CheckboxLabelStyle            tcell.Style
	CheckboxUncheckedStyle        tcell.Style
	CheckboxCheckedStyle          tcell.Style
	CheckboxFocusStyle            tcell.Style
	CheckboxCheckedString         string
	CheckboxUncheckedString       string
	CheckboxCursorCheckedString   string
	CheckboxCursorUncheckedString string

	// Input field
	InputFieldLabelColor                              tcell.Color
	InputFieldFieldBackgroundColor                    tcell.Color
	InputFieldFieldBackgroundFocusedColor             tcell.Color
	InputFieldFieldTextColor                          tcell.Color
	InputFieldFieldTextFocusedColor                   tcell.Color
	InputFieldPlaceholderTextColor                    tcell.Color
	InputFieldAutocompleteListTextColor               tcell.Color
	InputFieldAutocompleteListBackgroundColor         tcell.Color
	InputFieldAutocompleteListSelectedTextColor       tcell.Color
	InputFieldAutocompleteListSelectedBackgroundColor tcell.Color
	InputFieldAutocompleteSuggestionTextColor         tcell.Color
	InputFieldFieldNoteTextColor                      tcell.Color
	InputFieldLabelFocusedColor                       tcell.Color
	InputFieldPlaceholderTextFocusedColor             tcell.Color

	ListMainTextColor           tcell.Color
	ListSecondaryTextColor      tcell.Color
	ListShortcutColor           tcell.Color
	ListSelectedTextColor       tcell.Color
	ListScrollBarColor          tcell.Color
	ListSelectedBackgroundColor tcell.Color

	// Context menu
	ContextMenuPaddingTop    int
	ContextMenuPaddingBottom int
	ContextMenuPaddingLeft   int
	ContextMenuPaddingRight  int

	// Drop down
	DropDownAbbreviationChars string // The chars to show when the option's text gets shortened.
	DropDownSymbol            rune   // The symbol to draw at the end of the field when closed.
	DropDownOpenSymbol        rune   // The symbol to draw at the end of the field when opened.
	DropDownSelectedSymbol    rune   // The symbol to draw to indicate the selected list item.

	// Scroll bar
	ScrollBarColor tcell.Color

	// Window
	WindowMinWidth  int
	WindowMinHeight int
}

// Styles defines the appearance of an application. The default is for a black
// background and some basic colors: black, white, yellow, green, cyan, and
// blue.
var Styles = Theme{
	TitleColor:    tcell.ColorWhite.TrueColor(),
	BorderColor:   tcell.ColorWhite.TrueColor(),
	GraphicsColor: tcell.ColorWhite.TrueColor(),

	PrimaryTextColor:           tcell.ColorWhite.TrueColor(),
	SecondaryTextColor:         tcell.ColorYellow.TrueColor(),
	TertiaryTextColor:          tcell.ColorLimeGreen.TrueColor(),
	InverseTextColor:           tcell.ColorBlack.TrueColor(),
	ContrastPrimaryTextColor:   tcell.ColorBlack.TrueColor(),
	ContrastSecondaryTextColor: tcell.ColorLightSlateGray.TrueColor(),

	PrimitiveBackgroundColor:    tcell.ColorBlack.TrueColor(),
	ContrastBackgroundColor:     tcell.ColorGreen.TrueColor(),
	MoreContrastBackgroundColor: tcell.ColorDarkGreen.TrueColor(),

	ButtonCursorRune:              '◀',
	ButtonLabelColor:              tcell.ColorWhite.TrueColor(),
	ButtonLabelFocusedColor:       tcell.ColorWhite.TrueColor(),
	ButtonBackgroundColor:         tcell.ColorDarkGreen.TrueColor(),
	ButtonBackgroundFocusedColor:  tcell.ColorGreen.TrueColor(),
	ButtonBackgroundDisabledColor: tcell.ColorDarkGray.TrueColor(),
	ButtonLabelDisabledColor:      tcell.ColorBlack.TrueColor(),

	CheckboxLabelStyle:            tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()),
	CheckboxUncheckedStyle:        tcell.StyleDefault.Background(tcell.ColorGreen.TrueColor()).Foreground(tcell.ColorWhite.TrueColor()),
	CheckboxCheckedStyle:          tcell.StyleDefault.Background(tcell.ColorGreen.TrueColor()).Foreground(tcell.ColorWhite.TrueColor()),
	CheckboxFocusStyle:            tcell.StyleDefault.Background(tcell.ColorWhite.TrueColor()).Foreground(tcell.ColorGreen.TrueColor()),
	CheckboxCheckedString:         "[X]",
	CheckboxUncheckedString:       "[ ]",
	CheckboxCursorCheckedString:   ">X<",
	CheckboxCursorUncheckedString: "> <",

	InputFieldLabelColor:                              tcell.ColorYellow.TrueColor(),
	InputFieldFieldBackgroundColor:                    tcell.ColorDarkGreen.TrueColor(),
	InputFieldFieldBackgroundFocusedColor:             tcell.ColorGreen.TrueColor(),
	InputFieldFieldTextColor:                          tcell.ColorWhite.TrueColor(),
	InputFieldFieldTextFocusedColor:                   tcell.ColorWhite.TrueColor(),
	InputFieldPlaceholderTextColor:                    tcell.ColorLightSlateGray.TrueColor(),
	InputFieldAutocompleteListTextColor:               tcell.ColorBlack.TrueColor(),
	InputFieldAutocompleteListBackgroundColor:         tcell.ColorDarkGreen.TrueColor(),
	InputFieldAutocompleteListSelectedTextColor:       tcell.ColorBlack.TrueColor(),
	InputFieldAutocompleteListSelectedBackgroundColor: tcell.ColorWhite.TrueColor(),
	InputFieldAutocompleteSuggestionTextColor:         tcell.ColorLightSlateGray.TrueColor(),
	InputFieldFieldNoteTextColor:                      tcell.ColorYellow.TrueColor(),
	InputFieldLabelFocusedColor:                       ColorUnset,
	InputFieldPlaceholderTextFocusedColor:             ColorUnset,

	ListMainTextColor:           tcell.ColorWhite.TrueColor(),
	ListSecondaryTextColor:      tcell.ColorLimeGreen.TrueColor(),
	ListShortcutColor:           tcell.ColorYellow.TrueColor(),
	ListSelectedTextColor:       tcell.ColorBlack.TrueColor(),
	ListScrollBarColor:          tcell.ColorWhite.TrueColor(),
	ListSelectedBackgroundColor: tcell.ColorWhite.TrueColor(),

	ContextMenuPaddingTop:    0,
	ContextMenuPaddingBottom: 0,
	ContextMenuPaddingLeft:   1,
	ContextMenuPaddingRight:  1,

	DropDownAbbreviationChars: "...",
	DropDownSymbol:            '◀',
	DropDownOpenSymbol:        '▼',
	DropDownSelectedSymbol:    '▶',

	ScrollBarColor: tcell.ColorWhite.TrueColor(),

	WindowMinWidth:  4,
	WindowMinHeight: 3,
}
