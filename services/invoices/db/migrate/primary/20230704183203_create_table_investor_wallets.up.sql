CREATE TABLE IF NOT EXISTS investor_wallets (
  id UUID PRIMARY KEY,
  investor_id UUID NOT NULL,
  balance BIGINT DEFAULT 0,
  currency_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT uc_investor_wallets_investor_id UNIQUE (investor_id),
  FOREIGN KEY (investor_id) REFERENCES investors (id) ON DELETE CASCADE,
  FOREIGN KEY (currency_id) REFERENCES currency (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS investor_wallets_investor_id_IDX ON investor_wallets (investor_id);
