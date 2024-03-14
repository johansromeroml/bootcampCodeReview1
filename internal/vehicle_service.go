package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	FindByID(id int) (v Vehicle, err error)
	SaveVehicle(v Vehicle) (err error)
	DeleteVehicle(id int)
	FindByColorAndYear(color string, year int) (vehicles map[int]Vehicle, err error)
	FindByBrandAndYearRange(brand string, yearBegin int, yearEnd int) (vehicles map[int]Vehicle, err error)
	FindAvgSpeedByBrand(brand string) (average float64, err error)
	UpdateMaxSpeed(id int, newMaxSpeed float64) (err error)
}
