package aligo

import "os"

func GetAligoId() string {
	return os.Getenv("ALIGO_ID")
}

func GetAligoPhone() string {
	return os.Getenv("ALIGO_PHONE")
}

func GetAligoApiKey() string {
	return os.Getenv("ALIGO_APIKEY")
}

func GetAligoTest() string {
	return os.Getenv("ALIGO_ISTEST")
}

func GetTestPhone() string {
	return os.Getenv("TEST_PHONE")
}
