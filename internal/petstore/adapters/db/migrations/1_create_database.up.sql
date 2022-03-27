begin;

create table pet (
    id uuid primary key,
    name varchar(50) not null,
    tag varchar(50),
    created_at timestamp not null default now()
);

commit;