package fluentio

import "testing"

func TestFluentIO_Write(t *testing.T) {
	fluentIO, err := New(WithBasicConfig("127.0.0.1", 24224), WithTag("test"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = fluentIO.Write([]byte(`{"message": "hello world"}`))
	if err != nil {
		t.Fatal(err)
	}
}
