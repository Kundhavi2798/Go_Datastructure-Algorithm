package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

type Transaction struct {
	UserID    string
	Amount    float64
	Mode      string
	Timestamp time.Time
}

func generateUserids(n int) []string {
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		ids[i] = fmt.Sprintf("user%d", i)
	}
	return ids
}

func fetchdata(work string, db *sql.DB, startT, endT time.Time) ([]Transaction, error) {
	rows, err := db.Query(`
		SELECT user_id, amount, mode, timestamp
		FROM transactions
		WHERE user_id = $1 AND timestamp BETWEEN $2 AND $3
	`, work, startT, endT)
	if err != nil {
		log.Println("The fetch error", err)
		return nil, err
	}
	defer rows.Close()

	var result []Transaction

	for rows.Next() {
		var tx Transaction
		err := rows.Scan(&tx.UserID, &tx.Amount, &tx.Mode, &tx.Timestamp)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}
		result = append(result, tx)
	}
	return result, nil
}

func workerjobs(work <-chan string, db *sql.DB, wg *sync.WaitGroup, startT, endT time.Time) {
	defer wg.Done()
	for user := range work {
		data, err := fetchdata(user, db, startT, endT)
		if err != nil {
			log.Println("fetchdata error", err)
			continue
		}
		fmt.Printf("The user id: %s and the len of the data: %d\n", user, len(data))
	}
}

func main() {
	connec := "postgres://admin:kundhavi@localhost:5432/users?sslmode=disable"
	db, dbConErr := sql.Open("postgres", connec)
	if dbConErr != nil {
		log.Fatal("connection errors", dbConErr)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Ping error: %v", err)
	}
	defer db.Close()

	userID := generateUserids(1000)
	startTime := time.Now().Add(-24 * time.Hour)
	endTime := time.Now()

	var wg sync.WaitGroup
	workers := make(chan string, len(userID))
	workerSplit := 100

	// Start worker goroutines
	for i := 0; i < workerSplit; i++ {
		wg.Add(1)
		go workerjobs(workers, db, &wg, startTime, endTime)
	}

	for _, user := range userID {
		workers <- user
	}
	close(workers)

	wg.Wait()
}
