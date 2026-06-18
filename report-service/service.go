package main

import "context"

type ReportRepository interface {
	GetDailyReport(ctx context.Context, date string) (*DailyReport, error)
	GetProblems(ctx context.Context) ([]ProblemPackage, error)
	GetCourierPerformance(ctx context.Context, courierID int) (*CourierPerformance, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return s.repo.GetDailyReport(ctx, date)
}

func (s *ReportService) GetProblems(ctx context.Context) ([]ProblemPackage, error) {
	return s.repo.GetProblems(ctx)
}

func (s *ReportService) GetCourierPerformance(
	ctx context.Context,
	courierID int,
) (*CourierPerformance, error) {
	return s.repo.GetCourierPerformance(ctx, courierID)
}
