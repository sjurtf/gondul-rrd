package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const gondulURL = "https://nms.tg18.gathering.org"

func main() {

	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Graph
	r.GET("/graph", graphHandler)

	// Gondul
	// Proxy gondul api so i dont have to wait for CORS fix
	r.GET("/gondul/distro-tree", gondulDistroTree)
	r.GET("/gondul/switch-state", gondulSwitchState)

	// Web
	r.Static("/web", "./web/static")

	//	r.LoadHTMLFiles("web/index.html")
	//	r.GET("/", indexHandler)

	r.Run(":8080")
}

//func indexHandler(c *gin.Context) {
//	c.HTML(http.StatusOK, "index.html", nil)
//}

func gondulDistroTree(c *gin.Context) {
	distroURL := gondulURL + "/api/public/distro-tree"

	resp, err := http.Get(distroURL)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var d DistroTree
	json.Unmarshal(body, &d)
	c.JSON(http.StatusOK, d)

}

type DistroTree struct {
	Hash   string                       `json:"hash"`
	Time   int                          `json:"time"`
	Distro map[string]map[string]string `json:"distro-tree"`
}

func gondulSwitchState(c *gin.Context) {
	distroURL := gondulURL + "/api/public/switch-state"

	resp, err := http.Get(distroURL)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var s Gondul
	json.Unmarshal(body, &s)
	c.JSON(http.StatusOK, s)

}

type Gondul struct {
	Hash     string               `json:"hash"`
	Switches map[string]*switches `json:"switches"`
	Time     int64                `json:"time,omitempty"`
}

type switches struct {
	Time    string          `json:"time"`
	Temp    string          `json:"temp"`
	Ifs     map[string]*ifs `json:"ifs"`
	Clients clients         `json:"clients"`
	Vcp     vcp             `json:"vcp,omitempty"`
	Totals  totals          `json:"totals"`
	Uplinks uplinks         `json:"uplinks"`
}

type ifs struct {
	IfHCOutOctets uint64 `json:"ifHCOutOctets"`
	IfHCInOctets  uint64 `json:"ifHCInOctets"`
}

type vcp struct {
	VcpIntOut map[string]map[string]string `json:"jnxVirtualChassisPortOutOctets,omitempty"`
	VcpIntIn  map[string]map[string]string `json:"jnxVirtualChassisPortInOctets,omitempty"`
}

type clients struct {
	IfHCOutOctets uint64 `json:"ifHCOutOctets"`
	IfHCInOctets  uint64 `json:"ifHCInOctets"`
}

type totals struct {
	IfHCOutOctets uint64 `json:"ifHCOutOctets"`
	IfHCInOctets  uint64 `json:"ifHCInOctets"`
}

type uplinks struct {
	IfHCOutOctets uint64 `json:"ifHCOutOctets"`
	IfHCInOctets  uint64 `json:"ifHCInOctets"`
}
