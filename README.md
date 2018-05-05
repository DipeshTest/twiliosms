---
title: Send SMS
weight: 1
---

# Counter
This activity allows you to post a tweet using your twitter account credentials, the activity is built using Twilio API reference https://www.twilio.com/docs/usage/api

## Installation
### Flogo Web
This activity is built by AllStarts team
### Flogo CLI
```bash
flogo add activity github.com/DipeshTest/twiliosms
```

## Schema
Inputs and Outputs:

```json
{
"inputs":[
    {
		"name": "accountSid",
		"type": "string",
    "required": true
	},
	{
		"name": "authToken",
		"type": "string",
    "required": true
	},
	{
		"name": "urlString",
		"type": "string",
    "required": true
	},
  {
		"name": "method",
		"type": "string",
    "allowed": [
        "Send SMS",
        "Retrive Recipients"
      ],
      "value": "Send SMS",
      "required": true
	},
	{
		"name": "messageBody",
		"type": "string"
	},
	{
		"name": "to",
		"type": "string"
	},
	{
		"name": "from",
		"type": "string"
	}
  ],
  "outputs": [
    {
      "name": "statusCode",
      "type": "string"
    },
    {
      "name": "status",
      "type": "string"
    },
	{
      "name": "message",
      "type": "string"
    },
	{
      "name": "failedNumbers",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| accountSid | True     | The account SID of Twilio |         
| authToken   | True    | Auth Token form Twilio|
| urlString | True     | Twilio URI, eg : https://api.twilio.com/2010-04-01/Accounts/ |  
| method | True     | We expose 2 methods, one to send SMS to a recipient, second to retrieve the list of recipients verified @ your account , the response for the retrieve list will be a comma seperated list of values which can be used to directly send to a twilio sms activity to send smses to multiple recipients|  
| messageBody | True     | The body of SMS, limited to 1600 characters|  
| to       | True    |The number to send SMS to,the destination phone number for SMS/MMS or a Channel user address for other 3rd party channels. Destination phone numbers should be formatted with a '+' and country code e.g., +16175551212 (E.164 format).|
| from       | True    | Your Twilio number to send SMS using.A Twilio phone number (in E.164 format), alphanumeric sender ID or a Channel Endpoint address enabled for the type of message you wish to send. Phone numbers or short codes purchased from Twilio work here. You cannot (for example) spoof messages from your own cell phone number  |

## Examples
### Send SMS
The below example to send SMS to multiple recipients:

```json
{
	"AccountSid": "AC8112820c2969c0b9ba6",
	"AuthToken": "f2542c88dbfb58c494e642bf1",
	"UrlString": "https://api.twilio.com/2010-04-01/Accounts/",
	"method":"Send SMS",
	"MsgData": "Final Code",
	"To": "+9189908098098,+9189908098098",
	"From": "+14437433999"
}
```

### Send SMS
The below example to send SMS to single user:

```json
{
	"AccountSid": "AC8112820c2969c0b9",
	"AuthToken": "f2542c88dbfb58c49",
	"UrlString": "https://api.twilio.com/2010-04-01/Accounts/",
	"method":"Send SMS",
	"MsgData": "Final Code",
	"To": "+9189908098098",
	"From": "+14437433999"
}
```

### Retrive Recipients
The below example retrives the list of verified recipients:

```json
{
	"AccountSid": "AC8112820c2969c0b9ba6",
	"AuthToken": "f2542c88dbfb58c494e642b",
	"UrlString": "https://api.twilio.com/2010-04-01/Accounts/",
	"method":"Retrive Recipients"
}
```

## Response Codes
### Retrieve List
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request was successful and the response body contains the representation requested.|
|302 |FOUND| A common redirect response; you can GET the representation at the URI in the Location response header.|
|304 |NOT MODIFIED| Your client's cached version of the representation is still up to date|
|401 |UNAUTHORIZED| The supplied credentials, if any, are not sufficient to access the resource|
|404 |NOT FOUND| You know this one|
|429 |TOO MANY REQUESTS| Your application is sending too many simultaneous requests|
|500 |SERVER ERROR| We couldn't return the representation due to an internal server error|
|503 |SERVICE UNAVAILABLE| We are temporarily unable to return the representation. Please wait for a bit and try again.|
|901 |CONNECTIVITY ERROR| Unable to establish connection to Twilio URL specified.|
|1001| INTERNAL ERROR| Check the message field in response for more details on the error.|
|504| URL BLANK ERROR|Url String field is blank.|
|505|SID BLANK ERROR|Account SID field is blank.|
|506|AUTH TOKEN BLANK ERROR|Auth Token String field is blank.|

### Send SMS
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request was successful, we updated the resource and the response body contains the representation.|
|201 |CREATED| The request was successful, we created a new resource and the response body contains the representation.|
|400 |BAD REQUEST| The data given in the POST or PUT failed validation. Inspect the response body for details.|
|401 |UNAUTHORIZED| The supplied credentials, if any, are not sufficient to create or update the resource.|
|404 |NOT FOUND| You know this one.|
|405 |METHOD NOT ALLOWED| You can't POST or PUT to the resource.|
|429 |TOO MANY REQUESTS| Your application is sending too many simultaneous requests.|
|500 |SERVER ERROR| We couldn't create or update the resource.Please try again.|
|901 |CONNECTIVITY ERROR| Unable to establish connection to Twilio URL specified.|
|1001| INTERNAL ERROR| Check the message field in response for more details on the error.|
|504| URL BLANK ERROR|Url String field is blank.|
|505|SID BLANK ERROR|Account SID field is blank.|
|506|AUTH TOKEN BLANK ERROR|Auth Token String field is blank.|
|501|SMS BODY BLANK ERROR|SMS body is blank.|
|502|TO FIELD BLANK ERROR|To field is blank.|
|503|FROM FIELD BLANK ERROR|From field is blank.|