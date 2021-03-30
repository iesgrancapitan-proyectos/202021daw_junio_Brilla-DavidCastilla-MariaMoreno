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
// DesactivateUserQuery need one variable
// - @username
var DeactivateUserQuery string

//go:embed deleteBrillo.aql
//DeleteBrilloQuery need one variable
// -@brillo_key
var DeleteBrilloQuery string

//go:embed getBrilloById.aql
// GetBrilloByIdQuery need one variable
// -@id
var GetBrilloByIdQuery string

//go:embed getComments.aql
// GetCommentsQuery need one variable
// -@brillo
var GetCommentsQuery string

//go:embed interaction.aql
// InteractionQuery need 3 variables
// -@brilloid
// -@type
// -@userid
var InteractionQuery string

//go:embed newFollow.aql
// NewFollowQuery need 2 variables
// -@follower
// -@followed
var NewFollowQuery string

//go:embed rebrillo.aql
// RebrilloQuery need 2 variables
// -@brilloId
// -@userId
var RebrilloQuery string

//go:embed recomendationUser.aql
// RecomendationUserQuery need 1 variables
// -@username
var RecomendationUserQuery string

//go:embed commentBrillo.aql
// CommentQuery need 2 variables
// -@username
// -@content
// -@media
var CommentBrilloQuery string
