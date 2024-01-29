CREATE TABLE IF NOT EXISTS members (
    id varchar(36) NOT NULL,
    access_token varchar(1500) NOT NULL,
    id_token varchar(1500) NOT NULL,
    refresth_token varchar(2000) NOT NULL,
    is_used boolean default(0) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (id)
);
