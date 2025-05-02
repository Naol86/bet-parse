package main

type Odds struct {
	ID       string `json:"id"`
	Odds     string `json:"odds"`
	Name     string `json:"name"`
	Header   string `json:"header"`
	Handicap string `json:"handicap"`
}

type Market struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Odds  []Odds `json:"odds"`
	Open  int    `json:"open,omitempty"`
}