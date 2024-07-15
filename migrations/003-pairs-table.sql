-- create enum type pair exchange
CREATE TYPE pair_exchange AS ENUM ('binance', 'depreciated');

-- create enum type pair status
CREATE TYPE pair_status AS ENUM ('active', 'break', 'halt', 'end_of_day', 'delisted');

-- create pairs table
CREATE TABLE IF NOT EXISTS pairs (
    id         uuid PRIMARY KEY       DEFAULT uuid_generate_v4(),
    base       VARCHAR(24)    NOT NULL,
    quote      VARCHAR(24)    NOT NULL,
    exchange   pair_exchange NOT NULL,
    status     pair_status   NOT NULL,
    updated_at TIMESTAMPTZ   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_base_quote_exchange UNIQUE (base, quote, exchange)
);

-- add trigger function to automatically update the updated_at column
CREATE TRIGGER update_pairs_updated_at_column
    BEFORE UPDATE
    ON pairs
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
