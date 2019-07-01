package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User is a thing
type User struct {
	_ struct{} `collection:"users"`

	Name  string `bson:"name"`
	Nick  string `bson:"nick,omitempty"`
	Pass  string `bson:"-"`
	Thing string `bson:",omitempty"`
	Age   int    `bson:"age"`
}

func main() {
	u := User{Name: "Jacob", Nick: "Jake", Pass: "banana", Age: 32}

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	fmt.Println("Put stuff in this collection:", t.Field(0).Tag.Get("collection"))

	fields := make(bson.M)

	for i := 1; i < t.NumField(); i++ {
		field := t.Field(i)

		opts := strings.Split(field.Tag.Get("bson"), ",")

		var name string
		if len(opts) == 0 || strings.TrimSpace(opts[0]) == "" {
			name = strings.ToLower(field.Name)
		} else {
			name = opts[0]
		}

		// Do not include values whose name is just "-"
		if name == "-" {
			continue
		}

		var omitempty, inline, minsize bool
		for _, opt := range opts[1:] {
			switch opt {
			case "omitempty":
				omitempty = true
			case "inline":
				inline = true
			case "minsize":
				minsize = true
			}
		}
		_, _ = inline, minsize

		val := v.Field(i)

		if omitempty && isZero(val) {
			continue
		}

		fields[name] = val
	}

	fmt.Println(fields)
}

var (
	typeTime = reflect.TypeOf(time.Time{})
)

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct:
		if v.Type() == typeTime {
			return v.Interface().(time.Time).IsZero()
		}
		for i := v.NumField() - 1; i >= 0; i-- {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}

// MapFor creates a bson.M out of a given struct or map. Right now it does this
// by marshalling the value to bytes then unmarshalling those bytes to a
// bson.M. In the future it may be improved to go directly from the initial
// value to a bson.M.
func MapFor(val interface{}) (bson.M, error) {
	b, err := bson.Marshal(val)
	if err != nil {
		return nil, err
	}

	m := make(bson.M)
	if err := bson.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return m, nil
}
