export interface Transaction {
    id: number;
    title: string;
    amount: number;
    category_id: number;
    subcategory_id: number;
    currency: string;
    payment_method: string;
    exchange_rate: number;
    notes: string;
    date: string; // 2025-01-01 00:00:00
    installment_plan_id: number;
    recurring_payment_id: number;
    payment_number: number;

    category_name: string;
    category_icon: string;
    from_account_name: string;
    from_account_type: string;

    // Campos opcionales
    sub_category_name: string;
    sub_category_icon: string;
    to_account_name: string;
    to_account_type: string;
    installment_title: string;
    total_installments: number;
    recurring_title: string;
}

export interface SubCategory {
    id: number;
    name: string;
    icon: string;
    category_id: number;
}

export interface Category {
    id: number;
    name: string;
    type: string;
    icon: string;
    subcategories: SubCategory[];
}

export interface Accounts {
    id: number;
    name: string;
    type: string;
    current_balance: number;
    currency: string;
    Institution: string;
    is_active: boolean;
    created_at: string;
}
