-- +migrate Up
-- +migrate StatementBegin

ALTER TABLE perkuliahan ADD kode_perkuliahan INT NOT NULL, ADD PRIMARY KEY(kode_perkuliahan) 

-- +migrate StatementEnd

