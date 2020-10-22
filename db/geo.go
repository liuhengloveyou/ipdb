package db

import (
	"fmt"
	"math"
)

// ip库记录
type IpRecord struct {
	Start     uint32  `json:"-"`
	End       uint32  `json:"-"`
	IP        string  `json:"ip"`
	Country   string  `json:"country"`
	Province  string  `json:"province"`
	ISP       string  `json:"isp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var (
	defaultGeo Geo
	geos       map[string]GeoType = make(map[string]GeoType)
)

type Geo interface {
	FindIP(string) (*IpRecord, error)
}

type GeoType func(fn string) (Geo, error)

func init() {
	RegisterGeo("ipip", newIpipDB)
}

func RegisterGeo(name string, newFunc GeoType) {
	if newFunc == nil {
		panic("Register GEO nil.")
	}

	if _, ok := geos[name]; ok {
		panic("Register GEO duplicate for " + name)
	}

	geos[name] = newFunc
}

func NewGeo(t, fn string) (geo Geo, err error) {
	if newFunc, ok := geos[t]; ok {
		return newFunc(fn)
	}

	return nil, fmt.Errorf("No GEO types " + t)
}

func InitGEO(fmt, db string) error {
	var err error
	if defaultGeo, err = NewGeo(fmt, db); err != nil {
		return err
	}

	return nil
}

//IpRecord
func (p *IpRecord) GetLineArea() (line, area string) {
	return p.ISP, p.Province
}

///////////////////////////////////////////////////////////////////////////////
//// public interface
///////////////////////////////////////////////////////////////////////////////
func FindIP(ip string) (*IpRecord, error) {
	return defaultGeo.FindIP(ip)
}

func LatitudeLongitudeDistance(lat1, lon1, lat2, lon2 float64) (distance float64) {
	const RADIUS = 6378137 //赤道半径(单位m)

	radLat1 := lat1 * math.Pi / 180.0
	radLon1 := lon1 * math.Pi / 180.0
	radLat2 := lat2 * math.Pi / 180.0
	radLon2 := lon2 * math.Pi / 180.0

	dist := math.Acos(math.Sin(radLat1)*math.Sin(radLat2) + math.Cos(radLat1)*math.Cos(radLat2)*math.Cos(radLon2-radLon1))

	return dist * RADIUS
}
