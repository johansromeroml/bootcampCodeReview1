package service

import (
	"app/internal"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) FindByID(id int) (v internal.Vehicle, err error) {
	v, err = s.rp.FindByID(id)
	return
}

func (s *VehicleDefault) UpdateMaxSpeed(id int, newMaxSpeed float64) (err error) {
	err = s.rp.UpdateAttribute(id, "MaxSpeed", newMaxSpeed)
	return
}

func (s *VehicleDefault) SaveVehicle(v internal.Vehicle) (err error) {
	err = s.rp.SaveVehicle(v)
	return
}

func (s *VehicleDefault) DeleteVehicle(id int) {
	s.rp.DeleteVehicle(id)
}

func findCommon(vm1 map[int]internal.Vehicle, vm2 map[int]internal.Vehicle) (cvm map[int]internal.Vehicle) {
	cvm = make(map[int]internal.Vehicle)
	for key := range vm1 {
		if _, ok := vm2[key]; ok {
			cvm[key] = vm1[key]
		}
	}
	for key := range vm2 {
		if _, ok := vm1[key]; ok {
			cvm[key] = vm2[key]
		}
	}
	return
}

func (s *VehicleDefault) FindByColorAndYear(color string, year int) (vehicles map[int]internal.Vehicle, err error) {
	colorVehicles, err := s.rp.FindByAttribute("Color", color)
	if err != nil {
		return
	}
	yearVehicles, err := s.rp.FindByAttribute("FabricationYear", year)
	if err != nil {
		return
	}
	vehicles = findCommon(colorVehicles, yearVehicles)
	return
}

func (s *VehicleDefault) FindByBrandAndYearRange(brand string, yearBegin int, yearEnd int) (vehicles map[int]internal.Vehicle, err error) {
	brandVehicles, err := s.rp.FindByAttribute("Brand", brand)
	if err != nil {
		return
	}
	yearVehicles, err := s.rp.FindByAttributeRange("FabricationYear", yearBegin, yearEnd)
	if err != nil {
		return
	}
	vehicles = findCommon(brandVehicles, yearVehicles)
	return
}

func (s *VehicleDefault) FindAvgSpeedByBrand(brand string) (average float64, err error) {
	brandVehicles, err := s.rp.FindByAttribute("Brand", brand)
	if err != nil {
		return
	}
	sum := 0.0
	for key := range brandVehicles {
		sum += brandVehicles[key].MaxSpeed
	}
	average = sum / float64(len(brandVehicles))
	return
}
