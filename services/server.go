package services

import (
    "context"
    "net/http"

    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
    r := mux.NewRouter()
    r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

    r.Methods("POST").Path("/execution").Handler(httptransport.NewServer(
        endpoints.ExecutionMessageEndpoint,
        decodeExecutionMessageRequest,
        encodeResponse,
    ))

    r.Methods("POST").Path("/transaction").Handler(httptransport.NewServer(
        endpoints.TransactionMessageEndpoint,
        decodeTransactionMessageRequest,
        encodeResponse,
    ))

    r.Methods("POST").Path("/text").Handler(httptransport.NewServer(
        endpoints.TextMessageEndpoint,
        decodeTextMessageRequest,
        encodeResponse,
    ))

    return r
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

