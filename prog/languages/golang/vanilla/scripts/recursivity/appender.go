package main

func fetch(token string) ([]any, string) {
	switch token {
	case "":
		return []any{1, 2, 3}, "1"
	case "1":
		return []any{4, 5, 6}, "2"
	case "2":
		return []any{7, 8, 9}, "eof"
	}
	return nil, ""
}

func appender(token string) []any {
	if token == "eof" {
		return nil
	}
	response, token := fetch(token)
	return append(response, appender(token)...)
}
