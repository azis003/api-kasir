package repositories

import (
	"database/sql"
	"kasir-api/model"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetDailyReport() (*model.DailyReport, error) {
	var report model.DailyReport

	// 1. Total Revenue & Total Transactions
	queryStats := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(id) 
		FROM transactions 
		WHERE DATE(created_at) = CURRENT_DATE
	`
	err := r.db.QueryRow(queryStats).Scan(&report.TotalRevenue, &report.TotalTransactions)
	if err != nil {
		return nil, err
	}

	// 2. Produks Terlaris (Top 1)
	queryBestSeller := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`
	err = r.db.QueryRow(queryBestSeller).Scan(&report.BestSellingProduct.Name, &report.BestSellingProduct.QtyTerjual)

	if err == sql.ErrNoRows {
		// Jika belum ada penjualan hari ini
		report.BestSellingProduct = model.BestSellingProduct{Name: "-", QtyTerjual: 0}
	} else if err != nil {
		return nil, err
	}

	return &report, nil
}
