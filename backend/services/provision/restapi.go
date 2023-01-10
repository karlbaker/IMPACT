package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type device struct {
	ID			string `json:"id"`
	Manufacture	string `json:"manufacture"`
	Model		string `json:"model"`
	Serial		string `json:"serial"`
	MAC			string `json:"mac"`
}

var devices = []device {
	{ID: "dc424f54-91a9-4b4a-8f63-3af523d7f596", Manufacture: "Cisco Systems, Inc.", Model: "C220 M4", Serial: "FCH2112V1DE", MAC: "00A38E397AFF"},
	{ID: "dc424f54-91a9-4b4a-8f63-3af523d7f597", Manufacture: "Cisco Systems, Inc.", Model: "C220 M4", Serial: "FCH2112V2DE", MAC: "00A38E497AFF"},
	{ID: "dc424f54-91a9-4b4a-8f63-3af523d7f598", Manufacture: "Cisco Systems, Inc.", Model: "C220 M4", Serial: "FCH2112V3DE", MAC: "00A38E597AFF"},
}

func getDevices(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, devices)
}

func patchDevice(context *gin.Context) {
	id := context.Param("id")
	var Device device
	if err := context.ShouldBindJSON(&Device); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i := 0; i < len(devices); i++ {
		if devices[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Device not found",
		})
		return
	}
	devices[index] = Device
	context.JSON(http.StatusOK, Device)
}

func postDevice(context *gin.Context) {
	 var newDevice device

	 if err := context.BindJSON(&newDevice); err != nil {
		return
	 }
	 devices = append(devices, newDevice)
	 context.IndentedJSON(http.StatusCreated, newDevice)
}

func getDevicebyID(id string) (*device, error) {
	for i, p := range devices {
		if p.ID == id {
			return &devices[i], nil
		}
	}
	return nil, errors.New("Port not found.")
}

func getDevice(context *gin.Context) {
	id := context.Param("id")
	port, err := getDevicebyID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Device not found."})
	}

	context.IndentedJSON(http.StatusOK, port)
}

func addDevices(context *gin.Context) {
	var newDevice device

	if err := context.BindJSON(&newDevice); err != nil {
		return
	}

	devices = append(devices, newDevice)
	
	context.IndentedJSON(http.StatusCreated, newDevice)
}


func main() {
	router := gin.Default()

	router.GET("/api/provision/",getDevices)
	router.GET("/api/provision/:id",getDevice)
	router.PATCH("/api/provision/:id",patchDevice)
	router.POST("/api/provision/", addDevices)
	router.Run("localhost:8080")
}