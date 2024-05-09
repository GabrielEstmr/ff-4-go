package ff_utils

import (
	"encoding/json"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways_ws_resources "github.com/GabrielEstmr/ff-4-go/ff/gateways/ws/resources"
	"log"
	"net/http"
	"strconv"
)

const _ERROR_UTILS_MSG_ARCH_ISSUE = "Architecture application issue"

type HttpResponsesUtil struct {
}

func NewHttpResponsesUtil() *HttpResponsesUtil {
	return &HttpResponsesUtil{}
}

func (this HttpResponsesUtil) JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		json.NewEncoder(w).Encode(
			ff_gateways_ws_resources.NewErrorResponseSlgMsg(
				strconv.Itoa(statusCode),
				_ERROR_UTILS_MSG_ARCH_ISSUE))
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func (this HttpResponsesUtil) ERROR_APP(
	w http.ResponseWriter,
	appException ff_domains_exceptions.LibException,
) {
	r := ff_gateways_ws_resources.NewErrorResponse(
		strconv.Itoa(appException.GetCode()), appException.GetMessages())
	this.JSON(w, appException.GetCode(), r)
}
