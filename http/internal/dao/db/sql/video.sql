CREATE TABLE `video` (
  `id` int NOT NULL AUTO_INCREMENT,
  `create_time` int NOT NULL,
  `author_id` int NOT NULL,
  `play_url` varchar(1000) NOT NULL,
  `cover_url` varchar(1000) NOT NULL,
  `favorite_count` int NOT NULL DEFAULT '0',
  `comment_count` int NOT NULL DEFAULT '0',
  `is_favorite` tinyint(1) NOT NULL DEFAULT '0',
  `title` varchar(1000) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `author_id` (`author_id`),
  CONSTRAINT `video_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
