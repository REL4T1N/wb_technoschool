-- Таблица orders
INSERT INTO orders (
    order_uid, track_number, entry, locale, internal_signature, customer_id,
    delivery_service, shardkey, sm_id, date_created, oof_shard
) VALUES (
    'd983bc23e5f6787test',
    'WBILMTRACK003',
    'WBIL',
    'ru',
    '',
    'user3',
    'fedex',
    '3',
    102,
    '2024-06-01T08:45:00Z',
    '1'
);

-- Таблица delivery
INSERT INTO delivery (
    order_uid, name, phone, zip, city, address, region, email
) VALUES (
    'd983bc23e5f6787test',
    'Ivan Petrov',
    '+74951234567',
    '123456',
    'Moscow',
    'Lenina 25',
    'Moscow Oblast',
    'ivan.petrov@example.com'
);

-- Таблица payment
INSERT INTO payment (
    order_uid, transaction, request_id, currency, provider, amount, payment_dt,
    bank, delivery_cost, goods_total, custom_fee
) VALUES (
    'd983bc23e5f6787test',
    'd983bc23e5f6787test',
    '',
    'RUB',
    'sberpay',
    5000,
    1712125500,
    'sberbank',
    500,
    4500,
    0
);

-- Таблица items
INSERT INTO items (
    chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
    nm_id, brand, status
) VALUES (
    9934933,
    'd983bc23e5f6787test',
    'WBILMTRACK003',
    3000,
    'ef9876543210test',
    'Sneakers',
    0,
    '42',
    3000,
    2389215,
    'Puma',
    202
);

INSERT INTO items (
    chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
    nm_id, brand, status
) VALUES (
    9934934,
    'd983bc23e5f6787test',
    'WBILMTRACK003',
    1500,
    'ef9876543211test',
    'Socks',
    0,
    'L',
    1500,
    2389216,
    'Reebok',
    202
);
