// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

// Movie type represents The Movie's Database Movie
type Movie struct {
	Adult               bool    `json:"adult"`
	BackdropPath        string  `json:"backdrop_path"`
	Budget              int     `json:"budget"`
	Homepage            string  `json:"homepage"`
	ID                  int     `json:"id"`
	ImdbID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float64 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	Revenue             int     `json:"revenue"`
	ReleaseDate         string  `json:"release_date"`
	Runtime             int     `json:"runtime"`
	Status              string  `json:"status"`
	Tagline             string  `json:"tagline"`
	Title               string  `json:"title"`
	Video               bool    `json:"video"`
	VoteAverage         float64 `json:"vote_average"`
	VoteCount           int     `json:"vote_count"`
	BelongsToCollection struct {
		BackdropPath string `json:"backdrop_path"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
	} `json:"belongs_to_collection"`
	Credits struct {
		Cast []struct {
			CastID      int    `json:"cast_id"`
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
	Genres []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Images struct {
		Backdrops []struct {
			AspectRatio float64     `json:"aspect_ratio"`
			FilePath    string      `json:"file_path"`
			Height      int         `json:"height"`
			Iso639_1    interface{} `json:"iso_639_1"`
			VoteAverage float64     `json:"vote_average"`
			VoteCount   int         `json:"vote_count"`
			Width       int         `json:"width"`
		} `json:"backdrops"`
		Posters []struct {
			AspectRatio float64 `json:"aspect_ratio"`
			FilePath    string  `json:"file_path"`
			Height      int     `json:"height"`
			Iso639_1    string  `json:"iso_639_1"`
			VoteAverage float64 `json:"vote_average"`
			VoteCount   int     `json:"vote_count"`
			Width       int     `json:"width"`
		} `json:"posters"`
	} `json:"images"`
	ProductionCompanies []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string `json:"name"`
	} `json:"production_countries"`
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso_639_1"`
		Name     string `json:"name"`
	} `json:"spoken_languages"`
}
