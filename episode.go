// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Episode type represents The Movie's Database TV Shows Episode
type Episode struct {
	AirDate string `json:"air_date"`
	Credits struct {
		Cast []struct {
			Character   string `json:"character"`
			CreditID    string `json:"credit_id"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Order       int    `json:"order"`
			ProfilePath string `json:"profile_path"`
		} `json:"cast"`
		Crew []struct {
			CreditID    string      `json:"credit_id"`
			Department  string      `json:"department"`
			ID          int         `json:"id"`
			Job         string      `json:"job"`
			Name        string      `json:"name"`
			ProfilePath interface{} `json:"profile_path"`
		} `json:"crew"`
		GuestStars []struct {
			Character   string `json:"character"`
			CreditID    string `json:"credit_id"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Order       int    `json:"order"`
			ProfilePath string `json:"profile_path"`
		} `json:"guest_stars"`
	} `json:"credits"`
	Crew []struct {
		CreditID    string      `json:"credit_id"`
		Department  string      `json:"department"`
		ID          int         `json:"id"`
		Job         string      `json:"job"`
		Name        string      `json:"name"`
		ProfilePath interface{} `json:"profile_path"`
	} `json:"crew"`
	EpisodeNumber int `json:"episode_number"`
	ExternalIds   struct {
		FreebaseID  string `json:"freebase_id"`
		FreebaseMid string `json:"freebase_mid"`
		ImdbID      string `json:"imdb_id"`
		TvdbID      int    `json:"tvdb_id"`
		TvrageID    int    `json:"tvrage_id"`
	} `json:"external_ids"`
	GuestStars []struct {
		Character   string `json:"character"`
		CreditID    string `json:"credit_id"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Order       int    `json:"order"`
		ProfilePath string `json:"profile_path"`
	} `json:"guest_stars"`
	ID     int `json:"id"`
	Images struct {
		Stills []struct {
			AspectRatio float64     `json:"aspect_ratio"`
			FilePath    string      `json:"file_path"`
			Height      int         `json:"height"`
			Iso639_1    interface{} `json:"iso_639_1"`
			VoteAverage float64     `json:"vote_average"`
			VoteCount   int         `json:"vote_count"`
			Width       int         `json:"width"`
		} `json:"stills"`
	} `json:"images"`
	Name           string      `json:"name"`
	Overview       string      `json:"overview"`
	ProductionCode interface{} `json:"production_code"`
	SeasonNumber   int         `json:"season_number"`
	StillPath      string      `json:"still_path"`
	Videos         struct {
		Results []interface{} `json:"results"`
	} `json:"videos"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

// GetEpisode ...
func (tmdb *TMDB) GetEpisode(id string, snumber string, enumber string) (result Episode, err error) {
	s := fmt.Sprintf("%stv/%s/season/%s/episode/%s?api_key=%s&append_to_response=credits,external_ids,images,videos", tmdb.BaseURL, id, snumber, enumber, tmdb.APIKey)
	u, _ := url.Parse(s)
	body, err := fetchContent(u)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
