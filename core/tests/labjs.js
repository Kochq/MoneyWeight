const id = document.getElementById("id");
const add = document.getElementById("add");
const sub = document.getElementById("sub");

sub.addEventListener("click", () => id.value = parseInt(id.value) - 1)
add.addEventListener("click", () => id.value = parseInt(id.value) + 1)

const addTransaction = () => {
    fetch('http://serkq.org:8080/api/transactions', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            title: "titulo2",
            amount: 10002.2,
            category_id: 1,
            subcategory_id: 1,
            currency: "ARS",
            payment_method: "cosa",
            exchange_rate: 1500,
            notes: "esto es una nota",
            from_account_id: 1
        })
    })
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(err => console.error(err));
}

const deleteTransaction = () => {
    fetch(`http://localhost:8080/api/transactions/${id.value}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(err => console.error(err));
}

const deleteInstallment = () => {
    fetch(`http://localhost:8080/api/installments/${id.value}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(err => console.error(err));
}

const deleteRecurringPayment = () => {
    fetch(`http://localhost:8080/api/recurring/${id.value}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(err => console.error(err));
}

const addInstallment = () => {
    fetch('http://localhost:8080/api/installments', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            title: "installment1",
            total_amount: 10000,
            total_installments: 6,
            category_id: 1,
            installments_amount: 0,
            subcategory_id: 1,
            status: "nosesiestova",
            from_account_id: 1
        })
    })
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(err => console.error(err));
}

const addRecurringPayment = () => {
    fetch('http://localhost:8080/api/recurring', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            title: "recurring1",
            amount: 10000,
            is_active: true,
            category_id: 1,
            subcategory_id: 1,
            from_account_id: 1
        })
    })
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(err => console.error(err));
}
