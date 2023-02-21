package util

const (
	// USERS
	CREATE_USER = `INSERT INTO users(email, password) VALUES($1, $2) RETURNING user_id;`
	READ_USER   = `SELECT user_id, email, password FROM users WHERE user_id = $1;`
	UPDATE_USER = `UPDATE users SET email = $1, password = $2 WHERE user_id = $3;`
	DELETE_USER = `DELETE FROM users WHERE user_id = $1;`

	// CUSTOMERS
	CREATE_CUSTOMER = `INSERT INTO customers(user_id, name, balance, created_at) VALUES($1, $2, $3, $4) RETURNING customer_id;`
	READ_CUSTOMER   = `SELECT customer_id, user_id, name, balance, created_at FROM customers WHERE customer_id = $1;`
	UPDATE_CUSTOMER = `UPDATE customers SET name = $1, balance = $2 WHERE customer_id = $3;`
	DELETE_CUSTOMER = `DELETE FROM customers WHERE customer_id = $1;`

	// MERCHANTS
	CREATE_MERCHANT = `INSERT INTO merchants(user_id, name, balance, created_at) VALUES($1, $2, $3, $4) RETURNING merchant_id;`
	READ_MERCHANT   = `SELECT merchant_id, user_id, name, balance, created_at FROM merchants WHERE merchant_id = $1;`
	UPDATE_MERCHANT = `UPDATE merchants SET name = $1, balance = $2 WHERE merchant_id = $3;`
	DELETE_MERCHANT = `DELETE FROM merchants WHERE merchant_id = $1;`

	// BANKS
	CREATE_BANK = `INSERT INTO banks(bank_account_number, balance) VALUES($1, $2) RETURNING bank_id;`
	READ_BANK   = `SELECT bank_id, bank_account_number, balance FROM banks WHERE bank_id = $1;`
	UPDATE_BANK = `UPDATE banks SET balance = $1 WHERE bank_id = $2;`
	DELETE_BANK = `DELETE FROM banks WHERE bank_id = $1;`

	// PAYMENTS
	CREATE_PAYMENT = `INSERT INTO payments(sender_id, receiver_id, amount, bank_account_number) VALUES($1, $2, $3, $4) RETURNING payment_id;`
	READ_PAYMENT   = `SELECT payment_id, sender_id, receiver_id, amount, bank_account_number, created_at FROM payments WHERE payment_id = $1;`
	UPDATE_PAYMENT = `UPDATE payments SET sender_id = $1, receiver_id = $2, amount = $3, bank_account_number = $4 WHERE payment_id = $5;`
	DELETE_PAYMENT = `DELETE FROM payments WHERE payment_id = $1;`
)
