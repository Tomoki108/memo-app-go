CREATE TABLE memos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(255),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

INSERT INTO memos (title, body, created_at, updated_at)
VALUES
    ('Python', 'データサイエンスや競技プログラミングの分野で使われます', '2023-10-07 08:00:00', '2023-10-07 08:00:00'),
    ('JavaScript', '主にフロントエンド開発に用いられます', '2023-10-07 11:15:00', '2023-10-07 11:15:00')
;