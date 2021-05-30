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

//go:embed activateUser.aql
// ActivateUserQuery need one variable
// - @username
var ActivateUserQuery string

//go:embed deactivateUser.aql
// DeactivateUserQuery need one variable
// - @username
var DeactivateUserQuery string

//go:embed deleteBrillo.aql
//DeleteBrilloQuery need one variable
// - @brilloKey
var DeleteBrilloQuery string

//go:embed getBrilloById.aql
// GetBrilloByIdQuery need one variable
// - @id
var GetBrilloByIdQuery string

//go:embed getComments.aql
// GetCommentsQuery need one variable
// - @brillo
var GetCommentsQuery string

//go:embed interaction.aql
// InteractionQuery need 3 variables
// - @brilloKey
// - @type
// - @username
var InteractionQuery string

//go:embed createBrillo.aql
// NewFollowQuery need 2 variables
// - @content
// - @media
// - @username
var NewBrilloQuery string

//go:embed newFollow.aql
// NewFollowQuery need 2 variables
// - @follower
// - @followed
var NewFollowQuery string

//go:embed rebrillo.aql
// RebrilloQuery need 2 variables
// - @brilloKey
// - @username
var RebrilloQuery string

//go:embed recomendationUser.aql
// RecomendationUserQuery need 1 variables
// - @username
var RecomendationUserQuery string

//go:embed commentBrillo.aql
// CommentQuery need 2 variables
// - @username
// - @content
// - @media
// - @brilloKey
var CommentBrilloQuery string

//go:embed getBrillosByAuthor.aql
// GetBrillosByAuthorQuery need one variable
// - @username
var GetBrillosByAuthorQuery string

//go:embed getTimeline.aql
// GetTimelineQuery need 1 variables
// - @username
var GetTimelineQuery string

//go:embed isFollowing.aql
// IsFollowingQuery need 2 variables
// - @follower
// - @followed
var IsFollowingQuery string

//go:embed nFollowers.aql
// nfollowers need 1 variables
// - @username
var NFollowers string

//go:embed nFollowed.aql
// nFollowed need 1 variables
// - @username
var NFollowed string

//go:embed returnEmailExits.aql
// returnEmailExits need 1 variables
// - @email
var ReturnEmailExits string

//go:embed deleteRebrillo.aql
// deleteRebrillo need 1 variables
// - @idBrillo
var DeleteRebrilloQuery string

//go:embed deleteFollow.aql
// deleteFollow need 1 variables
// - @idBrillo
var DeleteFollowQuery string

//go:embed updateUser.aql
// updateUser need 1 variables
// - @idBrillo
var UpdateUserQuery string

//go:embed searchByContent.aql
// deleteRebrillo need 1 variables
// - @query
var SearchContentQuery string

//go:embed searchUser.aql
// deleteRebrillo need 1 variables
// - @query
var SearchUserQuery string
