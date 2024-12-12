-- Создание таблицы categories
CREATE TABLE if not exists categories (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL UNIQUE,
    public BOOLEAN DEFAULT FALSE
);

-- Создание таблицы posts
CREATE TABLE if not exists posts (
    id SERIAL PRIMARY KEY,
    categories_id INTEGER NOT NULL REFERENCES categories(id),
    logo_id INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL UNIQUE,
    public BOOLEAN DEFAULT FALSE
);

-- Создание таблицы analytics
CREATE TABLE IF NOT EXISTS analytics (
    id SERIAL PRIMARY KEY,
    posts_id INTEGER NOT NULL UNIQUE REFERENCES posts(id) ON DELETE CASCADE,
    search_query TEXT NOT NULL,
    parse_at TIMESTAMPTZ,
    vacancies_num INTEGER
);

-- Создание таблицы filters
CREATE TABLE if not exists filters (
    id SERIAL PRIMARY KEY,
    word VARCHAR(100) UNIQUE NOT NULL
);
