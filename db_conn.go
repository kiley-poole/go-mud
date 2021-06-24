func dbConnect(path := "world.db"){
	options :=
		"?" + "_busy_timeout=10000" +
			"&" + "_foreign_keys=ON"
	db, err := sql.Open("sqlite3", path+options)
	if err != nil {
		// handle the error here
	}   
}