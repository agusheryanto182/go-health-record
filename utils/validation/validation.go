package validation

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// func ValidateNipIt(fl validator.FieldLevel) bool {
// 	regexNipPattern := `^615[12](200[0-9]|201[0-9]|19[89][0-9]|20[2-9][0-9])(0[1-9]|1[0-2])([0-9]{3})[0-9]{3}$`
// 	nipInt := fl.Field().Interface().(int64)

// 	nipStr := strconv.FormatInt(nipInt, 10)

// 	if len(nipStr) < 13 || len(nipStr) > 15 {
// 		return false
// 	}

// 	rgx := regexp.MustCompile(regexNipPattern)
// 	return rgx.MatchString(nipStr)
// }

func ValidateNipIt(fl validator.FieldLevel) bool {
	nipInt := fl.Field().Interface().(int64)
	nipStr := strconv.FormatInt(nipInt, 10)

	if len(nipStr) < 13 || len(nipStr) > 15 {
		return false
	}

	// Check if the first 3 digits start with "615"
	if nipStr[:3] != "615" {
		return false
	}

	// Check if the fourth digit is either "1" (for male) or "2" (for female)
	genderDigit, err := strconv.Atoi(nipStr[3:4])
	if err != nil || (genderDigit != 1 && genderDigit != 2) {
		return false
	}

	// Check if the fifth to eighth digits represent a valid year (between 2000 and the current year)
	yearStr := nipStr[4:8]
	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 2000 || year > 2024 {
		return false
	}

	return err == nil
}

func ValidateNipNurse(fl validator.FieldLevel) bool {
	nipInt := fl.Field().Interface().(int64)
	nipStr := strconv.FormatInt(nipInt, 10)

	if len(nipStr) < 13 || len(nipStr) > 15 {
		return false
	}

	// Check if the first 3 digits start with "615"
	if nipStr[:3] != "303" {
		return false
	}

	// Check if the fourth digit is either "1" (for male) or "2" (for female)
	genderDigit, err := strconv.Atoi(nipStr[3:4])
	if err != nil || (genderDigit != 1 && genderDigit != 2) {
		return false
	}

	// Check if the fifth to eighth digits represent a valid year (between 2000 and the current year)
	yearStr := nipStr[4:8]
	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 2000 || year > 2024 {
		return false
	}

	return err == nil
}

// func ValidateNipNurse(fl validator.FieldLevel) bool {
// 	regexNipPattern := `^303[12](200[0-9]|201[0-9]|19[89][0-9]|20[2-9][0-9])(0[1-9]|1[0-2])([0-9]{3})$`
// 	nipInt := fl.Field().Interface().(int64)

// 	nipStr := strconv.FormatInt(nipInt, 10)

// 	rgx := regexp.MustCompile(regexNipPattern)
// 	return rgx.MatchString(nipStr)
// }

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
