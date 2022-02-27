package Middleware

import (
	"net/http"

	"github.com/nemesisdev2000/Nemesis/TeamServer/Jwt"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.RespondWriter, r *httpRequest) {
		if _, ok := r.Header["Token"]; !ok {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token missing"))
			return
		}

		token := r.Header["Token"][0]
		check, err := Jwt.ValidateToken(token, "somepassword")

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Token validation failed"))
			return
		}
		if !check {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Toke Invalid"))
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Authorized Token"))
	})
}
