package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func lt100(v interface{}) error {
	i, ok := v.(int)
	if !ok {
		return fmt.Errorf("not an int: %v", v)
	}
	if i >= 100 {
		return fmt.Errorf("%d >= 100", i)
	}
	return nil
}

func testChecks(t *testing.T) {
	type unpacked struct {
		post int `http:p,check:"lt100"`
	}
	tests := []struct {
		req       *http.Request
		want      unpacked
		errSubstr string
	}{
		{
			&http.Request{Form: url.Values{"p": []string{"120"}}},
			unpacked{},
			">= 100",
		},
		{
			&http.Request{Form: url.Values{"p": []string{"80"}}},
			unpacked{80},
			"",
		},
	}
	checks := map[string]Check{
		"lt100": lt100,
	}
	for _, test := range tests {
		var got unpacked
		err := Unpack(test.req, &got, checks)
		if test.errSubstr != "" && !strings.Contains(err.Error(), test.errSubstr) {
			t.Errorf("Unpack(%v), error %q doesn't contain %q",
				test.req, err, test.errSubstr)
			continue
		}
		if err != nil {
			t.Errorf("Unpack(%v): %s", test.req, err)
		}
		if reflect.DeepEqual(test.want, got) {
			t.Errorf("Unpack(%v), got %v, want %v", test.req, got, test.want)
		}
	}
}
