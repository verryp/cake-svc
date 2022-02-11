-- +migrate Up
CREATE TABLE IF NOT EXISTS table_name (
    id bigint not null AUTO_INCREMENT,

    primary key (id)
);

-- +migrate Down
DROP TABLE IF EXISTS table_name;
