create table users (
   id         serial primary key,
   name       varchar(255),
   password   varchar(255) not null,
   created_at timestamp not null
);