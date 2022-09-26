package sde

type SolarSystem struct {
	Border        bool             `yaml:"border" csv:"border" json:"border"`
	Center        Center           `yaml:"center" json:"center"`
	Corridor      bool             `yaml:"corridor" csv:"corridor" json:"corridor"`
	Fringe        bool             `yaml:"fringe" csv:"fringe" json:"fringe"`
	Hub           bool             `yaml:"hub" csv:"hub" json:"hub"`
	International bool             `yaml:"international" csv:"international" json:"international"`
	Max           Center           `yaml:"max" json:"max"`
	Min           Center           `yaml:"min" json:"min"`
	Planets       map[int]Planet   `yaml:"planets.omitEmpty" json:"planets,omitempty"`
	Radius        float64          `yaml:"radius" csv:"radius" json:"radius"`
	Regional      bool             `yaml:"regional" csv:"regional" json:"regional"`
	Security      float64          `yaml:"security" csv:"security" json:"security"`
	SecurityClass string           `yaml:"securityClass" csv:"securityClass" json:"securityClass"`
	SolarSystemID int              `yaml:"solarSystemID" csv:"solarSystemID" json:"solarSystemId"`
	Star          Star             `yaml:"star" json:"star"`
	Stargates     map[int]StarGate `yaml:"stargates" json:"stargates"`
}

type SolarSystemSummary struct {
	SolarSystemID   int     `csv:"solarSystemID" json:"solarSystemId"`
	SolarSystemName string  `csv:"solarSystemName" json:"solarSystemName"`
	SecurityStatus  float64 `csv:"security" json:"security"`
	Planets         int     `yaml:"planets" json:"planets,omitempty"`
	Moons           int     `yaml:"moons" json:"moons,omitempty"`
	AsteroidBelts   int     `yaml:"asteroidBelts" json:"asteroidBelts,omitempty"`
	Stargates       int     `yaml:"stargates" json:"stargates,omitempty"`
}

type Center struct {
	X float64 `yaml:"x" csv:"x" json:"x"`
	Y float64 `yaml:"y" csv:"y" json:"y"`
	Z float64 `yaml:"z" csv:"z" json:"z"`
}

type Planet struct {
	AsteroidBelts    map[int]AsteroidBelt `yaml:"asteroidBelts" json:"asteroidBelts,omitempty"`
	Moons            map[int]Moon         `yaml:"moons,omitempty" json:"moons,omitempty"`
	CelestialIndex   int                  `yaml:"celestialIndex" csv:"celestialIndex" json:"celestialIndex"`
	PlanetAttributes PlanetAttributes     `yaml:"planetAttributes" json:"planetAttributes"`
	Position         Center               `yaml:"position" json:"position"`
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
	Position   Center     `yaml:"position" json:"position"`
	Statistics Statistics `yaml:"statistics" json:"statistics"`
	TypeID     int        `yaml:"typeID" json:"typeId"`
}

type Moon struct {
	PlanetAttributes PlanetAttributes `yaml:"planetAttributes" json:"planetAttributes"`
	Position         Center           `yaml:"center" json:"center"`
	Radius           float64          `yaml:"radius" json:"radius"`
	Statistics       Statistics       `yaml:"statistics" json:"statistics"`
	TypeID           int              `yaml:"typeID" json:"typeId"`
}

type StarGate struct {
	Destination int    `yaml:"destination" json:"destination,omitempty"`
	Position    Center `yaml:"position" json:"position"`
	TypeID      int    `yaml:"typeID" json:"typeId"`
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
