package rpi2cube

import (
	"testing"
)

func TestTrue(t *testing.T){
  if ( false ){
		t.Errorf("nope")
	}
}

func TestNewSensor(t *testing.T){
  if (NewSensor("test").Name != "test"){
    t.Errorf("New Sensor not setting name.")
  }
}
