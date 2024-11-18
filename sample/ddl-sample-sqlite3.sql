CREATE TABLE customer (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT DEFAULT NULL,
    address TEXT DEFAULT NULL,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    parent_category_id INTEGER DEFAULT NULL,
    description TEXT DEFAULT NULL,
    sort_order INTEGER NOT NULL DEFAULT 0,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT DEFAULT NULL,
    price REAL NOT NULL CHECK(price >= 0),
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    category_id INTEGER NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES product_category(id)
);

CREATE TABLE supplier (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    contact_person TEXT DEFAULT NULL,
    phone TEXT DEFAULT NULL,
    email TEXT DEFAULT NULL,
    address TEXT DEFAULT NULL,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment_method (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT DEFAULT NULL,
    is_active INTEGER NOT NULL DEFAULT 1,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- customer テーブルの updated_at 更新トリガー
CREATE TRIGGER update_customer_updated_at
AFTER UPDATE ON customer
FOR EACH ROW
BEGIN
    UPDATE customer SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- product_category テーブルの updated_at 更新トリガー
CREATE TRIGGER update_product_category_updated_at
AFTER UPDATE ON product_category
FOR EACH ROW
BEGIN
    UPDATE product_category SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- product テーブルの updated_at 更新トリガー
CREATE TRIGGER update_product_updated_at
AFTER UPDATE ON product
FOR EACH ROW
BEGIN
    UPDATE product SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- supplier テーブルの updated_at 更新トリガー
CREATE TRIGGER update_supplier_updated_at
AFTER UPDATE ON supplier
FOR EACH ROW
BEGIN
    UPDATE supplier SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;

-- payment_method テーブルの updated_at 更新トリガー
CREATE TRIGGER update_payment_method_updated_at
AFTER UPDATE ON payment_method
FOR EACH ROW
BEGIN
    UPDATE payment_method SET updated_at = CURRENT_TIMESTAMP WHERE code = OLD.code;
END;
