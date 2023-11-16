package middleware

import (
	"crypto/subtle"
	"net/http"
	"os"
)

func Basicauth (h http.Handler)http.Handler{
    fn := func(w http.ResponseWriter, r *http.Request){
    
    osuser:= os.Getenv("BASIC_AUTH_USER_ID") 
    ospass:= os.Getenv("BASIC_AUTH_USER_PASSWORD")
    user,pass,ok := r.BasicAuth()
    if !ok || subtle.ConstantTimeCompare([]byte(user),[]byte(osuser)) !=1 || 
    subtle.ConstantTimeCompare([]byte(pass),[]byte(ospass)) !=1{
		w.WriteHeader(http.StatusUnauthorized)
	}
    h.ServeHTTP(w,r)
}
    return http.HandlerFunc(fn)

}