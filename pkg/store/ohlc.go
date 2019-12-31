package store

//ResultStore is the final result stored in the mongoDB
type ResultStore struct {
	IndicatorType  string
	InstrumentList []Instrument
}

//Instrument represents a partcular stock in BSE or NSE
type Instrument struct {
	Name     string  `json:"Name"`
	Exchange string  `json:"Exchange"`
	Symbol   string  `json:"Symbol"`
	Token    string  `json:"Token"`
	OHLC     *[]OHLC `json:"OHLC"`
}

//Result of the OHLC analysis
type Result struct {
	//Uptrend indicators
	BullishMarubuzoAfterDecline []Instrument `json:"BullishMarubuzoAfterDecline"`
	DoziAfterDecline            []Instrument `json:"DoziAfterDecline"`
	BullishHammerAfterDecline   []Instrument `json:"BullishHammerAfterDecline"`
	BearishHammerAfterDecline   []Instrument `json:"BearishHammerAfterDecline"`
	EndOfDecline                []Instrument `json:"EndOfDecline"`

	//Downtrend Indicators
	BearishMarubuzoAfterRally []Instrument `json:"BearishMarubuzoAfterRally"`
	DoziAfterRally            []Instrument `json:"DoziAfterRally"`
	ShootingStarAfterDecline  []Instrument `json:"ShootingStarAfterDecline"`
	ShootingStartAfterRally   []Instrument `json:"ShootingStartAfterRally"`
	EndOfRally                []Instrument `json:"EndOfRally"`

	//Others
	OpenLowHigh []Instrument `json:"OpenLowHigh"`

	//Other chart types
	BullishMarubuzo []Instrument `json:"BullishMarubuzo"`
	BearishMarubuzo []Instrument `json:"BearishMarubuzo"`
	Dozi            []Instrument `json:"Dozi"`
	Hammer          []Instrument `json:"Hammer"`
	ShootingStar    []Instrument `json:"ShootingStar"`

	//TimePeriod
	TimePeriod string `json:"TimePeriod"`
}

//OHLC is the open, high, low and close price for an instrument.
type OHLC struct {
	Open  float64 `json:"Open"`
	High  float64 `json:"High"`
	Low   float64 `json:"Low"`
	Close float64 `json:"Close"`
}
