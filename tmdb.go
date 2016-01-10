// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

import (
	"io/ioutil"
	"net/http"
)

const (
	apiKey  = "523587afe262c34af9ee7794c5f8de81"
	baseURL = "http://api.themoviedb.org/3/"
)

// TMDB is a type that implements a client to The Movie Database API
//
// More info : https://www.themoviedb.org/documentation/api
//						 http://docs.themoviedb.apiary.io/
//
type TMDB struct {
	APIKey  string
	BaseURL string
}

// fetchContent ...
func fetchContent(url string) (body []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}

// NewTMDB allocates and initializes a new TMDB.
//
func NewTMDB() *TMDB {
	return &TMDB{APIKey: apiKey, BaseURL: baseURL}
}
