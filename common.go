package IPFSClient

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
)

func fieldToString(typ reflect.Type, val reflect.Value) string {
	log.Println(typ, val)
	switch typ.Kind() {
	case reflect.String:
		return val.String()
	case reflect.Float32:
	case reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'e', 10, 8)
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())
	case TCidVersion:
		return strconv.FormatInt(val.Int(), 10)
	default:
		return val.String()
	}
	return val.String()
}

func StructToHttpData(dat interface{}) (query, form string) {

	//writer := multipart.NewWriter(body)

	v := reflect.ValueOf(dat)

	st := reflect.TypeOf(dat)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		// use of Lookup method
		if val, ok := field.Tag.Lookup("query"); ok {
			if len(val) > 0 {
				strVal := fieldToString(field.Type, v.FieldByName(field.Name))
				query = fmt.Sprintf("%s&%s=%v", query, url.QueryEscape(val), url.QueryEscape(strVal))
			}
		}

		if val, ok := field.Tag.Lookup("form"); ok {
			strVal := fieldToString(field.Type, v.FieldByName(field.Name))
			//form[url.QueryEscape(val)] =url.QueryEscape(strVal) // v.FieldByName(field.Name).String()
			form = fmt.Sprintf("%s&%s=%s", form, url.QueryEscape(val), url.QueryEscape(strVal))
		}

		// if val, ok := field.Tag.Lookup("form-file-name"); ok {
		// 	form = fmt.Sprintf("%s=%s&", form, url.QueryEscape(val))
		// 	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
		// }
	}

	return
}

func StructToHttpDataMap(dat interface{}) (query, form map[string]string) {

	query = make(map[string]string)
	form = make(map[string]string)
	//writer := multipart.NewWriter(body)

	v := reflect.ValueOf(dat)

	st := reflect.TypeOf(dat)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		// use of Lookup method
		if val, ok := field.Tag.Lookup("query"); ok {
			if len(val) > 0 {
				//query = fmt.Sprintf("%s&%s=%v", query, url.QueryEscape(val), v.FieldByName(field.Name))
				//query[url.QueryEscape(val)] = v.FieldByName(field.Name).String()
				strVal := fieldToString(field.Type, v.FieldByName(field.Name))
				query[url.QueryEscape(val)] = url.QueryEscape(strVal)
			}
		}

		if val, ok := field.Tag.Lookup("form"); ok {
			//form = fmt.Sprintf("%s&%s=%v", form, url.QueryEscape(val), v.FieldByName(field.Name))
			strVal := fieldToString(field.Type, v.FieldByName(field.Name))
			form[url.QueryEscape(val)] = url.QueryEscape(strVal) // v.FieldByName(field.Name).String()
		}

		// if val, ok := field.Tag.Lookup("form-file-name"); ok {
		// 	form = fmt.Sprintf("%s=%s&", form, url.QueryEscape(val))
		// 	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
		// }
	}

	return
}
