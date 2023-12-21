-- Insert seed data for artists
INSERT INTO artists (name) VALUES
    ('Artist 1'),
    ('Artist 2'),
    ('Artist 3');

-- Insert seed data for albums
INSERT INTO albums (title, artist_id, release_date) VALUES
    ('Album 1', (SELECT id FROM artists WHERE name = 'Artist 1'), '2022-01-01'),
    ('Album 2', (SELECT id FROM artists WHERE name = 'Artist 2'), '2022-02-01'),
    ('Album 3', (SELECT id FROM artists WHERE name = 'Artist 3'), '2022-03-01');

-- Insert seed data for users
INSERT INTO users (username, email, password) VALUES
    ('User1', 'user1@example.com', 'password1'),
    ('User2', 'user2@example.com', 'password2'),
    ('User3', 'user3@example.com', 'password3');

-- Insert seed data for songs
INSERT INTO songs (title, artist_id, album_id, duration) VALUES
    ('Song 1', (SELECT id FROM artists WHERE name = 'Artist 1'), (SELECT id FROM albums WHERE title = 'Album 1'), 180),
    ('Song 2', (SELECT id FROM artists WHERE name = 'Artist 2'), (SELECT id FROM albums WHERE title = 'Album 2'), 200),
    ('Song 3', (SELECT id FROM artists WHERE name = 'Artist 3'), (SELECT id FROM albums WHERE title = 'Album 3'), 160);

-- Insert seed data for playlists
INSERT INTO playlists (name, user_id) VALUES
    ('Playlist 1', (SELECT id FROM users WHERE username = 'User1')),
    ('Playlist 2', (SELECT id FROM users WHERE username = 'User2')),
    ('Playlist 3', (SELECT id FROM users WHERE username = 'User3'));
