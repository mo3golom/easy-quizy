create table if not exists easy_quizy_game (
    id UUID primary key not null,
    payload jsonb not null,
    created_at TIMESTAMPTZ not null default NOW()
)