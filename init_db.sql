-- Init DB
CREATE TABLE IF NOT EXISTS tenant (
  name TEXT
)

CREATE TABLE IF NOT EXISTS user (
  username TEXT,
  email TEXT NOT NULL UNIQUE
)

CREATE TABLE IF NOT EXISTS agent (
  tenant_id INTEGER,
  name TEXT,
  prompt TEXT,
  FOREIGN KEY (tenant_id) REFERENCES tenant(cid)
)

CREATE TABLE IF NOT EXISTS faq (
  tenant_id INTEGER,
  question TEXT,
  answer TEXT,
  priority INTEGER,
  FOREIGN KEY (tenant_id) REFERENCES tenant(cid)
)


CREATE TABLE IF NOT EXISTS tenant_user (
  tenant_id INTEGER,
  user_id INTEGER,
  role TEXT,
  FOREIGN KEY (tenant_id) REFERENCES tenant(cid)
  FOREIGN KEY (user_id) REFERENCES user(cid)
)

CREATE TABLE IF NOT EXISTS session (
  token TEXT PRIMARY KEY,
  user_email TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  expire_at DATETIME
);
