package main

type DailyReport struct {
	TotalPaket  int     `json:"total_paket"`
	Delivered   int     `json:"delivered"`
	Pending     int     `json:"pending"`
	Terlambat   int     `json:"terlambat"`
	RataRataETA float64 `json:"rata_rata_eta"`
}

type ProblemPackage struct {
	Resi   string `json:"resi"`
	Status string `json:"status"`
}

type CourierPerformance struct {
	CourierID       int     `json:"courier_id"`
	TotalPengiriman int     `json:"total_pengiriman"`
	Berhasil        int     `json:"berhasil"`
	Terlambat       int     `json:"terlambat"`
	Score           float64 `json:"score"`
}
