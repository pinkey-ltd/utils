package utils

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jonas-p/go-shp"
)

func ReadShp(path string) (string, error) {
	// open a shapefile for reading
	shape, err := shp.Open(path)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer func(shape *shp.Reader) {
		err := shape.Close()
		if err != nil {
			log.Println("close shape file error: ", err)
		}
	}(shape)

	// fields from the attribute table (DBF)
	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()

		// print feature
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}
		fmt.Println()
	}
	return "", nil
}
