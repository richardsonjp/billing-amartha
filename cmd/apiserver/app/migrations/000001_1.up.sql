CREATE TABLE borrower (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('normal', 'delinquent')),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE TABLE loan (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    borrower_id uuid NOT NULL REFERENCES borrower(id),
    amount BIGINT NOT NULL,
    interest BIGINT NOT NULL,
    total_amount BIGINT NOT NULL,
    weekly_installment BIGINT NOT NULL,
    total_weeks INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE TABLE billing (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    loan_id uuid NOT NULL REFERENCES loan(id),
    amount BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('outstanding', 'paid')),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);
