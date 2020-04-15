CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE room_type AS ENUM (
    'master',
    'flex',
    'standard'
);

CREATE TYPE job_type AS ENUM (
    'standard',
    'work from home',
    'self employed'
);

CREATE TYPE cleaning_schedule AS ENUM (
    'weekly',
    'bi-monthly',
    'monthly'
);

CREATE TABLE IF NOT EXISTS gender (
    id smallint PRIMARY KEY,
    gender varchar(255) NOT NULL
);

INSERT INTO gender (id, gender)
    VALUES (0, ''), (1, 'male'), (2, 'female'), (3, 'other');

CREATE TABLE IF NOT EXISTS sleep_wake_schedule (
    id smallint PRIMARY KEY,
    sleep_wake_schedule varchar(255) NOT NULL
);

INSERT INTO sleep_wake_schedule (id, sleep_wake_schedule)
    VALUES (0, ''), (1, 'early_bird'), (2, 'night_owl'), (3, 'neither');

CREATE TABLE IF NOT EXISTS users (
    id varchar(255) PRIMARY KEY,
    facebook_id varchar(255),
    google_id varchar(255),
    instagram_id varchar(255),
    spotify_id varchar(255),
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    birthday date,
    gender smallint,
    onboarded boolean NOT NULL,
    permission_level smallint NOT NULL,
    email_verified boolean NOT NULL,
    created_at timestamp WITH time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_preferences (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id varchar(255),
    budget float NOT NULL,
    move_in date NOT NULL,
    duration varchar(255) NOT NULL,
    job_type varchar(255),
    job_title varchar(255),
    created_at timestamp WITH time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.USERS (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_roommate_preferences (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id varchar(255),
    created_at timestamp WITH time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.USERS (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_photos (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id varchar(255),
    instagram_photo_id varchar(255),
    media_type varchar(255),
    media_link varchar(255),
    permalink varchar(255),
    profile_picture boolean,
    FOREIGN KEY (user_id) REFERENCES public.USERS (id) ON DELETE CASCADE
);

-- CREATE TABLE IF NOT EXISTS user_tracks (
--     user_id uuid,
--     track_id varchar(255),
--     track_name varchar(255),
--     album_id varchar(255),
--     album_name varchar(255),
--     artist_id varchar(255),
--     artist_name varchar(255),
--     image_url_640 varchar(255),
--     image_url_300 varchar(255),
--     image_url_64 varchar(255),
--     preview_url varchar(255),
--     PRIMARY KEY (user_id, track_id),
--     FOREIGN KEY (user_id) REFERENCES public.USERS (id) ON DELETE CASCADE
-- );

CREATE TABLE IF NOT EXISTS neighborhoods (
    id serial PRIMARY KEY,
    country varchar(255),
    state varchar(255),
    city varchar(255),
    district varchar(255),
    latitude float,
    longitude float,
    zip_code varchar(6)
);

