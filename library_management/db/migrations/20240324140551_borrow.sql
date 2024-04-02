-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS borrow (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL,
    status TEXT NOT NULL,
    borrow_date DATE NOT NULL,
    return_date DATE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop TABLE borrow;
-- +goose StatementEnd
