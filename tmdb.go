// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

import (
	"io/ioutil"
	"net/http"
	"net/url"
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
	// HTTP client used to communicate with the API.
	client *http.Client

	APIKey  string
	BaseURL *url.URL
}

// FetchContent ...
func (tmdb *TMDB) FetchContent(u *url.URL) (body []byte, resp *http.Response, err error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return body, nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err = tmdb.client.Do(req)
	if err != nil {
		return body, resp, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	req.Close = true
	body, err = ioutil.ReadAll(resp.Body)
	return body, resp, err
}

// New allocates and initializes a new TMDB.
//
func New() *TMDB {
	u, _ := url.Parse(baseURL)
	return &TMDB{client: http.DefaultClient, APIKey: apiKey, BaseURL: u}
}

// NewTMDB allocates and initializes a new TMDB.
//
func NewTMDB(apikey string, baseurl *url.URL) *TMDB {
	return &TMDB{client: http.DefaultClient, APIKey: apiKey, BaseURL: baseurl}
}
