CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `follow_count` int DEFAULT '0',
  `follower_count` int DEFAULT '0',
  `is_follow` tinyint(1) DEFAULT '0',
  `avatar` varchar(1000) NOT NULL DEFAULT 'https://cdn.acwing.com/media/user/profile/photo/160535_lg_e4534d8e65.jpg' COMMENT '用户头像',
  `background_image` varchar(1000) NOT NULL DEFAULT 'https://cdn.acwing.com/media/user/profile/photo/160535_lg_e4534d8e65.jpg' COMMENT '用户背景图',
  `signature` varchar(1000) DEFAULT NULL COMMENT '用户简介',
  `total_favorited` varchar(1000) DEFAULT '0' COMMENT '获赞数量',
  `work_count` int DEFAULT '0' COMMENT '作品数',
  `favorite_count` int DEFAULT '0' COMMENT '喜欢数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
