CREATE TABLE rates
(
    id         SERIAL PRIMARY KEY,
    timestamp  BIGINT  NOT NULL,
    ask_price  VARCHAR NOT NULL,
    bid_price  VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);