// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb_test

import (
	"testing"

	. "github.com/tvtio/tmdb"
)

func TestGetMovie(t *testing.T) {
	tmdb := NewTMDB()
	_, err := tmdb.GetMovie("603") // The Matrix
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestBadURLGetMovie(t *testing.T) {
	tmdb := NewTMDB()
	tmdb.BaseURL = "http://foo.bar/"
	_, err := tmdb.GetMovie("603") // The Matrix
	if err == nil {
		t.Errorf("It should fail because the BaseURL %s, does not exist.", tmdb.BaseURL)
	}
}
