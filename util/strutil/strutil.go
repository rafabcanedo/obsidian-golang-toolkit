package strutil

import (
	"strings"
	"unicode"
)

var accentReplacer = strings.NewReplacer(
	"á", "a", "à", "a", "ã", "a", "â", "a", "ä", "a",
	"é", "e", "è", "e", "ê", "e", "ë", "e",
	"í", "i", "ì", "i", "î", "i", "ï", "i",
	"ó", "o", "ò", "o", "õ", "o", "ô", "o", "ö", "o",
	"ú", "u", "ù", "u", "û", "u", "ü", "u",
	"ç", "c", "ñ", "n",
)

func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func Truncate(s string, max int) string {
	runes := []rune(s)

	if len(runes) <= max {
		return s
	}

	return string(runes[:max]) + "..."
}

func CamelToSnake(s string) string {
	var b strings.Builder

	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			b.WriteRune('_')
		}
		b.WriteRune(unicode.ToLower(r))
	}

	return b.String()
}

func Slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = accentReplacer.Replace(s)

	var b strings.Builder
	prevHyphen := false
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevHyphen = false
		} else if !prevHyphen {
			b.WriteRune('-')
			prevHyphen = true
		}
	}

	return strings.Trim(b.String(), "-")
}
