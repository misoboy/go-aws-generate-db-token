package rds

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	commonModel "github.com/misoboy/go-aws-generate-db-token/common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	"github.com/aws/aws-sdk-go/service/sts"
	"log"
	"strconv"
	"strings"
)

type RdsService struct {
	commonModel.AwsModel
	Window fyne.Window
	Logger *log.Logger
}

func (service *RdsService) GenerateRdsToken() {

	var hostname = service.Hostname + ":" + strconv.Itoa(service.Port) + "/"
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: service.Profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		service.Logger.Println(err)
		err := errors.New(err.Error())
		dialog.ShowError(err, service.Window)
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

	token, err := rdsutils.BuildAuthToken(hostname, *sess.Config.Region, service.Username, sess.Config.Credentials)
	if err != nil {
		service.Logger.Println(err)
		err := errors.New(err.Error())
		dialog.ShowError(err, service.Window)
		return
	}
	service.Logger.Println("########## Generate RDS Token Result ##########")
	service.Logger.Println("DbUser : " + service.Username)
	service.Logger.Println("DbPassword : " + token)
	service.Logger.Println("###############################################")

	w := fyne.CurrentApp().NewWindow("Complete")
	w.SetContent(
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			widget.NewButton("Username Show", func() {
				subWin := fyne.CurrentApp().NewWindow("Username")
				subWin.SetContent(widget.NewVBox(
					widget.NewLabel(service.Username),
					widget.NewButton("Close", func() {
						subWin.Close()
					}),
				))
				subWin.CenterOnScreen()
				subWin.Show()
			}),
			widget.NewButton("Copy", func() {
				w.Clipboard().SetContent(service.Username)
			}),
			widget.NewButton("Password Show", func() {
				subWin := fyne.CurrentApp().NewWindow("Password")
				subWin.SetContent(widget.NewVBox(
					widget.NewLabel(token),
					widget.NewButton("Close", func() {
						subWin.Close()
					}),
				))
				subWin.CenterOnScreen()
				subWin.Show()
			}),
			widget.NewButton("Copy", func() {
				w.Clipboard().SetContent(token)
			}),
		))
	w.Resize(fyne.NewSize(400, 200))
	w.SetFixedSize(true)
	w.CenterOnScreen()
	w.Show()
}
