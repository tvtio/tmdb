// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb_test

import (
	"net/url"
	"testing"

	. "github.com/tvtio/tmdb"
)

func TestGetSeason(t *testing.T) {
	tmdb := New()
	_, err := tmdb.GetSeason("456", "1") // The Simpsons Season 1
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestBadURLGetSeason(t *testing.T) {
	tmdb := New()
	tmdb.BaseURL, _ = url.Parse("http://foo.bar/")
	_, err := tmdb.GetSeason("456", "1") // The Simpsons Season 1
	if err == nil {
		t.Errorf("It should fail because the BaseURL %s, does not exist.", tmdb.BaseURL)
	}
}

func TestBadAPIKeyGetSeason(t *testing.T) {
	tmdb := New()
	tmdb.APIKey = "foo"
	_, err := tmdb.GetSeason("456", "1") // The Simpsons Season 1
	if err != nil && err.Error() != "401 Unauthorized" {
		t.Errorf("It should fail with '401 Unauthorized', but got: %v", err)
	}
}

func Test404GetSeason(t *testing.T) {
	tmdb := New()
	_, err := tmdb.GetSeason("foo", "bar") // Do not exist
	if err != nil && err.Error() != "404 Not Found" {
		t.Errorf("It should fail with '404 Not Found', but got: %v", err)
	}
}
