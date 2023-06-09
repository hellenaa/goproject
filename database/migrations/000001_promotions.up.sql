create table promotions (
    id                serial primary key       not null,
    promotion_id      varchar(250)             not null,
    price             float                    not null,
    expiration_date   varchar(250)             not null
)