package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	log *zap.Logger
)

func init() {
	l, _ := zap.NewProduction()
	log = l
}

func handle(ctx context.Context, request events.CognitoEventUserPoolsPostAuthentication) (events.CognitoEventUserPoolsPostAuthentication, error) {
	log.Info("a new request for cognito", zap.Any("request", request))

	data, err := json.Marshal(&request.Request)
	if err != nil {
		log.Error("cant marshall request", zap.Error(err))
		return request, fmt.Errorf("cant marshall request %w", err)
	}

	_, err = db.ExecContext(ctx, `insert into user_login_reports(id, user_pool_id, cognito_user_id, region, email, user_attributes, created_on_utc)  VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		uuid.New().String(), request.UserPoolID, request.UserName, request.Region, request.Request.UserAttributes["email"], string(data), time.Now().Format(time.RFC3339))

	if err != nil {
		log.Error("error while inserting", zap.Error(err))
	}

	return request, err
}

func main() {
	connStr := os.Getenv("POSTGRES_CONNECTION")

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("can not connect to database", zap.Error(err))
	}

	db = dbConn

	lambda.Start(handle)

}
