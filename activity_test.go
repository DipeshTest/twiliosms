package twiliosms

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestInvalidAuth_sendsms(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af414")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Send SMS")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "203")

}

func TestInvalidAuth_retrive(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af414")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Retrive Recipients")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "203")

}

func TestInvalidAccountSid_sendsms(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a406")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Retrive Recipients")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "203")

}

func TestInvalidAccountSid_retrive(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a406")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Send SMS")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "203")

}

func TestInvalidUrlString_sendsms(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio1234.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Send SMS")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "203")

}

func TestInvalidUrlString_retrive(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio1234.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Retrive Recipients")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "203")

}

func TestPartialSuccess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Send SMS")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623445")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "201")

}

/*
func TestSuccess_retrive(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Retrive Recipients")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	fmt.Println(tc.GetOutput("message"))

	assert.Equal(t, result, "200")
	tc.SetInput("to", tc.GetOutput("message"))
	tc.SetInput("method", "Send SMS")

	act.Eval(tc)
	result = tc.GetOutput("statusCode")

	assert.Equal(t, result, "200")

}

/*
func TestAllValidToNumbers(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("method", "Send SMS")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444,+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "200")

}


func TestEmptyAccountSid(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "203")

}


func TestEmptyAuthToken(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623444")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "203")

}


func TestEmptyUrlString(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623442,+919177623442,+919177623442,+919177623442,+919177623442")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "504")

}


func TestEmptyFrom(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623442,+919177623442,+919177623442,+919177623442,+919177623442")
	tc.SetInput("from", "")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "503")

}


func TestEmptyTo(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "502")

}


func TestEmptySMSBody(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("messageBody", "")
	tc.SetInput("to", "+919177623442,+919177623442,+919177623442,+919177623442,+919177623442")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "501")

}


func TestAllInvalidToNumbers(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accountSid", "AC8112820c2969c0b9ba6abac8ee6a4062")
	tc.SetInput("authToken", "f2542c88dbfb58c494e642bf10af4140")
	tc.SetInput("urlString", "https://api.twilio.com/2010-04-01/Accounts/")
	tc.SetInput("messageBody", "Hello from flogo")
	tc.SetInput("to", "+919177623442,+919177623442,+919177623442,+919177623442,+919177623442")
	tc.SetInput("from", "+14437433811")

	act.Eval(tc)

	//check result attr

	result := tc.GetOutput("statusCode")
	assert.Equal(t, result, "201")

}*/
