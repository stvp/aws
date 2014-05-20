package aws

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestDoError(t *testing.T) {
	v, err := DescribeInstances()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", v)
}

func TestUnmarshalError(t *testing.T) {
	body := []byte(`
		<?xml version="1.0" encoding="UTF-8"?>
		<Response>
			<Errors>
				<Error>
					<Code>AuthFailure</Code>
					<Message>AWS was not able to validate the provided access credentials</Message>
				</Error>
			</Errors>
			<RequestID>afc00dc9-0c19-46db-a987-f7de2a12a361</RequestID>
		</Response>
	`)

	type Error struct {
		Code    string
		Message string
	}

	type Response struct {
		RequestId string  `xml:"RequestID"`
		Errors    []Error `xml:"Errors>Error"`
	}

	got := new(Response)
	err := xml.Unmarshal(body, got)
	if err != nil {
		t.Fatal(err)
	}

	exp := &Response{
		RequestId: "afc00dc9-0c19-46db-a987-f7de2a12a361",
		Errors: []Error{
			{Code: "AuthFailure", Message: "AWS was not able to validate the provided access credentials"},
		},
	}

	if fmt.Sprintf("%#v", exp) != fmt.Sprintf("%#v", got) {
		t.Fatalf("Expected %#v, but got %#v", exp, got)
	}
}

func TestParamEncode(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected string
	}{
		{"foo", "bar", "foo=bar"},
		{"foo.1", "spaces here", "baz.1=spaces%20here"},
		{"foo", "punctuation's", "foo=punctuation%27s"},
	}

	for i, test := range tests {
		param := Param{Key: test.key, Val: test.value}
		got := param.Encode()
		if got != test.expected {
			t.Errorf("test[%d]: expected %s, got %s", i, test.expected, got)
		}
	}
}
