package favoritemovieserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	fm "test.com/go-favoritemovie-grpc/favoritemovie"
)

const API_KEY = "faf7e5bb&s"
const API_URL = "https://www.omdbapi.com/"

func SearhMovie(searchWord string, page int) (movies []Movie, err error) {
	if page == 0 {
		return movies, errors.New("spec: page is not valid, should be start from 1")
	}

	if searchWord == "" {
		return movies, errors.New("spec: searchword is not valid")
	}

	resp, err := http.Get(fmt.Sprintf("%v?apikey=%v&s=%v&page=%v", API_URL, API_KEY, searchWord, page))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var movieRes SearchMovieRes
	err = json.Unmarshal(body, &movieRes)
	if err != nil {
		return nil, err
	}

	for _, res := range movieRes.Search {
		movies = append(movies, Movie{
			Title:  res.Title,
			Year:   res.Year,
			Poster: res.Poster,
			ImdbID: res.ImdbID,
		})
	}

	return movies, err
}

func (svc *FavoriteMovieServer) SearchMovie(spec *fm.MovieSpec, resp fm.FavoriteMovieSearcher_SearchMovieServer) error {
	logrus.Infof("favoriteMovie.SearchMovie [Received request] searchword: %v", spec.GetSearchword())

	movies, err := SearhMovie(spec.Searchword, int(spec.Pagination))
	if err != nil {
		return err
	}
	for _, movie := range movies {
		response := fm.MovieDetail{
			Title:  movie.Title,
			Year:   movie.Year,
			ImdbId: movie.ImdbID,
			Type:   movie.Type,
			Poster: movie.Poster,
		}
		err := resp.Send(&response)
		if err != nil {
			logrus.Errorf("Error: %v", err)
			return err
		}

	}
	return nil
}
