drop type if exists jobs;
drop type if exists companyusers_jobs;
drop type if exists jobs_skills;
drop type if exists job_locations;
drop type if exists job_languages;



create table if not exists jobs
(
    id serial primary key,
    title varchar(512) not null,
    experience varchar(1024) not null,
    job_type varchar(128)  not null,
    employment_level varchar(128) ,
    overview varchar(1024)  not null,
    role_responibilities varchar(1024) not null,
    nice_to_have varchar(1024),
    candidate_description varchar(1024) ,
    location_type varchar(128) not null,
    salary_range varchar(128),
    benefits varchar(512),
    status boolean,
   start_date timestamp ,
   created_at  timestamp not null default CURRENT_TIMESTAMP
);


create table if not exists companyusers_jobs
(
    id serial primary key,
    job_id int not null,
    companyuser_id int not null,
    constraint fk_job foreign key (job_id) references jobs (id),
    constraint fk_companyuser foreign key (companyuser_id) references company_users (id)

);

create table if not exists jobs_locations
(
    id serial primary key,
    job_id int not null,
    city_id int not null,
    city_name varchar(128) not null,
    constraint fk_job foreign key (job_id) references jobs (id)
);


create table if not exists jobs_skills
(
    id serial primary key,
    job_id int not null,
    name varchar(64) not null,
    experience varchar(64)not null,
    constraint fk_job foreign key (job_id) references jobs (id)
);

create table if not exists jobs_languages
(
    id serial primary key,
    job_id int not null,
    language_id int not null,
    language_name varchar(128) not null,
    language_short_name varchar(128) not null,
    level string not null,
    constraint fk_job foreign key (job_id) references jobs (id)
);


 