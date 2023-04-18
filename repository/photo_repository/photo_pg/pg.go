package photo_pg

import (
	"database/sql"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/repository/photo_repository"
)

type photoPG struct {
	db *sql.DB
}

func NewPhotoPG(db *sql.DB) photo_repository.PhotoRepository {
	return &photoPG{
		db: db,
	}
}

const (
	createPhotoQuery = `
	INSERT INTO "photo" 
	(
		title, 
		caption,
		photo_url,
		user_id
	) 
	VALUES ($1,$2,$3,$4) 
	RETURNING 
		id, title, caption, photo_url, user_id 
	`

	getPhotosQuery = `
	SELECT 
		id, title, caption, photo_url, user_id, created_at, updated_at 
	FROM 
		"photo" 
	WHERE 
		user_id=$1
	ORDER BY
		updated_at DESC
	`

	getPhotoByIDQuery = `
	SELECT 
		id, title, caption, photo_url, user_id, created_at, updated_at 
	FROM 
		"photo" 
	WHERE 
		id=$1
	`

	updatePhotoByIdQuery = `
	UPDATE 
		"photo" 
	SET 
		title = $1,
		caption = $2,
		photo_url = $3,
		updated_at = $4 
	WHERE 
		id = $5
	`

	deletePhotoByIdQuery = `
	DELETE FROM 
		"photo"
	WHERE 
		id = $1
	`
)

func (p *photoPG) CreatePhoto(payload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	var photo entity.Photo

	row := p.db.QueryRow(createPhotoQuery, payload.Title, payload.Caption, payload.PhotoURL, payload.UserId)

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserId)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &photo, nil
}

func (p *photoPG) GetPhotos(userId int) ([]*entity.Photo, errs.MessageErr) {
	rows, err := p.db.Query(getPhotosQuery, userId)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var photos []*entity.Photo

	for rows.Next() {
		var photo entity.Photo

		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		photos = append(photos, &photo)
	}

	return photos, nil
}

func (p *photoPG) GetPhotoByID(photoID int) (*entity.Photo, errs.MessageErr) {
	row := p.db.QueryRow(getPhotoByIDQuery, photoID)

	var photo entity.Photo

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("photo not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &photo, nil
}

func (p *photoPG) UpdatePhotoById(payload entity.Photo) errs.MessageErr {
	_, err := p.db.Exec(updatePhotoByIdQuery, payload.Title, payload.Caption, payload.PhotoURL, payload.UpdatedAt, payload.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("photo not found")
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (p *photoPG) DeletePhotoById(photoId int) errs.MessageErr {
	_, err := p.db.Exec(deletePhotoByIdQuery, photoId)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("photo not found")
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
