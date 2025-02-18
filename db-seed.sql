CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';



CREATE TABLE public.users
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username   VARCHAR(255) NOT NULL,
    avatar     VARCHAR(1024),
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


DROP TABLE if exists user_auth_attempts;
DROP TABLE if exists mfa_security_tokens;
CREATE TABLE mfa_security_tokens
(
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id        UUID         NOT NULL,
    code_name      VARCHAR(1024),
    secret         VARCHAR(255) NOT NULL,
    recovery_codes TEXT[]           DEFAULT NULL,
    qr_code_url    TEXT             DEFAULT NULL,
    is_active      BOOLEAN          DEFAULT FALSE,
    is_init        BOOLEAN          DEFAULT TRUE,
    created_at     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    linked_at      TIMESTAMP        DEFAULT NULL,
    app_name       VARCHAR(100)     DEFAULT 'Google',
    device_info    TEXT             DEFAULT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);


DROP table if exists user_auth_attempts;
CREATE TABLE if not exists user_auth_attempts
(
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code_name       VARCHAR(1024),
    user_id         UUID             default NULL,
    attempt_type    VARCHAR(50) NOT NULL, -- Type of attempt (failed, backup_code, etc.)
    mfa_security_id VARCHAR(255),
    security_token  VARCHAR(16)      default NULL,
    ip_address      VARCHAR(45)      default NULL,
    user_agent      VARCHAR(255)     default NULL,
    last_attempt_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    is_successful   BOOLEAN          DEFAULT FALSE,
    created_at      TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE user_auth_attempts
    ALTER COLUMN user_id DROP NOT NULL,
    ALTER COLUMN user_id SET DEFAULT NULL;

CREATE TYPE auth_method_enum AS ENUM ('authenticator', 'password', 'recovery_code');
ALTER TABLE auth_sessions
    ADD COLUMN IF NOT EXISTS auth_method auth_method_enum default 'password';

ALTER TYPE auth_method_enum ADD VALUE 'recovery_code';
ALTER TYPE auth_provider ADD VALUE 'recovery_code';


DROP TABLE IF EXISTS user_profiles;
CREATE TABLE IF NOT EXISTS user_profiles
(
    user_id    UUID PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
    full_name  VARCHAR(100) DEFAULT NULL,
    birth_date DATE         DEFAULT NULL,
    gender     VARCHAR(10) CHECK (gender IN ('male', 'female', 'other', 'prefer not to say')),
    phone      VARCHAR(20)  DEFAULT NULL,
    location   VARCHAR(150) DEFAULT NULL,
    about_me   TEXT         DEFAULT NULL,
    website    TEXT         DEFAULT NULL,
    avatar     TEXT         DEFAULT NULL,
    cover      TEXT         DEFAULT NULL,

    -- Social Media Links
    facebook   TEXT         DEFAULT NULL,
    twitter    TEXT         DEFAULT NULL,
    linkedin   TEXT         DEFAULT NULL,
    instagram  TEXT         DEFAULT NULL,
    github     TEXT         DEFAULT NULL,
    youtube    TEXT         DEFAULT NULL,
    tiktok     TEXT         DEFAULT NULL,

    created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE if exists user_traffic;
CREATE TABLE if not exists user_traffic
(
    id            SERIAL PRIMARY KEY,
    route_path    TEXT,
    http_method   VARCHAR(10),
    user_agent    TEXT,
    ip_address    VARCHAR(45),
    request_time  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    response_time INTEGER
);


ALTER TABLE user_profiles
    DROP CONSTRAINT user_profiles_gender_check;



CREATE TABLE recovery_codes
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID        NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    code       VARCHAR(10) NOT NULL,
    is_used    BOOLEAN          DEFAULT FALSE,
    expires_at TIMESTAMP   NOT NULL,
    created_at TIMESTAMP        DEFAULT NOW(),
    updated_at TIMESTAMP        DEFAULT NOW()
);

CREATE INDEX idx_recovery_code ON recovery_codes (code);
CREATE INDEX idx_expires_at ON recovery_codes (expires_at);
CREATE INDEX idx_user_id ON recovery_codes (user_id);
