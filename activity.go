package crestrontomqtt_flogo

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/go-redis/redis"
	"strings"
	"strconv"
	"encoding/json"
)

var log = logger.GetLogger("crestrontomqtt_log")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}
type Variable struct {
	Key string `json:"key"`
	Desc  string `json:"desc"`
	Func string `json:"func"`
	Indexes []struct {
				Index int `json:"index"`
				Name string `json:"name"`
				Type string `json:"type"`
		} `json:"indexes"`
}

func MyFunction(a, b int) int {
  return a + b
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {
	address := context.GetInput("address").(string)
	dbNo := context.GetInput("dbNo").(int)
	message := context.GetInput("message").(string)
	log.Infof("Connecting to Redis: [%s]", address)
	log.Infof("DB no: [%s]",dbNo)

	splitted := strings.Split(message, ".")

	client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "", // no password set
			DB:       dbNo,  // use default DB
		})

		val, err := client.Get(splitted[0]).Result()
		mqtt_message := ""
		topic := ""
		m := make(map[string]interface{})

		if err != nil {
			log.Errorf("Error! Not found Redis var: %s", splitted[0] )
		} else {
			bytes := []byte(val)

			var v Variable
				err2 := json.Unmarshal(bytes, &v)
				if err2 != nil {
					panic(err2)
				}
			topic = v.Key
			m["desc"] = v.Desc
			if v.Func == "convertCrestronToFloat" {
				val,errs := strconv.Atoi(splitted[1])
				if errs != nil {
					log.Errorf("Error! Could not parse : %s", splitted[1] )
				}
				m["value"] = float32(val)/float32(10)
			}
			if v.Func == "convertCrestronToInt" {
				val,errs := strconv.Atoi(splitted[1])
				if errs != nil {
					log.Errorf("Error! Could not parse : %s", splitted[1] )
				}
				m["value"] = val
			}
		data, _ := json.Marshal(m)
		mqtt_message = string(data)
	}
	context.SetOutput("mqtt_message", mqtt_message)
	context.SetOutput("topic", topic)
	client.Close()
	return true, nil
}
