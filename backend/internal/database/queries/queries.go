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

//go:embed deactivateUser.aql
var DeactivateUserQuery string

//go:embed deleteBrillo.aql
var DeleteBrilloQuery string

//go:embed getBrilloById.aql
var GetBrilloByIdQuery string

//go:embed getComments.aql
var GetCommentsQuery string

//go:embed interaction.aql
var InteractionQuery string

//go:embed newFollow.aql
var NewFollowQuery string

//go:embed rebrillo.aql
var RebrilloQuery string

//go:embed recomendationUser.aql
var RecomendationUserQuery string

//go:embed commentBrillo.aql
var CommentBrilloQuery string
