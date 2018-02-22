package smsbackuprestore

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func GenerateSMSOutput(m *Messages) error {
	smsOutput, err := os.Create("sms.tsv")
	if err != nil {
		return fmt.Errorf("Unable to create file: sms.tsv\n%q", err)
	}
	defer smsOutput.Close()

	// print header row
	headers := []string{
		"SMS Index #",
		"Protocol",
		"Address",
		"Type",
		"Subject",
		"Body",
		"Service Center",
		"Status",
		"Read",
		"Date",
		"Locked",
		"Date Sent",
		"Readable Date",
		"Contact Name",
	}
	fmt.Fprintf(smsOutput, "%s\n", strings.Join(headers, "\t"))

	// iterate over sms
	for i, sms := range m.SMS {
		row := []string{
			strconv.Itoa(i),
			sms.Protocol,
			sms.Address.String(),
			sms.Type.String(),
			sms.Subject,
			CleanupMessageBody(sms.Body),
			sms.ServiceCenter.String(),
			sms.Status.String(),
			sms.Read.String(),
			sms.Date.String(),
			sms.Locked.String(),
			sms.DateSent.String(),
			sms.ReadableDate,
			RemoveCommasBeforeSuffixes(sms.ContactName),
		}
		fmt.Fprintf(smsOutput, "%s\n", strings.Join(row, "\t"))
	}

	return nil
}