DROP TABLE if exists test_issue.public.csvdata;

CREATE TABLE test_issue.public.csvdata
(
  "Id"	NUMERIC PRIMARY KEY,
  "compounds"	TEXT,
  "volume"	NUMERIC,
  created_at timestamp  NOT NULL  DEFAULT current_timestamp,
  updated_at timestamp  NOT NULL  DEFAULT current_timestamp
);
