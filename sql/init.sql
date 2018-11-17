DROP TABLE if exists csvdata;

CREATE TABLE csvdata
(
  "id"	NUMERIC PRIMARY KEY,
  "name"	TEXT,
  "email"	TEXT,
  "phone"	TEXT,
  created_at timestamp  NOT NULL  DEFAULT current_timestamp,
  updated_at timestamp  NOT NULL  DEFAULT current_timestamp
);
