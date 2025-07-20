-- create database vk;


-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
                                     id UUID PRIMARY KEY,
                                     username TEXT NOT NULL UNIQUE,
                                     email TEXT NOT NULL UNIQUE,
                                     password_hash TEXT NOT NULL,
                                     created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Таблица объявлений
CREATE TABLE IF NOT EXISTS ads (
                                   id UUID PRIMARY KEY,
                                   title TEXT NOT NULL,
                                   description TEXT NOT NULL,
                                   price NUMERIC NOT NULL,
                                   created_at TIMESTAMP NOT NULL DEFAULT now(),
                                   author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- Таблица изображений объявлений
CREATE TABLE IF NOT EXISTS ad_images (
                                         id UUID PRIMARY KEY,
                                         ad_id UUID NOT NULL REFERENCES ads(id) ON DELETE CASCADE,
                                         image_url TEXT NOT NULL,
                                         created_at TIMESTAMP NOT NULL DEFAULT now()
);



-- Таблица сессий (авторизация через токены)
CREATE TABLE IF NOT EXISTS sessions (
                                        token TEXT PRIMARY KEY,
                                        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                        created_at TIMESTAMP NOT NULL DEFAULT now(),
                                        expires_at TIMESTAMP NOT NULL
);
