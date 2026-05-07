ALTER TABLE users
    ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE credentials
    ADD CONSTRAINT credentials_email_key UNIQUE (email);

ALTER TABLE credentials
    DROP CONSTRAINT IF EXISTS credentials_username_key;

DROP TABLE sessions;