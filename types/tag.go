package types

import (
	"fmt"
	"strings"
)

type Tag []string

type TagType []string

// Append .
func (t Tag) Append(k string, v interface{}) {
	if len(t) == 0 {
		t = make([]string, 0)
	}

	switch v.(type) {
	case string:
		t = append(t, fmt.Sprintf(`%s:"%s"`, k, v))
	case TagType:
		var tt = v.(TagType)
		t = append(t, fmt.Sprintf(`%s:"%s"`, k, tt.String()))
	}
}

// Append .
func (tt TagType) Append(k string, v ...string) {
	if len(tt) == 0 {
		tt = make([]string, 0)
	}

	switch len(v) {
	case 0:
		tt = append(tt, k)
	default:
		tt = append(tt, fmt.Sprintf("%s:%s", k, v[0]))
	}
}

// String .
func (tt TagType) String() string {
	return strings.Join(tt, ";")
}

// String .
func (t Tag) String() string {
	return fmt.Sprintf("`%s`", strings.Join(t, " "))
}
