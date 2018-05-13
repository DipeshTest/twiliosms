package twiliosms

import (
	"strings"

	"github.com/DipeshTest/allstarsshared/twilio"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-twiliosms")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

const (
	SENDSMS      = "Send SMS"
	RETRIVE_LIST = "Retrive Recipients"
)

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	accountSid := context.GetInput("accountSid").(string)
	authToken := context.GetInput("authToken").(string)
	urlStr := context.GetInput("urlString").(string)
	method := context.GetInput("method").(string)
	smsText := context.GetInput("messageBody").(string)
	to := context.GetInput("to").(string)
	from := context.GetInput("from").(string)
	nums := strings.Split(to, ",")

	if len(strings.TrimSpace(urlStr)) == 0 {
		//	fmt.Println(twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
		context.SetOutput("statusCode", "504")
		context.SetOutput("status", "Failed")
		context.SetOutput("message", "Url String field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else if len(strings.TrimSpace(accountSid)) == 0 {
		//	fmt.Println(twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
		context.SetOutput("statusCode", "505")
		context.SetOutput("status", "Failed")
		context.SetOutput("message", "Account SID field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else if len(strings.TrimSpace(authToken)) == 0 {
		//	fmt.Println(twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
		context.SetOutput("statusCode", "506")
		context.SetOutput("status", "Failed")
		context.SetOutput("message", "Auth Token String field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else {

		switch method {

		case SENDSMS:
			{

				if len(strings.TrimSpace(smsText)) == 0 {
					//	fmt.Println(twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
					context.SetOutput("statusCode", "501")
					context.SetOutput("status", "Failed")
					context.SetOutput("message", "SMS body is blank")
					context.SetOutput("failedNumbers", to)

					//respond with this
				} else if len(strings.TrimSpace(to)) == 0 {
					//	fmt.Println(twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
					context.SetOutput("statusCode", "502")
					context.SetOutput("status", "Failed")
					context.SetOutput("message", "To field is blank")
					//context.SetOutput("failedNumbers", to)

					//respond with this
				} else if len(strings.TrimSpace(from)) == 0 {
					//	fmt.Println(twilio.ResponseData{501, "Failed", "Send input numbers back", "SMS body is blank"})
					context.SetOutput("statusCode", "503")
					context.SetOutput("status", "Failed")
					context.SetOutput("message", "From field is blank")
					//context.SetOutput("failedNumbers", to)

					//respond with this
				} else {
					log.Debugf("The Flogo engine says [%s] to [%s]", accountSid+authToken+urlStr, smsText+to+from)

					jobs := make(chan string)
					results := make(chan twilio.ResponseData)
					for i, _ := range nums {
						twilioData := twilio.Twilio{accountSid, authToken, urlStr + accountSid + "/Messages.json", smsText, nums[i], from}
						go twilio.SendSMS(twilioData, jobs, results)
						//fmt.Println("Creating go routinge number", nums[i])
					}
					for j, _ := range nums {
						jobs <- nums[j]
						//fmt.Println("Sending number to routine", nums[j])
					}
					close(jobs)
					respMap := make(map[int]string)
					endProc := false
					failedNumbers := ""
					faildNumMessage := ""
					for k, _ := range nums {
						respdata := <-results
						k = k + 1
						if respdata.StatusCode == 404 || respdata.StatusCode == 401 || respdata.StatusCode == 901 {
							// here we need to create a custom response and send out error code, use return statement
							respMap[respdata.StatusCode] = respdata.Num + ":" + respdata.Message
							//fmt.Println("Error from service" + respMap[respdata.StatusCode])
							context.SetOutput("statusCode", "203")
							context.SetOutput("status", "Failed")
							if respdata.StatusCode == 901 {
								context.SetOutput("message", respdata.ErrorData)
							} else {
								context.SetOutput("message", respdata.Message)
							}
							context.SetOutput("message", respdata.Message)
							context.SetOutput("failedNumbers", to)
							//return 3 param resp
							endProc = true
							break
						} else {
							respMap[respdata.StatusCode] = respdata.Num + ":" + respdata.Message
							if !(respdata.StatusCode >= 200 && respdata.StatusCode < 300) {
								failedNumbers = failedNumbers + respdata.Num + ":"
								faildNumMessage = faildNumMessage + respdata.Message + ":\n"
							}
							// /	fmt.Println("Response from service" + respMap[respdata.StatusCode])
						}
						//	fmt.Println("Response from ", k, <-results)

					}

					//close(results)
					respStat1001 := respMap[1001]
					respStat400 := respMap[400]
					respStat200 := respMap[200]
					if !endProc {
						if respStat200 == "" {
							context.SetOutput("statusCode", "202")
							context.SetOutput("status", "Failed")
							context.SetOutput("message", faildNumMessage)
							context.SetOutput("failedNumbers", failedNumbers)
						} else if respStat1001 == "" && respStat400 == "" {
							//fmt.Println(respMap[200])
							context.SetOutput("statusCode", "200")
							context.SetOutput("status", "Success")
							context.SetOutput("message", "SMS sent successfully to :"+to)
							//	context.SetOutput("failedNumbers", to)
						} else {
							context.SetOutput("statusCode", "201")
							context.SetOutput("status", "Partial-Success")
							context.SetOutput("message", faildNumMessage)
							context.SetOutput("failedNumbers", failedNumbers)
						}
					} else {

					}
					log.Debugf("The Flogo engine says [%s] to [%s]", context.GetOutput("statusCode"), err)

				}

			}

		case RETRIVE_LIST:
			{
				twilioData := twilio.Twilio{accountSid, authToken, urlStr + accountSid + "/OutgoingCallerIds.json", "", "", ""}
				respdata := twilio.RetrieveRecipientList(twilioData)
				if !(respdata.StatusCode >= 200 && respdata.StatusCode < 300) {
					context.SetOutput("statusCode", "203")
					context.SetOutput("status", "Failed")
					if respdata.StatusCode == 901 {
						context.SetOutput("message", respdata.ErrorData)
					} else {
						context.SetOutput("message", respdata.Message)
					}
				} else {
					context.SetOutput("statusCode", "200")
					context.SetOutput("status", "Success")
					context.SetOutput("message", respdata.Num)

				}

			}

		}

	}

	return true, err
}
