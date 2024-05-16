package validation

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func ValidateNipIt(fl validator.FieldLevel) bool {
	regexNipPattern := `^615[12](200[0-9]|201[0-9]|19[89][0-9]|20[2-9][0-9])(0[1-9]|1[0-2])([0-9]{3})$`
	nipInt := fl.Field().Interface().(int64)

	nipStr := strconv.FormatInt(nipInt, 10)

	rgx := regexp.MustCompile(regexNipPattern)
	return rgx.MatchString(nipStr)
}

func ValidateNipNurse(fl validator.FieldLevel) bool {
	regexNipPattern := `^303[12](200[0-9]|201[0-9]|19[89][0-9]|20[2-9][0-9])(0[1-9]|1[0-2])([0-9]{3})$`
	nipInt := fl.Field().Interface().(int64)

	nipStr := strconv.FormatInt(nipInt, 10)

	rgx := regexp.MustCompile(regexNipPattern)
	return rgx.MatchString(nipStr)
}

func ValidateImage(fl validator.FieldLevel) bool {
	image := fl.Field().String()

	pattern := `http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+\.(?:jpg|jpeg|png|gif|bmp|webp|svg)$`
	rgx := regexp.MustCompile(pattern)
	return rgx.MatchString(image)
}

func ValidatePhoneNumberFormat(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	pattern := `^\+62\d{10,15}$`
	rgx := regexp.MustCompile(pattern)
	return rgx.MatchString(phoneNumber)
}

func ValidateURL(fl validator.FieldLevel) bool {
	url := fl.Field().String()

	regex := `^(https?://)?([a-zA-Z0-9-]+\.){1,}[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=]*)?$`
	match, _ := regexp.MatchString(regex, url)
	return match
}
