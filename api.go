package main

type ApiConf struct {
	Format string
	Port   int
}

type Api struct {
}

func NewApi(conf *ApiConf) (*Api, error) {
	return nil, nil
}
