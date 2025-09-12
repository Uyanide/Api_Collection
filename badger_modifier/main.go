package main

var (
	dbPath = "../data/db"
)

func main() {
	db := GetDB()

	if err := db.Open(dbPath); err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Set("ip_requests", "114514"); err != nil {
		panic(err)
	}
}
