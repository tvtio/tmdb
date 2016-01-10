// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb_test

import (
	"net/url"
	"testing"

	. "github.com/tvtio/tmdb"
)

func TestGetTV(t *testing.T) {
	tmdb := New()
	_, err := tmdb.GetTV("456") // The Simpsons
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestBadURLGetTV(t *testing.T) {
	tmdb := New()
	tmdb.BaseURL, _ = url.Parse("http://foo.bar/")
	_, err := tmdb.GetTV("456") // The Simpsons
	if err == nil {
		t.Errorf("It should fail because the BaseURL %s, does not exist.", tmdb.BaseURL)
	}
}

func TestBadAPIKeyGetTV(t *testing.T) {
	tmdb := New()
	tmdb.APIKey = "foo"
	_, err := tmdb.GetTV("456") // The Matrix
	if err != nil && err.Error() != "401 Unauthorized" {
		t.Errorf("It should fail with '401 Unauthorized', but got: %v", err)
	}
}

func Test404GetTV(t *testing.T) {
	tmdb := New()
	_, err := tmdb.GetTV("foo") // Do not exist
	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("It should fail with '404 Not Found', but got: %v", err)
	}
}
