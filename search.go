// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// SearchMovieResult type represents The Movie Database Movies search results
type SearchMovieResult struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// SearchMultiResult type represents The Movie Database multiple media
// search result.
//
// * Search on Movies
// * Search on TV Shows
// * Search on Persons
type SearchMultiResult struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		ProfilePath      string   `json:"profile_path"`
		FirstAirDate     string   `json:"first_air_date"`
		GenreIds         []int    `json:"genre_ids"`
		ID               int      `json:"id"`
		MediaType        string   `json:"media_type"`
		Name             string   `json:"name"`
		Title            string   `json:"title"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		OriginalTitle    string   `json:"original_title"`
		Overview         string   `json:"overview"`
		Popularity       float64  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		ReleaseDate      string   `json:"release_date"`
		Video            bool     `json:"video"`
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int      `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// SearchTVResult type represents The Movie Database TV Shows search results
type SearchTVResult struct {
	Page    int `json:"page"`
	Results []struct {
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date"`
		GenreIds         []int    `json:"genre_ids"`
		ID               int      `json:"id"`
		Name             string   `json:"name"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float64  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int      `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// PopularMovie ...
func (tmdb *TMDB) PopularMovie() (result SearchMovieResult, err error) {
	s := fmt.Sprintf("%smovie/popular?api_key=%s", tmdb.BaseURL, tmdb.APIKey)
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := tmdb.FetchContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// PopularTV ...
func (tmdb *TMDB) PopularTV() (result SearchTVResult, err error) {
	s := fmt.Sprintf("%stv/popular?api_key=%s", tmdb.BaseURL, tmdb.APIKey)
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := tmdb.FetchContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// SearchMulti ...
func (tmdb *TMDB) SearchMulti(query string) (result SearchMultiResult, err error) {
	s := fmt.Sprintf("%ssearch/multi?api_key=%s&query=%s", tmdb.BaseURL, tmdb.APIKey, query)
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := tmdb.FetchContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// SearchMovie ...
func (tmdb *TMDB) SearchMovie(query string) (result SearchMovieResult, err error) {
	s := fmt.Sprintf("%ssearch/movie?api_key=%s&query=%s", tmdb.BaseURL, tmdb.APIKey, query)
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := tmdb.FetchContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
