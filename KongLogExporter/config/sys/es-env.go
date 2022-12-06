package sys

import "os"

func GetEsEnv() (string, string, string) {

	EsUrl := os.Getenv("ES_URL")
	EsName := os.Getenv("ES_USERNAME")
	EsPasswd := os.Getenv("ES_PASSWORD")

	return EsUrl, EsName, EsPasswd
}
