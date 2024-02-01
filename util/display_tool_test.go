package util

import (
	"testing"
)

func TestFormatSize(t *testing.T) {
	if text := FormatSize(124); text != "124b" {
		t.Errorf("expect 124b , but%s", text)
	}
	if text := FormatSize(2048); text != "2k" {
		t.Errorf("expect 2k , but%s", text)
	}
	if text := FormatSize(1024 * 1024); text != "1m" {
		t.Errorf("expect 1m , but%s", text)
	}
	if text := FormatSize(1024 * 3324); text != "3m" {
		t.Errorf("expect 1m , but%s", text)
	}

}
