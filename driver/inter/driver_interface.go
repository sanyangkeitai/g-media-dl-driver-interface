package inter

import "github.com/sanyangkeitai/g-media-dl-driver-interface/driver/model"

type Driver interface {
	DriverConf() *model.DriverConfig
	InvokeSearchFuzzy(keyword string) ([]*model.SearchMediaList, error)
	InvokeSearchDetail(fullKeyword string) (*model.MediaDetailInfo, error)
	InvokeSearchMediaPlayInfo(mediaIndexUrl string) (*model.MediaPlayInfo, error)
	InvokeSearchHlsIndexFile(hlsIndexUrl string) (string, error)
}

type RegisterFunc func(driver Driver)

var Register RegisterFunc
