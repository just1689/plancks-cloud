package controller

import (
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/plancks-cloud/plancks-docker/controller/pc-docker"
	pcmodel "github.com/plancks-cloud/plancks-docker/model"
	"time"
)

func StartHealthServer() {

	go func() {
		for {
			GetDesiredServiceStateAndFix()
			time.Sleep(60 * time.Second)
		}
	}()

}

func GetDesiredServiceStateAndFix() {
	svcChan := GetAllServices()
	dockerServices, err := pc_docker.GetAllServiceStates()
	if err != nil {
		fmt.Println(err)
		return
	}
	checkExistingServices(svcChan, dockerServices)

}

func checkExistingServices(desiredServices chan *model.Service, foundServices []pcmodel.ServiceState) {
	for d := range desiredServices {
		found := false
		for _, s := range foundServices {
			if d.Name == s.Name {
				found = true
			}
		}
		if !found {
			err := pc_docker.CreateService(d)
			if err != nil {
				fmt.Println("Could not start docker service ", err)
			}
		}
	}

}
