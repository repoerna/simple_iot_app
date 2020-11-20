package handler

import (
	"context"
	"encoding/json"
	"net/http"
	db "simpleiotapp/db/sqlc"
	"simpleiotapp/util"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler ...
type Handler struct {
	Queries *db.Queries
}

// GetDevice ...
func (s *Handler) GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	device, err := s.Queries.GetDevice(context.Background(), int64(id))
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(w, http.StatusOK, device)
}

// GetDevices ...
func (s *Handler) GetDevices(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.ParseInt(r.FormValue("limit"), 10, 32)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid Limit")
		return
	}
	offset, err := strconv.ParseInt(r.FormValue("offset"), 10, 32)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid Offset")
		return
	}

	arg := db.GetDevicesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	devices, err := s.Queries.GetDevices(context.Background(), arg)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(w, http.StatusOK, devices)
}

// UpdateDevice ...
func (s *Handler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var device *db.UpdateDeviceParams
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&device); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	arg := db.UpdateDeviceParams{
		ID:        int64(id),
		Name:      device.Name,
		Shortname: device.Shortname,
		Enabled:   device.Enabled,
	}

	res, err := s.Queries.UpdateDevice(context.Background(), arg)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, res)
}

// DeleteDevice ...
func (s *Handler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = s.Queries.DeleteDevice(context.Background(), int64(id))
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithMessage(w, http.StatusOK, "device deleted")
}

// CreateDevice ...
func (s *Handler) CreateDevice(w http.ResponseWriter, r *http.Request) {

	var device *db.CreateDeviceParams
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&device); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	arg := db.CreateDeviceParams{
		Name:      device.Name,
		Shortname: device.Shortname,
		Enabled:   device.Enabled,
	}

	res, err := s.Queries.CreateDevice(context.Background(), arg)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(w, http.StatusOK, res)

}
