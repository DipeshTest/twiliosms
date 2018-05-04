---
title: Post Tweet
weight: 1
---

# Counter
This activity allows you to post a tweet using your twitter account credentials.

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
| urlString | True     | Twilio URI |  
| method | True     | We expose 2 methods, one to send SMS to a recipient, second to retrieve the list of recipients verified @ your account |  
| messageBody | True     | The body of SMS |  
| to       | True    |Your Twilio number to send SMS using|
| from       | True    | The number to send SMS to |

## Examples
### Increment
The below example to send SMS to multiple recipients:

```json
{
	"AccountSid": "AC8112820c2969c0b9ba6abac8ee6a4062",
	"AuthToken": "f2542c88dbfb58c494e642bf10af4140",
	"UrlString": "https://api.twilio.com/2010-04-01/Accounts/",
	"method":"Send SMS",
	"MsgData": "Final Code",
	"To": "+9189908098098,+9189908098098",
	"From": "+14437433811"
}
```

### Get
The below example to send SMS to single user:

```json
{
	"AccountSid": "AC8112820c2969c0b9ba6abac8ee6a4062",
	"AuthToken": "f2542c88dbfb58c494e642bf10af4140",
	"UrlString": "https://api.twilio.com/2010-04-01/Accounts/",
	"method":"Send SMS",
	"MsgData": "Final Code",
	"To": "+9189908098098",
	"From": "+14437433811"
}
```

### Reset
The below example resets the 'messages' counter:

```json
{
	"AccountSid": "AC8112820c2969c0b9ba6abac8ee6a4062",
	"AuthToken": "f2542c88dbfb58c494e642bf10af4140",
	"UrlString": "https://api.twilio.com/2010-04-01/Accounts/",
	"method":"Retrive Recipients"
}
```