package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:example@tcp(172.16.1.5:3306)/testdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    // defer db.Close()

//=========perform a db.Query insert============
    // insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
    // if err != nil {
    //     panic(err.Error())
    // }
    // defer insert.Close()


//=========Execute the query============
type E_tb struct {
    ename   string `json:"ename"`
    epermit string `json:"epermit"`
    eposi   string `json:"eposi"`
}
    results, err := db.Query("SELECT ename, epermit, eposi FROM `pro_kk`.`e_tb` LIMIT 0,10")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var e_tb E_tb
        // for each row, scan the result into our e_tb composite object
        err = results.Scan(&e_tb.ename, &e_tb.epermit, &e_tb.eposi)
        if err != nil {
            //panic(err.Error()) // proper error handling instead of panic in your app
            log.Printf(err.Error())
        }
                // and then print out the e_tb's Name attribute
        log.Printf(e_tb.ename)
        log.Printf(e_tb.epermit)
        log.Printf(e_tb.eposi)
    }


defer db.Close()

}
