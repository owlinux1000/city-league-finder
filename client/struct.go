package client

import "net/http"

type Client struct {
	client   *http.Client
	Endpoint string
}

type EventSearchParams struct {
	Prefecture []string
	EventType  []string
	LeagueType []string
	Offset     string
	Accepting  string
	Order      string
}

type EventSearchResponse struct {
	Code       int     `json:"code"`
	Message    string  `json:"message"`
	EventCount int     `json:"eventCount"`
	Events     []Event `json:"event"`
}

type Event struct {
	Access            string  `json:"access"`
	Address           string  `json:"address"`
	BeginnerShopFlg   int     `json:"beginnerShopFlg"`
	CancelFlg         int     `json:"cancelFlg"`
	Capacity          int     `json:"capacity"`
	ChampionShopFlg   int     `json:"championShopFlg"`
	CspFlg            int     `json:"csp_flg"`
	DateID            int     `json:"date_id"`
	DeckCount         string  `json:"deck_count"`
	DiscontinuanceFlg int     `json:"discontinuance_flg"`
	Distance          float64 `json:"distance"`
	EntryRestartFlg   int     `json:"entryRestartFlg"`
	EntryStatus       string  `json:"entryStatus"`
	EntryStatusCode   int     `json:"entryStatusCode"`
	EntryFee          string  `json:"entry_fee"`
	EventAttrID       int     `json:"event_attr_id"`
	EventDate         string  `json:"event_date"`
	EventDateParams   string  `json:"event_date_params"`
	EventDateWeek     string  `json:"event_date_week"`
	EventEndedAt      string  `json:"event_ended_at"`
	EventHoldingID    int     `json:"event_holding_id"`
	EventLeague       int     `json:"event_league"`
	EventStartedAt    string  `json:"event_started_at"`
	EventTitle        string  `json:"event_title"`
	EventType         int     `json:"event_type"`
	FullOccupiedFlg   int     `json:"fullOccupiedFlg"`
	HolidayFlg        int     `json:"holiday_flg"`
	ID                int     `json:"id"`
	LeagueName        string  `json:"leagueName"`
	NoOfMyGymReg      int     `json:"noOfMyGymReg"`
	PrefectureName    string  `json:"prefecture_name"`
	RecruitFlg        int     `json:"recruitFlg"`
	Regulation        string  `json:"regulation"`
	ShopID            int     `json:"shop_id"`
	ShopName          string  `json:"shop_name"`
	ShopTerm          int     `json:"shop_term"`
	StrongShopFlg     int     `json:"strongShopFlg"`
	TrainersFlg       int     `json:"trainers_flg"`
	Venue             string  `json:"venue"`
	ZipCode           int     `json:"zip_code"`
}
