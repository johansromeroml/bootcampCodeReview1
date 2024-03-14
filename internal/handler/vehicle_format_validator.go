package handler

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"reflect"
)

type VehicleValidator struct {
}

func newVehicleValidator() *VehicleValidator {
	return &VehicleValidator{}
}

var defaultVehicleJSON = VehicleJSON{}

func (vv *VehicleValidator) VehicleJSON2Vehicle(vjson []byte) (v internal.Vehicle, err error) {
	var vj VehicleJSON
	err = json.Unmarshal(vjson, &vj)
	attributeTypes := reflect.ValueOf(defaultVehicleJSON).Type()
	vjAttributes := reflect.ValueOf(vj)
	for i := 0; i < attributeTypes.NumField(); i++ {
		fieldName := attributeTypes.Field(i).Name
		//fmt.Println(fieldName, "-", vjAttributes.Field(i))
		switch attributeTypes.Field(i).Type.Kind() {
		case reflect.Int:
			if vjAttributes.FieldByName(fieldName).Int() < 0 {
				err = fmt.Errorf(internal.ErrorEmptyField)
			}
		case reflect.Float64:
			if vjAttributes.FieldByName(fieldName).Float() < 0 {
				err = fmt.Errorf(internal.ErrorEmptyField)
			}
		case reflect.String:
			if vjAttributes.FieldByName(fieldName).String() == "" {
				err = fmt.Errorf(internal.ErrorEmptyField)
			}
		}
	}
	v = internal.Vehicle{
		Id: vj.ID,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           vj.Brand,
			Model:           vj.Model,
			Registration:    vj.Registration,
			Color:           vj.Color,
			FabricationYear: vj.FabricationYear,
			Capacity:        vj.Capacity,
			MaxSpeed:        vj.MaxSpeed,
			FuelType:        vj.FuelType,
			Transmission:    vj.Transmission,
			Weight:          vj.Weight,
			Dimensions: internal.Dimensions{
				Height: vj.Height,
				Length: vj.Length,
				Width:  vj.Width,
			},
		},
	}
	return
}

func (vv *VehicleValidator) Vehicle2VehicleJSON(v internal.Vehicle) (vj VehicleJSON) {
	vj = VehicleJSON{
		ID:              v.Id,
		Brand:           v.Brand,
		Model:           v.Model,
		Registration:    v.Registration,
		Color:           v.Color,
		FabricationYear: v.FabricationYear,
		Capacity:        v.Capacity,
		MaxSpeed:        v.MaxSpeed,
		FuelType:        v.FuelType,
		Transmission:    v.Transmission,
		Weight:          v.Weight,
		Height:          v.Height,
		Length:          v.Length,
		Width:           v.Width,
	}
	return
}

func (vv *VehicleValidator) VehicleMap2VehicleJSONMap(vehicles map[int]internal.Vehicle) (vj map[int]VehicleJSON) {
	vj = make(map[int]VehicleJSON)
	for key, value := range vehicles {
		vj[key] = VehicleJSON{
			ID:              value.Id,
			Brand:           value.Brand,
			Model:           value.Model,
			Registration:    value.Registration,
			Color:           value.Color,
			FabricationYear: value.FabricationYear,
			Capacity:        value.Capacity,
			MaxSpeed:        value.MaxSpeed,
			FuelType:        value.FuelType,
			Transmission:    value.Transmission,
			Weight:          value.Weight,
			Height:          value.Height,
			Length:          value.Length,
			Width:           value.Width,
		}
	}
	return
}
