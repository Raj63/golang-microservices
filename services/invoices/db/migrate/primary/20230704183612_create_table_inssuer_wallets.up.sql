CREATE TABLE IF NOT EXISTS issuer_wallets (
  id UUID PRIMARY KEY,
  issuer_id UUID NOT NULL,
  balance BIGINT DEFAULT 0,
  currency_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  CONSTRAINT uc_issuer_wallets_issuer_id UNIQUE (issuer_id),
  FOREIGN KEY (issuer_id) REFERENCES issuers (id) ON DELETE CASCADE,
  FOREIGN KEY (currency_id) REFERENCES currency (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS issuer_wallets_issuer_id_IDX ON issuer_wallets (issuer_id);
