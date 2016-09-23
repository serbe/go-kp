package gokp

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// GetFilm - Получить фильм
// [GET] /getFilm?filmID={filmID}
// Параметры:
// filmID (int) - ID фильма на КиноПоиск
func GetFilm(filmID int) (GetFilmResponse, error) {
	var (
		result GetFilmResponse
		u      url.URL
	)
	stringFilmID := strconv.Itoa(filmID)
	u.Scheme = "http"
	u.Host = "api.kinopoisk.cf"
	u.Path = "/getFilm"
	q := u.Query()
	q.Set("filmID", stringFilmID)
	u.RawQuery = q.Encode()
	body, err := getJSON(u.String())
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// GetGallery - Кадры фильма
// [GET] /getGallery?filmID={filmID}
// Параметры:
// filmID (int) - ID фильма на КиноПоиск
func GetGallery(filmID int) (GetGalleryResponse, error) {
	var (
		result GetGalleryResponse
		u      url.URL
	)
	stringFilmID := strconv.Itoa(filmID)
	u.Scheme = "http"
	u.Host = "api.kinopoisk.cf"
	u.Path = "/getGallery"
	q := u.Query()
	q.Set("filmID", stringFilmID)
	u.RawQuery = q.Encode()
	body, err := getJSON(u.String())
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// GetSimilar - Похожие фильмы
// [GET] /getSimilar?filmID={filmID}
// Параметры:
// filmID (int) - ID фильма на КиноПоиск
func GetSimilar(filmID int) (GetSimilarResponse, error) {
	var (
		result GetSimilarResponse
		u      url.URL
	)
	stringFilmID := strconv.Itoa(filmID)
	u.Scheme = "http"
	u.Host = "api.kinopoisk.cf"
	u.Path = "/getSimilar"
	q := u.Query()
	q.Set("filmID", stringFilmID)
	u.RawQuery = q.Encode()
	body, err := getJSON(u.String())
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// SearchFilms - Поиск фильмов
// [GET] /searchFilms?keyword={keyword}
// Параметры:
// keyword (string) - строка поиска
func SearchFilms(keyword string) (SearchFilmsResponse, error) {
	var (
		result SearchFilmsResponse
		u      url.URL
	)
	u.Scheme = "http"
	u.Host = "api.kinopoisk.cf"
	u.Path = "/searchFilms"
	q := u.Query()
	q.Set("keyword", keyword)
	u.RawQuery = q.Encode()
	body, err := getJSON(u.String())
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
