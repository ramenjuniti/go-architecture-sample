DROP TABLE IF EXISTS `todo`;

CREATE TABLE `todo` (
    id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL DEFAULT '',
    body TEXT,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO todo (title, body) VALUES
('test1', 'hoge'),
('test2', 'hogehoge'),
('test3', 'hogehogehoge');