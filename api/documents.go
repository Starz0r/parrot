package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anthonynsimon/parrot/errors"
	"github.com/anthonynsimon/parrot/model"
	"github.com/gorilla/mux"
)

func CreateDocument(w http.ResponseWriter, r *http.Request) (int, error) {
	projID, err := strconv.Atoi(mux.Vars(r)["projectID"])
	if err != nil {
		return http.StatusBadRequest, err
	}

	doc := &model.Document{}
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		return http.StatusBadRequest, err
	}

	doc.ProjectID = projID

	proj, err := store.GetProject(projID)
	if err != nil {
		if err == errors.ErrNotFound {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}

	doc.SyncKeys(proj.Keys, true)

	err = store.CreateDoc(doc)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return writeJSON(w, http.StatusOK, doc)
}

func ShowDocument(w http.ResponseWriter, r *http.Request) (int, error) {
	projID, err := strconv.Atoi(mux.Vars(r)["projectID"])
	if err != nil {
		return http.StatusBadRequest, err
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return http.StatusBadRequest, err
	}

	doc, err := store.GetProjectDoc(projID, id)
	if err != nil {
		if err == errors.ErrNotFound {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}

	return writeJSON(w, http.StatusOK, doc)
}

func UpdateDocument(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return http.StatusBadRequest, err
	}

	doc := &model.Document{}
	if err := json.NewDecoder(r.Body).Decode(&doc.Pairs); err != nil {
		return http.StatusBadRequest, err
	}
	doc.ID = id

	err = store.UpdateDoc(doc)
	if err != nil {
		if err == errors.ErrNotFound {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}

	return writeJSON(w, http.StatusOK, doc)
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return http.StatusBadRequest, err
	}

	resultID, err := store.DeleteDoc(id)
	if err != nil {
		if err == errors.ErrNotFound {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}

	return writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("deleted document with id %d", resultID),
	})
}
