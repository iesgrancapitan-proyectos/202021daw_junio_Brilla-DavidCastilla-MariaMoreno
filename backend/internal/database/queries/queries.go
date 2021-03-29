package queries

import _ "embed"

//go:embed insertUser.aql
// InsertUserQuery needs seven variables:
// - @username
// - @email
// - @password: Hashed password
// - @name
// - @bio
// - @birthday: Birthday in seconds
// - @profile_img: Path to profile_img
var InsertUserQuery string

//go:embed getUser.aql
// GetUserQuery needs one variable
// - @username: Username
var GetUserQuery string
