package handler

import (
	"app/internal"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id" default:"-1"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year" default:"-1"`
	Capacity        int     `json:"passengers" default:"-1"`
	MaxSpeed        float64 `json:"max_speed" default:"-1.1"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight" default:"-1.1"`
	Height          float64 `json:"height" default:"-1.1"`
	Length          float64 `json:"length" default:"-1.1"`
	Width           float64 `json:"width" default:"-1.1"`
}

type VehiclesJSON struct {
	Vehicles []VehicleJSON
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		vv := newVehicleValidator()
		data := vv.VehicleMap2VehicleJSONMap(v)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		// process
		v, err := h.sv.FindByID(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		// response
		vv := newVehicleValidator()
		data := vv.Vehicle2VehicleJSON(v)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) DeleteByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		// process
		h.sv.DeleteVehicle(id)
		// response
		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (h *VehicleDefault) PutUpdateMaxSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		payloadJSON, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		// process
		var payload map[string]any
		json.Unmarshal(payloadJSON, &payload)

		maxSpeed, ok := payload["max_speed"]
		if !ok {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": "new max speed in wrong format",
			})
			return
		}
		err = h.sv.UpdateMaxSpeed(id, float64(maxSpeed.(float64)))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
		})
	}
}

func (h *VehicleDefault) PostVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		vv := newVehicleValidator()
		newVehicle, err := vv.VehicleJSON2Vehicle(reqPayload)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		err = h.sv.SaveVehicle(newVehicle)
		if err != nil {
			response.JSON(w, http.StatusConflict, map[string]any{
				"error": err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
		})
	}
}

func (h *VehicleDefault) PostVehicleGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		vv := newVehicleValidator()
		var vehiclesJSON VehiclesJSON
		json.Unmarshal(reqPayload, &vehiclesJSON)
		for _, vehicleJSON := range vehiclesJSON.Vehicles {
			newVehicleJSON, _ := json.Marshal(vehicleJSON)
			newVehicle, err := vv.VehicleJSON2Vehicle(newVehicleJSON)
			if err != nil {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"error": err.Error(),
				})
				return
			}
			err = h.sv.SaveVehicle(newVehicle)
			if err != nil {
				response.JSON(w, http.StatusConflict, map[string]any{
					"error": err.Error(),
				})
				return
			}
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
		})
	}
}

func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		v, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		}

		// response
		vv := newVehicleValidator()
		data := vv.VehicleMap2VehicleJSONMap(v)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) GetByBrandAndYearRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		yearBegin, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		yearEnd, err := strconv.Atoi(chi.URLParam(r, "end_year"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
			return
		}
		v, err := h.sv.FindByBrandAndYearRange(brand, yearBegin, yearEnd)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		}

		// response
		vv := newVehicleValidator()
		data := vv.VehicleMap2VehicleJSONMap(v)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) GetAvgSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		v, err := h.sv.FindAvgSpeedByBrand(brand)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"error": err.Error(),
			})
			return
		}

		// response

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"average": v,
		})

	}
}
