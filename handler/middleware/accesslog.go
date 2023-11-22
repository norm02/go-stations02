package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type AccessLog struct{
	Timestamp time.Time `json:"timestamp"`
	Latency int64 `json:"latency"`
	Path string `json:"path"`
	OS string `json:"os"`
}
func AccessLogger (h http.Handler)http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		h.ServeHTTP(w,r)
		accesstime := time.Now()
		os,err := CtxOS(r.Context())
		if err!=nil{
			fmt.Println(err)
		}
		accesslog := AccessLog{
			Timestamp: accesstime,
			Latency: time.Since(accesstime).Milliseconds(),
			Path: r.URL.Path,
			OS: os,
		}
		fmt.Println(accesslog)
	}
	return http.HandlerFunc(fn)
}