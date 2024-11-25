package fstring

import (
	"errors"
)

var (
	ErrEmptyExpression       = errors.New("empty expression not allowed")
	ErrArgsNotDefined        = errors.New("args not defined")
	ErrLeftBracketNotClosed  = errors.New("single '{' is not allowed")
	ErrRightBracketNotClosed = errors.New("single '}' is not allowed")
)

// Format interpolates the given template with the given values by using
// f-string.
func Format(template string, values map[string]any, opts ...FormatOption) (string, error) {
	var options formatOptions

	for _, opt := range opts {
		opt(&options)
	}

	p := newParser(template, values, options.keyValidator)
	if err := p.parse(); err != nil {
		return "", err
	}
	return string(p.result), nil
}

type KeyValidatorFunc func(key string) bool

type formatOptions struct {
	keyValidator KeyValidatorFunc
}

type FormatOption func(o *formatOptions)

func WithKeyValidator(fn KeyValidatorFunc) FormatOption {
	return func(o *formatOptions) {
		o.keyValidator = fn
	}
}

func BasicKeyValidator(key string) bool {
	for _, r := range key {
		if !(('0' <= r && r <= '9') ||
			('A' <= r && r <= 'Z') ||
			('a' <= r && r <= 'z') ||
			r == '_' ||
			r == '-') {
			return false
		}
	}
	return true
}
