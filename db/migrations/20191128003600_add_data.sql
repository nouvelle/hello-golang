-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO gopher (id, name, img, count)
VALUES
  (1, 'diallo', 'diallo.png', 0),
  (2, 'eriko', 'eriko.png', 0),
  (3, 'fraser', 'fraser.png', 0),
  (4, 'heart', 'heart.png', 0);
-- +goose Down
  -- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM gopher;