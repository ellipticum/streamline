package storage

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func Find(database, collection string, resultType interface{}, filter interface{}, delay time.Duration) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), delay)

	defer cancel()

	if reflect.TypeOf(resultType).Kind() != reflect.Ptr || reflect.TypeOf(resultType).Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("resultType must be a pointer to a struct")
	}

	cursor, err := Client.Database(database).Collection(collection).Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	sliceType := reflect.SliceOf(reflect.TypeOf(resultType))

	results := reflect.MakeSlice(sliceType, 0, 0)

	for cursor.Next(ctx) {
		elem := reflect.New(reflect.TypeOf(resultType).Elem()).Interface()

		err := cursor.Decode(elem)

		if err != nil {
			return nil, err
		}

		results = reflect.Append(results, reflect.ValueOf(elem))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results.Interface(), nil
}
