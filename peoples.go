package gokp

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// GetStaff - Режисеры, актеры, операторы
// [GET] /getStaff?filmID={filmID}
// Параметры:
// filmID (int) - ID фильма на КиноПоиск
func GetStaff(filmID int) (GetStaffResponse, error) {
	var (
		result GetStaffResponse
		u      url.URL
	)
	stringFilmID := strconv.Itoa(filmID)
	u.Scheme = "http"
	u.Host = "api.kinopoisk.cf"
	u.Path = "/getStaff"
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

// SearchPeoples - Поиск людей
// [GET] /SearchPeoples?keyword={keyword}
// Параметры:
// keyword (string) - строка поиска
func SearchPeoples(keyword string) (SearchPeoplesResponse, error) {
	var (
		result SearchPeoplesResponse
		u      url.URL
	)
	u.Scheme = "http"
	u.Host = "api.kinopoisk.cf"
	u.Path = "/searchPeople"
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
