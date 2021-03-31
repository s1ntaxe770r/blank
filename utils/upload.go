package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type Undelivered struct {
	Status string `json:"Status"`
	Err    error  `json:"Err"`
}

type Parcel struct {
	Status string `json:"Status"`
	URL    string `json:"URL"`
}

func UploadFile(file *os.File, url string, creds Credentials) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url+"/push", file)
	if err != nil {
		panic("PANICED AT REQUEST")
		return errors.New(color.RedString("unable to make http request, ") + err.Error())
	}
	// req.SetBasicAuth(creds.Username, creds.Password)
	resp, err := client.Do(req)
	switch {
	case resp.StatusCode == http.StatusRequestEntityTooLarge:

		return errors.New(color.RedString("File size exceeds limit"))
	case resp.StatusCode == http.StatusBadRequest:
		errmessage := Undelivered{}
		json.NewDecoder(resp.Body).Decode(&errmessage)
		return errors.New(errmessage.Err.Error())
	case resp.StatusCode == http.StatusOK:
		file := Parcel{}
		json.NewDecoder(resp.Body).Decode(&file)
		logrus.Warn("HIIIIII")
		imageurl := fmt.Sprintf("http://%s", file.URL)
		fmt.Println(color.GreenString(imageurl))
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
