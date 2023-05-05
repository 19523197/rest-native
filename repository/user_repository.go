package repository

type UserRepository struct {
	BaseRepository BaseRepository
}

func (u *UserRepository) AuthenticateUser(username, password string) (bool, map[string]interface{}) {
	sql := u.BaseRepository.Sql
	err := sql.QueryRow("SELECT username FROM users WHERE username = ? AND password = ?", username, password).Scan(&username)
	switch {
	case err != nil:
		return false, nil
	case username == "":
		return false, nil
	default:
		m := make(map[string]interface{})
		m["username"] = username
		return true, m
	}
}
