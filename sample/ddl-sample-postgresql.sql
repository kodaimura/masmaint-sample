create function set_update_time() returns trigger AS '
    BEGIN
      new.updated_at := NOW();
      return new;
    END;
' language 'plpgsql';

CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT DEFAULT NULL,
    address TEXT DEFAULT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_category (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    parent_category_id INTEGER DEFAULT NULL,
    description TEXT DEFAULT NULL,
    sort_order INTEGER NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT DEFAULT NULL,
    price NUMERIC(10, 2) NOT NULL CHECK(price >= 0),
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    category_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES product_category(id)
);

CREATE TABLE supplier (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    contact_person TEXT DEFAULT NULL,
    phone TEXT DEFAULT NULL,
    email TEXT DEFAULT NULL,
    address TEXT DEFAULT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment_method (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT DEFAULT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- customer テーブルの updated_at 更新トリガー
CREATE TRIGGER trigger_update_customer_updated_at
BEFORE UPDATE ON customer
FOR EACH ROW
EXECUTE FUNCTION set_update_time();

-- product_category テーブルの updated_at 更新トリガー
CREATE TRIGGER trigger_update_product_category_updated_at
BEFORE UPDATE ON product_category
FOR EACH ROW
EXECUTE FUNCTION set_update_time();

-- product テーブルの updated_at 更新トリガー
CREATE TRIGGER trigger_update_product_updated_at
BEFORE UPDATE ON product
FOR EACH ROW
EXECUTE FUNCTION set_update_time();

-- supplier テーブルの updated_at 更新トリガー
CREATE TRIGGER trigger_update_supplier_updated_at
BEFORE UPDATE ON supplier
FOR EACH ROW
EXECUTE FUNCTION set_update_time();

-- payment_method テーブルの updated_at 更新トリガー
CREATE TRIGGER trigger_update_payment_method_updated_at
BEFORE UPDATE ON payment_method
FOR EACH ROW
EXECUTE FUNCTION set_update_time();