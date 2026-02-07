package services

import (
	"kasir-api/model"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport() (*model.DailyReport, error) {
	return s.repo.GetDailyReport()
}
