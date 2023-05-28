package pg

func ForUpdate(sql string) string {
	return sql + " for update"
}
