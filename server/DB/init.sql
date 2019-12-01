-- Exported from QuickDBD: https://www.quickdatabasediagrams.com/
-- Link to schema: https://app.quickdatabasediagrams.com/#/d/Owq6ib
-- NOTE! If you have used non-SQL datatypes in your design, you will have to change these here.

-- Modify this code to update the DB schema diagram.
-- To reset the sample schema, replace everything with
-- two dots ('..' - without quotes).

CREATE DATABASE IF NOT EXISTS orenotorero;

-- CREATE TABLE `orenotorero`.`user` (
--     `id` varchar(100)  NOT NULL ,
--     `name` varchar(255)  NOT NULL ,
--     `email` varchar(255)  NOT NULL ,
--     `password` varchar(255)  NOT NULL ,
--     -- S3 URL
--     `img` varchar(255)  NOT NULL ,
--     PRIMARY KEY (
--         `id`
--     )
-- );

-- CREATE TABLE `orenotorero`.`board` (
--     `id` int AUTO_INCREMENT NOT NULL ,
--     `created_user_id` varchar(255)  NOT NULL ,
--     `title` varchar(255)  NOT NULL ,
--     `publish` boolean  NOT NULL ,
--     -- S3リンク
--     `bg_img` varchar(255)  NOT NULL ,
--     -- 最終アクセス日
--     `last_access` datetime  NOT NULL ,
--     `invite_url` varchar(255)  NOT NULL ,
--     PRIMARY KEY (
--         `id`
--     )
-- );

-- CREATE TABLE `orenotorero`.`author` (
--     `id` int AUTO_INCREMENT NOT NULL ,
--     `borard_id` int  NOT NULL ,
--     `user_id` varchar(255)  NOT NULL ,
--     PRIMARY KEY (
--         `id`
--     )
-- );

-- CREATE TABLE `orenotorero`.`kanban` (
--     `id` int AUTO_INCREMENT NOT NULL ,
--     `board_id` int  NOT NULL ,
--     `title` varchar(255)  NOT NULL ,
--     `position` int  NOT NULL ,
--     PRIMARY KEY (
--         `id`
--     )
-- );

-- CREATE TABLE `orenotorero`.`card` (
--     `id` int AUTO_INCREMENT NOT NULL ,
--     `kanban_id` int  NOT NULL ,
--     `title` varchar(255)  NOT NULL ,
--     `describe` varchar(255)  NULL ,
--     `deadline` datetime  NULL ,
--     `position` int  NOT NULL ,
--     PRIMARY KEY (
--         `id`
--     )
-- );

-- CREATE TABLE `orenotorero`.`files` (
--     `id` int AUTO_INCREMENT NOT NULL ,
--     `card_id` int  NOT NULL ,
--     `name` varchar(255)  NOT NULL ,
--     -- S3リンク
--     `url` varchar(255)  NOT NULL ,
--     PRIMARY KEY (
--         `id`
--     )
-- );

-- ALTER TABLE `orenotorero`.`board` ADD CONSTRAINT `fk_board_created_user_id` FOREIGN KEY(`created_user_id`)
-- REFERENCES `orenotorero`.`user` (`id`);

-- ALTER TABLE `orenotorero`.`author` ADD CONSTRAINT `fk_author_borard_id` FOREIGN KEY(`borard_id`)
-- REFERENCES `orenotorero`.`board` (`id`);

-- ALTER TABLE `orenotorero`.`author` ADD CONSTRAINT `fk_author_user_id` FOREIGN KEY(`user_id`)
-- REFERENCES `orenotorero`.`user` (`id`);

-- ALTER TABLE `orenotorero`.`kanban` ADD CONSTRAINT `fk_kanban_board_id` FOREIGN KEY(`board_id`)
-- REFERENCES `orenotorero`.`board` (`id`);

-- ALTER TABLE `orenotorero`.`card` ADD CONSTRAINT `fk_card_kanban_id` FOREIGN KEY(`kanban_id`)
-- REFERENCES `orenotorero`.`kanban` (`id`);

-- ALTER TABLE `orenotorero`.`files` ADD CONSTRAINT `fk_files_card_id` FOREIGN KEY(`card_id`)
-- REFERENCES `orenotorero`.`card` (`id`);

