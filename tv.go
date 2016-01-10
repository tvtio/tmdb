// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

// TV type represents The Movie's Database TV Show
type TV struct {
	BackdropPath string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
	} `json:"created_by"`
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
			CreditID    string `json:"credit_id"`
			Department  string `json:"department"`
			ID          int    `json:"id"`
			Job         string `json:"job"`
			Name        string `json:"name"`
			ProfilePath string `json:"profile_path"`
		} `json:"crew"`
	} `json:"credits"`
	EpisodeRunTime []int  `json:"episode_run_time"`
	FirstAirDate   string `json:"first_air_date"`
	Genres         []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage     string   `json:"homepage"`
	ID           int      `json:"id"`
	InProduction bool     `json:"in_production"`
	Languages    []string `json:"languages"`
	LastAirDate  string   `json:"last_air_date"`
	Name         string   `json:"name"`
	Networks     []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"networks"`
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string   `json:"overview"`
	Popularity          float64  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"production_companies"`
	Seasons []struct {
		AirDate      interface{} `json:"air_date"`
		EpisodeCount int         `json:"episode_count"`
		ID           int         `json:"id"`
		PosterPath   string      `json:"poster_path"`
		SeasonNumber int         `json:"season_number"`
	} `json:"seasons"`
	Status      string  `json:"status"`
	Type        string  `json:"type"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}
