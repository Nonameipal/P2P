CREATE TABLE IF NOT EXISTS items
(
    id           SERIAL PRIMARY KEY,
    owner_id      INT NOT NULL,
    category_id    INT NOT NULL,
    title          VARCHAR(255) NOT NULL,
    description     TEXT NOT NULL,
    price_per_day   FLOAT NOT NULL,
    status          VARCHAR(50) NOT NULL DEFAULT 'ACTIVE',
    available_from  DATE NOT NULL,
    available_to    DATE NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);