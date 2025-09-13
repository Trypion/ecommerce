-- Create databases for each service
CREATE DATABASE order_service;
CREATE DATABASE payment_service;

-- Grant all privileges to the ecommerce user
GRANT ALL PRIVILEGES ON DATABASE order_service TO ecommerce;
GRANT ALL PRIVILEGES ON DATABASE payment_service TO ecommerce;

-- Connect to each database and create necessary extensions
\c order_service;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\c payment_service;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
