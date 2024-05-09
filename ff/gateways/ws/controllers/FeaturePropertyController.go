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

type FeaturePropertyController struct {
	createFeatureProperty        ff_usecases_interfaces.CreateFeatureProperty
	updateFeatureProperty        ff_usecases_interfaces.UpdateFeatureProperty
	deleteFeatureProperty        ff_usecases_interfaces.DeleteFeatureProperty
	findFeaturePropertyById      ff_usecases_interfaces.FindFeaturePropertyById
	addValueToFeatureProperty    ff_usecases_interfaces.AddValueToFeatureProperty
	removeValueToFeatureProperty ff_usecases_interfaces.RemoveValueToFeatureProperty
	enableFeatureProperty        ff_usecases_interfaces.EnableFeatureProperty
	disableFeatureProperty       ff_usecases_interfaces.DisableFeatureProperty
	httpResponsesUtil            ff_utils.HttpResponsesUtil
}

func NewFeaturePropertyController(
	createFeatureProperty ff_usecases_interfaces.CreateFeatureProperty,
	updateFeatureProperty ff_usecases_interfaces.UpdateFeatureProperty,
	deleteFeatureProperty ff_usecases_interfaces.DeleteFeatureProperty,
	findFeaturePropertyById ff_usecases_interfaces.FindFeaturePropertyById,
	addValueToFeatureProperty ff_usecases_interfaces.AddValueToFeatureProperty,
	removeValueToFeatureProperty ff_usecases_interfaces.RemoveValueToFeatureProperty,
	enableFeatureProperty ff_usecases_interfaces.EnableFeatureProperty,
	disableFeatureProperty ff_usecases_interfaces.DisableFeatureProperty,
) *FeaturePropertyController {
	return &FeaturePropertyController{
		createFeatureProperty:        createFeatureProperty,
		updateFeatureProperty:        updateFeatureProperty,
		deleteFeatureProperty:        deleteFeatureProperty,
		findFeaturePropertyById:      findFeaturePropertyById,
		addValueToFeatureProperty:    addValueToFeatureProperty,
		removeValueToFeatureProperty: removeValueToFeatureProperty,
		enableFeatureProperty:        enableFeatureProperty,
		disableFeatureProperty:       disableFeatureProperty,
		httpResponsesUtil:            *ff_utils.NewHttpResponsesUtil(),
	}
}

func (this *FeaturePropertyController) CreateFeatureProperty(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	var featurePropertyRequest ff_gateways_ws_resources.FeaturePropertyRequest
	if err = json.Unmarshal(requestBody, &featurePropertyRequest); err != nil {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	featureProperty := featurePropertyRequest.ToDomain()

	createdProperty, errApp := this.createFeatureProperty.Execute(featureProperty)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusCreated, *ff_gateways_ws_resources.NewFeaturePropertyResponse(createdProperty))
}

func (this *FeaturePropertyController) UpdateFeatureProperty(w http.ResponseWriter, r *http.Request) {

	key := r.PathValue("key")
	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	var featurePropertyRequest ff_gateways_ws_resources.UpdateFeaturePropertyRequest
	if err = json.Unmarshal(requestBody, &featurePropertyRequest); err != nil {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	featureProperty := featurePropertyRequest.ToDomain(key)
	updatedFeatureProperty, errApp := this.updateFeatureProperty.Execute(featureProperty)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK,
		*ff_gateways_ws_resources.NewFeaturePropertyResponse(updatedFeatureProperty))
}

func (this *FeaturePropertyController) DeleteFeatureProperty(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	errApp := this.deleteFeatureProperty.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusNoContent, nil)
}

func (this *FeaturePropertyController) FindFeaturePropertyById(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	feature, errApp := this.findFeaturePropertyById.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK,
		ff_gateways_ws_resources.NewFeaturePropertyResponse(feature))
}

func (this *FeaturePropertyController) AddValueToFeatureProperty(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	value := r.PathValue("value")
	updatedFeatureProperty, errApp := this.addValueToFeatureProperty.Execute(key, value)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK,
		ff_gateways_ws_resources.NewFeaturePropertyResponse(updatedFeatureProperty))
}

func (this *FeaturePropertyController) RemoveValueToFeatureProperty(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	value := r.PathValue("value")
	updatedFeatureProperty, errApp := this.removeValueToFeatureProperty.Execute(key, value)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK,
		ff_gateways_ws_resources.NewFeaturePropertyResponse(updatedFeatureProperty))
}

func (this *FeaturePropertyController) EnableFeatureProperty(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	updatedFeatureProperty, errApp := this.enableFeatureProperty.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK,
		ff_gateways_ws_resources.NewFeaturePropertyResponse(updatedFeatureProperty))
}

func (this *FeaturePropertyController) DisableFeatureProperty(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	updatedFeatureProperty, errApp := this.disableFeatureProperty.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK,
		ff_gateways_ws_resources.NewFeaturePropertyResponse(updatedFeatureProperty))
}
