package lab3

func TokenCheck(token string, tokens map[string]string) string {
	if _, exists := tokens[token]; exists {
		return tokens[token]
	}
	return ""
}
