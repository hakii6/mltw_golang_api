package main

const (
    host     = "localhost"
    database = "test2"
    user     = "user"
    password = "password"
    connect = user + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=utf8&parseTime=true"
)
