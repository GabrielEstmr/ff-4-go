package ff_gateways_ws_controllers

import (
	"encoding/json"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways_ws_resources "github.com/GabrielEstmr/ff-4-go/ff/gateways/ws/resources"
	ff_usecases_interfaces "github.com/GabrielEstmr/ff-4-go/ff/usecases/interfaces"
	ff_utils "github.com/GabrielEstmr/ff-4-go/ff/utils"
	"io"
	"net/http"
)

type FeatureController struct {
	createFeature     ff_usecases_interfaces.CreateFeatureFlag
	deleteFeature     ff_usecases_interfaces.DeleteFeatureFlag
	disableFeature    ff_usecases_interfaces.DisableFeatureFlag
	enableFeature     ff_usecases_interfaces.EnableFeatureFlag
	findFeatureByKey  ff_usecases_interfaces.FindFeatureFlagByKey
	isFeatureEnabled  ff_usecases_interfaces.IsFeatureFlagEnabled
	isFeatureDisabled ff_usecases_interfaces.IsFeatureFlagDisabled
	httpResponsesUtil ff_utils.HttpResponsesUtil
}

func NewFeatureController(
	createFeature ff_usecases_interfaces.CreateFeatureFlag,
	deleteFeature ff_usecases_interfaces.DeleteFeatureFlag,
	disableFeature ff_usecases_interfaces.DisableFeatureFlag,
	enableFeature ff_usecases_interfaces.EnableFeatureFlag,
	findFeatureByKey ff_usecases_interfaces.FindFeatureFlagByKey,
	isFeatureEnabled ff_usecases_interfaces.IsFeatureFlagEnabled,
	isFeatureDisabled ff_usecases_interfaces.IsFeatureFlagDisabled,
) *FeatureController {
	return &FeatureController{
		createFeature:     createFeature,
		deleteFeature:     deleteFeature,
		disableFeature:    disableFeature,
		enableFeature:     enableFeature,
		findFeatureByKey:  findFeatureByKey,
		isFeatureEnabled:  isFeatureEnabled,
		isFeatureDisabled: isFeatureDisabled,
		httpResponsesUtil: *ff_utils.NewHttpResponsesUtil(),
	}
}

func (this *FeatureController) CreateFeature(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	var featureRequest ff_gateways_ws_resources.FeatureFlagRequest
	if err = json.Unmarshal(requestBody, &featureRequest); err != nil {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	feature := featureRequest.ToDomain()

	createdFeature, errApp := this.createFeature.Execute(feature)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusCreated, *ff_gateways_ws_resources.NewFeatureResponse(createdFeature))
}

func (this *FeatureController) DeleteFeature(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	errApp := this.deleteFeature.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusNoContent, nil)
}

func (this *FeatureController) DisableFeature(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	feature, errApp := this.disableFeature.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewFeatureResponse(feature))
}

func (this *FeatureController) EnableFeature(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	feature, errApp := this.enableFeature.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewFeatureResponse(feature))
}

func (this *FeatureController) FindFeatureByKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	feature, errApp := this.findFeatureByKey.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewFeatureResponse(feature))
}

func (this *FeatureController) IsFeatureEnabled(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	isEnabled, errApp := this.isFeatureEnabled.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, isEnabled)
}

func (this *FeatureController) IsFeatureDisabled(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	isEnabled, errApp := this.isFeatureDisabled.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, isEnabled)
}
