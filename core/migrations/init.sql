CREATE DATABASE IF NOT EXISTS MoneyWeight;
USE MoneyWeight;

-- 1. Accounts (no foreign key dependencies)
CREATE TABLE `Accounts` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `type` enum('cash','bank','digital_wallet','credit_card','investment') DEFAULT NULL,
    `current_balance` decimal(19,2) NOT NULL,
    `currency` varchar(3) DEFAULT NULL,
    `institution` varchar(255) DEFAULT NULL,
    `is_active` tinyint(1) DEFAULT 1,
    `created_at` datetime DEFAULT current_timestamp(),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 2. Categories (no foreign key dependencies)
CREATE TABLE `Categories` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `icon` varchar(255) DEFAULT NULL,
    `type` enum('income','expense','investment') NOT NULL,
    `budget_limit` decimal(10,2) DEFAULT 0.00,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 3. SubCategories (depends on Categories)
CREATE TABLE `SubCategories` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `category_id` int(11) NOT NULL,
    `name` varchar(50) NOT NULL,
    `icon` varchar(255) DEFAULT NULL,
    `budget_limit` decimal(10,2) DEFAULT 0.00,
    PRIMARY KEY (`id`),
    KEY `category_id` (`category_id`),
    CONSTRAINT `SubCategories_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `Categories` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 4. InstallmentPlans (depends on Categories, SubCategories, and Accounts)
CREATE TABLE `InstallmentPlans` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `total_amount` decimal(10,2) NOT NULL,
    `total_installments` int(11) NOT NULL,
    `remaining_installments` int(11) DEFAULT NULL,
    `installment_amount` decimal(10,2) NOT NULL,
    `start_date` datetime NOT NULL,
    `next_payment_date` datetime DEFAULT NULL,
    `pay_date` datetime DEFAULT NULL,
    `status` varchar(50) DEFAULT NULL,
    `category_id` int(11) DEFAULT NULL,
    `from_account_id` int(11) DEFAULT NULL,
    `subcategory_id` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `category_id` (`category_id`),
    KEY `subcategory_id` (`subcategory_id`),
    KEY `from_account_id` (`from_account_id`),
    CONSTRAINT `InstallmentPlans_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `Categories` (`id`),
    CONSTRAINT `InstallmentPlans_ibfk_2` FOREIGN KEY (`subcategory_id`) REFERENCES `SubCategories` (`id`),
    CONSTRAINT `InstallmentPlans_ibfk_4` FOREIGN KEY (`from_account_id`) REFERENCES `Accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 5. RecurringPayments (depends on Categories, SubCategories, and Accounts)
CREATE TABLE `RecurringPayments` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `amount` decimal(10,2) NOT NULL,
    `frequency` varchar(50) NOT NULL,
    `start_date` datetime NOT NULL,
    `pay_date` datetime DEFAULT NULL,
    `status` tinyint(1) DEFAULT 1,
    `category_id` int(11) DEFAULT NULL,
    `from_account_id` int(11) DEFAULT NULL,
    `subcategory_id` int(11) DEFAULT NULL,
    `payments` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `category_id` (`category_id`),
    KEY `subcategory_id` (`subcategory_id`),
    KEY `from_account_id` (`from_account_id`),
    CONSTRAINT `RecurringPayments_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `Categories` (`id`),
    CONSTRAINT `RecurringPayments_ibfk_2` FOREIGN KEY (`subcategory_id`) REFERENCES `SubCategories` (`id`),
    CONSTRAINT `RecurringPayments_ibfk_4` FOREIGN KEY (`from_account_id`) REFERENCES `Accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 6. Transactions (depends on all other tables)
CREATE TABLE `Transactions` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `amount` decimal(10,2) NOT NULL,
    `category_id` int(11) DEFAULT NULL,
    `from_account_id` int(11) DEFAULT NULL,
    `to_account_id` int(11) DEFAULT NULL,
    `subcategory_id` int(11) DEFAULT NULL,
    `currency` varchar(3) NOT NULL,
    `payment_method` varchar(50) DEFAULT NULL,
    `exchange_rate` decimal(10,2) DEFAULT NULL,
    `notes` text DEFAULT NULL,
    `date` datetime NOT NULL,
    `status` varchar(50) DEFAULT NULL,
    `installment_plan_id` int(11) DEFAULT NULL,
    `recurring_payment_id` int(11) DEFAULT NULL,
    `payment_number` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `date` (`date`),
    KEY `category_id` (`category_id`),
    KEY `subcategory_id` (`subcategory_id`),
    KEY `installment_plan_id` (`installment_plan_id`),
    KEY `recurring_payment_id` (`recurring_payment_id`),
    KEY `from_account_id` (`from_account_id`),
    KEY `to_account_id` (`to_account_id`),
    CONSTRAINT `Transactions_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `Categories` (`id`),
    CONSTRAINT `Transactions_ibfk_2` FOREIGN KEY (`subcategory_id`) REFERENCES `SubCategories` (`id`),
    CONSTRAINT `Transactions_ibfk_3` FOREIGN KEY (`installment_plan_id`) REFERENCES `InstallmentPlans` (`id`),
    CONSTRAINT `Transactions_ibfk_4` FOREIGN KEY (`recurring_payment_id`) REFERENCES `RecurringPayments` (`id`),
    CONSTRAINT `Transactions_ibfk_6` FOREIGN KEY (`from_account_id`) REFERENCES `Accounts` (`id`),
    CONSTRAINT `Transactions_ibfk_7` FOREIGN KEY (`to_account_id`) REFERENCES `Accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 7. Mock data
INSERT INTO MoneyWeight.Accounts (name,`type`,current_balance,currency,institution,is_active,created_at) VALUES ('cuenta1','digital_wallet',10000.00,'ars','mercadolibre',1,'2025-01-18 18:03:02');
INSERT INTO MoneyWeight.Categories (name,icon,`type`,budget_limit) VALUES ('categoria1','nose','income',0.00);
INSERT INTO MoneyWeight.SubCategories (category_id,name,icon,budget_limit) VALUES (1,'subcat','nosetampoco',0.00);
