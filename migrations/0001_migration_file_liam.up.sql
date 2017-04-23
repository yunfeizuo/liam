CREATE TABLE IF NOT EXISTS shipment (
    id          SERIAL,
    ship_date   date NULL,
    status      varchar(32) NOT NULL DEFAULT 'new',
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS package (
    id              SERIAL,
    shipment_id     int NULL,
    barcode         varchar(128) NULL,
    weight          real NULL,
    name            varchar(128) NULL,
    address         varchar(256) NULL,
	cellphone       varchar(128) NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS orderitem (
    id              SERIAL,
    customer_name   varchar(128) NULL,
    name            varchar(128) NULL,
	desciption      varchar(512) NULL,
	brand           varchar(128) NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS orders (
    id              SERIAL,
    customer_id     int NULL,
    title           varchar(128) NULL,
    note            varchar(512) NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS products (
    id              uuid,
    title           varchar(128) NULL,
    img_urls        varchar(65536) NULL,
    PRIMARY KEY(id)
);