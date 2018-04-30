package SqlQuery

const (
	UserGetQuery = `
		SELECT id,name,email,password,phoneNumber
		FROM users
		where lower(email) = '%s'
	`

	UserExistsQuery = `
		SELECT count(1) as "count"
		FROM users where
		name like '%%%s%%' or email like '%%%s%%'
	`

	CreateUserQuery = `
		INSERT INTO
		users(name,email,password,phoneNumber)
		VALUES ('%s', '%s', '%s', '%s');
	`
)

