CREATE TABLE video_categories(
    video_id INTEGER NOT NULL REFERENCES videos(id) ON DELETE CASCADE,
    category_id INTEGER NOT NULL REFERENCES categories(id)
)