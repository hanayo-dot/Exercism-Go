package techpalace

import "strings"

func WelcomeMessage(name string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(name)
}

func AddBorder(message string, stars int) string {
	border := strings.Repeat("*", stars)
	return border + "\n" + message + "\n" + border
}

func CleanupMessage(message string) string {
	message = strings.ReplaceAll(message, "*", "")
	return strings.TrimSpace(message)
}
