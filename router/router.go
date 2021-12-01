package router

import (
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/cors"
  "net/http"
  "Accounts/api"
  "Accounts/impl"
)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler() http.Handler {
    serviceImpl := api.AccountsDelegate{
            AccountsDelegate: &impl.AccountsImpl{},
    }

  return RouterHandler(serviceImpl, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func RouterHandler(serviceImpl api.AccountsDelegate, r chi.Router) http.Handler {
    // This is here for local dev server development where the client consuming the API is NOT the same IP
    // and/or PORT as the server is running on. This allows local development from say, a web UI that runs
    // on a different port, accessing the API, so a reverse proxy is not needed.
    cors := cors.New(cors.Options{
        AllowedOrigins:     []string{"*"},
        AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
        AllowedHeaders:     []string{"Accept", "Authorization", "Content-Length", "Cache-Control", "Accept-Encoding", "Content-Type", "X-CSRF-Token"},
        AllowCredentials:   true,
        MaxAge:             300, // Maximum value not ignored by any of major browsers
        OptionsPassthrough: false,
        Debug:              true,
    })
    r.Use(cors.Handler)

    accountsWrapper := api.AccountsWrapper {
      AccountsDelegate: serviceImpl.AccountsDelegate,
    }

    r.Group(func(r chi.Router) {
      r.Post("/accounts/create", accountsWrapper.Create)
      r.Get("/accounts/{accountNumber}/overview", accountsWrapper.Overview)
      r.Get("/accounts/{accountNumber}/statement/date", accountsWrapper.Date)
      r.Get("/accounts/{accountNumber}/statement/latest", accountsWrapper.Latest)
    })

    return r
}
