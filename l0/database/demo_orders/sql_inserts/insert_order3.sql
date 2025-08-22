-- Таблица orders
INSERT INTO orders (
    order_uid, track_number, entry, locale, internal_signature, customer_id,
    delivery_service, shardkey, sm_id, date_created, oof_shard
) VALUES (
    'd985ijk7b2b84b8test',
    'WBILMTESTTRACK3',
    'WBIL',
    'en',
    '',
    'bob',
    'ups',
    '11',
    101,
    '2022-05-15T15:30:00Z',
    '3'
);

-- Таблица delivery
INSERT INTO delivery (
    order_uid, name, phone, zip, city, address, region, email
) VALUES (
    'd985ijk7b2b84b8test',
    'Bob Smith',
    '+1987654321',
    '90001',
    'Los Angeles',
    'Sunset Blvd 100',
    'CA',
    'bob@example.com'
);

-- Таблица payment
INSERT INTO payment (
    order_uid, transaction, request_id, currency, provider, amount, payment_dt,
    bank, delivery_cost, goods_total, custom_fee
) VALUES (
    'd985ijk7b2b84b8test',
    'd985ijk7b2b84b8test',
    '',
    'USD',
    'paypal',
    500,
    1680000000,
    'boa',
    50,
    450,
    0
);

-- Таблица items
INSERT INTO items (
    chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
    nm_id, brand, status
) VALUES (
    9934933,
    'd985ijk7b2b84b8test',
    'WBILMTESTTRACK3',
    450,
    '',
    'Foundation',
    0,
    'L',
    450,
    2389215,
    'Revlon',
    202
);
