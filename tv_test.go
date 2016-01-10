// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb_test

import (
	"testing"

	. "github.com/tvtio/tmdb"
)

func TestGetTV(t *testing.T) {
	tmdb := NewTMDB()
	_, err := tmdb.GetTV("456") // The Simpsons
	if err != nil {
		t.Errorf("It should not fail, but got an error: %s", err)
	}
}

func TestBadURLGetTV(t *testing.T) {
	tmdb := NewTMDB()
	tmdb.BaseURL = "http://foo.bar/"
	_, err := tmdb.GetTV("456") // The Simpsons
	if err == nil {
		t.Errorf("It should fail because the BaseURL %s, does not exist.", tmdb.BaseURL)
	}
}
