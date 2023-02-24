-- +migrate Up
-- +migrate StatementBegin

ALTER TABLE perkuliahan RENAME COLUMN kode_perkuliahan TO id;


ALTER TABLE perkuliahan
    ALTER COLUMN id TYPE INT;
-- +migrate StatementEnd