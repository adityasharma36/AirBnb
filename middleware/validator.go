package middleware

import (
	"context"
	"net/http"

	"goMod/dto"
    
	"goMod/utils/helper"
)



func LoginMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            next.ServeHTTP(w, r)
            return
        }

        var payload dto.LoginUserRequestDTO

        if err := helper.ReadJsonBody(r, &payload); err != nil {
            helper.WriteJsonErrResponse(
                w,
                http.StatusBadRequest,
                "Invalid request body",
                err,
            )
            return
        }

        if err := helper.ValidateStruct(payload); err != nil {
            helper.WriteJsonErrResponse(
                w,
                http.StatusBadRequest,
                "Validation failed",
                err,
            )
            return
        }

        ctx := context.WithValue(
            r.Context(),
            dto.PayloadKey,
            &payload,
        )

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func CreateUserMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            next.ServeHTTP(w, r)
            return
        }

        var payload dto.CreateUserRequestDTO

        if err := helper.ReadJsonBody(r, &payload); err != nil {
            helper.WriteJsonErrResponse(
                w,
                http.StatusBadRequest,
                "Invalid request body",
                err,
            )
            return
        }

        if err := helper.ValidateStruct(payload); err != nil {
            helper.WriteJsonErrResponse(
                w,
                http.StatusBadRequest,
                "Validation failed",
                err,
            )
            return
        }

        ctx := context.WithValue(
            r.Context(),
            dto.UserCreateKey,
            &payload,
        )

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
