package util_str

import (
	"database/sql"
	"github.com/deckarep/golang-set"
	"strconv"
	"strings"
)

func NewString(str string) *string {
	return &str
}

func SqlStringOrEmpty(str sql.NullString) string {
	if str.Valid {
		return str.String
	} else {
		return ""
	}
}
func EmptyOrDefault(ori, def string) string {
	if ori == "" {
		return def
	} else {
		return ori
	}
}

func MustToString(str interface{}) string {
	value, ok := str.(string)
	if !ok {
		panic("convert interface{} to string err")
	}
	return value
}

func JoinWith(set mapset.Set, sep string) string {
	a := set.ToSlice()
	switch len(a) {
	case 0:
		return ""
	case 1:
		return MustToString(a[0])
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(MustToString(a[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(MustToString(a[0]))
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(MustToString(s))
	}
	return b.String()
}

func GetInt64(ori string) (int64, error) {
	return strconv.ParseInt(ori, 10, 64)
}
func GetFloat64(ori string) (float64, error) {
	return strconv.ParseFloat(ori, 64)
}

func ToIntOrDefault(str string, def int) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return def
	}
	return i
}
func ToIntOrNil(str string) *int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return nil
	}
	return &i
}

func ToBoolOrDefault(str string, def bool) bool {
	b, e := strconv.ParseBool(str)
	if e != nil {
		return def
	}
	return b
}

func ToFloatOrDefault(str string, def float64) float64 {
	f, e := strconv.ParseFloat(str, 64)
	if e != nil {
		return def
	}
	return f
}

func ToBoolOrNil(str string) *bool {
	b, e := strconv.ParseBool(str)
	if e != nil {
		return nil
	}
	return &b
}

func StrOrDefault(str *string, def string) string {
	if str == nil {
		return def
	} else {
		return *str
	}
}
