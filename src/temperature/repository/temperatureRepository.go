package repository

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"

	"gopkg.in/mgo.v2/bson"

	"config/mongo"
	temperature "temperature"
)

type temperatureRepository struct {
}
type TemperatureEntity struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Date        time.Time
	Measurement float64
	Name        string
}

func (rep temperatureRepository) SaveTemperature(m temperature.Temperature) {
	sessionProv := mongo.MongoSessionProvider
	session := sessionProv.GetSession()
	defer session.Close()
	databaseName := viper.GetString("temperatureDatabaseInfo.database")         //measurements1
	collectionName := viper.GetString("temperatureDatabaseInfo.collectionName") //temperature
	c := session.DB(databaseName).C(collectionName)
	c.Insert(&TemperatureEntity{Date: m.Date, Measurement: m.Measurement, Name: m.Name})
}

func (rep temperatureRepository) GetTemperature(start time.Time, end time.Time) []temperature.Temperature {
	readQuery := viper.GetString("query.mongo.temperature.dateStartDateEnd")
	ssdate := fmt.Sprintf("%v", start)
	sedate := fmt.Sprintf("%v", end)
	myMap := map[string]string{"inicio": ssdate, "fin": sedate}
	replacedQuery := prepareQuery(readQuery, myMap)
	query := make(bson.M)
	bson.Unmarshal([]byte(replacedQuery), &query)
	// query2 := bson.M{"date": bson.M{"$gt": start, "$lt": end}}
	result := rep.GetTemperatureByQuery(query)
	return result
}

func NewTemperatureRepository() TemperatureRepository {
	return &temperatureRepository{}
}

func (rep temperatureRepository) GetTemperatureByQuery(query interface{}) []temperature.Temperature {
	sessionProv := mongo.MongoSessionProvider
	session := sessionProv.GetSession()

	defer session.Close()
	databaseName := viper.GetString("temperatureDatabaseInfo.database")         //measurements1
	collectionName := viper.GetString("temperatureDatabaseInfo.collectionName") //temperature

	c := session.DB(databaseName).C(collectionName) // var us userEntity
	result := []temperature.Temperature{}
	err := c.Find(query).All(&result)
	if err != nil {
		log.Printf("error en la consulta %v", err)
	}

	return result
}
func prepareQuery(query string, params map[string]string) string {
	for key := range params {
		value := params[key]
		query = strings.Replace(query, "$"+key, value, -1)
	}
	return query
}
