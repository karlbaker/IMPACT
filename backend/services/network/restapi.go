package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type port struct {
	ID			string `json:"id"`
	Switch		string `json:"switch"`
	Port		string `json:"port"`
	VLAN		string `json:"vlan"`
}

var ports = []port {
	{ID: "C10-R12-RU47-001", Switch: "192.168.50.221", Port: "GigabitEthernet1/0/1", VLAN: "Customer1-Data"},
	{ID: "C10-R12-RU47-002", Switch: "192.168.50.221", Port: "GigabitEthernet1/0/2", VLAN: "Customer2-Data"},
	{ID: "C10-R12-RU47-003", Switch: "192.168.50.221", Port: "GigabitEthernet1/0/3", VLAN: "Customer3-Data"},
}

func getPorts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, ports)
}

func patchPort(context *gin.Context) {
	id := context.Param("id")
	var Port port
	if err := context.ShouldBindJSON(&Port); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	index := -1
	for i := 0; i < len(ports); i++ {
		if ports[i].ID == id {
			index = 1
		}
	}
	if index == -1 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Port not found",
		})
		return
	}
	ports[index] = Port
	context.JSON(http.StatusOK, Port)
}

func postPort(context *gin.Context) {
	 var newPort port

	 if err := context.BindJSON(&newPort); err != nil {
		return
	 }
	 ports = append(ports, newPort)
	 context.IndentedJSON(http.StatusCreated, newPort)
}

func getPortbyID(id string) (*port, error) {
	for i, p := range ports {
		if p.ID == id {
			return &ports[i], nil
		}
	}
	return nil, errors.New("Port not found.")
}

func getPort(context *gin.Context) {
	id := context.Param("id")
	port, err := getPortbyID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Port not found."})
	}

	context.IndentedJSON(http.StatusOK, port)
}

func addPorts(context *gin.Context) {
	var newPort port

	if err := context.BindJSON(&newPort); err != nil {
		return
	}

	ports = append(ports, newPort)
	
	context.IndentedJSON(http.StatusCreated, newPort)
}


func main() {
	router := gin.Default()

	router.GET("/api/network/",getPorts)
	router.GET("/api/network/:id",getPort)
	router.PATCH("/api/network/:id",patchPort)
	router.POST("/api/network/", addPorts)
	router.Run("localhost:8181")
}