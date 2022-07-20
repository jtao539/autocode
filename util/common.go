package util

import "database/sql"

func CheckStringNULL(args ...string) bool {
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			return true
		}
	}
	return false
}

func IntToNullInt32(a int) sql.NullInt32 {
	return sql.NullInt32{Int32: int32(a), Valid: true}
}

func StringToNullString(a string) sql.NullString {
	return sql.NullString{String: a, Valid: true}
}
