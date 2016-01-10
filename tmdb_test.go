// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb_test

import (
	"testing"

	. "github.com/tvtio/tmdb"
)

func TestInstance(t *testing.T) {
	tmdb := NewTMDB()
	if tmdb.APIKey == "" {
		t.Errorf("APIKey should not be blank.")
	}
	if tmdb.BaseURL.String() == "" {
		t.Errorf("BaseURL should not be blank.")
	}
}
