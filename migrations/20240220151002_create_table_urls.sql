-- +goose Up
-- +goose StatementBegin
create table urls
(
    id  bigserial primary key ,
    url text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table urls
-- +goose StatementEnd
