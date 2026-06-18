package main

import (
	"encoding/json"
	"net/http"
)

type ReportHandler struct {
	svc *ReportService
}

func NewReportHandler(svc *ReportService) *ReportHandler {
	return &ReportHandler{svc: svc}
}

func (h *ReportHandler) DailyReport(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if date == "" {
		http.Error(w, "date required", http.StatusBadRequest)
		return
	}
	report, err := h.svc.GetDailyReport(r.Context(), date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) ProblemsReport(
	w http.ResponseWriter,
	r *http.Request,
) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, err := h.svc.GetProblems(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *ReportHandler) CourierPerformance(

	w http.ResponseWriter,
	r *http.Request,
) {

	id := 1
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := h.svc.GetCourierPerformance(
		r.Context(),
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
