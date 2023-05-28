package server

import (
	"backups/internal/backup"
	"backups/pkg/model"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	Manager *backup.BackupManager
}

type BackupResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewHandler(back *backup.BackupManager) (*Handler, error) {
	return &Handler{Manager: back}, nil
}

func (h *Handler) backupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Backup request")
	model.CreatedBackups.Inc()
	dumpName := r.FormValue("dump")

	log.Printf("dumpName %s", dumpName)

	if err := h.Manager.UploadDump(dumpName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	backupResponse := BackupResponse{
		Message: "Backup endpoint called successfully",
		Code:    http.StatusOK,
	}
	responseJSON, err := json.Marshal(backupResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(backupResponse.Code)
	w.Write(responseJSON)
}
