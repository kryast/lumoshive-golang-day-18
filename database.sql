-- table driver
CREATE TABLE driver (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    bike VARCHAR(255),
    police_number VARCHAR(255)
);

-- table customer
CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    phone INT
);


-- table orders

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    pick_up VARCHAR(255),
    destination VARCHAR(255),
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    customer_id INT,
    FOREIGN KEY (customer_id) REFERENCES customer(id)
);

-- Table order process
CREATE TABLE Order_Process (
    id SERIAL PRIMARY KEY,
    driver_id INT,
    order_id INT,
    FOREIGN KEY (driver_id) REFERENCES driver(id),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    status VARCHAR(50), 
    processed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

-- Session
CREATE TABLE session (
    id SERIAL PRIMARY KEY,
    customer_id INT,
    is_active BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (customer_id) REFERENCES customer(id)
);



INSERT INTO driver(name, bike, police_number)VALUES 
('syarifuddin 1', 'Beat Karbu', 'BG 1 A'),
('syarifuddin 2', 'Beat Karbu', 'BG 2 A'),
('syarifuddin 3', 'Beat Karbu', 'BG 3 A'),
('syarifuddin 4', 'Beat Karbu', 'BG 4 A'),
('syarifuddin 5', 'Beat Karbu', 'BG 5 A');


INSERT INTO customer(name, email, phone)VALUES 
('ahmad 1', 'test1@gmail.com', 123456781),
('ahmad 2', 'test2@gmail.com', 123456782),
('ahmad 3', 'test3@gmail.com', 123456783),
('ahmad 4', 'test4@gmail.com', 123456784),
('ahmad 5', 'test5@gmail.com', 123456785);

INSERT INTO Orders (pick_up, destination, price, customer_id) VALUES 
('Celentang', 'kuto', 20000, 1),
('Celentang', 'Lemabang', 15000, 1),
('Celentang', 'kuto', 25000, 1),
('Celentang', 'Lemabang', 30000, 2),
('Celentang', 'Kertapati', 20000, 3),
('Celentang', 'kuto', 18000, 1),
('Celentang', 'Lemabang', 22000, 2),
('Celentang', 'Plaju', 27000, 3);

INSERT INTO order_process (driver_id, order_id, status) VALUES
    (1, 1, 'completed'),
    (1, 2, 'completed'),
    (2, 3, 'completed'),
    (2, 4, 'in progress'),
    (3, 5, 'completed'),
    (3, 6, 'completed'),
    (3, 7, 'completed'),
    (1, 8, 'completed');

INSERT INTO session (customer_id, is_active) VALUES
(1, TRUE),
(2, FALSE),
(3, FALSE),
(4, TRUE),
(5, TRUE);

-- Total Order Setiap bulan

SELECT
TO_CHAR(o.created_at, 'Month') AS month,
COUNT(o.id) as total_order
FROM
order_process op

FULL JOIN
orders o ON op.order_id = o.id

GROUP BY
month;

-- Melihat Customer yang sering order setiap bulan
SELECT
	TO_CHAR(o.created_at, 'Month') AS month,
	COUNT(o.id) AS total_order,
	c.name AS customer_name
FROM
	orders o
JOIN
	customer c ON o.customer_id = c.id
LEFT JOIN
	order_process op ON o.id = op.order_id
GROUP BY
	month, c.name
ORDER BY
	total_order DESC
LIMIT 1;

-- Melihat daerah yg paling banyak Order
SELECT
    location,
    COUNT(*) AS total_orders
FROM (
    SELECT
        pick_up AS location
    FROM
        orders
    UNION ALL
    SELECT
        destination AS location
    FROM
        orders
) AS combined_locations
GROUP BY
    location
ORDER BY
    total_orders DESC;


-- Melihat waktu yang paling banyak order
SELECT
    order_hour,
    total_orders
FROM (
    SELECT
        EXTRACT(HOUR FROM created_at) AS order_hour,
        COUNT(*) AS total_orders
    FROM
        orders
    GROUP BY
        order_hour
) AS hourly_orders
ORDER BY
    total_orders DESC
LIMIT 1;


-- Melihat driver yang paling banyak mengambil order
SELECT
    d.name AS driver_name,
    COUNT(op.order_id) AS total_orders
FROM
    order_process op
JOIN
    driver d ON op.driver_id = d.id
GROUP BY
    d.name
ORDER BY
    total_orders DESC
LIMIT 1;

-- Melihat Jumlah customer yang login dan logout
SELECT
    CASE
        WHEN is_active THEN 'Logged In'
        ELSE 'Logged Out'
    END AS status,
    COUNT(*) AS total_customers
FROM
    session
GROUP BY
    is_active;
