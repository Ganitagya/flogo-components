# Send SQS Message

Send a message using Amazon Simple Queue Service (SQS)

## Installation

```bash
flogo install github.com/Ganitagya/flogo-components/activity/getactivespacestable
```
Link for flogo web:
```
https://github.com/Ganitagya/flogo-components/activity/getactivespacestable
```

## Schema
Inputs and Outputs:

```json
{
    "inputs": [
        {
            "name": "connectionURL",
            "type": "string",
            "required": true
        },
        {
            "name": "tableName",
            "type": "string",
            "required": true
        },
        {
            "name": "key",
            "type": "int",
            "required": true
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "string"
        }
    ]
}
```
## Inputs
| Input              | Description                                                                                                 |
|:-------------------|:------------------------------------------------------------------------------------------------------------|
| connectionURL      | ActiveSpaces connection URL (eg: "http://localhost:8080")                                                   |
| tableName          | Table name                                                                                                  |
| key                | Key of the value you wish to get                                                                            |


## Ouputs
| Output    | Description                       |
|:----------|:----------------------------------|
| result    | The JSON of the Key value pair    |
