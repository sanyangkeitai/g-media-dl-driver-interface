package model

type MediaDetailInfo struct {
	Title        string          `json:"title"`
	CoverUrl     string          `json:"coverUrl"`
	CategoryList []string        `json:"categoryList"`
	Year         string          `json:"year"`
	Director     string          `json:"director"`
	Introduction string          `json:"introduction"`
	PlayLineList [][]EpisodeLink `json:"playLineList"`
}
