CREATE TABLE `Games` (
  `game_key` varchar(20) NOT NULL,
  `player_name` varchar(255) NOT NULL,
  `difficulty` varchar(50) NOT NULL,
  `create_date` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `current_turn` tinyint(1) DEFAULT 0,
  `board` varchar(18) DEFAULT 'E|E|E|E|E|E|E|E|E',
  `winner` tinyint(2) DEFAULT 0 COMMENT '0: No Winner yet 1: Player Wins 2: AI Wins 3: Tie',
  `finish_date` datetime DEFAULT NULL,
  PRIMARY KEY (`game_key`),
  KEY `game_key` (`game_key`)
)