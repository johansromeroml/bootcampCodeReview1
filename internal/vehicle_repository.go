package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	FindByID(id int) (v Vehicle, err error)
	// SaveVehicle is a method that saves a vehicle in the repository
	SaveVehicle(v Vehicle) (err error)

	DeleteVehicle(id int)

	FindByAttribute(attribute string, value any) (vehicles map[int]Vehicle, err error)

	FindByAttributeRange(attribute string, begin any, end any) (vehicles map[int]Vehicle, err error)

	UpdateAttribute(id int, attribute string, value any) (err error)
}
