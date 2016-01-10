// Copyright 2016 The tmdb Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmdb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiKey  = "523587afe262c34af9ee7794c5f8de81"
	baseURL = "http://api.themoviedb.org/3/"
)

func fetchContent(url string) (body []byte, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}

// PopularMovie ...
func PopularMovie() (result SearchMovieResult, err error) {
	body, _ := fetchContent(baseURL + "movie/popular?api_key=" + apiKey)

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// PopularTV ...
func PopularTV() (result SearchTVResult, err error) {
	body, _ := fetchContent(baseURL + "tv/popular?api_key=" + apiKey)

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// SearchMulti ...
func SearchMulti(query string) (result SearchMultiResult, err error) {
	body, _ := fetchContent(baseURL + "search/multi?api_key=" + apiKey + "&query=" + query)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// SearchMovie ...
func SearchMovie(query string) (result SearchMovieResult, err error) {
	body, _ := fetchContent(baseURL + "search/movie?api_key=" + apiKey + "&query=" + query)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// GetMovie ...
func GetMovie(id string) (result Movie, err error) {
	body, _ := fetchContent(baseURL + "movie/" + id + "?api_key=" + apiKey + "&append_to_response=credits,images")
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}

	return result, nil
}

// GetPerson ...
func GetPerson(id string) (result Person, err error) {
	body, _ := fetchContent(baseURL + "person/" + id + "?api_key=" + apiKey + "&append_to_response=movie_credits,tv_credits")
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// GetTV ...
func GetTV(id string) (result TV, err error) {
	body, _ := fetchContent(baseURL + "tv/" + id + "?api_key=" + apiKey + "&append_to_response=credits")
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// GetSeason ...
func GetSeason(id string, snumber string) (result Season, err error) {
	body, _ := fetchContent(baseURL + "tv/" + id + "/season/" + snumber + "?api_key=" + apiKey + "&append_to_response=credits,external_ids,images,videos")
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}

// GetEpisode ...
func GetEpisode(id string, snumber string, enumber string) (result Episode, err error) {
	body, _ := fetchContent(baseURL + "tv/" + id + "/season/" + snumber + "/episode/" + enumber + "?api_key=" + apiKey + "&append_to_response=credits,external_ids,images,videos")
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}
