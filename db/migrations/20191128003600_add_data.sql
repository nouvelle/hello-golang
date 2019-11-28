-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO gopher (id, name, img, count)
VALUES
  (1, 'diaGopher', 'diaGopher.png', 0),
  (2, 'eriGopher', 'eriGopher.png', 0),
  (3, 'fraGopher', 'fraGopher.png', 0),
  (4, 'iGopher', 'iGopher.png', 0),
  (5, 'dereGopher', 'dereGopher.png', 0),
  (6, 'heartGopher', 'heartGopher.png', 0);
-- +goose Down
  -- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM gopher;