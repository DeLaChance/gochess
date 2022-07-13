CREATE TABLE `gochess`.`chess_player` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    `type` ENUM('AI', 'Human') DEFAULT 'AI'
);

CREATE TABLE `gochess`.`chess_game` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    white_player INT NOT NULL,
    black_player INT NOT NULL,
    FOREIGN KEY (`white_player`) REFERENCES `gochess`.`chess_player`(`id`) ON UPDATE RESTRICT ON DELETE CASCADE,
    FOREIGN KEY (`black_player`) REFERENCES `gochess`.`chess_player`(`id`) ON UPDATE RESTRICT ON DELETE CASCADE
);

CREATE TABLE `gochess`.`chess_game_moves` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    chess_game_id INT NOT NULL,
    from_position INT NOT NULL,
    to_position INT NOT NULL,
    FOREIGN KEY (`chess_game_id`) REFERENCES `gochess`.`chess_game`(`id`) ON UPDATE RESTRICT ON DELETE CASCADE
);
