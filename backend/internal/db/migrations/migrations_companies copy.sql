
drop type if exists company_additional_locations;
drop type if exists company_logos;
drop type if exists companies;


create table if not exists companies
(
    id serial primary key,
    company_name varchar(512) not null,
    linkedin_url varchar(512),
    kununu_url varchar(512),
    website_url varchar(512),
    contact_person_name varchar(512),
    email varchar(512),
    phone varchar(512),
    company_size default 'ANY',
    country varchar(128),
    address_line1 varchar(128),
    city varchar(128),
   postal_code varchar(64),
   street varchar(128)
   number_address varchar(64)
   mission varchar(1024)
   comapny_values varchar(1024)
   job_types varchar(512),
   created_at  timestamp    not null default CURRENT_TIMESTAMP
);

create table if not exists company_additional_locations
(
    id           serial primary key,
    company_id int          not null,
    city_id      int          not null,
    city_name    varchar(128) not null,
    constraint fk_company foreign key (company_id) references companies (id)
);

create table if not exists company_logos
(
    id        serial primary key,
    company_id   int          not null,
    logo_url varchar(512) not null,
    constraint fk_company foreign key (company_id) references companies (id)
);
