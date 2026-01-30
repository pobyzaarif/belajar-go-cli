CREATE TABLE fields (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  size VARCHAR(8) NOT NULL,
  location VARCHAR(32) NOT NULL,
  hourly_rate DECIMAL(10,2) NOT NULL,

  UNIQUE KEY uq_fields_name (name)
);

CREATE TABLE customers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(32) NOT NULL,
  last_name VARCHAR(32) NOT NULL,
  email VARCHAR(64) NOT NULL,
  phone_number VARCHAR(16) NOT NULL,

  UNIQUE KEY uq_customers_email (email),
  UNIQUE KEY uq_customers_phone (phone_number),
  INDEX idx_customers_name (last_name, first_name)
);

CREATE TABLE bookings (
  id INT AUTO_INCREMENT PRIMARY KEY,
  customer_id INT NOT NULL,
  field_id INT NOT NULL,
  booking_date DATE NOT NULL,
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  total_amount DECIMAL(10,2) NOT NULL,

  CONSTRAINT chk_time CHECK (end_time > start_time),

  CONSTRAINT fk_bookings_customer
    FOREIGN KEY (customer_id) REFERENCES customers(id)
    ON DELETE CASCADE,

  CONSTRAINT fk_bookings_field
    FOREIGN KEY (field_id) REFERENCES fields(id)
    ON DELETE RESTRICT,

  INDEX idx_booking_date (booking_date),
  INDEX idx_booking_field (field_id, booking_date),
  INDEX idx_booking_customer (customer_id)
);

CREATE TABLE payment_methods (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(16) NOT NULL UNIQUE
);

CREATE TABLE payments (
  id INT AUTO_INCREMENT PRIMARY KEY,
  booking_id INT NOT NULL,
  payment_method_id INT NOT NULL,
  payment_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  amount DECIMAL(10,2) NOT NULL,

  CONSTRAINT fk_payments_booking
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
    ON DELETE CASCADE,

  CONSTRAINT fk_payments_method
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id),

  INDEX idx_payments_booking (booking_id),
  INDEX idx_payments_date (payment_date)
);
