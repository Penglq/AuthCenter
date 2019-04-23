package main

import (
	"fmt"
	"github.com/Penglq/AuthCenter/routers"
)

func main()  {
	fmt.Println(routers.NewRouter().Run(":8989"))
}
