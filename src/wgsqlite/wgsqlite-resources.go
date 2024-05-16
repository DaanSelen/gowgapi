package wgsqlite

type UserQueryStruct struct {
	Username string
	Password string
	Role     string
	Salt     string
}

type InterfaceQueryStruct struct {
	Name          string
	Address       string
	Port          string
	Out_Interface string
	PrivKey       string
}

var (
	tableQueries = [...]string{
		`
		CREATE TABLE IF NOT EXISTS account (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(50) NOT NULL,
			salt VARCHAR(50) NOT NULL,
		);`,
		`
		CREATE TABLE IF NOT EXISTS iface (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(50) UNIQUE NOT NULL,
			addr VARCHAR(50) UNIQUE NOT NULL,
			port VARCHAR(20) UNIQUE NOT NULL,
			out_interface VARCHAR(50) NOT NULL,
			privkey VARCHAR(255) NOT NULL,
			description VARCHAR(255)
		);`,
		`
		CREATE TABLE IF NOT EXISTS client (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			client VARCHAR(50) UNIQUE NOT NULL,
			privkey VARCHAR(255) NOT NULL,
			pubkey VARCHAR(255) NOT NULL,
			pskkey VARCHAR(255) NOT NULL,
			iface VARCHAR(255) NOT NULL,
		);`,
	}
)
