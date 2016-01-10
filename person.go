// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

import "encoding/json"

// Person type represents The Movie's Database Person
type Person struct {
	Adult        bool          `json:"adult"`
	AlsoKnownAs  []interface{} `json:"also_known_as"`
	Biography    string        `json:"biography"`
	Birthday     string        `json:"birthday"`
	Deathday     string        `json:"deathday"`
	Homepage     string        `json:"homepage"`
	ID           int           `json:"id"`
	ImdbID       string        `json:"imdb_id"`
	MovieCredits struct {
		Cast []struct {
			Adult         bool   `json:"adult"`
			Character     string `json:"character"`
			CreditID      string `json:"credit_id"`
			ID            int    `json:"id"`
			OriginalTitle string `json:"original_title"`
			PosterPath    string `json:"poster_path"`
			ReleaseDate   string `json:"release_date"`
			Title         string `json:"title"`
		} `json:"cast"`
		Crew []struct {
			Adult         bool   `json:"adult"`
			CreditID      string `json:"credit_id"`
			Department    string `json:"department"`
			ID            int    `json:"id"`
			Job           string `json:"job"`
			OriginalTitle string `json:"original_title"`
			PosterPath    string `json:"poster_path"`
			ReleaseDate   string `json:"release_date"`
			Title         string `json:"title"`
		} `json:"crew"`
	} `json:"movie_credits"`
	Name         string  `json:"name"`
	PlaceOfBirth string  `json:"place_of_birth"`
	Popularity   float64 `json:"popularity"`
	ProfilePath  string  `json:"profile_path"`
	TvCredits    struct {
		Cast []struct {
			Character    string `json:"character"`
			CreditID     string `json:"credit_id"`
			EpisodeCount int    `json:"episode_count"`
			FirstAirDate string `json:"first_air_date"`
			ID           int    `json:"id"`
			Name         string `json:"name"`
			OriginalName string `json:"original_name"`
			PosterPath   string `json:"poster_path"`
		} `json:"cast"`
		Crew []struct {
			CreditID     string `json:"credit_id"`
			Department   string `json:"department"`
			EpisodeCount int    `json:"episode_count"`
			FirstAirDate string `json:"first_air_date"`
			ID           int    `json:"id"`
			Job          string `json:"job"`
			Name         string `json:"name"`
			OriginalName string `json:"original_name"`
			PosterPath   string `json:"poster_path"`
		} `json:"crew"`
	} `json:"tv_credits"`
}

// GetPerson ...
func (tmdb *TMDB) GetPerson(id string) (result Person, err error) {
	body, err := fetchContent(tmdb.BaseURL + "person/" + id + "?api_key=" + tmdb.APIKey + "&append_to_response=movie_credits,tv_credits")
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
