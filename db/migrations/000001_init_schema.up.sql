create table if not exists tasks
(
    id          uuid default gen_random_uuid() primary key,
    title       varchar   not null,
    description text,
    created_at  timestamp
);
