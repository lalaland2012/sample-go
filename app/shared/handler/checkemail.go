package handler

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CheckEmail(user string) bool {
	var check bool
	db, err := sql.Open("mysql", "root:@tcp(192.168.200.181:3306)/golang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT email FROM user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if user == value {
				check = true
				break
			} else {
				check = false
			}
		}
		if check == true {
			break
		}
	}
	if check == true {
		return true
	} else {
		return false
	}

}
