create table if not exists "user"
(
    id              integer         not null,
    first_name      varchar(255)    not null,
    last_name       varchar(255)    not null,
    username        varchar(255)    not null,
    language_code   varchar(2)      not null,
    created_at      timestamp       not null
);

create unique index user_id_uindex
    on "user" (id);

alter table "user"
    add constraint user_pk
        primary key (id);

