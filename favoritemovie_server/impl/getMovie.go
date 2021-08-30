package favoritemovieserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	fm "test.com/go-favoritemovie-grpc/favoritemovie"

	"github.com/sirupsen/logrus"
)

func GetMovie(id string) (movie Movie, err error) {
	if id == "" {
		return movie, errors.New("spec: searchword is not valid")
	}

	resp, err := http.Get(fmt.Sprintf("%v?apikey=%v&i=%v", API_URL, API_KEY, id))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return movie, errors.New("received non 200 response code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return movie, err
	}
	var movieRes SearchMovieRes
	err = json.Unmarshal(body, &movieRes)
	if err != nil {
		return movie, err
	}

	return
}

func (svc *FavoriteMovieServer) GetMovie(ctx context.Context, spec *fm.GetMovieSpec) (*fm.MovieDetail, error) {
	logrus.Infof("favoriteMovie.SearchMovie [Received request] searchword: %v", spec.GetIdMovie())

	movie, err := GetMovie(spec.IdMovie)
	if err != nil {
		return nil, err
	}

	return &fm.MovieDetail{
		Title:  movie.Title,
		Year:   movie.Year,
		ImdbId: movie.ImdbID,
		Type:   movie.Type,
		Poster: movie.Poster,
	}, nil
}
