// Package getactivespacestable gets Table data from ActiveSpaces
package getactivespacestable

import (

	"github.com/Ganitagya/activespacetest"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	ivconnectionURL = "connectionURL"
	ivtableName     = "tableName"
	ivkey           = "key"
	ovResult        = "result"
)

// log is the default package logger
var log1 = logger.GetLogger("activity-getactivespacestable")

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

// ExpressionAttribute is a structure representing the JSON payload for the expression syntax
type ExpressionAttribute struct {
	Name  string
	Value string
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	log1.Info("Yest")
	// Get the inputs
	
	activespacetest.Print()
	
	connectionURL := context.GetInput(ivconnectionURL).(string)
	tableName := context.GetInput(ivtableName).(string)
	key := context.GetInput(ivkey).(string)


	context.SetOutput(ovResult, "getRow")
	return true, nil

}
