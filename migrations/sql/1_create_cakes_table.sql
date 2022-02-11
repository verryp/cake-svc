-- +migrate Up
CREATE TABLE IF NOT EXISTS cakes (
    id int not null AUTO_INCREMENT,
    title varchar (200) not null,
    description text not null,
    rating int not null default 0,
    image varchar (255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,

    primary key (id)
);

-- +migrate Down
DROP TABLE IF EXISTS cakes;
