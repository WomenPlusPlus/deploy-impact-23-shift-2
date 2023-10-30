insert into associations (name, logo_url, website_url, focus)
values ('name', 'logo', 'website', 'focus');
insert into companies (name, contact_email, contact_phone, address, mission, values, job_types)
values ('name', 'contact_email', 'contact_phone', 'address', 'mission', 'values', 'job_types');

insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ADMIN', 'Shift 2', 'Team', 'SHIFT-ADMIN', 'shift2.deployimpact+admin@gmail.com', '999 000 333',
        '2023-10-15 00:00:00');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('CANDIDATE', 'Shift 2', 'Team', 'SHIFT-CANDIDATE', 'shift2.deployimpact+candidate@gmail.com', '999 000 333',
        '2023-10-15 00:00:00');
insert into candidates (user_id, years_of_experience, job_status, seek_location_type, work_permit)
values (2, 5, 'OPEN_TO', 'HYBRID', 'WORK_VISA');
insert into candidate_seek_locations (candidate_id, city_id, city_name)
values (1, 32832, 'Coimbra');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ASSOCIATION', 'Shift 2', 'Team', 'SHIFT-ASSOCIATION-ADMIN', 'shift2.deployimpact+assadmin@gmail.com',
        '999 000 333', '2023-10-15 00:00:00');
insert into association_users (user_id, association_id, role)
values (3, 1, 'ADMIN');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ASSOCIATION', 'Shift 2', 'Team', 'SHIFT-ASSOCIATION-USER', 'shift2.deployimpact+assuser@gmail.com',
        '999 000 333', '2023-10-15 00:00:00');
insert into association_users (user_id, association_id, role)
values (4, 1, 'USER');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('COMPANY', 'Shift 2', 'Team', 'SHIFT-COMPANY-ADMIN', 'shift2.deployimpact+compadmin@gmail.com', '999 000 333',
        '2023-10-15 00:00:00');
insert into company_users (user_id, company_id, role)
values (5, 1, 'ADMIN');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('COMPANY', 'Shift 2', 'Team', 'SHIFT-COMPANY-USER', 'shift2.deployimpact+compuser@gmail.com', '999 000 333',
        '2023-10-15 00:00:00');
insert into company_users (user_id, company_id, role)
values (6, 1, 'USER');

insert into invites (id, creator_id, kind, role, entity_id, email, state, ticket, expire_at, created_at)
values (default, 1, 'CANDIDATE', null, null, 'joaordev+candidate@gmail.com',
        'ACCEPTED', null, '2023-10-31 09:08:49.000000', default);
insert into invites (id, creator_id, kind, role, entity_id, email, state, ticket, expire_at, created_at)
values (default, 1, 'ASSOCIATION', 'ADMIN', null, 'joaordev+assadmin@gmail.com',
        'PENDING', null, '2023-10-31 18:09:57.000000', default);
