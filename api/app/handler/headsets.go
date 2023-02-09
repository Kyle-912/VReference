package handler

import (
	"encoding/json"
	"net/http"

	"github.com/VReference/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllHeadsets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	headsets := []model.Headset{}
	db.Find(&headsets)
	respondJSON(w, http.StatusOK, headsets)
}

func CreateHeadset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	headset := model.Headset{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&headset); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&headset).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, headset)
}

func GetHeadset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["title"]
	headset := getHeadsetOr404(db, name, w, r)
	if headset == nil {
		return
	}
	respondJSON(w, http.StatusOK, headset)
}

func UpdateHeadset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["title"]
	headset := getHeadsetOr404(db, name, w, r)
	if headset == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&headset); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&headset).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, headset)
}

func DeleteHeadset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["title"]
	headset := getHeadsetOr404(db, name, w, r)
	if headset == nil {
		return
	}
	if err := db.Unscoped().Delete(&headset).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableHeadset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["title"]
	headset := getHeadsetOr404(db, name, w, r)
	if headset == nil {
		return
	}
	headset.Disable()
	if err := db.Save(&headset).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, headset)
}

func EnableHeadset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["title"]
	headset := getHeadsetOr404(db, name, w, r)
	if headset == nil {
		return
	}
	headset.Enable()
	if err := db.Save(&headset).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, headset)
}


func getHeadsetOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Headset {
	headset := model.Headset{}
	otherset := model.Headset{Name: name}
	if err := db.Where(otherset).First(&headset).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &headset
}
