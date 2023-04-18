package socialmedia_pg

import (
	"database/sql"
	"kominfo/h8-myGram-finalproject/entity"
	"kominfo/h8-myGram-finalproject/pkg/errs"
	"kominfo/h8-myGram-finalproject/repository/socialmedia_repository"
)

type socialMediaPG struct {
	db *sql.DB
}

func NewSocialMediaPG(db *sql.DB) socialmedia_repository.SocialMediaRepository {
	return &socialMediaPG{
		db: db,
	}
}

const (
	createSocialmediaQuery = `
	INSERT INTO "socialmedia"
	(
		name,
		social_media_url,
		user_id
	)
	VALUES ($1, $2, $3)
	RETURNING 
		id, name, social_media_url, user_id
	`

	getSocialMediasQuery = `
	SELECT 
		id, name, social_media_url, user_id, created_at, updated_at 
	FROM 
		"socialmedia" 
	WHERE 
		user_id = $1
	ORDER BY
		updated_at DESC
	`

	getSocialMediaByIdQuery = `
	SELECT 
		id, name, social_media_url, user_id, created_at, updated_at 
	FROM 
		"socialmedia" 
	WHERE 
		id = $1
	`

	updateSocialMediaQuery = `
	UPDATE 
		"socialmedia"
	SET 
		name = $1,
		social_media_url = $2,
		updated_at = $3 
	WHERE 
		id = $4
	`

	deleteSocialMediaQuery = `
	DELETE FROM 
		"socialmedia"
	WHERE 
		id = $1
	`
)

func (s *socialMediaPG) CreateSocialMedia(payload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	row := s.db.QueryRow(createSocialmediaQuery, payload.Name, payload.SocialMediaURL, payload.UserId)

	var socialmedia entity.SocialMedia

	err := row.Scan(&socialmedia.Id, &socialmedia.Name, &socialmedia.SocialMediaURL, &socialmedia.UserId)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &socialmedia, nil
}

func (s *socialMediaPG) GetSocialMedias(userId int) ([]*entity.SocialMedia, errs.MessageErr) {
	rows, err := s.db.Query(getSocialMediasQuery, userId)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var socialMedias []*entity.SocialMedia

	for rows.Next() {
		var socialMedia entity.SocialMedia

		err := rows.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaURL, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		socialMedias = append(socialMedias, &socialMedia)
	}

	return socialMedias, nil
}

func (s *socialMediaPG) GetSocialMediabById(socialMediaId int) (*entity.SocialMedia, errs.MessageErr) {
	row := s.db.QueryRow(getSocialMediaByIdQuery, socialMediaId)

	var socialMedia entity.SocialMedia

	err := row.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaURL, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("photo not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &socialMedia, nil
}

func (s *socialMediaPG) UpdateSocialMediaById(payload *entity.SocialMedia) errs.MessageErr {
	_, err := s.db.Exec(updateSocialMediaQuery, payload.Name, payload.SocialMediaURL, payload.UpdatedAt, payload.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("product not found")
		}

		return errs.NewInternalServerError(err.Error())
	}

	return nil
}

func (s *socialMediaPG) DeleteSocialMediaById(socialMediaId int) errs.MessageErr {
	_, err := s.db.Exec(deleteSocialMediaQuery, socialMediaId)

	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NewNotFoundError("product not found")
		}

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
