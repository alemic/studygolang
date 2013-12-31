-- 角色表
INSERT INTO `role`(`name`) VALUES('站长');
INSERT INTO `role`(`name`) VALUES('副站长');
INSERT INTO `role`(`name`) VALUES('超级管理员');
INSERT INTO `role`(`name`) VALUES('社区管理员');
INSERT INTO `role`(`name`) VALUES('资源管理员');
INSERT INTO `role`(`name`) VALUES('酷站管理员');
INSERT INTO `role`(`name`) VALUES('高级会员');
INSERT INTO `role`(`name`) VALUES('中级会员');
INSERT INTO `role`(`name`) VALUES('初级会员');

-- 帖子节点表
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(0, '微言', '微言');

INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(1, '触动心扉', '触动你心扉的那句话');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(1, '缘分', '天注定');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(1, '微小说', '');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(1, '醍醐灌顶', '你在或者不在');


INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(0, '微语', '微语录露露');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(6, '语录', '');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(6, '1句话', '');

INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(0, '微博微信', '经典');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(9, '微博经典', '');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(9, '微信经典', '');

INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(0, '知乎豆瓣', '分享生活、学习、工作等方方面面');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(12, '知乎名言', '');
INSERT INTO `topics_node`(`parent`, `name`, `intro`) VALUES(12, '豆瓣名言', '');


INSERT INTO `resource_category`(`name`, `intro`) VALUES('精彩文章', '分享来自互联网的精彩文章');
INSERT INTO `resource_category`(`name`, `intro`) VALUES('其他资源', '分享有用的资源');
