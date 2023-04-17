create table if not exists orders (
    id          bigserial primary key,
    order_uid   varchar unique not null,
    data        jsonb not null
);

create index if not exists idx_orders_order_uid on orders (order_uid);