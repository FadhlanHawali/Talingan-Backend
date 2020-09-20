package helper

import "github.com/Talingan-Backend/v2/internal/pkg/shortid"


var sid *shortid.Shortid

func init() {
	sid, _ = shortid.New(1, shortid.DefaultABC, 44125)
}

func IdGenerator () string{
	shortid.SetDefault(sid)
	id, _ := shortid.Generate()
	return id
}

