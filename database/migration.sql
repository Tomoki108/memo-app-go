CREATE TABLE trees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);


INSERT INTO trees (title, description, created_at, updated_at)
VALUES
    ('Python', 'Pythonプログラミング言語の基本をマスターします。', '2023-10-07 08:00:00', '2023-10-07 08:00:00'),
    ('JavaScript', 'JavaScriptのフロントエンド開発を学びます。', '2023-10-07 11:15:00', '2023-10-07 11:15:00'),
    ('AWSクラウド', 'Amazon Web Servicesの基本を理解します。', '2023-10-07 12:45:00', '2023-10-07 12:45:00')
;