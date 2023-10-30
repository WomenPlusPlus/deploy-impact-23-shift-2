alter table associations
    add column if not exists deleted boolean default false;

alter table associations
    add column if not exists created_at timestamp default CURRENT_TIMESTAMP;
