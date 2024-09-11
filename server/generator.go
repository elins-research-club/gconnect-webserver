package server

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomData() string {
	const TEAM_ID = "1009"
	now := time.Now()
	CLOCK := fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())

	YAW := fmt.Sprintf("%.2f", rand.Float64()*360-180)
	PITCH := fmt.Sprintf("%.2f", rand.Float64()*360-180)
	ROLL := fmt.Sprintf("%.2f", rand.Float64()*360-180)

	lat := -7.773684
	long := 110.381798
	LATITUDE := fmt.Sprintf("%.6f", lat+rand.Float64()*0.0002-0.0001)
	LONGITUDE := fmt.Sprintf("%.6f", long+rand.Float64()*0.0002-0.0001)

	VOLTAGE := fmt.Sprintf("%.2f", rand.Float64()*12)
	PRESSURE := fmt.Sprintf("%.2f", rand.Float64()*100)
	ALTITUDE := fmt.Sprintf("%.2f", rand.Float64()*700)

	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s;",
		TEAM_ID, CLOCK, YAW, PITCH, ROLL, LATITUDE, LONGITUDE, VOLTAGE, PRESSURE, ALTITUDE)
}
