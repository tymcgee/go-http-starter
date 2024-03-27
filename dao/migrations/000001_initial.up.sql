create table if not exists author (
    id integer primary key,
    name text not null
);

create table if not exists book (
    id integer primary key,
    author_id integer not null,
    title text not null,
    description text
);
