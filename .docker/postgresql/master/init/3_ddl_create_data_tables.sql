create table if not exists users
(
    id serial constraint users_pk primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    login varchar(255) not null,
    created_at timestamp not null,
    is_deleted boolean not null DEFAULT false
);

alter table users owner to postgres_user;
create unique index if not exists users_login_uindex on users (login);

create table if not exists subtitles
(
    id serial constraint subtitle_pk primary key,
    name varchar(255) not null,
    text text not null,
    user_id int not null constraint subtitle_user_id_fk references users,
    created_at timestamp not null,
    is_deleted boolean not null DEFAULT false
);

alter table subtitles owner to postgres_user;

create table if not exists phrases
(
    id serial constraint phrase_pk primary key,
    text text not null,
    subtitle_id int not null constraint phrase_subtitle_id_fk references subtitles,
    created_at timestamp not null,
    is_deleted boolean not null DEFAULT false
);

alter table phrases owner to postgres_user;

create table if not exists phrase_translations
(
    id serial constraint phrase_pk primary key,
    text text not null,
    phrase_id int not null constraint translate_phrase_id_fk references phrases,
    created_at timestamp not null,
    is_deleted boolean not null DEFAULT false
);

alter table phrases owner to postgres_user;