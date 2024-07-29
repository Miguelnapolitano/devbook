package repositories

import (
	"api/src/models"
	"database/sql"

)


//publications is publications repo
type Publications struct {
	db *sql.DB
}

//Creates a publications repo
func NewPublicationsRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

//Inserts an publication at database
func (repo Publications) Creates(publication models.Publication) (uint64, error) {
	statement, err := repo.db.Prepare("insert into publications (title, content, author_id) values (?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(insertedID), nil

}

func  (repo Publications) RetrieveByID(publicationID uint64) (models.Publication, error) {
	row, err := repo.db.Query(`
		select p.*, u.nick from
		publications p join users u
		on u.id = p.author_id 
		where p.id = ?`,
		publicationID,
	)
	if err != nil{
		return models.Publication{}, err
	}
	defer row.Close()

	var publication models.Publication

	if row.Next(){
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.Created_at,
			&publication.AuthorNick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}

func  (repo Publications) List(userID uint64) ([]models.Publication, error) {
	rows, err := repo.db.Query(`
		select distinct p.*, u.nick 
		from publications p 
		join users u
		on u.id = p.author_id 
		join followers f
		on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?
		order by 1 desc`,
		userID, userID,
	)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next(){
		var publication models.Publication
		
		if err = rows.Scan(
				&publication.ID,
				&publication.Title,
				&publication.Content,
				&publication.AuthorID,
				&publication.Likes,
				&publication.Created_at,
				&publication.AuthorNick,
			); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}


func (repo Publications) Updates(publicationID uint64, publication models.Publication) error {
	statement, err := repo.db.Prepare("update publications set title = ?, content = ? where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}
	
	return nil

}

func (repo Publications) Delete(publicationID uint64) error {
	statement, err := repo.db.Prepare("delete from publications where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}
	
	return nil

}

func (repo Publications) ListByUser(userID uint64) ([]models.Publication, error) {
	rows, err := repo.db.Query(`
		select distinct p.*, u.nick 
		from publications p 
		join users u
		on u.id = p.author_id
		where p.author_id = ?
		order by 1 desc`,
		userID,
	)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next(){
		var publication models.Publication
		
		if err = rows.Scan(
				&publication.ID,
				&publication.Title,
				&publication.Content,
				&publication.AuthorID,
				&publication.Likes,
				&publication.Created_at,
				&publication.AuthorNick,
			); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil

}

func (repo Publications) Like(publicationID uint64) error {
	statement, err := repo.db.Prepare("update publications set likes = likes + 1 where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}
	
	return nil

}

func (repo Publications) Unlike(publicationID uint64) error {
	statement, err := repo.db.Prepare(`update publications set likes = 
	CASE WHEN likes > 0 THEN likes - 1
	ELSE likes END
	where id = ?`)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}
	
	return nil

}