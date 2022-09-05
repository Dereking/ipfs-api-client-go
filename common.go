package IPFSClient

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type ErrMessage struct {
	//{"Message":"invalid path \"/\": invalid ipfs path","Code":0,"Type":"error"}
	Message string
	Code    int
	Type    string
}

func (e ErrMessage) String() string {
	return fmt.Sprintf(`ErrMessage: code=%d, msg=%s, type=%s`, e.Code, e.Message, e.Type)
}

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
	q, f, err := StructToHttpDataMap(dat, false)
	if err != nil {
		return "", "", err
	}

	// fetch all keys for q
	var queryKeys []string
	for k := range q {
		queryKeys = append(queryKeys, k)
	}
	sort.Strings(queryKeys)
	for _, key := range queryKeys {
		vals := q[key]
		for _, v := range vals {
			query = fmt.Sprintf("%s&%s=%v", query, key, v)
		}
	}

	// fetch all keys for f
	var formKeys []string
	for k := range q {
		formKeys = append(formKeys, k)
	}
	sort.Strings(formKeys)
	for _, key := range formKeys {
		vals := f[key]
		for _, v := range vals {
			form = fmt.Sprintf("%s&%s=%s", form, key, v)
		}
	}
	return
}

func StructToHttpDataMap(dat interface{}, includeEmptyStr bool) (query, form map[string][]string, err error) {

	query = make(map[string][]string)
	form = make(map[string][]string)

	v := reflect.ValueOf(dat)
	st := reflect.TypeOf(dat)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		// 检查是否 required。如果是 true
		required := false
		strReq, exists := field.Tag.Lookup("required")
		if exists && strings.Compare(strReq, "true") == 0 {
			log.Println(field.Name + " is required")
			required = true
		}

		// use of Lookup method
		if key, ok := field.Tag.Lookup("query"); ok {
			if len(key) > 0 {
				key = url.QueryEscape(key)

				if _, ok := query[key]; ok {

				} else {
					query[key] = make([]string, 0)
				}

				strVal := fieldToString(field.Type, v.FieldByIndex([]int{i})) //v.FieldByName(field.Name))
				if required && len(strVal) == 0 {
					err = errors.New(fmt.Sprintf("'%s (%s)' is required, but missed.", field.Name, key))
					return
				}
				if len(strVal) == 0 {
					if includeEmptyStr {
						//query = fmt.Sprintf("%s&%s=%v", query, url.QueryEscape(key), url.QueryEscape(strVal))
						query[key] = append(query[key], url.QueryEscape(strVal))
					}
				} else {
					//query = fmt.Sprintf("%s&%s=%v", query, url.QueryEscape(key), url.QueryEscape(strVal))
					query[key] = append(query[key], url.QueryEscape(strVal))
				}
			} else {
				err = errors.New(fmt.Sprintf("'%s' haven't a valid tag 'query'.", key))
				return
			}
		}

		if key, ok := field.Tag.Lookup("form"); ok {
			if len(key) > 0 {
				key = url.QueryEscape(key)

				if _, ok := form[key]; ok {

				} else {
					form[key] = make([]string, 0)
				}
				strVal := fieldToString(field.Type, v.FieldByIndex([]int{i})) //v.FieldByName(field.Name))

				form[key] = append(form[key], url.QueryEscape(strVal))
			} else {
				err = errors.New(fmt.Sprintf("'%s' haven't a valid tag 'form'.", key))
				return
			}
		}
	}
	return
}
