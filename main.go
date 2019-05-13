package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kabesan/go-test-sample/infrastructure"
	"github.com/kabesan/go-test-sample/model"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"log"
	"os"
)

func main() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("KABESAMPLE_MYSQL_USER"),
		os.Getenv("KABESAMPLE_MYSQL_PASSWORD"),
		os.Getenv("KABESAMPLE_MYSQL_HOST"),
		os.Getenv("KABESAMPLE_MYSQL_PORT"),
		os.Getenv("KABESAMPLE_MYSQL_DB_NAME"),
	)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := &infrastructure.RDBUserRepository{
		DB: db,
	}

	stripe.DefaultLeveledLogger = &stripe.LeveledLogger{
		Level: 0,
	}
	stripeKey := os.Getenv("KABESAMPLE_STRIPE_KEY")
	stripeAPI := client.New(stripeKey, nil)
	cusRepo := &infrastructure.StripeCustomerRepository{
		StripeAPI: stripeAPI,
	}
	subRepo := &infrastructure.StripeSubscriptionRepository{
		StripeAPI: stripeAPI,
	}

	m := &model.SampleModel{
		UserRepo:         userRepo,
		CustomerRepo:     cusRepo,
		SubscriptionRepo: subRepo,
	}

	userID := os.Args[1]
	planID := os.Args[2]
	result, err := m.RegisterNewSubscription(userID, planID)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result.SubscriptionID)
}
