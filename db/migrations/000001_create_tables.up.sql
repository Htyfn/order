CREATE TABLE IF NOT EXISTS order (
    Id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    price       MONEY NOT NULL,
    curr        INT NOT NULL,
    createDate  DATE,
    updateDate  DATE
);

CREATE TABLE IF NOT EXISTS order_user (
    id          INT PRIMARY KEY,
    user        INT NOT NULL
)

CREATE TABLE IF NOT EXISTS order_status_log (
    id          INT NOT NULL,
    status      INT NOT NULL,
    comment     VARCHAR(400),
    statusDate  DATE
)

CREATE INDEX IF NOT EXISTS order_status_log_id ON order_status_log(id);
