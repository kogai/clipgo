package main

import (
	"testing"
)

func TestTrimExtension(t *testing.T) {
	trimed := trimExtension("foo.ext")
	if trimed != "foo" {
		t.Error("拡張子が削除できていない")
		t.Fail()
	}
}
