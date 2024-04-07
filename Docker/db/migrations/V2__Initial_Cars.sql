CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    reg_num VARCHAR(20) UNIQUE NOT NULL,
    mark VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    year INTEGER,
    owner_id INTEGER,
    FOREIGN KEY (owner_id) REFERENCES people(id)
);
