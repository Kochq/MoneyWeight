package routes

type Transaction struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Amount             float64 `json:"amount"`
	CategoryID         int     `json:"category_id"`
	SubCategoryID      int     `json:"subcategory_id"`
	Currency           string  `json:"currency"`
	PaymentMethod      string  `json:"payment_method"`
	ExchangeRate       float64 `json:"exchange_rate"`
	Notes              string  `json:"notes"`
	Date               string  `json:"date"` // Podríamos usar time.Time
	InstallmentPlanID  *int    `json:"installment_plan_id,omitempty"`
	RecurringPaymentID *int    `json:"recurring_payment_id,omitempty"`
	PaymentNumber      *int    `json:"payment_number,omitempty"`
}

type Installment struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	TotalAmount        float64 `json:"total_amount"`
	TotalInstallments  int     `json:"total_installments"`
	InstallmentsAmount float64 `json:"installment_amount"`
	StartDate          string  `json:"start_date"` // Podríamos usar time.Time
	PayDate            string  `json:"pay_date"`   // Podríamos usar time.Time
	Status             string  `json:"status"`
	CategoryID         int     `json:"category_id"`
	SubCategoryID      int     `json:"subcategory_id"`
}

type RecurringPayment struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Amount        float64 `json:"amount"`
	CategoryID    int     `json:"category_id"`
	SubCategoryID int     `json:"subcategory_id"`
	IsActive      bool    `json:"is_active"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	Frequency     string  `json:"frequency"`
}
