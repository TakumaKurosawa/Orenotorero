-- Exported from QuickDBD: https://www.quickdatabasediagrams.com/
-- Link to schema: https://app.quickdatabasediagrams.com/#/d/Owq6ib
-- NOTE! If you have used non-SQL datatypes in your design, you will have to change these here.

-- Modify this code to update the DB schema diagram.
-- To reset the sample schema, replace everything with
-- two dots ('..' - without quotes).

CREATE TABLE `user` (
    `id` varchar(100)  NOT NULL ,
    `name` varchar(255)  NOT NULL ,
    `email` varchar(255)  NOT NULL ,
    `password` varchar(255)  NOT NULL ,
    PRIMARY KEY (
        `id`
    )
);

CREATE TABLE `board` (
    `id` int AUTO_INCREMENT NOT NULL ,
    `created_user_id` string  NOT NULL ,
    `title` string  NOT NULL ,
    `publish` boolean  NOT NULL ,
    -- S3リンク
    `bg_img` string  NOT NULL ,
    -- 最終アクセス日
    `last_access` datetime  NOT NULL ,
    `invite_url` string  NOT NULL ,
    PRIMARY KEY (
        `id`
    )
);

CREATE TABLE `author` (
    `id` int AUTO_INCREMENT NOT NULL ,
    `borard_id` int  NOT NULL ,
    `user_id` string  NOT NULL ,
    PRIMARY KEY (
        `id`
    )
);

CREATE TABLE `kanban` (
    `id` int AUTO_INCREMENT NOT NULL ,
    `board_id` int  NOT NULL ,
    `title` string  NOT NULL ,
    `position` int  NOT NULL ,
    PRIMARY KEY (
        `id`
    )
);

CREATE TABLE `card` (
    `id` int AUTO_INCREMENT NOT NULL ,
    `kanban_id` int  NOT NULL ,
    `title` string  NOT NULL ,
    `describe` string  NULL ,
    `deadline` datetime  NULL ,
    `position` int  NOT NULL ,
    PRIMARY KEY (
        `id`
    )
);

CREATE TABLE `files` (
    `id` int AUTO_INCREMENT NOT NULL ,
    `card_id` int  NOT NULL ,
    `name` string  NOT NULL ,
    -- S3リンク
    `url` string  NOT NULL ,
    PRIMARY KEY (
        `id`
    )
);

ALTER TABLE `board` ADD CONSTRAINT `fk_board_created_user_id` FOREIGN KEY(`created_user_id`)
REFERENCES `user` (`id`);

ALTER TABLE `author` ADD CONSTRAINT `fk_author_borard_id` FOREIGN KEY(`borard_id`)
REFERENCES `board` (`id`);

ALTER TABLE `author` ADD CONSTRAINT `fk_author_user_id` FOREIGN KEY(`user_id`)
REFERENCES `user` (`id`);

ALTER TABLE `kanban` ADD CONSTRAINT `fk_kanban_board_id` FOREIGN KEY(`board_id`)
REFERENCES `board` (`id`);

ALTER TABLE `card` ADD CONSTRAINT `fk_card_kanban_id` FOREIGN KEY(`kanban_id`)
REFERENCES `kanban` (`id`);

ALTER TABLE `files` ADD CONSTRAINT `fk_files_card_id` FOREIGN KEY(`card_id`)
REFERENCES `card` (`id`);

