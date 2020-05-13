package bst_models

import "net/http"

type Error struct {
	Code int
	CorrespondingHttpCode int
	Message string
}

// common errors 0-99
var (
	ErrorOK = Error {
		Code: 0,
		CorrespondingHttpCode: http.StatusOK,
		Message: "OK",
	}
	ErrorJwt = Error {
		Code: 1,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "invalid token supplied",
	}
	ErrorJwtProfile = Error {
		Code: 1,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "token did not contain correct profile structure",
	}

	ErrorDBConnection = Error {
		Code: 5,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "database connection failed",
	}

	ErrorBadHeader = Error {
		Code: 20,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "eader did not contain correct contents",
	}
	ErrorBadBody = Error {
		Code: 21,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "request body did not contain correct contents",
	}
	ErrorBadQuery = Error {
		Code: 22,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "query did not contain correct contents",
	}
)

// api errors 100-199
var (
	ErrorPrivateUser = Error {
		Code: 110,
		CorrespondingHttpCode: http.StatusForbidden,
		Message: "user does not allow data views",
	}
	ErrorUnknownUser = Error {
		Code: 111,
		CorrespondingHttpCode: http.StatusNotFound,
		Message: "user does not exist in this context",
	}

	ErrorNoUserCache = Error {
		Code: 120,
		CorrespondingHttpCode: http.StatusNotFound,
		Message: "no cache exists for requested user",
	}
	ErrorNoSuppliedUserCache = Error {
		Code: 121,
		CorrespondingHttpCode: http.StatusBadRequest,
		Message: "no user supplied for cache request",
	}
	ErrorWriteUserCache = Error {
		Code: 122,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "could not write profile changes to db",
	}
)

// eagate errors 200-299
var (
	ErrorNoCookie = Error {
		Code: 200,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "no eagate cookie for user",
	}
	ErrorBadCookie = Error {
		Code: 201,
		CorrespondingHttpCode: http.StatusUnauthorized,
		Message: "bad or expired eagate cookie",
	}
)

// ddr errors 300-399
var (
	ErrorDdrSongIds = Error {
		Code: 300,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr song ids from eagate",
	}

	ErrorDdrSongIdsDbWrite = Error {
		Code: 330,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to write ddr song ids to db",
	}

	ErrorDdrSongIdsDbRead = Error {
		Code: 360,
		CorrespondingHttpCode: http.StatusInternalServerError,
		Message: "failed to get ddr song ids from db",
	}
)

// drs errors 400-499