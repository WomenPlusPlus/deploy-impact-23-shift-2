create table if not exists jobs
(
    id                    serial primary key,
    title                 varchar(512)  not null,
    creator_id            int           not null,
    experience_from       int,
    experience_to         int,
    job_type              varchar(128)  not null,
    employment_level_from int,
    employment_level_to   int,
    overview              varchar(4096) not null,
    role_responsibilities varchar(4096) not null,
    candidate_description varchar(4096) not null,
    location_type         location_type not null,
    salary_range_from     int,
    salary_range_to       int,
    benefits              varchar(4096) not null,
    deleted               boolean       not null default false,
    start_date            timestamp,
    created_at            timestamp     not null default CURRENT_TIMESTAMP,
    constraint fk_creator foreign key (creator_id) references users (id)
);

create table if not exists job_locations
(
    id        serial primary key,
    job_id    int          not null,
    city_id   int          not null,
    city_name varchar(128) not null,
    constraint fk_job foreign key (job_id) references jobs (id)
);


create table if not exists job_skills
(
    id     serial primary key,
    job_id int         not null,
    name   varchar(64) not null,
    years  int         not null,
    constraint fk_job foreign key (job_id) references jobs (id)
);

create table if not exists job_languages
(
    id                  serial primary key,
    job_id              int          not null,
    language_id         int          not null,
    language_name       varchar(128) not null,
    language_short_name varchar(128) not null,
    constraint fk_job foreign key (job_id) references jobs (id)
);


