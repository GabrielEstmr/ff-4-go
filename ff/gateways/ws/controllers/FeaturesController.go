package ff_gateways_ws_controllers

import (
	ff_configs_resources "baseapplicationgo/main/configs/ff/lib/configs/resources"
	ff_domains_exceptions "baseapplicationgo/main/configs/ff/lib/domains/exceptions"
	ff_gateways_ws_resources "baseapplicationgo/main/configs/ff/lib/gateways/ws/resources"
	ff_usecases_interfaces "baseapplicationgo/main/configs/ff/lib/usecases/interfaces"
	ff_utils "baseapplicationgo/main/configs/ff/lib/utils"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const _FEATURES_CONTROLLER_PATH_SUFFIX_ENABLE = "/enable"
const _FEATURES_CONTROLLER_PATH_SUFFIX_DISABLE = "/disable"
const _FEATURES_CONTROLLER_PATH_SUFFIX_VERIFY_ENABLE = "/verify-enabled"

type FeaturesController struct {
	_FEATURES_CONTROLLER_PATH_PREFIX string
	createFeature                    ff_usecases_interfaces.CreateFeature
	deleteFeature                    ff_usecases_interfaces.DeleteFeature
	disableFeature                   ff_usecases_interfaces.DisableFeature
	enableFeature                    ff_usecases_interfaces.EnableFeature
	findFeatureByKey                 ff_usecases_interfaces.FindFeatureByKey
	isFeatureEnabled                 ff_usecases_interfaces.IsFeatureEnabled
	httpResponsesUtil                ff_utils.HttpResponsesUtil
}

func NewFeaturesController(
	clientArgs ff_configs_resources.FfClientArgs,
	createFeature ff_usecases_interfaces.CreateFeature,
	deleteFeature ff_usecases_interfaces.DeleteFeature,
	disableFeature ff_usecases_interfaces.DisableFeature,
	enableFeature ff_usecases_interfaces.EnableFeature,
	findFeatureByKey ff_usecases_interfaces.FindFeatureByKey,
	isFeatureEnabled ff_usecases_interfaces.IsFeatureEnabled,
) *FeaturesController {
	return &FeaturesController{
		_FEATURES_CONTROLLER_PATH_PREFIX: clientArgs.GetBaseRoutePath() + "/v1/features/",
		createFeature:                    createFeature,
		deleteFeature:                    deleteFeature,
		disableFeature:                   disableFeature,
		enableFeature:                    enableFeature,
		findFeatureByKey:                 findFeatureByKey,
		isFeatureEnabled:                 isFeatureEnabled,
		httpResponsesUtil:                *ff_utils.NewHttpResponsesUtil(),
	}
}

func (this *FeaturesController) CreateFeature(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
	}

	var featureRequest ff_gateways_ws_resources.FeatureRequest
	if err = json.Unmarshal(requestBody, &featureRequest); err != nil {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
	}

	feature := featureRequest.ToDomain()

	createdFeature, errApp := this.createFeature.Execute(feature)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
	}
	this.httpResponsesUtil.JSON(w, http.StatusCreated, *ff_gateways_ws_resources.NewFeatureResponse(createdFeature))
}

func (this *FeaturesController) DeleteFeature(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimSuffix(r.URL.Path, this._FEATURES_CONTROLLER_PATH_PREFIX)
	errApp := this.deleteFeature.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
	}
	this.httpResponsesUtil.JSON(w, http.StatusNoContent, nil)
}

func (this *FeaturesController) DisableFeature(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, this._FEATURES_CONTROLLER_PATH_PREFIX),
		_FEATURES_CONTROLLER_PATH_SUFFIX_DISABLE)
	feature, errApp := this.disableFeature.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewFeatureResponse(feature))
}

func (this *FeaturesController) EnableFeature(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, this._FEATURES_CONTROLLER_PATH_PREFIX),
		_FEATURES_CONTROLLER_PATH_SUFFIX_ENABLE)
	feature, errApp := this.enableFeature.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewFeatureResponse(feature))
}

func (this *FeaturesController) FindFeatureByKey(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, this._FEATURES_CONTROLLER_PATH_PREFIX)
	feature, errApp := this.findFeatureByKey.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewFeatureResponse(feature))
}

func (this *FeaturesController) IsFeatureEnabled(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, this._FEATURES_CONTROLLER_PATH_PREFIX),
		_FEATURES_CONTROLLER_PATH_SUFFIX_VERIFY_ENABLE)
	isEnabled, errApp := this.isFeatureEnabled.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, isEnabled)
}
