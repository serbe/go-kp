package gokp

// GetFilmResponse - JSON response from getFilm
type GetFilmResponse struct {
	HasSimilarFilms int `json:"hasSimilarFilms"`
	ReviewsCount    int `json:"reviewsCount"`
	RatingData      struct {
		RatingGoodReview           string `json:"ratingGoodReview"`
		RatingGoodReviewVoteCount  int    `json:"ratingGoodReviewVoteCount"`
		Rating                     string `json:"rating"`
		RatingVoteCount            string `json:"ratingVoteCount"`
		RatingAwait                string `json:"ratingAwait"`
		RatingAwaitCount           string `json:"ratingAwaitCount"`
		RatingIMDb                 string `json:"ratingIMDb"`
		RatingIMDbVoteCount        string `json:"ratingIMDbVoteCount"`
		RatingFilmCritics          string `json:"ratingFilmCritics"`
		RatingFilmCriticsVoteCount string `json:"ratingFilmCriticsVoteCount"`
		RatingRFCritics            string `json:"ratingRFCritics"`
		RatingRFCriticsVoteCount   int    `json:"ratingRFCriticsVoteCount"`
	} `json:"ratingData"`
	HasSequelsAndPrequelsFilms int    `json:"hasSequelsAndPrequelsFilms"`
	HasRelatedFilms            int    `json:"hasRelatedFilms"`
	FilmID                     string `json:"filmID"`
	WebURL                     string `json:"webURL"`
	NameRU                     string `json:"nameRU"`
	NameEN                     string `json:"nameEN"`
	PosterURL                  string `json:"posterURL"`
	Year                       string `json:"year"`
	FilmLength                 string `json:"filmLength"`
	Country                    string `json:"country"`
	Genre                      string `json:"genre"`
	Slogan                     string `json:"slogan"`
	Description                string `json:"description"`
	VideoURL                   struct {
		Hd  string `json:"hd"`
		Sd  string `json:"sd"`
		Low string `json:"low"`
	} `json:"videoURL"`
	RatingMPAA      string `json:"ratingMPAA"`
	IsIMAX          int    `json:"isIMAX"`
	Is3D            int    `json:"is3D"`
	RatingAgeLimits string `json:"ratingAgeLimits"`
	RentData        struct {
		PremiereRU           string `json:"premiereRU"`
		Distributors         string `json:"Distributors"`
		PremiereWorld        string `json:"premiereWorld"`
		PremiereWorldCountry string `json:"premiereWorldCountry"`
		PremiereDVD          string `json:"premiereDVD"`
		PremiereBluRay       string `json:"premiereBluRay"`
		DistributorRelease   string `json:"distributorRelease"`
	} `json:"rentData"`
	BudgetData struct {
		GrossRU    string `json:"grossRU"`
		GrossUSA   string `json:"grossUSA"`
		GrossWorld string `json:"grossWorld"`
		Budget     string `json:"budget"`
	} `json:"budgetData"`
	Gallery []struct {
		Preview string `json:"preview"`
	} `json:"gallery"`
	Creators []struct {
		Num0 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"0"`
		Num1 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"1,omitempty"`
		Num2 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"2,omitempty"`
	} `json:"creators"`
	TopNewsByFilm struct {
		ID                  string `json:"ID"`
		NewsDate            string `json:"newsDate"`
		NewsImagePreviewURL string `json:"newsImagePreviewURL"`
		NewsTitle           string `json:"newsTitle"`
		NewsDescription     string `json:"newsDescription"`
	} `json:"topNewsByFilm"`
	TriviaData []string `json:"triviaData"`
	ImdbID     string   `json:"imdbID"`
}

// GetGalleryResponse - JSON response from getGallery
type GetGalleryResponse struct {
	Gallery struct {
		Kadr []struct {
			Image     string `json:"image"`
			Preview   string `json:"preview"`
			SocialURL string `json:"socialURL"`
		} `json:"kadr"`
		KadrSp []struct {
			Image     string `json:"image"`
			Preview   string `json:"preview"`
			SocialURL string `json:"socialURL"`
		} `json:"kadr_sp"`
		Poster []struct {
			Image     string `json:"image"`
			Preview   string `json:"preview"`
			SocialURL string `json:"socialURL"`
		} `json:"poster"`
	} `json:"gallery"`
}

// GetStaffResponse - JSON response from getStaff
type GetStaffResponse struct {
	Creators []struct {
		Num0 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"0"`
		Num1 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"1,omitempty"`
		Num2 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"2,omitempty"`
		Num3 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"3,omitempty"`
		Num4 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"4,omitempty"`
		Num5 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"5,omitempty"`
		Num6 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"6,omitempty"`
		Num7 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"7,omitempty"`
		Num8 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"8,omitempty"`
		Num9 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"9,omitempty"`
		Num10 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"10,omitempty"`
		Num11 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"11,omitempty"`
		Num12 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"12,omitempty"`
		Num13 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"13,omitempty"`
		Num14 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"14,omitempty"`
		Num15 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"15,omitempty"`
		Num16 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"16,omitempty"`
		Num17 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"17,omitempty"`
		Num18 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"18,omitempty"`
		Num19 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"19,omitempty"`
		Num20 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"20,omitempty"`
		Num21 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"21,omitempty"`
		Num22 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"22,omitempty"`
		Num23 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"23,omitempty"`
		Num24 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"24,omitempty"`
		Num25 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"25,omitempty"`
		Num26 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"26,omitempty"`
		Num27 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"27,omitempty"`
		Num28 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"28,omitempty"`
		Num29 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"29,omitempty"`
		Num30 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"30,omitempty"`
		Num31 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"31,omitempty"`
		Num32 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"32,omitempty"`
		Num33 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"33,omitempty"`
		Num34 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"34,omitempty"`
		Num35 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"35,omitempty"`
		Num36 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"36,omitempty"`
		Num37 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"37,omitempty"`
		Num38 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"38,omitempty"`
		Num39 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"39,omitempty"`
		Num40 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"40,omitempty"`
		Num41 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"41,omitempty"`
		Num42 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"42,omitempty"`
		Num43 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"43,omitempty"`
		Num44 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"44,omitempty"`
		Num45 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"45,omitempty"`
		Num46 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"46,omitempty"`
		Num47 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"47,omitempty"`
		Num48 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"48,omitempty"`
		Num49 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"49,omitempty"`
		Num50 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"50,omitempty"`
		Num51 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"51,omitempty"`
		Num52 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"52,omitempty"`
		Num53 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"53,omitempty"`
		Num54 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"54,omitempty"`
		Num55 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"55,omitempty"`
		Num56 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"56,omitempty"`
		Num57 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"57,omitempty"`
		Num58 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"58,omitempty"`
		Num59 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"59,omitempty"`
		Num60 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"60,omitempty"`
		Num61 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"61,omitempty"`
		Num62 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"62,omitempty"`
		Num63 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"63,omitempty"`
		Num64 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"64,omitempty"`
		Num65 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"65,omitempty"`
		Num66 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"66,omitempty"`
		Num67 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"67,omitempty"`
		Num68 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"68,omitempty"`
		Num69 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"69,omitempty"`
		Num70 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"70,omitempty"`
		Num71 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"71,omitempty"`
		Num72 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"72,omitempty"`
		Num73 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"73,omitempty"`
		Num74 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"74,omitempty"`
		Num75 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"75,omitempty"`
		Num76 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"76,omitempty"`
		Num77 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"77,omitempty"`
		Num78 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"78,omitempty"`
		Num79 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"79,omitempty"`
		Num80 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"80,omitempty"`
		Num81 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"81,omitempty"`
		Num82 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"82,omitempty"`
		Num83 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"83,omitempty"`
		Num84 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"84,omitempty"`
		Num85 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"85,omitempty"`
		Num86 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"86,omitempty"`
		Num87 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"87,omitempty"`
		Num88 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"88,omitempty"`
		Num89 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"89,omitempty"`
		Num90 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"90,omitempty"`
		Num91 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"91,omitempty"`
		Num92 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"92,omitempty"`
		Num93 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"93,omitempty"`
		Num94 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"94,omitempty"`
		Num95 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"95,omitempty"`
		Num96 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"96,omitempty"`
		Num97 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"97,omitempty"`
		Num98 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"98,omitempty"`
		Num99 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"99,omitempty"`
		Num100 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"100,omitempty"`
		Num101 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"101,omitempty"`
		Num102 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"102,omitempty"`
		Num103 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"103,omitempty"`
		Num104 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"104,omitempty"`
		Num105 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"105,omitempty"`
		Num106 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"106,omitempty"`
		Num107 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"107,omitempty"`
		Num108 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"108,omitempty"`
		Num109 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"109,omitempty"`
		Num110 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"110,omitempty"`
		Num111 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"111,omitempty"`
		Num112 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"112,omitempty"`
		Num113 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"113,omitempty"`
		Num114 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"114,omitempty"`
		Num115 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"115,omitempty"`
		Num116 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"116,omitempty"`
		Num117 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"117,omitempty"`
		Num118 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"118,omitempty"`
		Num119 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"119,omitempty"`
		Num120 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"120,omitempty"`
		Num121 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"121,omitempty"`
		Num122 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"122,omitempty"`
		Num123 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"123,omitempty"`
		Num124 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"124,omitempty"`
		Num125 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"125,omitempty"`
		Num126 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"126,omitempty"`
		Num127 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"127,omitempty"`
		Num128 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"128,omitempty"`
		Num129 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"129,omitempty"`
		Num130 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"130,omitempty"`
		Num131 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"131,omitempty"`
		Num132 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"132,omitempty"`
		Num133 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"133,omitempty"`
		Num134 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			PosterURL      string `json:"posterURL"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"134,omitempty"`
		Num135 struct {
			ID             string `json:"id"`
			Type           string `json:"type"`
			NameRU         string `json:"nameRU"`
			NameEN         string `json:"nameEN"`
			Description    string `json:"description"`
			ProfessionText string `json:"professionText"`
			ProfessionKey  string `json:"professionKey"`
		} `json:"135,omitempty"`
	} `json:"creators"`
}

// GetSimilarResponse - JSON response from getSimilar
type GetSimilarResponse struct {
	Items []struct {
		Type         string `json:"type"`
		ID           string `json:"id"`
		NameRU       string `json:"nameRU"`
		NameEN       string `json:"nameEN"`
		Year         string `json:"year"`
		Rating       string `json:"rating"`
		PosterURL    string `json:"posterURL"`
		FilmLength   string `json:"filmLength"`
		Country      string `json:"country"`
		Genre        string `json:"genre"`
		FilmTypeText string `json:"filmTypeText"`
		VideoURL     struct {
			Sd  string `json:"sd"`
			Low string `json:"low"`
		} `json:"videoURL,omitempty"`
	} `json:"items"`
	PagesCount int `json:"pagesCount"`
}

// SearchFilmsResponse - JSON response from searchFilms
type SearchFilmsResponse struct {
	Keyword     string `json:"keyword"`
	PagesCount  int    `json:"pagesCount"`
	SearchFilms []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		NameRU      string `json:"nameRU"`
		NameEN      string `json:"nameEN"`
		Description string `json:"description"`
		PosterURL   string `json:"posterURL,omitempty"`
		FilmLength  string `json:"filmLength,omitempty"`
		Year        string `json:"year,omitempty"`
		Country     string `json:"country"`
		Genre       string `json:"genre"`
		Rating      string `json:"rating"`
		VideoURL    struct {
			Hd  string `json:"hd"`
			Sd  string `json:"sd"`
			Low string `json:"low"`
		} `json:"videoURL,omitempty"`
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
