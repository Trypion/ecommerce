-- Drop tables in reverse order (due to foreign key constraints)

-- Drop refunds table first (has foreign key to payments)
DROP TABLE IF EXISTS refunds;

-- Drop payments table
DROP TABLE IF EXISTS payments;
