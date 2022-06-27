package utils

const InsertNewUserSQL = "INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
const SelectUsernameSQL = "SELECT username FROM users WHERE username=$1"
const SelectEmailSQL = "SELECT email FROM users WHERE email=$1"
