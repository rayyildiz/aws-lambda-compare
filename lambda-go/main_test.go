package main

import (
	"context"
	"database/sql"
	"github.com/aws/aws-lambda-go/events"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	memDb, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
          t.Fatalf("error while opening in memory, %#v", err)
     }
     ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
     defer cancel()
     _, err = memDb.ExecContext(ctx, `create table user_login_reports
(
     id                 text            not null primary key,
     user_pool_id     text,
     cognito_user_id text,
     region            text,
     email             text,
     user_attributes text,
     created_on_utc  timestamp      not null
);`)
     if err != nil {
          t.Fatalf("error while creating table, %#v", err)
     }

     db = memDb

     resp, err := handle(context.TODO(), events.CognitoEventUserPoolsPostAuthentication{
          CognitoEventUserPoolsHeader: events.CognitoEventUserPoolsHeader{
               UserName:   "uuid",
               UserPoolID: "poolId",
          },
          Request: events.CognitoEventUserPoolsPostAuthenticationRequest{
               UserAttributes: map[string]string{
                    "email": "test@example.com",
               },
          },
          Response: events.CognitoEventUserPoolsPostAuthenticationResponse{},
     })
     if err != nil {
          t.Fatalf("eroor while invoking handler, %v", err)
     }

     if resp.UserName != "uuid" {
          t.Errorf("expected %s, got %s", "uuid", resp.UserName)
     }

     if resp.Request.UserAttributes["email"] != "test@example.com" {
          t.Errorf("expected %s, got %s", "test@example.com", resp.Request.UserAttributes["email"])
     }
}

