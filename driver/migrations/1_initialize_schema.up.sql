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

CREATE TYPE sleep_wake_schedule AS ENUM (
    'early bird',
    'night owl'
);

CREATE TYPE cleaning_schedule AS ENUM (
    'weekly',
    'bi-monthly',
    'monthly'
);

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    facebook_id varchar(255),
    google_id varchar(255),
    instagram_id varchar(255),
    spotify_id varchar(255),
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    birthday date,
    gender varchar(255),
    onboarded boolean,
    permission_level smallint,
    created_at timestamp WITH time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
    -- updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_preferences (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id uuid,
    budget FLOAT,
    move_in date,
    duration varchar(255),
    room_type varchar(255),
    job_type varchar(255),
    job_title varchar(255),
    created_at timestamp WITH time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
    FOREIGN KEY (user_id) REFERENCES public.USERS (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_roommate_preferences (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id uuid,
    FOREIGN KEY (user_id) REFERENCES public.USERS (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_photos (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id uuid,
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
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    country varchar(255),
    state varchar(255),
    city varchar(255),
    district varchar(255),
    latitude FLOAT,
    longitude FLOAT,
    zip_code varchar(6)
);

