package http

import "net/http"

func (s *Server) swagger() {
	fs := http.FileServer(http.Dir("www/swagger-ui"))

	s.httpMux.HandleFunc("/swagger.json", serveSwagger)
	s.httpMux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger-ui/swagger.json")
}
