-- Create databases for each service
CREATE DATABASE order_service;
CREATE DATABASE payment_service;
CREATE DATABASE user_service;

-- Grant all privileges to the ecommerce user
GRANT ALL PRIVILEGES ON DATABASE order_service TO ecommerce;
GRANT ALL PRIVILEGES ON DATABASE payment_service TO ecommerce;
GRANT ALL PRIVILEGES ON DATABASE user_service TO ecommerce;