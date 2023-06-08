create table urls
(
    id bigserial primary key unique not null,
    url text not null,
    short_url text not null
);