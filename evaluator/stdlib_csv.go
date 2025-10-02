package evaluator

import (
	"encoding/csv"
	"os"

	"github.com/LabMarket/roach/object"
)

func csvRead(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	path, ok := args[0].(*object.String)
	if !ok {
		return newError("argument to `csv.read` must be STRING, got %s", args[0].Type())
	}

	file, err := os.Open(path.Value)
	if err != nil {
		return newError("could not open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return newError("could not read csv file: %s", err)
	}

	rows := make([]object.Object, len(records))
	for i, record := range records {
		cols := make([]object.Object, len(record))
		for j, col := range record {
			cols[j] = &object.String{Value: col}
		}
		rows[i] = &object.Array{Elements: cols}
	}

	return &object.Array{Elements: rows}
}

func csvWrite(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}
	path, ok := args[0].(*object.String)
	if !ok {
		return newError("first argument to `csv.write` must be STRING, got %s", args[0].Type())
	}
	data, ok := args[1].(*object.Array)
	if !ok {
		return newError("second argument to `csv.write` must be ARRAY, got %s", args[1].Type())
	}

	file, err := os.Create(path.Value)
	if err != nil {
		return newError("could not create file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, rowObj := range data.Elements {
		row, ok := rowObj.(*object.Array)
		if !ok {
			return newError("data must be an array of arrays")
		}

		record := make([]string, len(row.Elements))
		for i, colObj := range row.Elements {
			col, ok := colObj.(*object.String)
			if !ok {
				return newError("data must be an array of arrays of strings")
			}
			record[i] = col.Value
		}

		if err := writer.Write(record); err != nil {
			return newError("error writing record to csv: %s", err)
		}
	}

	return NULL
}

func init() {
	RegisterBuiltin("csv.read",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (csvRead(args...))
		})
	RegisterBuiltin("csv.write",
		func(env *object.Environment, args ...object.Object) object.Object {
			return (csvWrite(args...))
		})
}
