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
//next handlerできていないので実装する
func AccessLogger (h http.Handler)http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		accesstime := time.Now()
		h.ServeHTTP(w,r)
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