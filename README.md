# crestrontomqtt_flogo
This activity convert a Crestron mqtt message to a new splitted and evaluated message


## Installation

```bash
flogo install github.com/ultrafab/crestrontomqtt_flogo
```
Link for flogo web:
```
https://github.com/ultrafab/crestrontomqtt_flogo
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "address",
      "type": "string"
    },
    {
      "name": "dbNo",
      "type": "integer"
    },
    {
      "name": "message",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "mqtt_message",
      "type": "string"
    },
    {
      "name": "topic",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| host    | the Redis address + port |
| dbNo    | the Redis database number |
| message    | the Crestron sentence |

## Ouputs
| Output   | Description    |
|:----------|:---------------|
| mqtt_message    | the composed mqtt messsage |
| topic    | the topic |
