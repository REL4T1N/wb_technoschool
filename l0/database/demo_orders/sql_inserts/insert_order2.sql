-- Таблица orders
INSERT INTO orders (
    order_uid, track_number, entry, locale, internal_signature, customer_id,
    delivery_service, shardkey, sm_id, date_created, oof_shard
) VALUES (
    'c874fgh7b2b84b7test',
    'WBILMTESTTRACK2',
    'WBIL',
    'en',
    '',
    'alice',
    'fedex',
    '10',
    100,
    '2022-02-10T10:00:00Z',
    '2'
);

-- Таблица delivery
INSERT INTO delivery (
    order_uid, name, phone, zip, city, address, region, email
) VALUES (
    'c874fgh7b2b84b7test',
    'Alice Johnson',
    '+1234567890',
    '10001',
    'New York',
    '5th Avenue 21',
    'NY',
    'alice@example.com'
);

-- Таблица payment
INSERT INTO payment (
    order_uid, transaction, request_id, currency, provider, amount, payment_dt,
    bank, delivery_cost, goods_total, custom_fee
) VALUES (
    'c874fgh7b2b84b7test',
    'c874fgh7b2b84b7test',
    '',
    'USD',
    'stripe',
    2300,
    1670000000,
    'chase',
    500,
    1800,
    0
);

-- Таблица items
INSERT INTO items (
    chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
    nm_id, brand, status
) VALUES (
    9934931,
    'c874fgh7b2b84b7test',
    'WBILMTESTTRACK2',
    900,
    '',
    'Lipstick',
    10,
    'M',
    810,
    2389213,
    'MAC',
    202
);

INSERT INTO items (
    chrt_id, order_uid, track_number, price, rid, name, sale, size, total_price,
    nm_id, brand, status
) VALUES (
    9934932,
    'c874fgh7b2b84b7test',
    'WBILMTESTTRACK2',
    990,
    '',
    'Eyeliner',
    5,
    'S',
    940,
    2389214,
    'Maybelline',
    202
);
