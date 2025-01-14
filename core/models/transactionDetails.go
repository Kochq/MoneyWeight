package models

import "database/sql"

type TransactionDetails struct {
	Transaction
	// Campos que siempre estarán presentes
	CategoryName    string `json:"category_name"`
	CategoryIcon    string `json:"category_icon"`
	FromAccountName string `json:"from_account_name"`
	FromAccountType string `json:"from_account_type"`

	// Campos opcionales
	SubCategoryName   *string `json:"subcategory_name,omitempty"`
	SubCategoryIcon   *string `json:"subcategory_icon,omitempty"`
	ToAccountName     *string `json:"to_account_name,omitempty"`
	ToAccountType     *string `json:"to_account_type,omitempty"`
	InstallmentTitle  *string `json:"installment_title,omitempty"`
	TotalInstallments *int    `json:"total_installments,omitempty"`
	RecurringTitle    *string `json:"recurring_title,omitempty"`
}

func (td TransactionDetails) GetSelectQuery() string {
	return `
    SELECT t.id, t.title, t.amount, t.category_id, t.from_account_id,
    t.to_account_id, t.subcategory_id, t.currency, t.payment_method,
    t.exchange_rate, t.notes, t.date, t.status, t.installment_plan_id,
    t.recurring_payment_id, t.payment_number,

    -- Campos obligatorios
    c.name as category_name, c.icon as category_icon,
    sc.name as subcategory_name, sc.icon as subcategory_icon,
    fa.name as from_account_name, fa.type as from_account_type,

    -- Campos opcionales
    ta.name as to_account_name, ta.type as to_account_type,
    ip.title as installment_title, ip.total_installments,
    rp.title as recurring_title

    FROM Transactions t

    -- JOINs obligatorios
    INNER JOIN Categories c ON t.category_id = c.id
    INNER JOIN Accounts fa ON t.from_account_id = fa.id

    -- JOINs opcionales
    LEFT JOIN SubCategories sc ON t.subcategory_id = sc.id
    LEFT JOIN Accounts ta ON t.to_account_id = ta.id
    LEFT JOIN InstallmentPlans ip ON t.installment_plan_id = ip.id
    LEFT JOIN RecurringPayments rp ON t.recurring_payment_id = rp.id
    ORDER BY t.date DESC
    LIMIT ? OFFSET ?`
}

func (td *TransactionDetails) Scan(rows *sql.Rows) error {
	// Variables temporales para campos nullables
	var (
		toAccountID        sql.NullInt64
		installmentPlanID  sql.NullInt64
		recurringPaymentID sql.NullInt64
		paymentNumber      sql.NullInt64

		subcategoryName   sql.NullString
		subcategoryIcon   sql.NullString
		toAccountName     sql.NullString
		toAccountType     sql.NullString
		installmentTitle  sql.NullString
		totalInstallments sql.NullInt64
		recurringTitle    sql.NullString
	)

	err := rows.Scan(
		&td.ID, &td.Title, &td.Amount, &td.CategoryID, &td.FromAccountID,
		&toAccountID, &td.SubCategoryID, &td.Currency, &td.PaymentMethod,
		&td.ExchangeRate, &td.Notes, &td.Date, &td.Status, &installmentPlanID,
		&recurringPaymentID, &paymentNumber,

		// Campos obligatorios de relaciones
		&td.CategoryName, &td.CategoryIcon, &td.FromAccountName,
		&td.FromAccountType,

		// Campos opcionales de relaciones
		&subcategoryName, &subcategoryIcon, &toAccountName, &toAccountType,
		&installmentTitle, &totalInstallments,
		&recurringTitle,
	)

	if err != nil {
		return err
	}

	// Asignar los valores nullables solo si son válidos
	if toAccountID.Valid {
		td.ToAccountID = &toAccountID.Int64
	}
	if installmentPlanID.Valid {
		temp := int(installmentPlanID.Int64)
		td.InstallmentPlanID = &temp
	}
	if recurringPaymentID.Valid {
		temp := int(recurringPaymentID.Int64)
		td.RecurringPaymentID = &temp
	}
	if paymentNumber.Valid {
		temp := int(paymentNumber.Int64)
		td.PaymentNumber = &temp
	}

	// Asignar campos opcionales de las relaciones
	if subcategoryName.Valid {
		td.SubCategoryName = &subcategoryName.String
	}
	if subcategoryIcon.Valid {
		td.SubCategoryIcon = &subcategoryIcon.String
	}
	if toAccountName.Valid {
		td.ToAccountName = &toAccountName.String
	}
	if toAccountType.Valid {
		td.ToAccountType = &toAccountType.String
	}
	if installmentTitle.Valid {
		td.InstallmentTitle = &installmentTitle.String
	}
	if totalInstallments.Valid {
		temp := int(totalInstallments.Int64)
		td.TotalInstallments = &temp
	}
	if recurringTitle.Valid {
		td.RecurringTitle = &recurringTitle.String
	}

	return nil
}
