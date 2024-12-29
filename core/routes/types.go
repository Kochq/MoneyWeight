package routes

type Transaction struct {
    ID                int      `json:"id"`
    Title            string    `json:"title"`
    Amount           float64   `json:"amount"`
    CategoryID       int       `json:"category_id"`     // Nombres consistentes
    SubCategoryID    int       `json:"subcategory_id"`
    Currency         string    `json:"currency"`
    PaymentMethod    string    `json:"payment_method"`
    ExchangeRate     float64   `json:"exchange_rate"`
    Notes            string    `json:"notes"`
    Date             string    `json:"date"`            // Podríamos usar time.Time
    InstallmentPlanID *int     `json:"installment_plan_id,omitempty"`
    RecurringPaymentID *int    `json:"recurring_payment_id,omitempty"`
    PaymentNumber    *int      `json:"payment_number,omitempty"`
}
