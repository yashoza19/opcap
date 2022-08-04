package capability

import (
	"fmt"
	"os"
	"time"
)

func (ca CapAudit) Report(opts ...ReportOption) error {

	for _, opt := range opts {

		err := opt.report(ca)
		if err != nil {
			logger.Debugw("Unable to generate report for %T", opt, "Error: %s", err)
		}

	}
	return nil
}

type ReportOption interface {
	report(ca CapAudit) error
}

type RptOptionPrint struct{}

func (RptOptionPrint) report(ca CapAudit) error {

	fmt.Println()
	fmt.Println("opcap report:")
	fmt.Println("-----------------------------------------")
	fmt.Printf("Report Date: %s\n", time.Now())
	fmt.Printf("OpenShift Version: %s\n", ca.OcpVersion)
	fmt.Printf("Package Name: %s\n", ca.Subscription.Package)
	fmt.Printf("Channel: %s\n", ca.Subscription.Channel)
	fmt.Printf("Catalog Source: %s\n", ca.Subscription.CatalogSource)
	fmt.Printf("Install Mode: %s\n", ca.Subscription.InstallModeType)

	if !ca.CsvTimeout {
		fmt.Printf("Result: %s\n", ca.Csv.Status.Phase)
	} else {
		fmt.Println("Result: timeout")
	}

	fmt.Printf("Message: %s\n", ca.Csv.Status.Message)
	fmt.Printf("Reason: %s\n", ca.Csv.Status.Reason)
	fmt.Println("-----------------------------------------")

	return nil
}

type RptOptionFile struct {
	FilePath string
}

func (opt RptOptionFile) report(ca CapAudit) error {

	file, err := os.OpenFile(opt.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		file.Close()
		return err
	}
	defer file.Close()

	if !ca.CsvTimeout {

		file.WriteString("{\"level\":\"info\",\"message\":\"" + string(ca.Csv.Status.Phase) + "\",\"package\":\"" + ca.Subscription.Package + "\",\"channel\":\"" + ca.Subscription.Channel + "\",\"installmode\":\"" + string(ca.Subscription.InstallModeType) + "\"}\n")
	} else {

		file.WriteString("{\"level\":\"info\",\"message\":\"" + "timeout" + "\",\"package\":\"" + ca.Subscription.Package + "\",\"channel\":\"" + ca.Subscription.Channel + "\",\"installmode\":\"" + string(ca.Subscription.InstallModeType) + "\"}\n")
	}

	return nil
}
