drop type if exists invite_state;
create type invite_state as enum ('CREATED','PENDING','ERROR','ACCEPTED','CANCELLED');

create table if not exists invites
(
    id         serial primary key,
    creator_id int          not null,
    kind       user_kind    not null,
    role       user_role,
    entity_id  int,
    email      varchar(512) not null unique,
    state      invite_state not null default 'CREATED',
    ticket     varchar(512),
    expire_at  timestamp    not null,
    created_at timestamp    not null default CURRENT_TIMESTAMP,
    constraint fk_creator foreign key (creator_id) references users (id)
)
