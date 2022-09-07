create table if not exists user_login_reports
(
    id              text not null primary key,
    user_pool_id    text,
    cognito_user_id text,
    region          text,
    email           text,
    user_attributes text,
    created_on_utc  timestamp default now()
);