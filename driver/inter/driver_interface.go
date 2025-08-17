package inter

type Driver interface {
	DriverCode() string
	InvokeSearchFuzzy(keyword string) string
	InvokeSearchExact(fullName string) string
}

type RegisterFunc func(driver Driver)

var Register RegisterFunc
