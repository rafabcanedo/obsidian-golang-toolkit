package binding

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func JSON(r *http.Request, dst any) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return fmt.Errorf("invalid body: %w", err)
	}
	return nil
}

func Query(r *http.Request, dst any) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("dst must be a pointer to a struct")
	}

	v = v.Elem()
	t := v.Type()
	params := r.URL.Query()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("form")
		if tag == "" {
			continue
		}

		raw := params.Get(tag)
		if raw == "" {
			continue
		}

		if err := setField(v.Field(i), raw); err != nil {
			return fmt.Errorf("param %q: %w", tag, err)
		}
	}

	return nil
}

func setField(fv reflect.Value, raw string) error {
	switch fv.Kind() {
	case reflect.String:
		fv.SetString(raw)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return fmt.Errorf("expected integer, got %q", raw)
		}
		fv.SetInt(n)

	case reflect.Bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return fmt.Errorf("expected bool, got %q", raw)
		}
		fv.SetBool(b)

	default:
		return fmt.Errorf("unsupported type: %s", fv.Kind())
	}
	return nil
}
