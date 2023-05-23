CREATE TABLE jobs (
  id          UUID          PRIMARY KEY,
  company     VARCHAR(100)  NOT NULL,
  title       VARCHAR(100)  NOT NULL,
  description TEXT          NOT NULL,
  location    VARCHAR(100)  NOT NULL,
  salary      FLOAT         NOT NULL,
  role        VARCHAR(100)  NOT NULL,
  remote      BOOLEAN       NOT NULL DEFAULT TRUE,
  created_at  TIMESTAMP     NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMP     NOT NULL DEFAULT NOW(),
  deleted_at  TIMESTAMP
);