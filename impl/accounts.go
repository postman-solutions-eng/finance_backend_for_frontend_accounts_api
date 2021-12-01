package impl

import (
  "net/http"
  "Accounts/types")

type AccountsImpl struct {}

                    // Create
                // Post /accounts/create
func (accounts *AccountsImpl) Create(w http.ResponseWriter, r *http.Request) {
  // Implement your logic here
 
  w.WriteHeader(200) 
}

                    // Overview
                // Get /accounts/{accountNumber}/overview
func (accounts *AccountsImpl) Overview(w http.ResponseWriter, r *http.Request, params types.OverviewParams) {
  // Implement your logic here
 
  w.WriteHeader(200) 
}

                    // Date
                // Get /accounts/{accountNumber}/statement/date
func (accounts *AccountsImpl) Date(w http.ResponseWriter, r *http.Request, params types.DateParams) {
  // Implement your logic here
  w.Header().Set("Content-Type", "text/plain")

  w.WriteHeader(200)                                                                 
}

                    // Latest
                // Get /accounts/{accountNumber}/statement/latest
func (accounts *AccountsImpl) Latest(w http.ResponseWriter, r *http.Request, params types.LatestParams) {
  // Implement your logic here
 
  w.WriteHeader(200) 
}
