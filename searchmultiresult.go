// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

// SearchMultiResult type represents The Movie's Database multiple media
// search result.
//
// * Search on Movies
// * Search on TV Shows
// * Search on Persons
type SearchMultiResult struct {
	Page    int `json:"page"`
	Results []struct {
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
		Overview         string   `json:"overview"`
		Popularity       float64  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int      `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}
