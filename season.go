// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Season type represents The Movie's Database TV Show Season
type Season struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
	AirDate      string `json:"air_date"`
	Credits      struct {
		Cast []struct {
			Character   string `json:"character"`
			CreditID    string `json:"credit_id"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Order       int    `json:"order"`
			ProfilePath string `json:"profile_path"`
		} `json:"cast"`
		Crew []struct {
			CreditID    string `json:"credit_id"`
			Department  string `json:"department"`
			ID          int    `json:"id"`
			Job         string `json:"job"`
			Name        string `json:"name"`
			ProfilePath string `json:"profile_path"`
		} `json:"crew"`
	} `json:"credits"`
	Episodes []struct {
		AirDate string `json:"air_date"`
		Crew    []struct {
			CreditID    string      `json:"credit_id"`
			Department  string      `json:"department"`
			ID          int         `json:"id"`
			Job         string      `json:"job"`
			Name        string      `json:"name"`
			ProfilePath interface{} `json:"profile_path"`
		} `json:"crew"`
		EpisodeNumber int `json:"episode_number"`
		GuestStars    []struct {
			Character   string `json:"character"`
			CreditID    string `json:"credit_id"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Order       int    `json:"order"`
			ProfilePath string `json:"profile_path"`
		} `json:"guest_stars"`
		ID             int         `json:"id"`
		Name           string      `json:"name"`
		Overview       string      `json:"overview"`
		ProductionCode interface{} `json:"production_code"`
		SeasonNumber   int         `json:"season_number"`
		StillPath      string      `json:"still_path"`
		VoteAverage    float64     `json:"vote_average"`
		VoteCount      int         `json:"vote_count"`
	} `json:"episodes"`
	ExternalIds struct {
		FreebaseID  string      `json:"freebase_id"`
		FreebaseMid string      `json:"freebase_mid"`
		TvdbID      interface{} `json:"tvdb_id"`
		TvrageID    interface{} `json:"tvrage_id"`
	} `json:"external_ids"`
	Images struct {
		Posters []struct {
			AspectRatio float64     `json:"aspect_ratio"`
			FilePath    string      `json:"file_path"`
			Height      int         `json:"height"`
			Iso639_1    interface{} `json:"iso_639_1"`
			VoteAverage float64     `json:"vote_average"`
			VoteCount   int         `json:"vote_count"`
			Width       int         `json:"width"`
		} `json:"posters"`
	} `json:"images"`
	Videos struct {
		Results []interface{} `json:"results"`
	} `json:"videos"`
}

// GetSeason ...
func (tmdb *TMDB) GetSeason(id string, snumber string) (result Season, err error) {
	s := fmt.Sprintf("%stv/%s/season/%s?api_key=%s&append_to_response=credits,external_ids,images,videos", tmdb.BaseURL, id, snumber, tmdb.APIKey)
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, err := fetchContent(u)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
