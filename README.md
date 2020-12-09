# Stream

![GitHub CI](https://github.com/MufidJamaluddin/stream/workflows/Go/badge.svg)

Do you want to fetch data, process it, and send it or collect it later without store all initial data in memory? Stream is the best choice for you!
 
## Instalation

Add to your Go project by `go get github.com/mufidjamaluddin/stream`

## Feature

1. Map & Filter Data

You can map your data to another type or filter it with your criteria

2. Multiple Collector

You can collect result to multiple destination which implemented ICollector interface


## Usage

```
	inStream := &Stream{
		enter: func(feedFunc func(interface{})) error {
		  
			rows, err := db.Query(`SELECT u.id, u.price FROM products u WHERE YEAR(created_at) = 2020 AND type IN ('laptop', 'pc');`)
			if err != nil { 
			     log.Fatal(err) 
			}
			defer rows.Close()

			for i := 0; rows.Next(); i++ {
			     if err = rows.Scan(&product.Id, &product.Price); err != nil {
				feedFunc(product)
			     }
			}

			return err
		},
	}

	err := inStream.
		Filter(func(item interface{}) interface{} {
			return item.(product).Price > 1000 && item.(product).Price < 10000
		}).
		Map(func(item interface{}) interface{} {
			return MapToComputerModel(item)
		}).
		Filter(func(item interface{}) interface{} {
			return IsComputerPriceUnnatural(item.(computer))
		}).
		Map(func(item interface{}) interface{} {
			return MapToComputerScamModel(item)
		}).
        	Collect(ScamCollector, ScamChannel, ProductDBUpdate).
		Run()
```

## Notes

Inspired by Java Streams but simplest and most widely used
