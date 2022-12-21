package encode

import "golang.org/x/text/encoding/charmap"

// ISO8859_1 converts strings from UTF-8 to ISO 8859-1 (Latin-1).
// Facebook incorrectly exports data to Latin-1 instead of UTF-8,
// causing problems with non-English characters, emojis, etc.
func ISO8859_1(s *string) {
	ret, err := charmap.ISO8859_1.NewEncoder().String(*s)
	if err == nil {
		*s = ret
	}
}
