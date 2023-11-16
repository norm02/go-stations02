package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mileusna/useragent"
)

type storedKey struct{}
var k storedKey

func StoreOS(h http.Handler)http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := r.UserAgent()
		os := useragent.Parse(ua).OS

		ctx := context.WithValue(r.Context(), k, os)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func CtxOS(ctx context.Context) (string,error){
	v := ctx.Value(k)
	if v == nil {
		return "", fmt.Errorf("os not found: %s", k)
	}
	os, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("invalid value for key %s: %T", k, v)
	}
	return os, nil
}