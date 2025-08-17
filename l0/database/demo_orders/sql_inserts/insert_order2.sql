-- Таблица orders
INSERT INTO orders (
    order_uid, track_number, entry, locale, internal_signature, customer_id,
    delivery_service, shardkey, sm_id, date_created, oof_shard
) VALUES (
    'c872ab12d4e3456test',
    'WBILMTRACK002',
    'WBIL',
    'en',
    '',
    'user2',
    'dhl',
    '5',
    101,
    '2023-02-15T12:30:00Z',
    '2'
);

-- Таблица delivery
INSERT INTO delivery (
    order_uid, name, phone, zip, city, address, region, email
) VALUES (
    'c872ab12d4e3456test',
    'Alice Johnson',
    '+12025550123',
    '10001',
    'New York',
    '5th Avenue 10',
    'NY',
    'alice@example.com'
);

-- Таблица payment
INSERT INTO payment (
    order_uid, transaction, request_id, currency, provider, amount, payment_dt,
    bank, delivery_cost, goods_total, custom_fee
) VALUES (
    'c872ab12d4e3456test',
    'c872ab12d4e3456test',
    '',
    'USD',
    'paypal',
    250,
    1676457600,
    'chase',
    20,
    230,
    0
);

-- Таблица items
INSERT INTO items (
    chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
    nm_id, brand, status
) VALUES (
    9934931,
    'c872ab12d4e3456test',
    'WBILMTRACK002',
    150,
    'cd1234567890test',
    'T-Shirt',
    10,
    'M',
    135,
    2389213,
    'Nike',
    202
);

INSERT INTO items (
    chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
    nm_id, brand, status
) VALUES (
    9934932,
    'c872ab12d4e3456test',
    'WBILMTRACK002',
    100,
    'cd1234567891test',
    'Cap',
    5,
    'L',
    95,
    2389214,
    'Adidas',
    202
);
