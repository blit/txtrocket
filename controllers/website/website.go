package website

import (
	"net/http"

	"github.com/blit/txtrocket/controllers"
)

// Index home/root route
func Index(w http.ResponseWriter, req *http.Request) {
	controllers.RenderHTML(w, http.StatusOK, "website/index", nil, "website/layout")
}
