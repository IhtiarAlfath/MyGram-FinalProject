package comment_pg

import (
	"database/sql"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/repository/comment_repository"
)

type commentPG struct {
	db *sql.DB
}

func NewCommentPG(db *sql.DB) comment_repository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

const (
	createCommentQuery = `
	INSERT INTO "comment" 
	(
		user_id,
		photo_id,
		message
	)
	VALUES ($1,$2,$3) 
	RETURNING 
		id, user_id, photo_id, message
	`

	getCommentsQuery = `
	SELECT 
		id, user_id, photo_id, message, created_at, updated_at
	FROM 
		"comment" 
	WHERE 
		photo_id = $1
	ORDER BY
		updated_at DESC
	`

	getCommentsByIdQuery = `
	SELECT 
		id, user_id, photo_id, message, created_at, updated_at
	FROM 
		"comment" 
	WHERE 
		photo_id = $1 AND
		id = $2
	`

	updateCommentQuery = `
	UPDATE
		"comment"
	SET
		message = $1,
		updated_at = $2
	WHERE 
		photo_id = $3 AND
		id = $4
	`

	deleteCommentQuery = `
	DELETE FROM
		"comment"
	WHERE
		photo_id = $1 AND
		id = $2
	`
)

func (c *commentPG) CreateComment(payload entity.Comment) (*entity.Comment, errs.MessageErr) {
	row := c.db.QueryRow(createCommentQuery, payload.UserId, payload.PhotoId, payload.Message)

	var newComment entity.Comment

	err := row.Scan(&newComment.Id, &newComment.UserId, &newComment.PhotoId, &newComment.Message)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &newComment, nil
}

func (c *commentPG) GetComments(photoId int) ([]*entity.Comment, errs.MessageErr) {
	rows, err := c.db.Query(getCommentsQuery, photoId)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	var comments []*entity.Comment

	for rows.Next() {
		var comment entity.Comment

		err := rows.Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		comments = append(comments, &comment)
	}

	return comments, nil
}

func (c *commentPG) GetCommentById(photoId, commentId int) (*entity.Comment, errs.MessageErr) {
	row := c.db.QueryRow(getCommentsByIdQuery, photoId, commentId)

	var comment entity.Comment

	err := row.Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("comment not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) UpdateCommentById(payload entity.Comment) errs.MessageErr {
	_, err := c.db.Exec(updateCommentQuery, payload.Message, payload.UpdatedAt, payload.PhotoId, payload.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("comment not found")
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (c *commentPG) DeleteCommentById(photoId, commentId int) errs.MessageErr {
	_, err := c.db.Exec(deleteCommentQuery, photoId, commentId)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("comment not found")
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
