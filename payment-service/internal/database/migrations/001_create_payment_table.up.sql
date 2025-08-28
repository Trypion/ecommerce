-- payments table
CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    method VARCHAR(50) NOT NULL,
    
    provider VARCHAR(50),
    provider_id VARCHAR(255),
    provider_response TEXT,
    
    refunded_amount DECIMAL(10,2) DEFAULT 0,
    is_refunded BOOLEAN DEFAULT FALSE,
    
    failure_code VARCHAR(50),
    failure_message VARCHAR(500),
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    processed_at TIMESTAMP NULL,
    
    INDEX idx_payments_order_id (order_id),
    INDEX idx_payments_user_id (user_id),
    INDEX idx_payments_status (status),
    INDEX idx_payments_provider_id (provider_id),
    INDEX idx_payments_deleted_at (deleted_at)
);

-- refunds table
CREATE TABLE refunds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    payment_id UUID NOT NULL,
    
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    reason VARCHAR(500),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    
    provider_id VARCHAR(255),
    provider_response TEXT,
    processed_by VARCHAR(255),
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    processed_at TIMESTAMP NULL,
    
    FOREIGN KEY (payment_id) REFERENCES payments(id) ON DELETE CASCADE,
    INDEX idx_refunds_payment_id (payment_id),
    INDEX idx_refunds_status (status),
    INDEX idx_refunds_deleted_at (deleted_at)
);