package model

import "time"

// IncomeTransactionModel 收入交易模型
type IncomeTransactionModel struct {
	BaseModel
	MentorID        string     `json:"mentor_id" gorm:"not null"`
	StudentID       string     `json:"student_id" gorm:"not null"`
	TransactionType string     `json:"transaction_type" gorm:"not null"`
	Amount          float64    `json:"amount" gorm:"type:decimal(10,2);not null"`
	PlatformFee     float64    `json:"platform_fee" gorm:"type:decimal(10,2);not null;default:0"`
	NetIncome       float64    `json:"net_income" gorm:"type:decimal(10,2);not null"`
	Status          string     `json:"status" gorm:"default:'pending'"`
	Description     string     `json:"description"`
	CourseID        *string    `json:"course_id"`
	AppointmentID   *string    `json:"appointment_id"`
	CompletedAt     *time.Time `json:"completed_at"`

	// 关联关系
	Mentor      *Mentor           `json:"mentor,omitempty" gorm:"foreignKey:MentorID"`
	Student     *User             `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Course      *Course           `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	Appointment *AppointmentModel `json:"appointment,omitempty" gorm:"foreignKey:AppointmentID"`
}

// TableName 指定表名
func (IncomeTransactionModel) TableName() string {
	return "income_transactions"
}

// WithdrawalModel 提现模型
type WithdrawalModel struct {
	BaseModel
	MentorID    string     `json:"mentor_id" gorm:"not null"`
	Amount      float64    `json:"amount" gorm:"type:decimal(10,2);not null"`
	Fee         float64    `json:"fee" gorm:"type:decimal(10,2);not null;default:0"`
	NetAmount   float64    `json:"net_amount" gorm:"type:decimal(10,2);not null"`
	Status      string     `json:"status" gorm:"default:'pending'"`
	BankAccount string     `json:"bank_account" gorm:"not null"`
	BankName    string     `json:"bank_name" gorm:"not null"`
	CompletedAt *time.Time `json:"completed_at"`

	// 关联关系
	Mentor *Mentor `json:"mentor,omitempty" gorm:"foreignKey:MentorID"`
}

// TableName 指定表名
func (WithdrawalModel) TableName() string {
	return "withdrawals"
}
