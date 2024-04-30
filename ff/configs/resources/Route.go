package ff_configs_resources

import "net/http"

type Route struct {
	URI            string
	Method         string
	ControllerFunc func(w http.ResponseWriter, r *http.Request)
}
