package main

import (
	"fmt"

	"github.com/goSamples/mysqlSample/service"
)

func main() {
	r := service.NewUserPointRepository()
	result := r.GetTotalPointByUserID(1)
	fmt.Println(result)
}
