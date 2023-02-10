package sheets

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

const PAGE_URI = "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999"

func UpdateGoogleSheets(creds []byte, sheetId string) error {
	srv, err := GetService(creds)
	if err != nil {
		return err
	}

	values := []interface{}{fmt.Sprintf("=IMPORTHTML(\"%s\";\"table\")", PAGE_URI)}
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, values)

	_, err = srv.Spreadsheets.Values.Update(sheetId, "A1", &vr).ValueInputOption("USER_ENTERED").Do()
	return err
}
