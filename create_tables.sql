create table users
(
    id         integer not null
        constraint users_pk
            primary key,
    amount     integer,
    bookamount integer
);

create table orders
(
    id         integer not null
        constraint orders_pk
            primary key,
    fromuserid integer REFERENCES users (id),
    touserid integer REFERENCES users (id),
    serviceid integer,
    orderid integer,
    amount     integer,
    date date,
    description varchar(50)
);