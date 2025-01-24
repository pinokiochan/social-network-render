package utils

import (
	"regexp"
	"unicode"
	"strings"
)

func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsValidEmail(email string) bool {
    // Регулярное выражение для базовой проверки email
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$`)
    if !re.MatchString(email) {
        return false
    }

    // Проверка на недопустимые символы (например, ".." в домене)
    if strings.Contains(email, "..") {
        return false
    }

    // Проверка на длину TLD (ограничим до 6 символов для большинства случаев)
    parts := strings.Split(email, ".")
    if len(parts) < 2 { // Если нет части TLD
        return false
    }

    tld := parts[len(parts)-1]
    if len(tld) > 6 {
        return false
    }

    return true
}
