
CREATE DATABASE topup;
USE topup;

CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);

-- Tambah data awal kategori
INSERT INTO categories (name) VALUES
('Game'),
('Pulsa'),
('Streaming');

-- Tambah data awal produk
INSERT INTO products (name, price, category_id) VALUES
('Mobile Legends 86 Diamond', 20000, 1),
('Pulsa Telkomsel 20K', 21000, 2),
('Langganan Netflix 1 Bulan', 50000, 3);
