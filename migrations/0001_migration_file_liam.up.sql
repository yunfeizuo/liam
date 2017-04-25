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

CREATE TABLE IF NOT EXISTS orderitems (
    id              uuid,
    order_id        uuid,
    image_url       varchar(1024) NULL,
    title           varchar(512) NULL,
	note            varchar(1024) NULL,
	brand           varchar(128) NULL,
    category        varchar(128) NULL,
    price           NUMERIC(2) NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS orders (
    id              uuid,
    seller_id       uuid NULL,
    buyer_id        uuid NULL,
    buyer_name      varchar(128) NULL,
    buyer_address   varchar(256) NULL,
    buyer_cellphone varchar(32) NULL,
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