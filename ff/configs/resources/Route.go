package ff_configs_resources

import (
	"fmt"
	"net/http"
)

type Route struct {
	URI            string
	Method         string
	ControllerFunc func(w http.ResponseWriter, r *http.Request)
}

func (this Route) GetPattern() string {
	return fmt.Sprintf("%s %s", this.Method, this.URI)
}
