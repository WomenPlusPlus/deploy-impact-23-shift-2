insert into associations (id, name, logo_url, website_url, focus)
values (1, 'name', 'logo', 'website', 'focus');
insert into companies (id)
values (1);

insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ADMIN', 'Shift 2', 'Team', 'SHIFT-ADMIN', 'shift2.deployimpact+admin@gmail.com', '999 000 333', '2023-10-15 00:00:00');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('CANDIDATE', 'Shift 2', 'Team', 'SHIFT-CANDIDATE', 'shift2.deployimpact+candidate@gmail.com', '999 000 333', '2023-10-15 00:00:00');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ASSOCIATION', 'Shift 2', 'Team', 'SHIFT-ASSOCIATION-ADMIN', 'shift2.deployimpact+assadmin@gmail.com', '999 000 333', '2023-10-15 00:00:00');
insert into association_users (user_id, association_id, role)
values (3, 1, 'ADMIN');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ASSOCIATION', 'Shift 2', 'Team', 'SHIFT-ASSOCIATION-USER', 'shift2.deployimpact+assuser@gmail.com', '999 000 333', '2023-10-15 00:00:00');
insert into association_users (user_id, association_id, role)
values (4, 1, 'USER');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('COMPANY', 'Shift 2', 'Team', 'SHIFT-COMPANY-ADMIN', 'shift2.deployimpact+compadmin@gmail.com', '999 000 333', '2023-10-15 00:00:00');
insert into company_users (user_id, company_id, role)
values (5, 1, 'ADMIN');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('COMPANY', 'Shift 2', 'Team', 'SHIFT-COMPANY-USER', 'shift2.deployimpact+compuser@gmail.com', '999 000 333', '2023-10-15 00:00:00');
insert into company_users (user_id, company_id, role)
values (6, 1, 'USER');
