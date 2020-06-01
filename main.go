package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	customEntry "generateIamDbToken/kr/misoboy/common/entry"
	"generateIamDbToken/kr/misoboy/common/model"
	customTheme "generateIamDbToken/kr/misoboy/common/theme"
	rdsService "generateIamDbToken/kr/misoboy/rds"
	redshiftService "generateIamDbToken/kr/misoboy/redshift"
	iniUtil "github.com/c4pt0r/ini"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	logPath                  = "./log"
	logFileName              = "application.log"
	myLogger                 *log.Logger
	dev_eu_profile           *string
	dev_eu_rdsHostname       *string
	dev_eu_rdsPort           *int
	dev_eu_rdsUsername       *string
	dev_eu_redshiftClusterId *string
	dev_eu_redshiftUsername  *string

	dev_cn_profile           *string
	dev_cn_rdsHostname       *string
	dev_cn_rdsPort           *int
	dev_cn_rdsUsername       *string
	dev_cn_redshiftClusterId *string
	dev_cn_redshiftUsername  *string

	stg_eu_profile           *string
	stg_eu_rdsHostname       *string
	stg_eu_rdsPort           *int
	stg_eu_rdsUsername       *string
	stg_eu_redshiftClusterId *string
	stg_eu_redshiftUsername  *string

	stg_cn_profile           *string
	stg_cn_rdsHostname       *string
	stg_cn_rdsPort           *int
	stg_cn_rdsUsername       *string
	stg_cn_redshiftClusterId *string
	stg_cn_redshiftUsername  *string

	prd_eu_profile           *string
	prd_eu_mfaArn            *string
	prd_eu_rdsHostname       *string
	prd_eu_rdsPort           *int
	prd_eu_rdsUsername       *string
	prd_eu_redshiftClusterId *string
	prd_eu_redshiftUsername  *string

	prd_cn_profile           *string
	prd_cn_mfaArn            *string
	prd_cn_rdsHostname       *string
	prd_cn_rdsPort           *int
	prd_cn_rdsUsername       *string
	prd_cn_redshiftClusterId *string
	prd_cn_redshiftUsername  *string
)

func main() {

	fmt.Println("")
	fmt.Println("    #     # ###  #####  ####### ######  ####### #     #")
	fmt.Println("    ##   ##  #  #     # #     # #     # #     #  #   #")
	fmt.Println("    # # # #  #  #       #     # #     # #     #   # #")
	fmt.Println("    #  #  #  #   #####  #     # ######  #     #    #")
	fmt.Println("    #     #  #        # #     # #     # #     #    #")
	fmt.Println("    #     #  #  #     # #     # #     # #     #    #")
	fmt.Println("    #     # ###  #####  ####### ######  #######    #")
	fmt.Println("")

	makeLog()

	absPath, err := filepath.Abs("./conf/env.conf")
	if err != nil {
		log.Println(err)
		return
	}

	var conf = iniUtil.NewConf(absPath)

	dev_eu_profile = conf.String("dev-eu", "PROFILE", "default")
	dev_eu_rdsHostname = conf.String("dev-eu", "RDS_HOSTNAME", "")
	dev_eu_rdsPort = conf.Int("dev-eu", "RDS_PORT", 5306)
	dev_eu_rdsUsername = conf.String("dev-eu", "RDS_USERNAME", "")
	dev_eu_redshiftClusterId = conf.String("dev-eu", "REDSHIFT_CLUSTER_ID", "")
	dev_eu_redshiftUsername = conf.String("dev-eu", "REDSHIFT_USERNAME", "")

	dev_cn_profile = conf.String("dev-cn", "PROFILE", "default")
	dev_cn_rdsHostname = conf.String("dev-cn", "RDS_HOSTNAME", "")
	dev_cn_rdsPort = conf.Int("dev-cn", "RDS_PORT", 5306)
	dev_cn_rdsUsername = conf.String("dev-cn", "RDS_USERNAME", "")
	dev_cn_redshiftClusterId = conf.String("dev-cn", "REDSHIFT_CLUSTER_ID", "")
	dev_cn_redshiftUsername = conf.String("dev-cn", "REDSHIFT_USERNAME", "")

	stg_eu_profile = conf.String("stg-eu", "PROFILE", "default")
	stg_eu_rdsHostname = conf.String("stg-eu", "RDS_HOSTNAME", "")
	stg_eu_rdsPort = conf.Int("stg-eu", "RDS_PORT", 5306)
	stg_eu_rdsUsername = conf.String("stg-eu", "RDS_USERNAME", "")
	stg_eu_redshiftClusterId = conf.String("stg-eu", "REDSHIFT_CLUSTER_ID", "")
	stg_eu_redshiftUsername = conf.String("stg-eu", "REDSHIFT_USERNAME", "")

	stg_cn_profile = conf.String("stg-cn", "PROFILE", "default")
	stg_cn_rdsHostname = conf.String("stg-cn", "RDS_HOSTNAME", "")
	stg_cn_rdsPort = conf.Int("stg-cn", "RDS_PORT", 5306)
	stg_cn_rdsUsername = conf.String("stg-cn", "RDS_USERNAME", "")
	stg_cn_redshiftClusterId = conf.String("stg-cn", "REDSHIFT_CLUSTER_ID", "")
	stg_cn_redshiftUsername = conf.String("stg-cn", "REDSHIFT_USERNAME", "")

	prd_eu_profile = conf.String("prd-eu", "PROFILE", "default")
	prd_eu_mfaArn = conf.String("prd-eu", "MFA_ARN", "")
	prd_eu_rdsHostname = conf.String("prd-eu", "RDS_HOSTNAME", "")
	prd_eu_rdsPort = conf.Int("prd-eu", "RDS_PORT", 5306)
	prd_eu_rdsUsername = conf.String("prd-eu", "RDS_USERNAME", "")
	prd_eu_redshiftClusterId = conf.String("prd-eu", "REDSHIFT_CLUSTER_ID", "")
	prd_eu_redshiftUsername = conf.String("prd-eu", "REDSHIFT_USERNAME", "")

	prd_cn_profile = conf.String("prd-cn", "PROFILE", "default")
	prd_cn_mfaArn = conf.String("prd-cn", "MFA_ARN", "")
	prd_cn_rdsHostname = conf.String("prd-cn", "RDS_HOSTNAME", "")
	prd_cn_rdsPort = conf.Int("prd-cn", "RDS_PORT", 5306)
	prd_cn_rdsUsername = conf.String("prd-cn", "RDS_USERNAME", "")
	prd_cn_redshiftClusterId = conf.String("prd-cn", "REDSHIFT_CLUSTER_ID", "")
	prd_cn_redshiftUsername = conf.String("prd-cn", "REDSHIFT_USERNAME", "")
	conf.Parse()

	app := app.New()
	app.Settings().SetTheme(customTheme.NewCustomTheme())

	w := app.NewWindow("Generate AWS DB Token Tool")
	w.Resize(fyne.NewSize(600, 500))
	//w.SetFixedSize(true)
	w.CenterOnScreen()

	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("Information", "Email : misoboy.kor@gmail.com", w)
		}),
	)))
	w.SetMaster()

	w.SetContent(
		widget.NewTabContainer(
			widget.NewTabItem("DEV", makeDevTab(w)),
			widget.NewTabItem("STG", makeStgTab(w)),
			widget.NewTabItem("PRD", makePrdTab(w)),
		))

	w.ShowAndRun()
}

func makeDevTab(w fyne.Window) fyne.Widget {
	return widget.NewVBox(
		widget.NewButton("dev-eu-rds", func() {
			var awsModel = model.AwsModel{Profile: *dev_eu_profile, Hostname: *dev_eu_rdsHostname, Port: *dev_eu_rdsPort, Username: *dev_eu_rdsUsername}
			var service = rdsService.RdsService{AwsModel: awsModel, Window: w, Logger: myLogger}
			service.GenerateRdsToken()
		}),
		widget.NewButton("dev-eu-redshift", func() {
			var awsModel = model.AwsModel{Profile: *dev_eu_profile, Username: *dev_eu_redshiftUsername, ClusterId: *dev_eu_redshiftClusterId}
			var service = redshiftService.RedshiftService{AwsModel: awsModel, Window: w, Logger: myLogger}
			service.GenerateRedshiftToken()
		}),
		widget.NewButton("dev-cn-rds", func() {
			var awsModel = model.AwsModel{Profile: *dev_cn_profile, Hostname: *dev_cn_rdsHostname, Port: *dev_cn_rdsPort, Username: *dev_cn_rdsUsername}
			var service = rdsService.RdsService{AwsModel: awsModel, Window: w, Logger: myLogger}
			service.GenerateRdsToken()
		}),
		widget.NewButton("dev-cn-redshift", func() {
			var awsModel = model.AwsModel{Profile: *dev_cn_profile, Username: *dev_cn_redshiftUsername, ClusterId: *dev_cn_redshiftClusterId}
			var service = redshiftService.RedshiftService{AwsModel: awsModel, Window: w, Logger: myLogger}
			service.GenerateRedshiftToken()
		}),
	)
}

func makeStgTab(w fyne.Window) fyne.Widget {
	return widget.NewVBox(
		widget.NewButton("stg-eu-rds", func() {
			var awsModel = model.AwsModel{Profile: *stg_eu_profile, Hostname: *stg_eu_rdsHostname, Port: *stg_eu_rdsPort, Username: *stg_eu_rdsUsername}
			var service = rdsService.RdsService{AwsModel: awsModel, Window: w, Logger: myLogger}
			service.GenerateRdsToken()
		}),
		widget.NewButton("stg-eu-redshift", func() {
			var awsModel = model.AwsModel{Profile: *stg_eu_profile, Username: *stg_eu_redshiftUsername, ClusterId: *stg_eu_redshiftClusterId}
			var service = redshiftService.RedshiftService{AwsModel: awsModel, Window: w, Logger: myLogger}
			service.GenerateRedshiftToken()
		}),
		widget.NewButton("stg-cn-rds", func() {
			var awsModel = model.AwsModel{Profile: *stg_cn_profile, Hostname: *stg_cn_rdsHostname, Port: *stg_cn_rdsPort, Username: *stg_cn_rdsUsername}
			var service = rdsService.RdsService{AwsModel: awsModel, Window: w, Logger: myLogger}
			service.GenerateRdsToken()
		}),
		widget.NewButton("stg-cn-redshift", func() {
			var awsModel = model.AwsModel{Profile: *stg_cn_profile, Username: *stg_cn_redshiftUsername, ClusterId: *stg_cn_redshiftClusterId}
			var service = redshiftService.RedshiftService{AwsModel: awsModel, Window: w, Logger: myLogger}
			service.GenerateRedshiftToken()
		}),
	)
}

func makePrdTab(w fyne.Window) fyne.Widget {
	return widget.NewVBox(
		widget.NewButton("prd-eu-rds", func() {
			makeOtpWindowForm(func(OtpNum string) {
				var awsModel = model.AwsModel{Profile: *prd_eu_profile, MfaArn: *prd_eu_mfaArn, OtpNum: OtpNum, Hostname: *prd_eu_rdsHostname, Port: *prd_eu_rdsPort, Username: *prd_eu_rdsUsername}
				var service = rdsService.RdsService{AwsModel: awsModel, Window: w, Logger: myLogger}
				service.GenerateRdsToken()
			})
		}),
		widget.NewButton("prd-eu-redshift", func() {
			makeOtpWindowForm(func(otpNum string) {
				var awsModel = model.AwsModel{Profile: *prd_eu_profile, MfaArn: *prd_eu_mfaArn, OtpNum: otpNum, Username: *prd_eu_redshiftUsername, ClusterId: *prd_eu_redshiftClusterId}
				var service = redshiftService.RedshiftService{AwsModel: awsModel, Window: w, Logger: myLogger}
				service.GenerateRedshiftToken()
			})
		}),
		widget.NewButton("prd-cn-rds", func() {
			makeOtpWindowForm(func(otpNum string) {
				var awsModel = model.AwsModel{Profile: *prd_cn_profile, MfaArn: *prd_cn_mfaArn, OtpNum: otpNum, Hostname: *prd_cn_rdsHostname, Port: *prd_cn_rdsPort, Username: *prd_cn_rdsUsername}
				var service = rdsService.RdsService{AwsModel: awsModel, Window: w, Logger: myLogger}
				service.GenerateRdsToken()
			})
		}),
		widget.NewButton("prd-cn-redshift", func() {
			makeOtpWindowForm(func(otpNum string) {
				var awsModel = model.AwsModel{Profile: *prd_cn_profile, MfaArn: *prd_cn_mfaArn, OtpNum: otpNum, Username: *prd_cn_redshiftUsername, ClusterId: *prd_cn_redshiftClusterId}
				var service = redshiftService.RedshiftService{AwsModel: awsModel, Window: w, Logger: myLogger}
				service.GenerateRedshiftToken()
			})
		}),
	)
}

func makeOtpWindowForm(callback func(OtpNum string)) {
	w := fyne.CurrentApp().NewWindow("Type to OptCode")
	otpCodeEntry := customEntry.NewEnterEntry()
	otpCodeEntry.SetKeyName(fyne.KeyReturn)
	otpCodeEntry.SetPlaceHolder("type to OtpCode")

	form := &widget.Form{
		OnCancel: func() {
			w.Close()
		},
		OnSubmit: func() {
			if otpCodeEntry.Text == "" {
				err := errors.New("Required OtpCode")
				dialog.ShowError(err, w)
			} else {
				w.Close()
				callback(otpCodeEntry.Text)
			}
		},
	}

	otpCodeEntry.SetCallback(func(value string) {
		form.OnSubmit()
	})

	form.Append("OtpCode", otpCodeEntry)
	w.SetContent(form)
	w.Resize(fyne.NewSize(200, 170))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.Show()

}

func makeLog() {
	os.MkdirAll(logPath, os.ModePerm)
	if !FileExists(logPath + "/" + logFileName) {
		CreateFile(logPath + "/" + logFileName)
	}
	fpLog, err := os.OpenFile(logPath+"/"+logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//defer fpLog.Close()
	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
