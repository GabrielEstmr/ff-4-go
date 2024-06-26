/*
 * Copyright (c) 2024. Licensed under the Apache License, Version 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 */

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

type RolloutController struct {
	createRollout           ff_usecases_interfaces.CreateRollout
	updateRollout           ff_usecases_interfaces.UpdateRollout
	deleteRollout           ff_usecases_interfaces.DeleteRollout
	findRolloutById         ff_usecases_interfaces.FindRolloutById
	addTargetToRollout      ff_usecases_interfaces.AddTargetToRollout
	removeTargetFromRollout ff_usecases_interfaces.RemoveTargetFromRollout
	enableRolloutToAll      ff_usecases_interfaces.EnableRolloutToAll
	disableRolloutToAll     ff_usecases_interfaces.DisableRolloutToAll
	isTargetInRollout       ff_usecases_interfaces.VerifyIsEnabledAllOrTargetInRollout
	httpResponsesUtil       ff_utils.HttpResponsesUtil
}

func NewRolloutController(
	createRollout ff_usecases_interfaces.CreateRollout,
	updateRollout ff_usecases_interfaces.UpdateRollout,
	deleteRollout ff_usecases_interfaces.DeleteRollout,
	findRolloutById ff_usecases_interfaces.FindRolloutById,
	addTargetToRollout ff_usecases_interfaces.AddTargetToRollout,
	removeTargetFromRollout ff_usecases_interfaces.RemoveTargetFromRollout,
	enableRolloutToAll ff_usecases_interfaces.EnableRolloutToAll,
	disableRolloutToAll ff_usecases_interfaces.DisableRolloutToAll,
	isTargetInRollout ff_usecases_interfaces.VerifyIsEnabledAllOrTargetInRollout,
) *RolloutController {
	return &RolloutController{
		createRollout:           createRollout,
		updateRollout:           updateRollout,
		deleteRollout:           deleteRollout,
		findRolloutById:         findRolloutById,
		addTargetToRollout:      addTargetToRollout,
		removeTargetFromRollout: removeTargetFromRollout,
		enableRolloutToAll:      enableRolloutToAll,
		disableRolloutToAll:     disableRolloutToAll,
		isTargetInRollout:       isTargetInRollout,
		httpResponsesUtil:       *ff_utils.NewHttpResponsesUtil(),
	}
}

func (this *RolloutController) CreateRollout(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	var rolloutRequest ff_gateways_ws_resources.RolloutRequest
	if err = json.Unmarshal(requestBody, &rolloutRequest); err != nil {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	rollout := rolloutRequest.ToDomain()

	createdRollout, errApp := this.createRollout.Execute(rollout)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusCreated, *ff_gateways_ws_resources.NewRolloutResponse(createdRollout))
}

func (this *RolloutController) UpdateRollout(w http.ResponseWriter, r *http.Request) {

	key := r.PathValue("key")
	requestBody, err := io.ReadAll(r.Body)
	if err != nil || len(requestBody) == 0 {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	var rolloutRequest ff_gateways_ws_resources.UpdateRolloutRequest
	if err = json.Unmarshal(requestBody, &rolloutRequest); err != nil {
		this.httpResponsesUtil.ERROR_APP(
			w,
			ff_domains_exceptions.NewBadRequestExceptionSglMsg("invalid request body"),
		)
		return
	}

	rollout := rolloutRequest.ToDomain(key)
	updatedRollout, errApp := this.updateRollout.Execute(rollout)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, *ff_gateways_ws_resources.NewRolloutResponse(updatedRollout))
}

func (this *RolloutController) DeleteRollout(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	errApp := this.deleteRollout.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusNoContent, nil)
}

func (this *RolloutController) FindRolloutByKey(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	rollout, errApp := this.findRolloutById.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewRolloutResponse(rollout))
}

func (this *RolloutController) AddTargetToRollout(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	target := r.PathValue("target")
	rollout, errApp := this.addTargetToRollout.Execute(key, target)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewRolloutResponse(rollout))
}

func (this *RolloutController) RemoveTargetFromRollout(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	target := r.PathValue("target")
	rollout, errApp := this.removeTargetFromRollout.Execute(key, target)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewRolloutResponse(rollout))
}

func (this *RolloutController) EnableToAll(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	rollout, errApp := this.enableRolloutToAll.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewRolloutResponse(rollout))
}

func (this *RolloutController) DisableToAll(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	rollout, errApp := this.disableRolloutToAll.Execute(key)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, ff_gateways_ws_resources.NewRolloutResponse(rollout))
}

func (this *RolloutController) VerifyIsTargetInRollout(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	target := r.PathValue("target")
	isTargetInRollout, errApp := this.isTargetInRollout.Execute(key, target)
	if errApp != nil {
		this.httpResponsesUtil.ERROR_APP(w, errApp)
		return
	}
	this.httpResponsesUtil.JSON(w, http.StatusOK, isTargetInRollout)
}
