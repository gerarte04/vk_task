package response

import (
	"log"
	"marketplace/internal/repository"
	"marketplace/internal/usecases"
	pkgErrors "marketplace/pkg/errors"
	"net/http"
)

var (
	errorCodes = map[error]int{
		repository.ErrUserExists: http.StatusBadRequest,
		repository.ErrUserNotFound: http.StatusUnauthorized,
		
		usecases.ErrWrongPassword: http.StatusUnauthorized,
	}
)

func ProcessCreatingRequestError(w http.ResponseWriter, err error, debugMode bool) {
	log.Print(err.Error())

	if !debugMode {
		err = pkgErrors.UnwrapAll(err)
	}

	http.Error(w, err.Error(), http.StatusBadRequest)
}

func ProcessError(w http.ResponseWriter, err error, debugMode bool) {
	log.Print(err.Error())

	if !debugMode {
		err = pkgErrors.UnwrapAll(err)
	}

	code := http.StatusInternalServerError

	if docCode, ok := errorCodes[pkgErrors.UnwrapAll(err)]; ok {
		code = docCode
	}

	http.Error(w, err.Error(), code)
}
