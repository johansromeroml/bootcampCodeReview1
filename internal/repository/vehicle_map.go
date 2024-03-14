package repository

import (
	"app/internal"
	"fmt"
	"reflect"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) FindByID(id int) (v internal.Vehicle, err error) {
	v, ok := r.db[id]
	if !ok {
		err = fmt.Errorf(internal.ErrorNotFound)
	}
	return
}

func (r *VehicleMap) UpdateAttribute(id int, attribute string, value any) (err error) {
	v, ok := r.db[id]
	if !ok {
		err = fmt.Errorf(internal.ErrorNotFound)
	}
	//fmt.Println(v.MaxSpeed)
	reflect.ValueOf(&v).Elem().FieldByName(attribute).Set(reflect.ValueOf(value))
	//fmt.Println(v.MaxSpeed)
	r.db[id] = v
	return
}

func (r *VehicleMap) SaveVehicle(v internal.Vehicle) (err error) {
	_, ok := r.db[v.Id]
	if !ok {
		r.db[v.Id] = v
	} else {
		err = fmt.Errorf(internal.ErrorVehicleAlreadyCreated)
	}
	return
}

func (r *VehicleMap) DeleteVehicle(id int) {
	delete(r.db, id)
}

func (r *VehicleMap) FindByAttribute(attribute string, value any) (vehicles map[int]internal.Vehicle, err error) {
	vehicles = make(map[int]internal.Vehicle)
	// copy db
	for key, vehic := range r.db {
		vehicAttributes := reflect.ValueOf(vehic)
		if vehicAttributes.FieldByName(attribute).Equal(reflect.ValueOf(value)) {
			vehicles[key] = vehic
		}
	}
	if len(vehicles) == 0 {
		err = fmt.Errorf(internal.ErrorNotFound)
	}
	return
}

func (r *VehicleMap) FindByAttributeRange(attribute string, begin any, end any) (vehicles map[int]internal.Vehicle, err error) {
	vehicles = make(map[int]internal.Vehicle)
	// copy db
	for key, vehic := range r.db {
		vehicAttributes := reflect.ValueOf(vehic)
		field := vehicAttributes.FieldByName(attribute)
		switch {
		case field.CanInt():
			vehicAttrib := field.Int()
			if vehicAttrib >= int64(begin.(int)) && vehicAttrib <= int64(end.(int)) {
				vehicles[key] = vehic
			}
		case field.CanFloat():
			vehicAttrib := field.Float()
			if vehicAttrib >= float64(begin.(float64)) && vehicAttrib <= float64(end.(float64)) {
				vehicles[key] = vehic
			}
		}

	}
	if len(vehicles) == 0 {
		err = fmt.Errorf(internal.ErrorNotFound)
	}
	return
}
