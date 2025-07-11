package colors

const (
	reset = "\033[0m"

	// Text colors
	primaryColor   = "\033[32m"
	secondaryColor = "\033[34m"
	redColor       = "\033[31m"

	// Background colors
	bgPrimaryColor   = "\033[42m"
	bgSecondaryColor = "\033[44m"
	bgRedColor       = "\033[41m"

	// Other Styles
	boldCode   = "\033[1m"
	italicCode = "\033[3m"
)

// Normal colors
func Primary(s string) string {
	return primaryColor + s + reset
}

func Secondary(s string) string {
	return secondaryColor + s + reset
}

func Red(s string) string {
	return redColor + s + reset
}

// Bold colors
func PrimaryBold(s string) string {
	return boldCode + primaryColor + s + reset
}

func SecondaryBold(s string) string {
	return boldCode + secondaryColor + s + reset
}

func RedBold(s string) string {
	return boldCode + redColor + s + reset
}

// Italic colors
func PrimaryItalic(s string) string {
	return italicCode + primaryColor + s + reset
}

func SecondaryItalic(s string) string {
	return italicCode + secondaryColor + s + reset
}

func RedItalic(s string) string {
	return italicCode + redColor + s + reset
}

// Bold + Italic colors
func PrimaryBoldItalic(s string) string {
	return boldCode + italicCode + primaryColor + s + reset
}

func SecondaryBoldItalic(s string) string {
	return boldCode + italicCode + secondaryColor + s + reset
}

func RedBoldItalic(s string) string {
	return boldCode + italicCode + redColor + s + reset
}

// Highlighted background
func PrimaryBG(s string) string {
	return bgPrimaryColor + s + reset
}

func SecondaryBG(s string) string {
	return bgSecondaryColor + s + reset
}

func RedBG(s string) string {
	return bgRedColor + s + reset
}

// Bold + Highlighted
func PrimaryBoldBG(s string) string {
	return boldCode + bgPrimaryColor + s + reset
}

func SecondaryBoldBG(s string) string {
	return boldCode + bgSecondaryColor + s + reset
}

func RedBoldBG(s string) string {
	return boldCode + bgRedColor + s + reset
}

// Italic + Highlighted
func PrimaryItalicBG(s string) string {
	return italicCode + bgPrimaryColor + s + reset
}

func SecondaryItalicBG(s string) string {
	return italicCode + bgSecondaryColor + s + reset
}

func RedItalicBG(s string) string {
	return italicCode + bgRedColor + s + reset
}

// Bold + Italic + Highlighted
func PrimaryBoldItalicBG(s string) string {
	return boldCode + italicCode + bgPrimaryColor + s + reset
}

func SecondaryBoldItalicBG(s string) string {
	return boldCode + italicCode + bgSecondaryColor + s + reset
}

func RedBoldItalicBG(s string) string {
	return boldCode + italicCode + bgRedColor + s + reset
}

// Bold text only
func Bold(s string) string {
	return boldCode + s + reset
}

// Italic text only
func Italic(s string) string {
	return italicCode + s + reset
}

// Bold + Italic text only
func BoldItalic(s string) string {
	return boldCode + italicCode + s + reset
}
