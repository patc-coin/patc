// Copyright 2018 The go-pc Authors
// This file is part of the go-pc library.
//
// The go-pc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-pc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-pc library. If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"testing"
)

func TestURLParsing(t *testing.T) {
	url, err := parseURL("https://pc.dp.tc")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "pc.dp.tc" {
		t.Errorf("expected: %v, got: %v", "pc.dp.tc", url.Path)
	}

	_, err = parseURL("pc.dp.tc")
	if err == nil {
		t.Error("expected err, got: nil")
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "pc.dp.tc"}
	if url.String() != "https://pc.dp.tc" {
		t.Errorf("expected: %v, got: %v", "https://pc.dp.tc", url.String())
	}

	url = URL{Scheme: "", Path: "pc.dp.tc"}
	if url.String() != "pc.dp.tc" {
		t.Errorf("expected: %v, got: %v", "pc.dp.tc", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "pc.dp.tc"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://pc.dp.tc\"" {
		t.Errorf("expected: %v, got: %v", "\"https://pc.dp.tc\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	url := &URL{}
	err := url.UnmarshalJSON([]byte("\"https://pc.dp.tc\""))
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "pc.dp.tc" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	tests := []struct {
		urlA   URL
		urlB   URL
		expect int
	}{
		{URL{"https", "pc.dp.tc"}, URL{"https", "pc.dp.tc"}, 0},
		{URL{"http", "pc.dp.tc"}, URL{"https", "pc.dp.tc"}, -1},
		{URL{"https", "pc.dp.tc/a"}, URL{"https", "pc.dp.tc"}, 1},
		{URL{"https", "abc.org"}, URL{"https", "pc.dp.tc"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}
