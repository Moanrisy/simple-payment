CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL REFERENCES users(user_id),
	name VARCHAR(255),
    balance DECIMAL(10,0) NOT NULL DEFAULT 0,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE merchants (
    merchant_id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL REFERENCES users(user_id),
	name VARCHAR(255),
    balance DECIMAL(10,0) NOT NULL DEFAULT 0,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE banks (
    bank_id SERIAL PRIMARY KEY,
	name VARCHAR(255),
	bank_account_number TEXT NOT NULL UNIQUE,
    balance DECIMAL(10,0) NOT NULL DEFAULT 0
);

CREATE TABLE payments (
  payment_id SERIAL PRIMARY KEY,
  sender_id INT NOT NULL REFERENCES customers(user_id) CHECK (sender_id <> receiver_id),
  receiver_id INT NOT NULL REFERENCES merchants(user_id),
  amount DECIMAL(10, 0) NOT NULL CHECK (amount > 0),
  bank_account_number TEXT REFERENCES banks(bank_account_number),
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS payments, banks, merchants, customers, users;

WITH models AS (
  WITH data AS (
    SELECT
        replace(initcap(table_name::text), '_', '') table_name,
        replace(initcap(column_name::text), '_', '') column_name,
        CASE data_type
        WHEN 'timestamp without time zone' THEN 'time.Time'
        WHEN 'timestamp with time zone' THEN 'time.Time'
        WHEN 'boolean' THEN 'bool'
        -- add your own type converters as needed or it will default to 'string'
        ELSE 'string'
        END AS type_info,
	    '`' ||
        'json:"' || column_name ||'" '  
	    'db:"' || column_name ||'"' || 
	    '`' AS annotation    
    FROM information_schema.columns
    WHERE table_schema IN ('dvs_app', 'dvs_system', 'public')
    ORDER BY table_schema, table_name, ordinal_position
  )
    SELECT table_name, STRING_AGG(E'\t' || column_name || E'\t' || type_info || E'\t' || annotation, E'\n') fields
    FROM data
    GROUP BY table_name
)
SELECT 'type ' || table_name || E' struct {\n' || fields || E'\n}' models
FROM models ORDER BY 1

INSERT INTO users (email, password) VALUES ('johndoe@mail.com', 'qwerty123');
INSERT INTO users (email, password) VALUES ('johndoe2@mail.com', 'qwerty123');
SELECT * FROM customers;