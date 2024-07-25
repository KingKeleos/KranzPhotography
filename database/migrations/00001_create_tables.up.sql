create table locations (
    id bigserial not null primary key,
    name varchar(25) not null,
    position varchar(255) not null,
    created_at time default current_timestamp,
    updated_at time default current_timestamp
);

create table events (
    id bigserial not null primary key,
    name varchar(30) not null,
    location_id bigserial not null,
    created_at time default current_timestamp,
    updated_at time default current_timestamp,
    constraint fk_location foreign key(location_id) references locations(id) on delete set null
);

create table projects (
    id bigserial not null primary key,
    name varchar(255) not null,
    rating decimal not null default 0,
    public boolean not null default False,
    category varchar(30) not null,
    tags varchar(255),
    event_id bigserial,
    created_at time default current_timestamp,
    updated_at time default current_timestamp,
    constraint fk_event foreign key(event_id) references events(id) on delete set null
);

create table images (
    id bigserial not null primary key,
    name varchar(25) not null,
    project_id bigserial not null,
    link varchar not null,
    approved boolean not null default False,
    rating decimal not null default 0,
    public boolean not null default False,
    created_at time default current_timestamp,
    updated_at time default current_timestamp,
    constraint fk_project foreign key(project_id) references projects(id) on delete set null
);

create table users (
    id bigserial not null primary key,
    name varchar(32) not null,
    password varchar(255) not null,
    created_at time default current_timestamp,
    updated_at time default current_timestamp
);

create table participants (
    id bigserial not null primary key,
    user_id bigserial not null,
    project_id bigserial not null,
    created_at time default current_timestamp,
    updated_at time default current_timestamp,
    constraint fk_project foreign key(project_id) references projects(id) on delete set null,
    constraint fk_user foreign key(user_id) references users(id) on delete set null
);