package server

import (
	"net/http"
	"net/url"
)

type Server struct {
	Packages []Package
}

type Package struct {
	Vanity url.URL
	Hosted url.URL
}

func (s *Server) Handler() http.Handler {
	return nil
}
