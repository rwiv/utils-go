package web

import (
	"fmt"
)

func JoinUrl(endpoint string, params string) string {
	return fmt.Sprintf(
		"%s%s",
		endpoint,
		paramsWithDelimiter(params),
	)
}

func JoinUrlWithQs(endpoint string, params string, queryString string) string {
	return fmt.Sprintf(
		"%s%s%s",
		endpoint,
		paramsWithDelimiter(params),
		queryStringWithDelimiter(queryString),
	)
}

func paramsWithDelimiter(params string) string {
	if params[0] == '/' {
		return params
	} else {
		return "/" + params
	}
}

func queryStringWithDelimiter(queryString string) string {
	if queryString[0] == '?' {
		return queryString
	} else {
		return "?" + queryString
	}
}
