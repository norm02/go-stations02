package middleware

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Basicauth (h http.Handler, realm string)http.Handler{
    fn := func(w http.ResponseWriter, r *http.Request){

    err := godotenv.Load(".env")
    if err!=nil{
        fmt.Println("envファイルを読み込みできませんでした。")
    }
    
    osuser:= os.Getenv("BASIC_AUTH_USER_ID") 
    ospass:= os.Getenv("BASIC_AUTH_USER_PASSWORD")

    user,pass,ok := r.BasicAuth()
    //アクセス時に User ID, Password を送信しなかった場合、Basic 認証が失敗し HTTP Status Code が 401 で返却されているかどうか。
    if !ok || 
    //空の User ID, Password を送信した場合、 Basic 認証が失敗し HTTP Status Code が 401 で返却されているかどうか。
    user == "" || pass == "" ||
    //間違った User ID, Password を送信した場合、 Basic 認証が失敗しHTTP Status Code が 401 で返却されているかどうか。
    subtle.ConstantTimeCompare([]byte(user),[]byte(osuser)) !=1 || 
    subtle.ConstantTimeCompare([]byte(pass),[]byte(ospass)) !=1{
        w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("userid or password is denied"))
        return
	}
    //対象のAPIのみBasic認証をクリアし、アクセスできるかどうか
    h.ServeHTTP(w,r)
}
    return http.HandlerFunc(fn)

}