package Controller

import (
	"net/http"
	"path/filepath"
)

func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	const _landingPageDir = "/home/dist/dog-hub-user-interface/browser"
	filePath := filepath.Join(_landingPageDir, r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, filePath)
}
