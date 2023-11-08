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
				s := fmt.Sprintf("%v", err)
				m := map[string]string{"status": s}
				if err := json.NewEncoder(w).Encode(m); err != nil {
					fmt.Println(err)
				}
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}