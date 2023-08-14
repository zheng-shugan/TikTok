CREATE TABLE `user_favorite` (
  `user_id` int DEFAULT NULL,
  `video_id` int DEFAULT NULL,
  KEY `user_favorite_ibfk_1` (`user_id`),
  KEY `user_favorite_ibfk_2` (`video_id`),
  CONSTRAINT `user_favorite_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `user_favorite_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci