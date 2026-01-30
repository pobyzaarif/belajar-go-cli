INSERT INTO fields (name, size, location, hourly_rate) VALUES
('Ronaldo Arena', '5v5',  'Jkt', 50.00),
('Messi Park', 'medium', 'Bali', 70.00),
('Neymar Ground', 'small',  'Sby', 45.00);

INSERT INTO customers (first_name, last_name, email, phone_number) VALUES
('Andi', 'Saputra', 'andi@mail.com', '081234567890'),
('Budi', 'Pratama', 'budi@mail.com', '081234567891'),
('Citra', 'Lestari', 'citra@mail.com', '081234567892'),
('Eko', 'Satrio', 'eko@mail.com', '081234567893'),
('Agus', 'Wiratmoko', 'agus@mail.com', '081234567894');

INSERT INTO payment_methods (name) VALUES
('CASH'),
('TRANSFER'),
('E-WALLET');

INSERT INTO bookings
(customer_id, field_id, booking_date, start_time, end_time, total_amount)
VALUES
(1, 1, '2026-02-01', '18:00:00', '20:00:00', 100.00),
(2, 2, '2026-02-02', '19:00:00', '21:00:00', 140.00),
(3, 3, '2026-02-03', '16:00:00', '19:00:00', 135.00),
(4, 3, '2026-02-03', '11:00:00', '12:00:00', 45.00),
(5, 3, '2026-02-03', '12:00:00', '13:00:00', 45.00);

INSERT INTO payments (booking_id, payment_method_id, amount) VALUES
(1, 2, 300000.00),
(2, 3, 240000.00),
(3, 2, 270000.00);

-- query generate revenue report
SELECT
  f.name AS field_name,
  COUNT(b.id) AS total_bookings,
  COALESCE(SUM(b.total_amount), 0) AS total_revenue
FROM fields f
LEFT JOIN bookings b ON b.field_id = f.id
GROUP BY f.id, f.name
ORDER BY total_revenue DESC;

-- query list customer without payment
SELECT
  CONCAT(c.first_name, ' ', c.last_name) AS customer_name,
  b.id AS booking_id,
  b.booking_date
FROM bookings b
JOIN customers c ON c.id = b.customer_id
LEFT JOIN payments p ON p.booking_id = b.id
WHERE p.id IS NULL
ORDER BY b.booking_date ASC;
