package db

import (
  "gorm.io/driver/postgres"
  "github.com/ikennarichard/essentials/lib"
  "gorm.io/gorm"
  "os"
)



func ConnectToDb() {
  var dberror error

  dsn := os.Getenv("DATABASE_URL")
  lib.DB, dberror = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if dberror != nil {
    panic("Error connecting to database")
  }
}
