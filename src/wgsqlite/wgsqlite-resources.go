package wgsqlite

type UserQueryStruct struct {
	Username string
	Password string
	Role     string
	Salt     string
}

var (
	tableQueries = [...]string{
		`
		CREATE TABLE IF NOT EXISTS account (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(50) NOT NULL,
			salt VARCHAR(50) NOT NULL
		);`,
		`
		CREATE TABLE IF NOT EXISTS iface (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			addr VARCHAR(50) UNIQUE NOT NULL,
			port VARCHAR(20) UNIQUE NOT NULL,
			privkey VARCHAR(255) NOT NULL
		);`,
		`
		CREATE TABLE IF NOT EXISTS client (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			client VARCHAR(50) UNIQUE NOT NULL,
			privkey VARCHAR(255) NOT NULL,
			pubkey VARCHAR(255) NOT NULL,
			pskkey VARCHAR(255) NOT NULL,
			iface VARCHAR(255) NOT NULL
		);`,
	}
)
