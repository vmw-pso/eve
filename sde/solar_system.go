package sde

import (
	"encoding/json"
	"io"
	"os"
)

type SolarSystem struct {
	Border              bool             `yaml:"border" csv:"border" json:"border"`
	Center              []float64        `yaml:"center" json:"center"`
	Corridor            bool             `yaml:"corridor" csv:"corridor" json:"corridor"`
	Fringe              bool             `yaml:"fringe" csv:"fringe" json:"fringe"`
	Hub                 bool             `yaml:"hub" csv:"hub" json:"hub"`
	International       bool             `yaml:"international" csv:"international" json:"international"`
	Max                 []float64        `yaml:"max" json:"max"`
	Min                 []float64        `yaml:"min" json:"min"`
	Planets             map[int]Planet   `yaml:"planets.omitEmpty" json:"planets,omitempty"`
	Radius              float64          `yaml:"radius" csv:"radius" json:"radius"`
	Regional            bool             `yaml:"regional" csv:"regional" json:"regional"`
	Security            float64          `yaml:"security" csv:"security" json:"security"`
	SecurityClass       string           `yaml:"securityClass" csv:"securityClass" json:"securityClass"`
	SolarSystemID       int              `yaml:"solarSystemID" csv:"solarSystemID" json:"solarSystemId"`
	SolarSystemName     string           `csv:"solarSystemName" json:"solarSystemName"`
	SolarSystemTypeName string           `csv:"solarSystemTypeName,omitempty" json:"solarSystemTypeName,omitempty"`
	RegionName          string           `csv:"regionName,omitempty" json:"regionName,omitempty"`
	ConstellationName   string           `csv:"constellationName,omitEmpty" json:"constellationName,omitempty"`
	Star                Star             `yaml:"star" json:"star"`
	Stargates           map[int]StarGate `yaml:"stargates,omitempty" json:"stargates,omitempty"`
}

type SolarSystemSummary struct {
	SolarSystemID     int     `csv:"solarSystemID" json:"solarSystemId"`
	SolarSystemName   string  `csv:"solarSystemName" json:"solarSystemName"`
	SpaceTypeName     string  `csv:"spaceTypeName,omitempty" json:"spaceTypeName,omitempty"`
	RegionName        string  `csv:"regionName,omitempty" json:"regionName,omitempty"`
	ConstellationName string  `csv:"constellationName,omitEmpty" json:"constellationName,omitempty"`
	Security          float64 `csv:"security" json:"security"`
	Planets           int     `yaml:"planets" json:"planets,omitempty"`
	Moons             int     `yaml:"moons" json:"moons,omitempty"`
	AsteroidBelts     int     `yaml:"asteroidBelts" json:"asteroidBelts,omitempty"`
	Stargates         int     `yaml:"stargates" json:"stargates,omitempty"`
}

type Planet struct {
	AsteroidBelts    map[int]AsteroidBelt `yaml:"asteroidBelts" json:"asteroidBelts,omitempty"`
	Moons            map[int]Moon         `yaml:"moons,omitempty" json:"moons,omitempty"`
	CelestialIndex   int                  `yaml:"celestialIndex" csv:"celestialIndex" json:"celestialIndex"`
	PlanetAttributes PlanetAttributes     `yaml:"planetAttributes" json:"planetAttributes"`
	Position         []float64            `yaml:"position" json:"position"`
	Radius           float64              `yaml:"radius" json:"radius"`
	Statistics       Statistics           `yaml:"statistics" json:"statistics"`
	TypeID           int                  `yaml:"typeID" json:"typeId"`
}

type PlanetAttributes struct {
	HeightMap1   int  `yaml:"heightMap1" json:"heightMap1"`
	HeightMap2   int  `yaml:"heightMap2" json:"heightMap2"`
	Population   bool `yaml:"population" json:"population"`
	ShaderPreset int  `yaml:"shaderPreset" json:"shaderPreset"`
}

type Statistics struct {
	Density        float64 `yaml:"statistics" json:"statistics"`
	Eccentricity   float64 `yaml:"eccentricity" json:"eccentricity"`
	EscapeVelocity float64 `yaml:"escapeVelocity" json:"escapeVelocity"`
	Fragmented     bool    `yaml:"fragmented" json:"fragmented"`
	Life           float64 `yaml:"life" json:"life"`
	Locked         bool    `yaml:"locked" json:"locked"`
	MassDust       float64 `yaml:"massDust" json:"massDust"`
	MassGas        float64 `yaml:"massGas" json:"massGas"`
	OrbitPeriod    float64 `yaml:"orbitPeriod" json:"orbitPeriod"`
	OrbitRadius    float64 `yaml:"orbitRadius" json:"orbitRadious"`
	Pressure       float64 `yaml:"pressure" json:"pressure"`
	Radius         float64 `yaml:"radius" json:"radius"`
	RotationRate   float64 `yaml:"rotationRate" json:"rotationRate"`
	SpectralClass  string  `yaml:"spectralClass" json:"spectralClass"`
	SurfaceGravity float64 `yaml:"surfaceGravity" json:"surfaceGravity"`
	Temperature    float64 `yaml:"temperature" json:"temperature"`
}

type AsteroidBelt struct {
	Position   []float64  `yaml:"position" json:"position"`
	Statistics Statistics `yaml:"statistics" json:"statistics"`
	TypeID     int        `yaml:"typeID" json:"typeId"`
}

type Moon struct {
	PlanetAttributes PlanetAttributes `yaml:"planetAttributes" json:"planetAttributes"`
	Position         []float64        `yaml:"center" json:"center"`
	Radius           float64          `yaml:"radius" json:"radius"`
	Statistics       Statistics       `yaml:"statistics" json:"statistics"`
	TypeID           int              `yaml:"typeID" json:"typeId"`
}

type StarGate struct {
	Destination int       `yaml:"destination" json:"destination,omitempty"`
	Position    []float64 `yaml:"position" json:"position"`
	TypeID      int       `yaml:"typeID" json:"typeId"`
}

type Star struct {
	StarID int     `yaml:"id" json:"starId"`
	Radius float64 `yaml:"radius" json:"radius"`
}

type StartStatistics struct {
	Age           float64 `yaml:"age" json:"age"`
	Life          float64 `yaml:"life" json:"life"`
	Locked        bool    `yaml:"locked" json:"locaked"`
	Luminosity    float64 `yaml:"luminosity" json:"luminosity"`
	Radius        float64 `yaml:"radius" json:"radius"`
	SpectralClass string  `yaml:"spectralClass" json:"spectralClass"`
	Temperature   float64 `yaml:"temperature" json:"temperature"`
}

type cluster struct {
	solarSystems []SolarSystem
}

func NewFromJSON(filename string) (*cluster, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var c cluster

	err = json.Unmarshal(data, &c.solarSystems)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *cluster) SystemsWithSecurity(min, max float64) cluster {
	var systems cluster
	for _, system := range c.solarSystems {
		if system.Security > min && system.Security <= max {
			systems.solarSystems = append(systems.solarSystems, system)
		}
	}
	return systems
}

func (c *cluster) HighsecSystems() cluster {
	return c.SystemsWithSecurity(0.45, 1.0)
}

func (c *cluster) LowsecSystems() cluster {
	return c.SystemsWithSecurity(0.0, 0.45)
}

func (c *cluster) NullsecSystems() cluster {
	kspace := c.KSpaceSystems()
	var systems cluster
	for _, system := range kspace.solarSystems {
		if system.Security <= 0.0 && !isJove(system) {
			systems.solarSystems = append(systems.solarSystems, system)
		}
	}
	return systems
}

func (c *cluster) SystemCount() int {
	return len(c.solarSystems)
}

func (c *cluster) RegionCount() int {
	regions := make(map[string]int)
	for _, system := range c.solarSystems {
		regions[system.RegionName] = 1
	}
	return len(regions)
}

func (c *cluster) KSpaceSystems() cluster {
	var systems []SolarSystem
	for _, system := range c.solarSystems {
		if system.SolarSystemTypeName == "eve" {
			systems = append(systems, system)
		}
	}
	return cluster{solarSystems: systems}
}

func (c *cluster) JSpaceSystems() cluster {
	var systems []SolarSystem
	for _, system := range c.solarSystems {
		if system.SolarSystemTypeName == "wormhole" {
			systems = append(systems, system)
		}
	}
	return cluster{solarSystems: systems}
}

func isJove(solarSystem SolarSystem) bool {
	jove := []string{"A821-A", "J7HZ-F", "UUA-F4"}
	for _, region := range jove {
		if solarSystem.RegionName == region {
			return true
		}
	}
	return false
}
