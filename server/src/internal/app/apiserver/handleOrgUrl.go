package apiserver

import (
	"net/http"
	// "github.com/google/go-github/github"
	// "context"
)

func (s *APIServer) HandleOrgUrl() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(("Hello on Hello page")))
	}
}
