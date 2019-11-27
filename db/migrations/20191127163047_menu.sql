-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE gopher(
  id SERIAL,
  name varchar(20),
  img varchar(50),
  count INTEGER
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE gopher;