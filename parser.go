package fstring

import (
	"fmt"
	"strconv"
	"strings"
)

type parser struct {
	data         []rune
	result       []rune
	idx          int
	values       map[string]any
	keyValidator KeyValidatorFunc
}

func newParser(s string, values map[string]any, keyValidator KeyValidatorFunc) *parser {
	if len(values) == 0 {
		values = map[string]any{}
	}
	return &parser{
		data:         []rune(s),
		result:       nil,
		idx:          0,
		values:       values,
		keyValidator: keyValidator,
	}
}

func (r *parser) parse() error {
	lastLeftCurlyBracketIdx := -1
	for ; r.hasMore(); r.idx++ {
		s := r.get()

		switch s {
		case '{':
			if lastLeftCurlyBracketIdx >= 0 {
				r.result = append(r.result, r.data[lastLeftCurlyBracketIdx:r.idx]...)
			}
			lastLeftCurlyBracketIdx = r.idx
		case '}':
			if lastLeftCurlyBracketIdx >= 0 {
				if lastLeftCurlyBracketIdx < r.idx {
					key := strings.TrimSpace(string(r.data[lastLeftCurlyBracketIdx+1 : r.idx]))
					if r.keyValidator == nil || r.keyValidator(string(key)) {
						val, ok := r.values[string(key)]
						if !ok {
							return fmt.Errorf("%w: %s", ErrArgsNotDefined, string(key))
						}
						r.result = append(r.result, []rune(toString(val))...)
						lastLeftCurlyBracketIdx = -1
						continue
					}
				}
				r.result = append(r.result, r.data[lastLeftCurlyBracketIdx:r.idx+1]...)
				lastLeftCurlyBracketIdx = -1
				continue
			}
			r.result = append(r.result, s)
		default:
			if lastLeftCurlyBracketIdx == -1 {
				r.result = append(r.result, s)
			}
		}
	}
	return nil
}

func (r *parser) hasMore() bool {
	return r.idx < len(r.data)
}

func (r *parser) get() rune {
	return r.data[r.idx]
}

// nolint: cyclop
func toString(val any) string {
	if val == nil {
		return "nil" // f'None' -> "None"
	}
	switch val := val.(type) {
	case string:
		return val
	case []rune:
		return string(val)
	case []byte:
		return string(val)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		return fmt.Sprintf("%s", val)
	}
}
