package wgsqlite

func SaveAccount(username, password string) {
	
	prep := wgdb.Prepare("INSERT INTO user (username, password, role, salt) 
	VALUES ('?', '?', '?', '?');")
	wgdb.Exec()
}
