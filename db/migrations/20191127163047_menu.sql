-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE menu(
  id SERIAL,
  name varchar(20),
  price INTEGER
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE menu;