package repositories

import (
	"DevBookAPI/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.Users) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO users(Name, Nick, Email, Password) VALUES(?,?,?,?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (repository Users) FindAllUsers(nameOrNick string) ([]models.Users, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	stmt, err := repository.db.Query("SELECT Id, Name, Nick, Email, CreatedAt FROM users where Name like ? or Nick like ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var users []models.Users
	for stmt.Next() {
		var user models.Users

		err := stmt.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository Users) FindOneUser(id uint64) (models.Users, error) {
	stmt, err := repository.db.Query("SELECT Id, Name, Nick, Email, CreatedAt FROM users WHERE Id = ?", id)
	if err != nil {
		return models.Users{}, err
	}
	defer stmt.Close()
	var user models.Users

	if stmt.Next() {
		err = stmt.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)

		if err != nil {
			return models.Users{}, err
		}
	}

	if user.Id == 0 {
		return models.Users{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (repository Users) UpdateUser(id uint64, user models.Users) (err error) {
	stmt, err := repository.db.Prepare("update users set Name = ?, Nick = ?, Email = ? where Id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Nick, user.Email, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository Users) DeleteUser(id uint64) error {
	stmt, err := repository.db.Prepare("delete FROM users where Id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (repository Users) FindByEmail(email string) (models.Users, error) {
	stmt, err := repository.db.Query("SELECT Id, Password FROM users WHERE Email = ?", email)
	if err != nil {
		return models.Users{}, err
	}
	defer stmt.Close()

	var user models.Users
	if stmt.Next() {
		err = stmt.Scan(&user.Id, &user.Password)
		if err != nil {
			return models.Users{}, err
		}
	}

	if user.Id == 0 {
		return models.Users{}, fmt.Errorf("user not found")
	}

	return user, nil

}

func (repository Users) Follow(userId uint64, followerId uint64) error {
	stmt, err := repository.db.Prepare("insert ignore into followers(user_id, follower_id) values(?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, followerId)
	if err != nil {
		return err
	}

	return nil

}

func (repository Users) Unfollow(userId uint64, followerId uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM followers WHERE follower_id = ? AND user_id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(followerId, userId)
	if err != nil {
		return err
	}

	return nil

}

func (repository Users) FindFollowers(userId uint64) ([]models.Users, error) {
	stmt, err := repository.db.Query("SELECT  f.user_id, u.Nick FROM"+
		" users as u, followers as f "+
		" WHERE u.Id = f.user_id and f.follower_id = ?", userId)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var users []models.Users
	for stmt.Next() {
		var user models.Users

		err := stmt.Scan(&user.Id, &user.Nick)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

func (repository Users) FindById(userId uint64) (string, error) {
	stmt, err := repository.db.Query("SELECT Id, Password FROM users WHERE Id = ?", userId)

	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var user models.Users

	if stmt.Next() {
		err = stmt.Scan(&user.Id, &user.Password)
		if err != nil {
			return "", err
		}
	}

	return user.Password, nil

}

func (repository Users) UpdatePassword(userId uint64, passwordHash []byte) error {
	stmt, err := repository.db.Prepare("UPDATE users SET PASSWORD = ? WHERE Id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(passwordHash, userId)
	if err != nil {
		return err
	}

	return nil
}
