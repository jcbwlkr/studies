package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `{"_id":"51e873b3b199b3cb9318d981","station_id":"vqsp4","condition":{"air_temperature_celsius":30.5,"air_temperature_fahrenheit":86.9,"average_wave_period_sec":0,"date":"2016-07-22T18:54:00Z","dewpoint_temperature_celsius":0,"dewpoint_temperature_fahrenheit":32,"dominant_wave_period_sec":0,"dominant_wave_prediction_degnorth":0,"guest_wind_speed_kilohour":18.359999656677246,"guest_wind_speed_metersec":5.099999904632568,"guest_wind_speed_milehour":11.408393786668778,"gust_wind_speed_knots":9.913583814620972,"latitude":18.152999877929688,"longitude":-65.44400024414062,"pressure_tendency_hpa":0,"sea_level_pressure_hpa":0,"sea_surface_temperature_fahrenheit":32,"sea_surfae_temperature_celsius":0,"station_id":"vqsp4","station_visibility_nmi":0,"tide_water_level_feet":0,"tide_water_level_meter":0,"update_date":"0001-01-01T00:00:00Z","wave_height_feet":0,"wave_height_meter":0,"wind_direction_degnorth":80,"wind_speed_kilohour":9.359999656677246,"wind_speed_knots":5.053983814620972,"wind_speed_metersec":2.5999999046325684,"wind_speed_milehour":5.816043786668778},"name":"Station VQSP4 - 9752619","region":"Caribbean","location_desc":"Isabel Segunda, Vieques, PR","location":{"type":"Point","coordinates":[-65.444,18.153]},"image":"http://www.ndbc.noaa.gov/images/stations/vqsp4.jpg","image_small":"http://www.ndbc.noaa.gov/images/stations/vqsp4_mini.jpg","owned":"Puerto Rico Seismic Network","station_group":"station","station_type":"Water Level Observation Network","surface":"coastline","geo_near_distance":0}`

func main() {
	var d map[string]interface{}
	if err := json.Unmarshal([]byte(data), &d); err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	x := make(map[interface{}]interface{})

	x["foo"] = "F"
	x[123] = 456

	fmt.Println(x)
}
