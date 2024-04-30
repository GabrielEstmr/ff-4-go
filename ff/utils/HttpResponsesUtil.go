package ff_utils

import (
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways_ws_resources "baseapplicationgo/main/configs/ff/lib/gateways/ws/resources"
	"encoding/json"
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
