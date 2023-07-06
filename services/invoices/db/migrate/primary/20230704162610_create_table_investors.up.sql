CREATE TABLE IF NOT EXISTS investors (
  id UUID PRIMARY KEY,
  name VARCHAR NOT NULL,
  vat_number VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS investors_name_IDX ON investors (name);
CREATE INDEX IF NOT EXISTS investors_vat_number_IDX ON investors (vat_number);
ALTER TABLE investors ADD CONSTRAINT uc_investors_vat_number UNIQUE (vat_number);