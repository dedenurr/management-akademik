-- +migrate Up
-- +migrate StatementBegin
SELECT MAX(id)+1 FROM perkuliahan;

CREATE SEQUENCE perkuliahan_id_seq MINVALUE 0;

ALTER TABLE perkuliahan ALTER id SET DEFAULT nextval('perkuliahan_id_seq');

ALTER SEQUENCE perkuliahan_id_seq OWNED BY perkuliahan.id;

-- +migrate StatementEnd