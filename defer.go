package main

import (
	"fmt"
  "database/sql"
)



func main(){

  func doSomeInsertsInDB(ctx context.Context, db *sql.DB, insertion1, insertion2 string) (err error){
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
      return err
    }
    defer func() { //Aquesta funcio sexecutara la ultima
      if err == nil {
        err = tx.Commit() // si res ha fallat guardem els canvis que hem fet a la DB
      }
      if err != nil {
        tx.Rollback() // Si alguna insercio de les que venen posteriors falla, tot es tira enrere
      }
    }()
    _, err = tx.ExecContext(ctx, "Insert into foo (val) insertions $1", insertion1) //insereix a la DB la insercio1
    if err != nil {
      return err
    }
  }
}
