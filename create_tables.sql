create table users
(
    id         serial constraint users_pk primary key,
    amount     integer,
    bookamount integer
);


create table orders
(
    id          serial constraint orders_pk primary key,
    fromuserid  integer references users,
    touserid    integer references users,
    serviceid   integer,
    amount      integer,
    date        date,
    description varchar(50),
    orderid     integer
);
