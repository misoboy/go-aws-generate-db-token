package redshift

import (
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	commonModel "generateIamDbToken/kr/misoboy/common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
	"strings"
)

type RedshiftService struct {
	commonModel.AwsModel
	Window fyne.Window
	Logger *log.Logger
}

func (service *RedshiftService) GenerateRedshiftToken() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: service.Profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		service.Logger.Println(err)
		return
	}

	if service.OtpNum != "" {
		otpNum := strings.TrimSuffix(service.OtpNum, "\r\n")

		stsClient := sts.New(sess)
		sessionTokenInput := sts.GetSessionTokenInput{SerialNumber: &service.MfaArn, TokenCode: &otpNum}
		resp, err := stsClient.GetSessionToken(&sessionTokenInput)
		if err != nil {
			service.Logger.Println(err)
			err := errors.New(err.Error())
			dialog.ShowError(err, service.Window)
			return
		}

		staticProvider := credentials.NewStaticCredentials(*resp.Credentials.AccessKeyId, *resp.Credentials.SecretAccessKey, *resp.Credentials.SessionToken)
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{Credentials: staticProvider, Region: sess.Config.Region},
			SharedConfigState: session.SharedConfigEnable,
		})
		if err != nil {
			service.Logger.Println(err)
			err := errors.New(err.Error())
			dialog.ShowError(err, service.Window)
			return
		}
	}

	redshiftClient := redshift.New(sess)
	credentialsInput := redshift.GetClusterCredentialsInput{ClusterIdentifier: &service.ClusterId, DbUser: &service.Username, AutoCreate: &[]bool{false}[0]}
	resp, err := redshiftClient.GetClusterCredentials(&credentialsInput)
	if err != nil {
		fmt.Println(err)
		err := errors.New(err.Error())
		dialog.ShowError(err, service.Window)
		return
	}
	service.Logger.Println("########## Generate Redshift Token Result ##########")
	service.Logger.Println("DbUser : " + *resp.DbUser)
	service.Logger.Println("DbPassword : " + *resp.DbPassword)
	service.Logger.Println("####################################################")

	w := fyne.CurrentApp().NewWindow("Complete")
	w.SetContent(
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			widget.NewButton("Username Show", func() {
				subWin := fyne.CurrentApp().NewWindow("Username")
				subWin.SetContent(widget.NewVBox(
					widget.NewLabel(*resp.DbUser),
					widget.NewButton("Close", func() {
						subWin.Close()
					}),
				))
				subWin.CenterOnScreen()
				subWin.Show()
			}),
			widget.NewButton("Copy", func() {
				w.Clipboard().SetContent(*resp.DbUser)
			}),
			widget.NewButton("Password Show", func() {
				subWin := fyne.CurrentApp().NewWindow("Password")
				subWin.SetContent(widget.NewVBox(
					widget.NewLabel(*resp.DbPassword),
					widget.NewButton("Close", func() {
						subWin.Close()
					}),
				))
				subWin.CenterOnScreen()
				subWin.Show()
			}),
			widget.NewButton("Copy", func() {
				w.Clipboard().SetContent(*resp.DbPassword)
			}),
		))
	w.Resize(fyne.NewSize(400, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.Show()
}
