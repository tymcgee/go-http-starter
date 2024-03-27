-- name: GetAuthor :one
select * from author where id = ? limit 1;

-- name: ListAuthors :many
select * from author order by name;
