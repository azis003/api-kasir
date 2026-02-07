package model

type DailyReport struct {
	TotalRevenue       int                `json:"total_revenue"`
	TotalTransactions  int                `json:"total_transaksi"`
	BestSellingProduct BestSellingProduct `json:"produk_terlaris"`
}

type BestSellingProduct struct {
	Name       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
