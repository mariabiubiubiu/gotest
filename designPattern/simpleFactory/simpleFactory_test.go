package simpleFactory

import "testing"

func TestCar_Create(t *testing.T) {
	api := NewFactory(1)
	s := api.Create()
	if s != "Create Car Done"{
		t.Fatal("failed")
	}
}


func TestBike_Create(t *testing.T) {
	api := NewFactory(2)
	s := api.Create()
	if s != "Create Bike Done"{
		t.Fatal("failed")
	}
}