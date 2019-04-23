package sphere

import (
	"math"
	"reflect"
	"testing"

	"github.com/twpayne/go-kml"
)

func TestCircle(t *testing.T) {
	for _, tc := range []struct {
		center kml.Coordinate
		radius float64
		maxErr float64
		want   []kml.Coordinate
	}{
		{
			center: kml.Coordinate{Lon: 0, Lat: 0, Alt: 100},
			radius: 1000,
			maxErr: 1,
			want: []kml.Coordinate{
				{Lon: 0, Lat: 0.008983152841195214, Alt: 100},
				{Lon: 0.0011258876022698656, Lat: 0.008912317997331125, Alt: 100},
				{Lon: 0.002234019283634706, Lat: 0.008700930573621838, Alt: 100},
				{Lon: 0.0033069191447876595, Lat: 0.008352324276246322, Alt: 100},
				{Lon: 0.004327666913493653, Lat: 0.007871996835109306, Alt: 100},
				{Lon: 0.005280164787461206, Lat: 0.007267523301307612, Alt: 100},
				{Lon: 0.006149391306330964, Lat: 0.0065484365839908, Alt: 100},
				{Lon: 0.006921638249063812, Lat: 0.005726077110629261, Alt: 100},
				{Lon: 0.007584726820717381, Lat: 0.004813413981645677, Alt: 100},
				{Lon: 0.008128199719224507, Lat: 0.003824840439915492, Alt: 100},
				{Lon: 0.00854348605317887, Lat: 0.0027759468807100067, Alt: 100},
				{Lon: 0.008824036509791956, Lat: 0.001683274981853487, Alt: 100},
				{Lon: 0.008965426641358826, Lat: 0.0005640568316080706, Alt: 100},
				{Lon: 0.008965426641358826, Lat: -0.0005640568316080696, Alt: 100},
				{Lon: 0.008824036509791956, Lat: -0.0016832749818534878, Alt: 100},
				{Lon: 0.00854348605317887, Lat: -0.002775946880710006, Alt: 100},
				{Lon: 0.008128199719224507, Lat: -0.003824840439915493, Alt: 100},
				{Lon: 0.007584726820717381, Lat: -0.004813413981645679, Alt: 100},
				{Lon: 0.006921638249063812, Lat: -0.00572607711062926, Alt: 100},
				{Lon: 0.006149391306330961, Lat: -0.006548436583990801, Alt: 100},
				{Lon: 0.005280164787461206, Lat: -0.007267523301307612, Alt: 100},
				{Lon: 0.004327666913493655, Lat: -0.007871996835109304, Alt: 100},
				{Lon: 0.0033069191447876573, Lat: -0.008352324276246324, Alt: 100},
				{Lon: 0.002234019283634706, Lat: -0.008700930573621838, Alt: 100},
				{Lon: 0.001125887602269864, Lat: -0.008912317997331125, Alt: 100},
				{Lon: 1.1001189463363886e-18, Lat: -0.008983152841195214, Alt: 100},
				{Lon: -0.0011258876022698621, Lat: -0.008912317997331125, Alt: 100},
				{Lon: -0.002234019283634708, Lat: -0.008700930573621838, Alt: 100},
				{Lon: -0.003306919144787659, Lat: -0.008352324276246324, Alt: 100},
				{Lon: -0.004327666913493654, Lat: -0.007871996835109304, Alt: 100},
				{Lon: -0.005280164787461205, Lat: -0.007267523301307612, Alt: 100},
				{Lon: -0.00614939130633096, Lat: -0.006548436583990801, Alt: 100},
				{Lon: -0.006921638249063812, Lat: -0.00572607711062926, Alt: 100},
				{Lon: -0.007584726820717378, Lat: -0.004813413981645681, Alt: 100},
				{Lon: -0.00812819971922451, Lat: -0.0038248404399154876, Alt: 100},
				{Lon: -0.00854348605317887, Lat: -0.002775946880710008, Alt: 100},
				{Lon: -0.008824036509791956, Lat: -0.001683274981853488, Alt: 100},
				{Lon: -0.008965426641358826, Lat: -0.0005640568316080777, Alt: 100},
				{Lon: -0.008965426641358826, Lat: 0.0005640568316080745, Alt: 100},
				{Lon: -0.008824036509791954, Lat: 0.0016832749818534924, Alt: 100},
				{Lon: -0.00854348605317887, Lat: 0.002775946880710005, Alt: 100},
				{Lon: -0.008128199719224507, Lat: 0.003824840439915492, Alt: 100},
				{Lon: -0.0075847268207173855, Lat: 0.004813413981645672, Alt: 100},
				{Lon: -0.006921638249063809, Lat: 0.005726077110629263, Alt: 100},
				{Lon: -0.0061493913063309594, Lat: 0.006548436583990802, Alt: 100},
				{Lon: -0.005280164787461206, Lat: 0.007267523301307612, Alt: 100},
				{Lon: -0.004327666913493653, Lat: 0.007871996835109306, Alt: 100},
				{Lon: -0.0033069191447876655, Lat: 0.008352324276246322, Alt: 100},
				{Lon: -0.002234019283634703, Lat: 0.00870093057362184, Alt: 100},
				{Lon: -0.0011258876022698613, Lat: 0.008912317997331125, Alt: 100},
				{Lon: 0, Lat: 0.008983152841195214, Alt: 100},
			},
		},
		{
			center: kml.Coordinate{Lon: 13.631333, Lat: 46.438500},
			radius: 50,
			maxErr: 1,
			want: []kml.Coordinate{
				{Lon: 13.631333, Lat: 46.43894915764205},
				{Lon: 13.631658888465811, Lat: 46.43888898146551},
				{Lon: 13.631897453677293, Lat: 46.43872457743259},
				{Lon: 13.631984772278717, Lat: 46.438499998148764},
				{Lon: 13.631897449024441, Lat: 46.43827541979054},
				{Lon: 13.63165888381296, Lat: 46.43811101760886},
				{Lon: 13.631333, Lat: 46.43805084235793},
				{Lon: 13.63100711618704, Lat: 46.43811101760886},
				{Lon: 13.630768550975558, Lat: 46.43827541979054},
				{Lon: 13.630681227721285, Lat: 46.438499998148764},
				{Lon: 13.630768546322708, Lat: 46.43872457743259},
				{Lon: 13.63100711153419, Lat: 46.43888898146551},
				{Lon: 13.631333, Lat: 46.43894915764205},
			},
		},
	} {
		if got := WGS84.Circle(tc.center, tc.radius, tc.maxErr); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("WGS84.Circle(%v, %v, %v) == %+v, want %+v", tc.center, tc.radius, tc.maxErr, got, tc.want)
		}
		for _, c := range tc.want {
			distance := WGS84.HaversineDistance(tc.center, c)
			delta := math.Abs(distance - tc.radius)
			threshold := 1e-5
			if math.Abs(delta) > threshold {
				t.Errorf("math.Abs(WGS84.HaversineDistance(%v, %v)-%f) == %f, want <=%f", tc.center, c, tc.radius, delta, threshold)
			}
		}
	}
}

func TestSphereHaversineDistance(t *testing.T) {
	for _, tc := range []struct {
		sphere    T
		c1        kml.Coordinate
		c2        kml.Coordinate
		want      float64
		threshold float64
	}{
		{
			sphere:    FAI,
			c1:        kml.Coordinate{Lon: -108.6180554, Lat: 35.4325002},
			c2:        kml.Coordinate{Lon: -108.61, Lat: 35.43},
			want:      781,
			threshold: 1e-3,
		},
	} {
		distance := tc.sphere.HaversineDistance(tc.c1, tc.c2)
		delta := math.Abs(distance - tc.want)
		if delta > tc.threshold {
			t.Errorf("math.Abs(tc.sphere.HaversineDistance(%v, %v)-%f) == %f, want <=%f", tc.c1, tc.c2, tc.want, delta, tc.threshold)
		}
	}
}
