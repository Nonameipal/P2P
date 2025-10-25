CREATE TYPE user_role AS ENUM ('USER', 'ADMIN');
UPDATE users SET role = 'ADMIN' WHERE username = 'your';
