package validation

import (
	"Musicstreamuser/dto"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateUser(user *dto.RegisterRequest) error {
	validate := validator.New()
	validate.RegisterValidation("must_contain_sepcial_character", isSpecialCharacterPresent)
	return validate.Struct(user)
}

func isSpecialCharacterPresent(f validator.FieldLevel) bool {
	specialCharactersAllowed := `[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",<>\./\?\\|~]`
	regex := regexp.MustCompile(specialCharactersAllowed)
	return regex.MatchString(f.Field().String())

}
