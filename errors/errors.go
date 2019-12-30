package errors

import "net/url"

type InvalidRequestError *CustomError

// try to RFC7808
// Content-Type: application/problem+json
// しばらくはtypeはabout:blankで
// instanceはいらないかもだけどtry
type CustomError struct {
	Type     string   // URI to error detail page
	Title    string   // Error reason
	Instance *url.URL // Error location api uri
}

func NewInvalidRequestError(ins *url.URL) InvalidRequestError {
	return &CustomError{
		Type:     "about:blank",
		Title:    "Invalid request parameter",
		Instance: ins,
	}
}
