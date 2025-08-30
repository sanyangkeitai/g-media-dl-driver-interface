package model

import "github.com/sanyangkeitai/g-media-dl-driver-interface/driver/constant"

type MediaDetailInfo struct {
	Title        string           `json:"title"`
	CoverUrl     string           `json:"coverUrl"`
	CategoryList []string         `json:"categoryList"`
	Year         string           `json:"year"`
	Director     string           `json:"director"`
	Introduction string           `json:"introduction"`
	PlayLineList [][]*EpisodeLink `json:"playLineList"`
}

type MediaPlayInfo struct {
	State    string                 `json:"state"`
	PlayUrl  string                 `json:"playUrl"`
	PlayType constant.MediaPlayType `json:"playType"`
}
