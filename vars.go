package gokp

type videoURL struct {
	Hd  string `json:"hd"`
	Sd  string `json:"sd"`
	Low string `json:"low"`
}

type image struct {
	Image     string `json:"image"`
	Preview   string `json:"preview"`
	SocialURL string `json:"socialURL"`
}

type creator struct {
	ID             string `json:"id"`
	NameEN         string `json:"nameEN"`
	NameRU         string `json:"nameRU"`
	PosterURL      string `json:"posterURL"`
	ProfessionKey  string `json:"professionKey"`
	ProfessionText string `json:"professionText"`
	Type           string `json:"type"`
}

// GetFilmResponse - JSON response from getFilm
type GetFilmResponse struct {
	BudgetData struct {
		Budget     string `json:"budget"`
		GrossRU    string `json:"grossRU"`
		GrossUSA   string `json:"grossUSA"`
		GrossWorld string `json:"grossWorld"`
		Marketing  string `json:"marketing"`
	} `json:"budgetData"`
	Country                    string      `json:"country"`
	Creators                   [][]creator `json:"creators"`
	Description                string      `json:"description"`
	FilmID                     string      `json:"filmID"`
	FilmLength                 string      `json:"filmLength"`
	Gallery                    []image     `json:"gallery"`
	Genre                      string      `json:"genre"`
	HasRelatedFilms            int         `json:"hasRelatedFilms"`
	HasSequelsAndPrequelsFilms int         `json:"hasSequelsAndPrequelsFilms"`
	HasSimilarFilms            int         `json:"hasSimilarFilms"`
	ImdbID                     string      `json:"imdbID"`
	NameEN                     string      `json:"nameEN"`
	NameRU                     string      `json:"nameRU"`
	PosterURL                  string      `json:"posterURL"`
	RatingAgeLimits            string      `json:"ratingAgeLimits"`
	RatingData                 struct {
		Rating                     string `json:"rating"`
		RatingFilmCritics          string `json:"ratingFilmCritics"`
		RatingFilmCriticsVoteCount string `json:"ratingFilmCriticsVoteCount"`
		RatingGoodReview           string `json:"ratingGoodReview"`
		RatingGoodReviewVoteCount  int    `json:"ratingGoodReviewVoteCount"`
		RatingIMDb                 string `json:"ratingIMDb"`
		RatingIMDbVoteCount        string `json:"ratingIMDbVoteCount"`
		RatingRFCritics            string `json:"ratingRFCritics"`
		RatingRFCriticsVoteCount   int    `json:"ratingRFCriticsVoteCount"`
		RatingVoteCount            string `json:"ratingVoteCount"`
	} `json:"ratingData"`
	RatingMPAA string `json:"ratingMPAA"`
	RentData   struct {
		Distributors         string `json:"Distributors"`
		DistributorRelease   string `json:"distributorRelease"`
		PremiereBluRay       string `json:"premiereBluRay"`
		PremiereDVD          string `json:"premiereDVD"`
		PremiereRU           string `json:"premiereRU"`
		PremiereWorld        string `json:"premiereWorld"`
		PremiereWorldCountry string `json:"premiereWorldCountry"`
	} `json:"rentData"`
	ReviewsCount  int    `json:"reviewsCount"`
	Slogan        string `json:"slogan"`
	TopNewsByFilm struct {
		ID                  string `json:"ID"`
		NewsDate            string `json:"newsDate"`
		NewsDescription     string `json:"newsDescription"`
		NewsImagePreviewURL string `json:"newsImagePreviewURL"`
		NewsTitle           string `json:"newsTitle"`
	} `json:"topNewsByFilm"`
	TriviaData []string `json:"triviaData"`
	WebURL     string   `json:"webURL"`
	Year       string   `json:"year"`
}

// GetGalleryResponse - JSON response from getGallery
type GetGalleryResponse struct {
	Gallery struct {
		Kadr   []image `json:"kadr"`
		KadrSp []image `json:"kadr_sp"`
		Poster []image `json:"poster"`
	} `json:"gallery"`
}

// GetStaffResponse - JSON response from getStaff
type GetStaffResponse struct {
	Creators [][]creator `json:"creators"`
}

// GetSimilarResponse - JSON response from getSimilar
type GetSimilarResponse struct {
	Items []struct {
		Country      string   `json:"country"`
		FilmLength   string   `json:"filmLength"`
		FilmTypeText string   `json:"filmTypeText"`
		Genre        string   `json:"genre"`
		ID           string   `json:"id"`
		NameEN       string   `json:"nameEN"`
		NameRU       string   `json:"nameRU"`
		PosterURL    string   `json:"posterURL"`
		Rating       string   `json:"rating"`
		Type         string   `json:"type"`
		VideoURL     videoURL `json:"videoURL"`
		Year         string   `json:"year"`
	} `json:"items"`
	PagesCount int `json:"pagesCount"`
}

// GetGenresResponse - JSON response from getGenres
type GetGenresResponse struct {
	GenreData []struct {
		GenreID   string `json:"genreID"`
		GenreName string `json:"genreName"`
	} `json:"genreData"`
}

// SearchFilmsResponse - JSON response from searchFilms
type SearchFilmsResponse struct {
	Keyword     string `json:"keyword"`
	PagesCount  int    `json:"pagesCount"`
	SearchFilms []struct {
		Country     string   `json:"country"`
		Description string   `json:"description"`
		FilmLength  string   `json:"filmLength"`
		Genre       string   `json:"genre"`
		ID          string   `json:"id"`
		NameEN      string   `json:"nameEN"`
		NameRU      string   `json:"nameRU"`
		PosterURL   string   `json:"posterURL"`
		Rating      string   `json:"rating"`
		Type        string   `json:"type"`
		VideoURL    videoURL `json:"videoURL"`
		Year        string   `json:"year"`
	} `json:"searchFilms"`
	SearchFilmsCountResult int `json:"searchFilmsCountResult"`
}

// SearchPeoplesResponse - JSON response from SearchPeople
type SearchPeoplesResponse struct {
	Keyword      string `json:"keyword"`
	PagesCount   int    `json:"pagesCount"`
	SearchPeople []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		NameRU      string `json:"nameRU"`
		NameEN      string `json:"nameEN"`
		Description string `json:"description"`
		PosterURL   string `json:"posterURL"`
	} `json:"searchPeople"`
	SearchPeoplesCountResult int `json:"searchPeoplesCountResult"`
}
