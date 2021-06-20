package ge

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Number float64

// CASTING FUNCTIONS

func (n Number) Float64() float64 {
	return float64(n)
}

func (n Number) Int() int {
	return int(n)
}

func (n Number) String() string {

	s := fmt.Sprintf("%f", n)
	s = strings.TrimRightFunc(s, func(r rune) bool {
		return r == '0'
	})
	return s
}

func (n Number) Bool() bool {
	return n != 0.0
}

// UnmarshalJSON implements Unmarshaler interface
func (n *Number) UnmarshalJSON(bs []byte) error {
	var value interface{}
	if err := json.Unmarshal(bs, &value); err != nil {
		return err
	}
	switch value.(type) {
	case float64:
		*n = Number(value.(float64))

	case string:
		v := value.(string)
		r, _ := regexp.Match("^\\d*\\.?(?:\\d*)?$", []byte(v))
		if !r || v == "" {
			return errors.New("value " + value.(string) + " is not numeric")
		}
		f, _ := strconv.ParseFloat(v, 64)
		*n = Number(f)
	case float32:
		*n = Number(value.(float32))
	case bool:
		if value.(bool) {
			*n = 1.
		} else {
			*n = 0.
		}
	default:
		return errors.New("invalid type")
	}

	return nil
}

// MarshalJSON implements Marshaler interface
func (n Number) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Float64())
}
