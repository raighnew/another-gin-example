CREATE TABLE
  `courses` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `course_id` varchar(100) DEFAULT NULL,
    `name` longtext,
    `lessons` int DEFAULT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_courses_course_id` (`course_id`),
    KEY `idx_courses_deleted_at` (`deleted_at`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci

INSERT INTO courses (course_id, name, lessons) VALUES
('CS101', 'Intro to Computer Science', 24),
('CS102', 'Intro to Computer Science 2', 24),
('CS103', 'Intro to Computer Science 3', 24),
('CS104', 'Intro to Computer Science 4', 24),
('CS105', 'Intro to Computer Science 5', 24),
('MATH255', 'Calculus I', 10),
('PHYS150', 'General Physics', 15);