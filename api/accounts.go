package api

import (
  "encoding/json"
  "fmt"
  "github.com/go-chi/chi/v5"
  "net/http"
  "Accounts/types"
)

// ServerInterface represents all server handlers.
type AccountsStub interface {
                // Create
            Create(w http.ResponseWriter, r *http.Request)

                // Overview
            Overview(w http.ResponseWriter, r *http.Request, params types.OverviewParams)

                // Date
            Date(w http.ResponseWriter, r *http.Request, params types.DateParams)

                // Latest
            Latest(w http.ResponseWriter, r *http.Request, params types.LatestParams)

}

type AccountsWrapper struct {
  AccountsDelegate AccountsStub
}


                // Create
            func(stub *AccountsWrapper) Create(w http.ResponseWriter, r *http.Request) {


    decoder := json.NewDecoder(r.Body)
    var body types.CreateInlineReqJson
    err := decoder.Decode(&body)

    if err != nil {
      fmt.Println(err)
    }

  
  stub.AccountsDelegate.Create(w, r)
}

                // Overview
            func(stub *AccountsWrapper) Overview(w http.ResponseWriter, r *http.Request) {


  overviewParams := types.OverviewParams {
    AccountNumber: chi.URLParam(r, "AccountNumber"),
  }  

  stub.AccountsDelegate.Overview(w, r, overviewParams)
}

                // Date
            func(stub *AccountsWrapper) Date(w http.ResponseWriter, r *http.Request) {


  dateParams := types.DateParams {
    AccountNumber: chi.URLParam(r, "AccountNumber"),
  }  

  stub.AccountsDelegate.Date(w, r, dateParams)
}

                // Latest
            func(stub *AccountsWrapper) Latest(w http.ResponseWriter, r *http.Request) {


  latestParams := types.LatestParams {
    AccountNumber: chi.URLParam(r, "AccountNumber"),
  }  

  stub.AccountsDelegate.Latest(w, r, latestParams)
}
