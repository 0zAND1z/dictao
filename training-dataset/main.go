package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

// Tuple defines the structure of the JSON objects
// generated from https://www.json-generator.com/
type Tuple struct {
	ID        string `json:"id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Timestamp uint   `json:"timestamp"`
	Infection uint8  `json:"infection"`
}

// TupleNew defines the refined format of the JSON objects
// needed to train the DecisionTreeClassifier model
type TupleNew struct {
	ID        string  `json:"id"`
	Timestamp uint    `json:"timestamp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Infection uint8   `json:"infection"`
}

func main() {
	fileBytes, err := ioutil.ReadFile("./dataset.json")
	if err != nil {
		log.Panicln("Unable to read the dataset.json file: ", err)
	}
	var TuplesObject []Tuple
	json.Unmarshal(fileBytes, &TuplesObject)

	NewTuplesObject := ConvToFloat(TuplesObject)
	// fmt.Println(NewTuplesObject)

	var prev uint
	// count := 0
	for index, tuple := range NewTuplesObject {
		if tuple.ID == "BELLA" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if NewTuplesObject[index+1].ID != "BELLA" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 600

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.0000118
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.0000118
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index].Infection = 1
				NewTuplesObject[index+1].Infection = 1
			}
			fmt.Println("Found ", count, "entries of \"BELLA\"")
		} else if tuple.ID == "BUDDY" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if NewTuplesObject[index+1].ID != "BUDDY" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 600

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.0000118
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.0000118
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index].Infection = 1
				NewTuplesObject[index+1].Infection = 1
			}
			fmt.Println("Found ", count, "entries of \"BUDDY\"")
		} else if tuple.ID == "COCO" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if NewTuplesObject[index+1].ID != "COCO" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 600

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.0000118
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.0000118
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index].Infection = 1
				NewTuplesObject[index+1].Infection = 1
			}
			fmt.Println("Found ", count, "entries of \"COCO\"")
		} else if tuple.ID == "LUCY" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if index+1 < len(NewTuplesObject) {
				if NewTuplesObject[index+1].ID != "LUCY" {
					NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 600

					NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.0000118
					fmt.Println("Converting...")
					s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
					fmt.Println("s lat = ", s)
					if f, err := strconv.ParseFloat(s, 64); err == nil {
						fmt.Println(f) // 3.14159265
						NewTuplesObject[index+1].Latitude = f
					}

					NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.0000118
					fmt.Println("Converting...")
					s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
					fmt.Println("s long = ", s)
					if f, err := strconv.ParseFloat(s, 64); err == nil {
						fmt.Println(f) // 3.14159265
						NewTuplesObject[index+1].Longitude = f
					}
					NewTuplesObject[index].Infection = 1
					NewTuplesObject[index+1].Infection = 1
				}
			}
			fmt.Println("Found ", count, "entries of \"LUCY\"")
		} else if tuple.ID == "MAX" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if NewTuplesObject[index+1].ID != "MAX" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 600

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.0000118
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.0000118
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index].Infection = 1
				NewTuplesObject[index+1].Infection = 1
			}
			fmt.Println("Found ", count, "entries of \"MAX\"")
		} else if tuple.ID == "MILO" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if index+1 >= len(NewTuplesObject) {
				continue
			} else if NewTuplesObject[index+1].ID != "MILO" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 90000

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.01
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.01
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index+1].Infection = 0
			}
			fmt.Println("Found ", count, "entries of \"MILO\"")
		} else if tuple.ID == "MOLLY" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if index+1 >= len(NewTuplesObject) {
				continue
			} else if NewTuplesObject[index+1].ID != "MOLLY" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 90000

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.01
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.01
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index+1].Infection = 0
			}
			fmt.Println("Found ", count, "entries of \"MOLLY\"")
		} else if tuple.ID == "OSCAR" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if index+1 >= len(NewTuplesObject) {
				continue
			} else if NewTuplesObject[index+1].ID != "OSCAR" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 90000

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.01
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.01
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index+1].Infection = 0
			}
			fmt.Println("Found ", count, "entries of \"OSCAR\"")
		} else if tuple.ID == "RUBY" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if index+1 >= len(NewTuplesObject) {
				continue
			} else if NewTuplesObject[index+1].ID != "RUBY" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 90000

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.01
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.01
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index+1].Infection = 0
			}
			fmt.Println("Found ", count, "entries of \"RUBY\"")
		} else if tuple.ID == "SCOOBY" {
			count := 0
			s := fmt.Sprintf("%.7f", NewTuplesObject[index].Latitude)
			fmt.Println("s lat = ", s)
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Latitude = f
			}

			t := fmt.Sprintf("%.7f", NewTuplesObject[index].Longitude)
			fmt.Println("s long = ", t)
			if f, err := strconv.ParseFloat(t, 64); err == nil {
				fmt.Println(f) // 3.14159265
				NewTuplesObject[index].Longitude = f
			}

			// fmt.Println("At ", index, "found Tuple: ", tuple)
			if prev == 0 {
				tuple.Timestamp += 3600
			} else {
				tuple.Timestamp = prev + 3600
			}
			prev = tuple.Timestamp
			fmt.Println("At ", index, "found Tuple: ", tuple)
			count++
			if index+1 >= len(NewTuplesObject) {
				continue
			} else if NewTuplesObject[index+1].ID != "SCOOBY" && index+1 < len(NewTuplesObject) {
				NewTuplesObject[index+1].Timestamp = NewTuplesObject[index].Timestamp + 90000

				NewTuplesObject[index+1].Latitude = tuple.Latitude - 0.01
				fmt.Println("Converting...")
				s := fmt.Sprintf("%.7f", NewTuplesObject[index+1].Latitude)
				fmt.Println("s lat = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Latitude = f
				}

				NewTuplesObject[index+1].Longitude = tuple.Longitude - 0.01
				fmt.Println("Converting...")
				s = fmt.Sprintf("%.7f", NewTuplesObject[index+1].Longitude)
				fmt.Println("s long = ", s)
				if f, err := strconv.ParseFloat(s, 64); err == nil {
					fmt.Println(f) // 3.14159265
					NewTuplesObject[index+1].Longitude = f
				}
				NewTuplesObject[index+1].Infection = 0
			}
			fmt.Println("Found ", count, "entries of \"SCOOBY\"")
		}

	}

	groomedData, err := json.Marshal(NewTuplesObject)
	// fmt.Println("NEW JSON: ", groomedData)
	err = ioutil.WriteFile("./dataset_new.json", groomedData, 0644)
	if err != nil {
		log.Panicln("Failed to write the new dataset at dataset_new.json: ", err)
	}
}

// ConvToFloat is used to structure the dataset as per the new Data definition
func ConvToFloat(tupleObject []Tuple) []TupleNew {
	tupleNewObjects := make([]TupleNew, len(tupleObject))

	for index := range tupleObject {
		tupleNewObjects[index].ID = tupleObject[index].ID

		LatFloatValue, err := strconv.ParseFloat(tupleObject[index].Latitude, 32)
		if err != nil {
			log.Panicln("Unable to convert to float!")
		}
		tupleNewObjects[index].Latitude = LatFloatValue

		LongFloatValue, err := strconv.ParseFloat(tupleObject[index].Longitude, 32)
		if err != nil {
			log.Panicln("Unable to convert to float!")
		}
		tupleNewObjects[index].Longitude = LongFloatValue

		tupleNewObjects[index].Infection = tupleObject[index].Infection
		tupleNewObjects[index].Timestamp = tupleObject[index].Timestamp

	}

	sort.Slice(tupleNewObjects[:], func(i, j int) bool {
		return tupleNewObjects[i].ID < tupleNewObjects[j].ID
	})

	UpdatedtupleNewObjects := fixDuplicateInSlice(tupleNewObjects)

	sort.Slice(UpdatedtupleNewObjects[:], func(i, j int) bool {
		return UpdatedtupleNewObjects[i].Timestamp < UpdatedtupleNewObjects[j].Timestamp
	})

	return UpdatedtupleNewObjects
}

func fixDuplicateInSlice(arr []TupleNew) []TupleNew {
	newArr := make([]TupleNew, len(arr), len(arr))
	users := []string{"BELLA", "BUDDY", "COCO", "LUCY", "MAX", "MILO", "MOLLY", "OSCAR", "RUBY", "SCOOBY"}
	timestamps := []uint{1593373946, 1593377546, 1593381146, 1593384746, 1593388346, 1593391946, 1593395546, 1593399146, 1593402746, 1593406346, 1593409946, 1593413546, 1593417146, 1593420746, 1593424346, 1593427946, 1593431546, 1593435146, 1593438746, 1593442346, 1593445946, 1593449546, 1593453146, 1593456746, 1593460346}
	duplicateMapping := make(map[string][]uint, 0)
	for _, user := range users {
		for _, timestamp := range timestamps {
			count := uint(0)
			for index, v := range arr {
				if v.ID == user && v.Timestamp == timestamp {
					if timestampAlreadyExists(duplicateMapping, user, timestamp) {
						v.Timestamp += (90000 + count)
						newArr[index].ID = v.ID
						newArr[index].Timestamp = v.Timestamp
						newArr[index].Latitude = v.Latitude
						newArr[index].Longitude = v.Longitude
						newArr[index].Infection = v.Infection
						count += 3600
					} else {
						duplicateMapping[user] = append(duplicateMapping[user], v.Timestamp)
						newArr[index].ID = v.ID
						newArr[index].Timestamp = v.Timestamp
						newArr[index].Latitude = v.Latitude
						newArr[index].Longitude = v.Longitude
						newArr[index].Infection = v.Infection
					}
				}

			}
		}
	}
	fmt.Println("duplicateMapping: ", duplicateMapping)
	return newArr
}

func timestampAlreadyExists(duplicateMapping map[string][]uint, user string, timestamp uint) bool {
	for _, v := range duplicateMapping[user] {
		if v == timestamp {
			return true
		}
	}
	return false
}

func getAllIDs(arr []TupleNew) []string {
	var allIDs []string
	for _, tuple := range arr {
		if !checkIDExists(allIDs, tuple.ID) {
			allIDs = append(allIDs, tuple.ID)
		}
	}
	return allIDs
}

func checkIDExists(allIDs []string, ID string) bool {
	for _, id := range allIDs {
		if id == ID {
			return true
		}
	}
	return false
}
