-- create candles table
CREATE TABLE IF NOT EXISTS candles (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    pair_id uuid NOT NULL,
    open_time TIMESTAMPTZ NOT NULL,
    close_time TIMESTAMPTZ NOT NULL,
    open NUMERIC(18, 10) NOT NULL,
    close NUMERIC(18, 10) NOT NULL,
    high NUMERIC(18, 10) NOT NULL,
    low NUMERIC(18, 10) NOT NULL,
    volume NUMERIC(18, 10),
    no_trades INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, open_time),
    CONSTRAINT fk_pair FOREIGN KEY(pair_id) REFERENCES pairs(id),
    CONSTRAINT unique_pair_open_time UNIQUE (pair_id, open_time)
);

-- create hypertable for candles table
SELECT create_hypertable('candles', by_range('open_time'));
