USE ticket_test;

CREATE TABLE `sale` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `amount` float DEFAULT NULL,
  `sale_type` varchar(45) DEFAULT NULL,
  `date_created` datetime DEFAULT NULL,
  `event_id` int(11) DEFAULT NULL,
  `country_id` int(11) DEFAULT NULL,
  `country_name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

