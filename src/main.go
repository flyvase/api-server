package main

import (
	"fmt"
	"harvest/src/adapter/gateway/sql"
	"harvest/src/application/repository"
	"harvest/src/domain/value"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sqlDriver := sql.NewDriver()

	repo := repository.SpaceImpl{
		SqlDriver: sqlDriver,
	}

	space, _ := repo.Fetch(value.SpaceId{
		Value: 1,
	})
	fmt.Println(space)

	// res, _ := repo.List()
	// for _, v := range res {
	// 	fmt.Println(v.Id)
	// 	fmt.Println(v.Headline)
	// 	fmt.Println(v.Access)
	// 	fmt.Println(v.NumberOfVisitors)
	// 	fmt.Println(v.CustomerSegment)
	// 	fmt.Println(v.Price)
	// 	for _, i := range v.Images {
	// 		fmt.Println(i)
	// 	}
	// }
}
