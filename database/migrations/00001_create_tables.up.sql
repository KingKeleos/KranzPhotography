--create the tables
create table if not exists images (
    id bigserial not null primary key,
    name varchar(25) not null,
    project_id bigserial not null,
    link varchar not null, --link to Lightroom image
    approved boolean not null default False,
    rating decimal not null default 0,
    public boolean not null default False,
    created_at time default current_timestamp,
    updated_at time default current_timestamp
);

create table if not exists projects (
    id bigserial not null primary key,
    name varchar(255) not null,
    header bigserial,
    rating decimal not null default 0,
    public boolean not null decimal False,
    category varchar(30) not null,
    tags varchar(255),
    created_at time default current_timestamp,
    updated_at time default current_timestamp
);

create table if not exists participants (
    id bigserial no null primary key,
    user_id bigserial not null,
    project_id bigserial not null,
    created_at time default current_timestamp,
    updated_at time default current_timestamp
);

create table if not exists users (
    id bigserial not null primary key,
    name string not null,
    password password not null,
    profile_picture 
    created_at time default current_timestamp,
    updated_at time default current_timestamp
);

-- link the tables