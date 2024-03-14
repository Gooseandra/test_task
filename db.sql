create table coin_price(
    id serial primary key,
    name varchar(15),
    price float,
    time timestamp
)

create table watching_coins(
    if serial primary key,
    name varchar(15)
)