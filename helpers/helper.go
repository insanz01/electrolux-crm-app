package helpers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var customerMandatory = []string{"Service Order No", "Mobile No", "Product Line Alpha - Numeric", "Date Of Purchase", "City", "State", "Country", "Call Type", "Model Description"}

var giftClaimMandatory = []string{"ID Claim", "Nama", "No HP", "No HP2", "Toko", "Tanggal Pembelian", "Pembelian", "Tanggal Klaim", "Hadiah", "QTY"}

var dateType = []string{"date_of_purchase", "service_order_date", "appointment_date", "closure_date", "attended_date"}

func IsHeaderMandatory(header, category string) bool {

	mandatoryString := customerMandatory

	if category == "gift_claim" {
		mandatoryString = giftClaimMandatory
	}

	for _, v := range mandatoryString {
		if v == header {
			return true
		}
	}

	return false
}

func IsIndexContains(mandatoryIndex []int, index int) bool {
	for _, v := range mandatoryIndex {
		if v == index {
			return true
		}
	}

	return false
}

func IsHeaderContains(headers []string, header string) bool {
	for _, v := range headers {
		if v == header {
			return true
		}
	}

	return false
}

func RowIndexToExcel(index, row int) string {
	var asciiChar string

	asciiNumber := index + 65

	if asciiNumber > 90 {
		asciiNumberCheck := asciiNumber % 90

		asciiChar = "A"
		asciiNumberResidual := asciiNumberCheck + 64

		if asciiNumber > 116 {
			asciiChar = "B"
			asciiNumberResidual = asciiNumberResidual % 90
			asciiNumberResidual = asciiNumberResidual + 64
		}

		asciiChar = fmt.Sprintf("%s%c", asciiChar, asciiNumberResidual)
	} else {
		asciiChar = fmt.Sprintf("%c", asciiNumber)
	}

	return asciiChar + strconv.Itoa(row+1)
}

func DownloadFile(url, filepath string) error {
	// Mengirim permintaan HTTP GET ke server untuk mengunduh file
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Membuka file tujuan untuk menulis data yang diunduh
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Menyalin data yang diunduh ke file tujuan
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func CheckDate(keyName string) bool {
	for _, v := range dateType {
		if v == keyName {
			return true
		}
	}

	return false
}

func WrapCloser(closeFn func() error) {
	if err := closeFn(); err != nil {
		log.Println(err)
	}
}

// GetTokenFromHeader get token from header authorization with bearer token
func GetTokenFromHeader(req *http.Request) string {
	_headerAuthorization := "Authorization"
	_authScheme := "Bearer"
	authHeader := strings.Split(req.Header.Get(_headerAuthorization), " ")

	if len(authHeader) != 2 || authHeader[0] != _authScheme {
		return ""
	}

	return strings.TrimSpace(authHeader[1])
}

// GenerateOpenAPIResponse generate response for open api in case of a success response
func GenerateOpenAPIResponse(data interface{}) map[string]any {
	return map[string]any{
		"status":  1,
		"message": "",
		"data":    data,
	}
}

// GenerateOpenAPIErrorResponse generate response for open api in case of an error response
func GenerateOpenAPIErrorResponse(httpErrorCode int) map[string]any {
	type empty struct{}
	return map[string]any{
		"status":  0,
		"message": http.StatusText(httpErrorCode),
		"data":    empty{},
	}
}
