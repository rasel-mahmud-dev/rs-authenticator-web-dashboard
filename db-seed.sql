CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';



CREATE TABLE public.users
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username   VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ      DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX users_email_key ON users USING btree (email);

CREATE TABLE public.two_fa_secrets
(
    id                  UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id             UUID NOT NULL REFERENCES public.users (id) ON DELETE CASCADE,
    secret              TEXT NOT NULL,                      -- Store encrypted or hashed secret here
    failed_attempts     INT                      DEFAULT 0, -- Tracks number of failed attempts
    last_failed_attempt TIMESTAMP WITH TIME ZONE,           -- Timestamp for last failed attempt
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id)                                        -- Ensure only one 2FA secret per user
);


CREATE TABLE public.user_login_attempts
(
    id           UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id      UUID    NOT NULL REFERENCES public.users (id) ON DELETE CASCADE,
    attempt_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    success      BOOLEAN NOT NULL, -- Whether the login attempt was successful
    ip_address   VARCHAR(45),      -- Store IP address (IPv6 support)
    user_agent   VARCHAR(255),     -- Store the user agent string
    description  VARCHAR,          -- Store the user agent string
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE public.auth_sessions
(
    id            UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id       UUID    NOT NULL REFERENCES public.users (id) ON DELETE CASCADE,
    ip_address    VARCHAR(45),
    user_agent    VARCHAR(255),
    access_token  VARCHAR NOT NULL,
    refresh_token VARCHAR NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_used_at  TIMESTAMP WITH TIME ZONE,
    is_revoked    BOOLEAN                  DEFAULT FALSE
);

CREATE INDEX idx_auth_sessions_access_token ON public.auth_sessions (access_token);
CREATE INDEX idx_auth_sessions_refresh_token ON public.auth_sessions (refresh_token);


-- Archived sessions table (history)
CREATE TABLE public.auth_session_history
(
    id            UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id       UUID NOT NULL REFERENCES public.users (id) ON DELETE CASCADE,
    ip_address    VARCHAR(45),
    user_agent    VARCHAR(255),
    access_token  VARCHAR,
    refresh_token VARCHAR,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE,
    status        VARCHAR(20),
    last_used_at  TIMESTAMP WITH TIME ZONE,
    is_revoked    BOOLEAN,
    archived_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);