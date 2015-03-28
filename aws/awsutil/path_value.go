package awsutil

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var indexRe = regexp.MustCompile(`(.+)\[(-?\d+)?\]$`)

func rValuesAtPath(i interface{}, path string, create bool) []reflect.Value {
	values := []reflect.Value{reflect.Indirect(reflect.ValueOf(i))}
	components := strings.Split(path, ".")
	for len(values) > 0 && len(components) > 0 {
		var index *int64
		var indexStar bool
		c := components[0]
		if c == "" { // no actual component, illegal syntax
			return nil
		} else if c != "*" && strings.ToLower(c[0:1]) == c[0:1] {
			// TODO normalize case for user
			return nil // don't support unexported fields
		}

		// parse this component
		if m := indexRe.FindStringSubmatch(c); m != nil {
			c = m[1]
			if m[2] == "" {
				index = nil
				indexStar = true
			} else {
				i, _ := strconv.ParseInt(m[2], 10, 32)
				index = &i
				indexStar = false
			}
		}

		nextvals := []reflect.Value{}
		for _, value := range values {
			// pull component name out of struct member
			if value.Kind() != reflect.Struct {
				continue
			}

			if c == "*" { // pull all members
				for i := 0; i < value.NumField(); i++ {
					if f := reflect.Indirect(value.Field(i)); f.IsValid() {
						nextvals = append(nextvals, f)
					}
				}
				continue
			}

			value = value.FieldByName(c)
			if create && value.Kind() == reflect.Ptr && value.IsNil() {
				value.Set(reflect.New(value.Type().Elem()))
				value = value.Elem()
			} else {
				value = reflect.Indirect(value)
			}

			if value.IsValid() {
				nextvals = append(nextvals, value)
			}
		}
		values = nextvals

		if indexStar || index != nil {
			nextvals = []reflect.Value{}
			for _, value := range values {
				value := reflect.Indirect(value)
				if value.Kind() != reflect.Slice {
					continue
				}

				if indexStar { // grab all indices
					for i := 0; i < value.Len(); i++ {
						idx := reflect.Indirect(value.Index(i))
						if idx.IsValid() {
							nextvals = append(nextvals, idx)
						}
					}
					continue
				}

				// pull out index
				i := int(*index)
				if i >= value.Len() { // check out of bounds
					if create {
						// TODO resize slice
					} else {
						continue
					}
				} else if i < 0 { // support negative indexing
					i = value.Len() + i
				}
				value = reflect.Indirect(value.Index(i))

				if value.IsValid() {
					nextvals = append(nextvals, value)
				}
			}
			values = nextvals
		}

		components = components[1:]
	}
	return values
}

// ValuesAtPath returns a list of objects at the lexical path inside of a structure
func ValuesAtPath(i interface{}, path string) []interface{} {
	if rvals := rValuesAtPath(i, path, false); rvals != nil {
		vals := make([]interface{}, len(rvals))
		for i, rval := range rvals {
			vals[i] = rval.Interface()
		}
		return vals
	}
	return nil
}

// SetValueAtPath sets an object at the lexical path inside of a structure
func SetValueAtPath(i interface{}, path string, v interface{}) {
	if rvals := rValuesAtPath(i, path, true); rvals != nil {
		for _, rval := range rvals {
			rval.Set(reflect.ValueOf(v))
		}
	}
}
