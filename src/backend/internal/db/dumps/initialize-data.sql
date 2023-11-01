INSERT INTO associations (name, logo_url, website_url, focus)
VALUES ('Tech For All', 'associations/1/logo/image.jpg', 'https://techforall.org', 'Promoting tech education in underserved areas');

INSERT INTO companies (name, logo_url, contact_email, contact_phone, address, mission, values, job_types)
VALUES ('NexaTech Solutions', 'companies/1/logo/image.jpg', 'info@nexatech.com', '+1-123-456-7890', '123 Tech Lane, Silicon Valley, CA', 'To provide innovative tech solutions that empower businesses', 'Integrity, Innovation, Teamwork', 'FULL_TIME, PART_TIME, INTERNSHIP');

insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ADMIN', 'Shift 2', 'Team', 'Admin', 'shift2.deployimpact+admin@gmail.com', '999 000 333',
        '2023-10-15 00:00:00');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('CANDIDATE', 'Shift 2', 'Team', 'Candidate', 'shift2.deployimpact+candidate@gmail.com', '999 000 333',
        '2023-10-15 00:00:00');
insert into candidates (user_id, years_of_experience, job_status, seek_location_type, work_permit)
values (22, 5, 'OPEN_TO', 'HYBRID', 'WORK_VISA');
insert into candidate_seek_locations (candidate_id, city_id, city_name)
values (7, 32832, 'Coimbra');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ASSOCIATION', 'Shift 2', 'Team', 'Association Admin', 'shift2.deployimpact+assadmin@gmail.com',
        '999 000 333', '2023-10-15 00:00:00');
insert into association_users (user_id, association_id, role)
values (23, 21, 'ADMIN');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('ASSOCIATION', 'Shift 2', 'Team', 'Association User', 'shift2.deployimpact+assuser@gmail.com',
        '999 000 333', '2023-10-15 00:00:00');
insert into association_users (user_id, association_id, role)
values (24, 21, 'USER');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('COMPANY', 'Shift 2', 'Team', 'Company Admin', 'shift2.deployimpact+compadmin@gmail.com', '999 000 333',
        '2023-10-15 00:00:00');
insert into company_users (user_id, company_id, role)
values (25, 21, 'ADMIN');
insert into users (kind, first_name, last_name, preferred_name, email, phone_number, birth_date)
values ('COMPANY', 'Shift 2', 'Team', 'Company User', 'shift2.deployimpact+compuser@gmail.com', '999 000 333',
        '2023-10-15 00:00:00');
insert into company_users (user_id, company_id, role)
values (26, 21, 'USER');
