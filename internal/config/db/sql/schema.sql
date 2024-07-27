-- Tabela de Usuários
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    login VARCHAR(255) NOT NULL UNIQUE,
    cpf CHAR(11) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    wallet DOUBLE DEFAULT 0
);

-- Tabela de Tutores
CREATE TABLE IF NOT EXISTS tutors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    login VARCHAR(255) NOT NULL UNIQUE,
    cpf CHAR(11) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    wallet DOUBLE DEFAULT 0
);

-- Tabela de Cursos
CREATE TABLE IF NOT EXISTS courses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price INT NOT NULL,
    tutor_id INT,
    FOREIGN KEY (tutor_id) REFERENCES tutors(id)
);

-- Tabela de Questionários
CREATE TABLE IF NOT EXISTS quizzes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    subject VARCHAR(255) NOT NULL,
    creator_id INT,
    course_id INT,
    FOREIGN KEY (creator_id) REFERENCES tutors(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

-- Tabela de Testes
CREATE TABLE IF NOT EXISTS tests (
    id INT AUTO_INCREMENT PRIMARY KEY,
    question TEXT NOT NULL,
    response TEXT NOT NULL,
    quiz_id INT,
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id)
);

-- Tabela de Associação de Cursos em Andamento para Usuários
CREATE TABLE IF NOT EXISTS user_courses (
    user_id INT,
    course_id INT,
    PRIMARY KEY (user_id, course_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

-- Tabela de Associação de Cursos Oferecidos para Tutores
CREATE TABLE IF NOT EXISTS tutor_courses (
    tutor_id INT,
    course_id INT,
    PRIMARY KEY (tutor_id, course_id),
    FOREIGN KEY (tutor_id) REFERENCES tutors(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

-- Tabela de Associação de Questionários para Cursos
CREATE TABLE IF NOT EXISTS course_quizzes (
    course_id INT,
    quiz_id INT,
    PRIMARY KEY (course_id, quiz_id),
    FOREIGN KEY (course_id) REFERENCES courses(id),
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id)
);
