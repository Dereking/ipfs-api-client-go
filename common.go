package IPFSClient

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func fieldToString(typ reflect.Type, val reflect.Value) string {
	//log.Println(typ, val)
	kind := typ.Kind()

	switch kind {
	case reflect.String:
		return val.String()
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'e', 10, 64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())
	default:
		return val.String()
	}
}

// func StructToHttpData2(dat interface{}) (query, form string, err error) {

// 	//writer := multipart.NewWriter(body)

// 	v := reflect.ValueOf(dat)
// 	st := reflect.TypeOf(dat)
// 	for i := 0; i < st.NumField(); i++ {
// 		field := st.Field(i)

// 		// 检查是否 required。如果是 true
// 		required := false
// 		strReq, exists := field.Tag.Lookup("required")
// 		if exists && strings.Compare(strReq, "true") == 0 {
// 			required = true
// 		}

// 		// use of Lookup method
// 		if key, ok := field.Tag.Lookup("query"); ok {
// 			if len(key) > 0 {
// 				strVal := fieldToString(field.Type, v.FieldByIndex([]int{i})) //v.FieldByName(field.Name))
// 				if required && len(strVal) == 0 {
// 					err = errors.New(fmt.Sprintf("'%s' is required, but missed.", key))
// 					return
// 				}
// 				query = fmt.Sprintf("%s&%s=%v", query, url.QueryEscape(key), url.QueryEscape(strVal))
// 			} else {
// 				err = errors.New(fmt.Sprintf("'%s' haven't a valid tag 'query'.", key))
// 				return
// 			}
// 		}

// 		if key, ok := field.Tag.Lookup("form"); ok {
// 			if len(key) > 0 {
// 				strVal := fieldToString(field.Type, v.FieldByIndex([]int{i})) //v.FieldByName(field.Name))
// 				//form[url.QueryEscape(val)] =url.QueryEscape(strVal) // v.FieldByName(field.Name).String()
// 				form = fmt.Sprintf("%s&%s=%s", form, url.QueryEscape(key), url.QueryEscape(strVal))
// 			} else {
// 				err = errors.New(fmt.Sprintf("'%s' haven't a valid tag 'form'.", key))
// 				return
// 			}
// 		}
// 	}

// 	return
// }

func StructToHttpData(dat interface{}) (query, form string, err error) {
	q, f, err := StructToHttpDataMap(dat)
	if err != nil {
		return "", "", err
	}

	var keys []string
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		query = fmt.Sprintf("%s&%s=%v", query, k, q[k])
	}

	var keyForms []string
	for k := range f {
		keyForms = append(keyForms, k)
	}
	sort.Strings(keyForms)
	for _, k := range f {
		form = fmt.Sprintf("%s&%s=%s", form, k, f[k])
	}
	return
}

func StructToHttpDataMap(dat interface{}) (query, form map[string]string, err error) {

	query = make(map[string]string)
	form = make(map[string]string)

	v := reflect.ValueOf(dat)
	st := reflect.TypeOf(dat)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		// 检查是否 required。如果是 true
		required := false
		strReq, exists := field.Tag.Lookup("required")
		if exists && strings.Compare(strReq, "true") == 0 {
			required = true
		}

		// use of Lookup method
		if key, ok := field.Tag.Lookup("query"); ok {
			if len(key) > 0 {
				strVal := fieldToString(field.Type, v.FieldByIndex([]int{i})) //v.FieldByName(field.Name))
				if required && len(strVal) == 0 {
					err = errors.New(fmt.Sprintf("'%s' is required, but missed.", key))
					return
				}
				//query = fmt.Sprintf("%s&%s=%v", query, url.QueryEscape(key), url.QueryEscape(strVal))
				query[url.QueryEscape(key)] = url.QueryEscape(strVal)
			} else {
				err = errors.New(fmt.Sprintf("'%s' haven't a valid tag 'query'.", key))
				return
			}
		}

		if key, ok := field.Tag.Lookup("form"); ok {
			if len(key) > 0 {
				strVal := fieldToString(field.Type, v.FieldByIndex([]int{i})) //v.FieldByName(field.Name))
				form[url.QueryEscape(key)] = url.QueryEscape(strVal)
				//form = fmt.Sprintf("%s&%s=%s", form, url.QueryEscape(key), url.QueryEscape(strVal))
			} else {
				err = errors.New(fmt.Sprintf("'%s' haven't a valid tag 'form'.", key))
				return
			}
		}
	}
	return
}
