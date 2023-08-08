DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '作品的id',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '作品标题',
  `author_id` int(11) NOT NULL COMMENT '作者的id',
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '视频资源的url',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '封面的url',
  `favorite_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞数量',
  `comment_count` int(11) NOT NULL DEFAULT 0 COMMENT '评论数量',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '逻辑删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `title`(`title`) USING BTREE COMMENT '针对视频标题添加索引',
  UNIQUE KEY `author_id`(`author_id`) USING BTREE COMMENT '针对视频作者设置索引',
  UNIQUE KEY `id`(`id`, `author_id`) USING BTREE COMMENT '针对id和作者id设置联合索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '短视频表';
