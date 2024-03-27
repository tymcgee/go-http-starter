-- name: GetBooksByAuthor :many
select * from book where author_id = ?;

-- name: GetBooks :many
select * from book;
