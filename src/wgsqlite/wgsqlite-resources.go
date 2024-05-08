package wgsqlite

type UserQueryStruct struct {
	Username string
	Password string
	Role     string
	Salt     string
}

const (
	userTab string = `
	CREATE TABLE IF NOT EXISTS account (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(50) NOT NULL,
		salt VARCHAR(50) NOT NULL
	);`
	ifaceTab string = `
	CREATE TABLE IF NOT EXISTS iface (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		addr VARCHAR(50) UNIQUE NOT NULL,
		port VARCHAR(20) UNIQUE NOT NULL,
		privkey varchar(255) NOT NULL
	)`
	clientTab string = `
	CREATE TABLE IF NOT EXISTS client (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client VARCHAR(50) UNIQUE NOT NULL,
		privkey varchar(255) NOT NULL,
		pubkey varchar(255) NOT NULL,
		pskkey varchar(255) NOT NULL,
		iface varchat(255) NOT NULL
 	);`
)
