package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Recover(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: ここに実装する
		// deferをつかって、panic処理より先にrecoverできるようにする
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusNotFound)
				status := fmt.Sprintf("%v", err)
				msg := map[string]string{"status": status}
				m,err:=json.Marshal(msg)
				if err!=nil{
					fmt.Println(m)
				}
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}