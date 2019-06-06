package insight_goclient

import (
	"bytes"
	"strconv"
)

type StringBool bool

func (c StringBool) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(strconv.Quote(strconv.FormatBool(bool(c))))
	return buf.Bytes(), nil
}

func (c *StringBool) UnmarshalJSON(in []byte) error {
	value := string(in)
	unquoted, err := strconv.Unquote(value)
	if err != nil {
		return err
	}
	var b bool
	b, err = strconv.ParseBool(unquoted)
	if err != nil {
		return err
	}
	res := StringBool(b)
	*c = res
	return nil
}
