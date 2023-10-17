drop type if exists user_kind;
create type user_kind as enum ('ADMIN','ASSOCIATION','COMPANY','CANDIDATE');

drop type if exists user_role;
create type user_role as enum ('ADMIN','USER');

drop type if exists user_state;
create type user_state as enum ('ACTIVE','ANONYMOUS','DELETED');

drop type if exists job_status;
create type job_status as enum ('SEARCHING','OPEN_TO','NOT_SEARCHING');

drop type if exists job_type;
create type job_type as enum ('ANY','FULL_TIME','PART_TIME','INTERNSHIP','TEMPORARY');

drop type if exists company_size;
create type company_size as enum ('ANY','SMALL','MEDIUM','LARGE');

drop type if exists location_type;
create type location_type as enum ('REMOTE','HYBRID','ON_SITE');

drop type if exists work_permit;
create type work_permit as enum ('CITIZEN','PERMANENT_RESIDENT','WORK_VISA','STUDENT_VISA','TEMPORARY_RESIDENT','NO_WORK_PERMIT','OTHER');

drop table if exists users;
create table users
(
    id             serial primary key,
    kind           user_kind    not null,
    first_name     varchar(128) not null,
    last_name      varchar(128) not null,
    preferred_name varchar(256),
    email          varchar(512) not null unique,
    phone_number   varchar(20)  not null,
    birth_date     timestamp,
    image_url      varchar(512),
    linkedin_url   varchar(512),
    github_url     varchar(512),
    portfolio_url  varchar(512),
    state          user_state   not null default 'ACTIVE',
    created_at     timestamp    not null default CURRENT_TIMESTAMP
);

create table if not exists candidates
(
    id                  serial primary key,
    user_id             int           not null,
    cv_url              varchar(512),
    video_url           varchar(512),
    years_of_experience int           not null,
    job_status          job_status    not null,
    seek_job_type       job_type     default 'ANY',
    seek_company_size   company_size default 'ANY',
    seek_location_type  location_type not null,
    seek_salary         int,
    seek_values         varchar(1024),
    work_permit         work_permit   not null,
    notice_period       int,
    constraint fk_user foreign key (user_id) references users (id)
);

create table if not exists companies
(
    id serial primary key
    -- TODO: temporary, replace with what Adrianna has on her branch.
);

create table if not exists company_users
(
    id         serial primary key,
    user_id    int       not null,
    company_id int       not null,
    role       user_role not null,
    constraint fk_user foreign key (user_id) references users (id),
    constraint fk_company foreign key (company_id) references companies (id)
);

create table if not exists associations
(
    id          serial primary key,
    name        varchar(100)  not null,
    logo_url    varchar(512)  not null,
    website_url varchar(512)  not null,
    focus       varchar(1024) not null
);

create table if not exists association_users
(
    id             serial primary key,
    user_id        int       not null,
    association_id int       not null,
    role           user_role not null,
    constraint fk_user foreign key (user_id) references users (id),
    constraint fk_association foreign key (association_id) references associations (id)
);

create table if not exists candidate_skills
(
    id           serial primary key,
    candidate_id int         not null,
    name         varchar(64) not null,
    years        int         not null,
    constraint fk_candidate foreign key (candidate_id) references candidates (id)
);

create table if not exists candidate_spoken_languages
(
    id                  serial primary key,
    candidate_id        int          not null,
    language_id         int          not null,
    language_name       varchar(128) not null,
    language_short_name varchar(128) not null,
    level               int          not null,
    constraint fk_candidate foreign key (candidate_id) references candidates (id)
);

create table if not exists candidate_seek_locations
(
    id           serial primary key,
    candidate_id int          not null,
    city_id      int          not null,
    city_name    varchar(128) not null,
    constraint fk_candidate foreign key (candidate_id) references candidates (id)
);

create table if not exists candidate_attachments
(
    id             serial primary key,
    candidate_id   int          not null,
    attachment_url varchar(512) not null,
    constraint fk_candidate foreign key (candidate_id) references candidates (id)
);

create table if not exists candidate_education_history
(
    id           serial primary key,
    candidate_id int          not null,
    title        varchar(128) not null,
    description  varchar(512) not null,
    entity       varchar(128) not null,
    from_date    timestamp    not null,
    to_date      timestamp,
    constraint fk_candidate foreign key (candidate_id) references candidates (id)
);

create table if not exists candidate_employment_history
(
    id           serial primary key,
    candidate_id int          not null,
    title        varchar(128) not null,
    description  varchar(512) not null,
    company      varchar(128) not null,
    from_date    timestamp    not null,
    to_date      timestamp,
    constraint fk_candidate foreign key (candidate_id) references candidates (id)
);
