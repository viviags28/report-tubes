package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
)

type MySQLReportRepository struct {
	db *sql.DB
}

func (r MySQLReportRepository) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return &DailyReport{
		TotalPaket:  10,
		Delivered:   8,
		Pending:     1,
		Terlambat:   1,
		RataRataETA: 2.5,
	}, nil
}

func (r MySQLReportRepository) GetProblems(ctx context.Context) ([]ProblemPackage, error) {
	return []ProblemPackage{
		{
			Resi:   "RESI001",
			Status: "terlambat",
		},
		{
			Resi:   "RESI002",
			Status: "pending",
		},
	}, nil
}

func (r MySQLReportRepository) GetCourierPerformance(
	ctx context.Context,
	courierID int,
) (*CourierPerformance, error) {

	return &CourierPerformance{
		CourierID:       courierID,
		TotalPengiriman: 100,
		Berhasil:        95,
		Terlambat:       5,
		Score:           95,
	}, nil
}

func main() {

	ConnectDB()

	repo := MySQLReportRepository{db: DB}
	svc := NewReportService(repo)
	h := NewReportHandler(svc)

	http.HandleFunc("/report/daily", h.DailyReport)
	http.HandleFunc("/report/problems", h.ProblemsReport)
	http.HandleFunc("/report/courier-performance", h.CourierPerformance)

	log.Println("Running on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
