package aligo

import "os"

func GetEmail() string {
	return os.Getenv("SC_EMAIL")
}

func GetPassword() string {
	return os.Getenv("SC_PW")
}

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

func GetLoadingPage() string {
	return os.Getenv("LOADING_PAGE")
}

func GetLastReservedId() string {
	return os.Getenv("LAST_RESERVED_ID")
}
