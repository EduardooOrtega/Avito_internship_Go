package data

import "database/sql"

type Transaction struct {
	TransactionId   int64   `json:"transaction_id"`
	SenderId        int64   `json:"sender_id"`
	ReceiverId      int64   `json:"receiver_id"`
	OperationId     int64   `json:"operation_id"`
	TransactionTime string  `json:"transaction_time"`
	Price           float64 `json:"price"`
}

type TransactionHistory struct {
	TransactionId    int64   `json:"transaction_id"`
	SenderId         int64   `json:"sender_id"`
	ReceiverId       int64   `json:"receiver_id"`
	TransactionTime  string  `json:"transaction_time"`
	TransactionPrice float64 `json:"transaction_price"`
	OperationType    string  `json:"operation_type"`
}

type TransactionHistoryResponse struct {
	SenderId         int64   `json:"sender_id"`
	ReceiverId       int64   `json:"receiver_id"`
	TransactionTime  string  `json:"transaction_time"`
	TransactionPrice float64 `json:"transaction_price"`
	OperationType    string  `json:"operation_type"`
}

type Account struct {
	AccountId           int64   `json:"account_id"`
	AccountCash         float64 `json:"account_cash"`
	AccountReservedCash float64 `json:"account_reserved_cash"`
}

type Service struct {
	ServiceId    int64   `json:"service_id"`
	ServiceName  string  `json:"service_name"`
	ServicePrice float64 `json:"service_price"`
}

type Filters struct {
	Page     int
	PageSize int
	Sort     string
}

type Metadata struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	FirstPage    int `json:"first_page"`
	LastPage     int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

type Report struct {
	AccountId  int64  `json:"account_id"`
	ReportId   int64  `json:"report_id"`
	ServiceId  int64  `json:"service_id"`
	ReportTime string `json:"report_time"`
}

type ReportResult struct {
	ServiceId int64  `json:"service_id"`
	Revenue   string `json:"report_revenue"`
}

type UserReportHistoryResponse struct {
	AccountId    int64   `json:"account_id"`
	ReportTime   string  `json:"report_time"`
	ServicePrice float64 `json:"service_price"`
	ServiceName  string  `json:"service_name"`
}

// Models

type AccountModel struct {
	DB *sql.DB
}

type ReportModel struct {
	DB *sql.DB
}

type TransactionHistoryModel struct {
	DB *sql.DB
}

type TransactionModel struct {
	DB *sql.DB
}

type ServiceModel struct {
	DB *sql.DB
}

//-------------------------------

type Models struct {
	Account            AccountModel
	Report             ReportModel
	TransactionHistory TransactionHistoryModel
	Transaction        TransactionModel
	Service            ServiceModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Account:            AccountModel{DB: db},
		Report:             ReportModel{DB: db},
		TransactionHistory: TransactionHistoryModel{DB: db},
		Transaction:        TransactionModel{DB: db},
		Service:            ServiceModel{DB: db},
	}
}
