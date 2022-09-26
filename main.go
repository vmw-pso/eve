package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type SolarSystem struct {
	Border        bool             `yaml:"border" csv:"border" json:"border"`
	Center        []float64        `yaml:"center" json:"center"`
	Corridor      bool             `yaml:"corridor" csv:"corridor" json:"corridor"`
	Fringe        bool             `yaml:"fringe" csv:"fringe" json:"fringe"`
	Hub           bool             `yaml:"hub" csv:"hub" json:"hub"`
	International bool             `yaml:"international" csv:"international" json:"international"`
	Max           []float64        `yaml:"max" json:"max"`
	Min           []float64        `yaml:"min" json:"min"`
	Planets       map[int]Planet   `yaml:"planets,omitempty" json:"planets,omitempty"`
	Radius        float64          `yaml:"radius" csv:"radius" json:"radius"`
	Regional      bool             `yaml:"regional" csv:"regional" json:"regional"`
	Security      float64          `yaml:"security" csv:"security" json:"security"`
	SecurityClass string           `yaml:"securityClass" csv:"securityClass" json:"securityClass"`
	SolarSystemID int              `yaml:"solarSystemID" csv:"solarSystemID" json:"solarSystemId"`
	Star          Star             `yaml:"star" json:"star"`
	Stargates     map[int]StarGate `yaml:"stargates" json:"stargates"`
}

type SolarSystemData struct {
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

func main() {
	// archive := "./assets/202209_sde.zip" // file not exported to github as can be downloaded from CCP. Remove this had coding during development
	// sde := sde.SDE{Filename: archive}

	// sde.WriteFileStructure("202209_sde_content.txt")

	// systems := joveStaticData("./assets/202209_sde.zip")

	// var solarSystemDump []SolarSystemData

	// for name, system := range systems {
	// 	systemData := SolarSystemData{
	// 		SolarSystemID:   system.SolarSystemID,
	// 		SolarSystemName: name,
	// 		SecurityStatus:  system.Security,
	// 		Planets:         len(system.Planets),
	// 		Moons:           systemMoonCount(system.Planets),
	// 		AsteroidBelts:   systemAsteroidBeltCount(system.Planets),
	// 		Stargates:       systemStargateCount(system.Stargates),
	// 	}
	// 	solarSystemDump = append(solarSystemDump, systemData)
	// }

	// jsonData, _ := json.Marshal(solarSystemDump)
	// os.WriteFile("jove.json", jsonData, os.ModePerm)

	var kspace []SolarSystemData
	file, _ := os.Open("solarsystems.json")
	data, _ := io.ReadAll(file)
	json.Unmarshal(data, &kspace)

	highsec := highsecSystems(kspace)
	lowsec := lowsecSystems(kspace)
	nullsec := nullsecSystems(kspace)

	var wormholes []SolarSystemData
	file, _ = os.Open("wormholes.json")
	data, _ = io.ReadAll(file)
	json.Unmarshal(data, &wormholes)

	var jove []SolarSystemData
	file, _ = os.Open("jove.json")
	data, _ = io.ReadAll(file)
	json.Unmarshal(data, &jove)

	fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Highsec", len(highsec), totalPlanets(highsec), totalMoons(highsec), totalAsteroidBelts(highsec), totalStargates(highsec))
	fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Lowsec", len(lowsec), totalPlanets(lowsec), totalMoons(lowsec), totalAsteroidBelts(lowsec), totalStargates(lowsec))
	fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Nullsec", len(nullsec), totalPlanets(nullsec), totalMoons(nullsec), totalAsteroidBelts(nullsec), totalStargates(nullsec))
	fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Jove", len(jove), totalPlanets(jove), totalMoons(jove), totalAsteroidBelts(jove), totalStargates(jove))
	fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Wormholes", len(wormholes), totalPlanets(wormholes), totalMoons(wormholes), totalAsteroidBelts(wormholes), 0)
}

func securityTypeCount(solarSystems []SolarSystemData) (int, int, int) {
	highsec := 0
	lowsec := 0
	nullsec := 0
	for _, system := range solarSystems {
		if system.SecurityStatus >= 0.45 {
			highsec++
		} else if system.SecurityStatus > 0.0 {
			lowsec++
		} else {
			nullsec++
		}
	}
	return highsec, lowsec, nullsec
}

func highsecSystems(cluster []SolarSystemData) []SolarSystemData {
	var highsec []SolarSystemData
	for _, system := range cluster {
		if system.SecurityStatus >= 0.45 {
			highsec = append(highsec, system)
		}
	}
	return highsec
}

func lowsecSystems(cluster []SolarSystemData) []SolarSystemData {
	var lowsec []SolarSystemData
	for _, system := range cluster {
		if system.SecurityStatus > 0.0 && system.SecurityStatus < 0.45 {
			lowsec = append(lowsec, system)
		}
	}
	return lowsec
}

func nullsecSystems(cluster []SolarSystemData) []SolarSystemData {
	var nullsec []SolarSystemData
	for _, system := range cluster {
		if system.SecurityStatus <= 0.0 && system.SecurityStatus < 0.45 {
			nullsec = append(nullsec, system)
		}
	}
	return nullsec
}

func systemStargateCount(stargates map[int]StarGate) int {
	return len(stargates)
}

func systemPlanetCount(planets map[int]Planet) int {
	return len(planets)
}

func systemMoonCount(planets map[int]Planet) int {
	moons := 0
	for _, planet := range planets {
		moons += len(planet.Moons)
	}
	return moons
}

func totalPlanets(systems []SolarSystemData) int {
	planets := 0
	for _, system := range systems {
		planets += system.Planets
	}
	return planets
}

func totalMoons(systems []SolarSystemData) int {
	moons := 0
	for _, system := range systems {
		moons += system.Moons
	}
	return moons
}

func totalAsteroidBelts(systems []SolarSystemData) int {
	belts := 0
	for _, system := range systems {
		belts += system.AsteroidBelts
	}
	return belts
}

func totalStargates(systems []SolarSystemData) int {
	gates := 0
	for _, system := range systems {
		gates += system.Stargates
	}
	return gates
}

func systemAsteroidBeltCount(planets map[int]Planet) int {
	asteroidBelts := 0
	for _, planet := range planets {
		asteroidBelts += len(planet.AsteroidBelts)
	}
	return asteroidBelts
}

func readAll(file *zip.File) ([]byte, error) {
	fc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fc.Close()

	content, err := io.ReadAll(fc)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func Add(regions []string, region string) []string {
	for _, name := range regions {
		if name == region {
			return regions
		}
	}
	regions = append(regions, region)
	return regions
}

func readUniverse(zf *zip.ReadCloser) {
	space := []string{}
	regions := make(map[string][]string)
	for _, file := range zf.File {
		name := file.Name
		if strings.Contains(name, "sde/fsd/universe") {
			spaceSubstring := name[17:]
			index := strings.Index(spaceSubstring, "/")
			space = Add(space, spaceSubstring[:index])
		}
		for _, s := range space {
			if strings.Contains(name, s) {
				index := 18 + len(s)
				region := name[index:]
				indexSlash := strings.Index(region, "/")
				regions[s] = Add(regions[s], region[:indexSlash])
			}
		}
	}
	// fmt.Println(space)
	// fmt.Println(regions)

	for k, v := range regions {
		fmt.Println(k, len(v))
	}
	fmt.Println()
	fmt.Println(regions["eve"])
}

func dumpFilenames(archive string, saveFilename string) error {
	zf, err := zip.OpenReader(archive)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	defer zf.Close()

	file, err := os.OpenFile(saveFilename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, document := range zf.File {
		writer.WriteString(document.Name + "\n")
	}
	return nil
}

func getName(filename string) string {
	lastIndex := strings.LastIndex(filename, "/")
	directories := filename[0:lastIndex]
	lastIndex = strings.LastIndex(directories, "/")
	return directories[lastIndex+1:]
}

func isJove(filename string) bool {
	jove := []string{"A821-A", "J7HZ-F", "UUA-F4"}
	for _, region := range jove {
		if strings.Contains(filename, region) {
			return true
		}
	}
	return false
}

func kSpaceStaticData(zipFilename string, kspace bool) map[string]SolarSystem {
	solarSystems := make(map[string]SolarSystem)
	zf, err := zip.OpenReader(zipFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	defer zf.Close()

	for _, file := range zf.File {
		name := file.Name
		solarSystemName := getName(name)
		if strings.Contains(name, "sde/fsd/universe/eve") {
			if kspace && isJove(name) {
				continue
			}
			if strings.Contains(name, "solarsystem.staticdata") {
				data, err := file.Open()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: %s\n", name, err.Error())
					os.Exit(1)
				}

				content, err := io.ReadAll(data)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s read error: %s\n", name, err.Error())
					os.Exit(1)
				}

				var solarSystem SolarSystem
				if err = yaml.Unmarshal(content, &solarSystem); err != nil {
					if err != nil {
						fmt.Fprintf(os.Stderr, "unmarshal error in %s: %s\n", name, err.Error())
						os.Exit(1)
					}
				}
				solarSystems[solarSystemName] = solarSystem
			}
		}
	}
	return solarSystems
}

func wormholeStaticData(zipFilename string) map[string]SolarSystem {
	solarSystems := make(map[string]SolarSystem)
	zf, err := zip.OpenReader(zipFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	defer zf.Close()

	for _, file := range zf.File {
		name := file.Name
		solarSystemName := getName(name)
		if strings.Contains(name, "sde/fsd/universe/wormhole") {
			if strings.Contains(name, "solarsystem.staticdata") {
				data, err := file.Open()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: %s\n", name, err.Error())
					os.Exit(1)
				}

				content, err := io.ReadAll(data)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s read error: %s\n", name, err.Error())
					os.Exit(1)
				}

				var solarSystem SolarSystem
				if err = yaml.Unmarshal(content, &solarSystem); err != nil {
					if err != nil {
						fmt.Fprintf(os.Stderr, "unmarshal error in %s: %s\n", name, err.Error())
						os.Exit(1)
					}
				}
				solarSystems[solarSystemName] = solarSystem
			}
		}
	}
	return solarSystems
}

func joveStaticData(zipFilename string) map[string]SolarSystem {
	solarSystems := make(map[string]SolarSystem)
	zf, err := zip.OpenReader(zipFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	defer zf.Close()

	for _, file := range zf.File {
		name := file.Name
		solarSystemName := getName(name)
		if strings.Contains(name, "sde/fsd/universe/eve") {
			if !isJove(name) {
				continue
			}
			if strings.Contains(name, "solarsystem.staticdata") {
				data, err := file.Open()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: %s\n", name, err.Error())
					os.Exit(1)
				}

				content, err := io.ReadAll(data)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s read error: %s\n", name, err.Error())
					os.Exit(1)
				}

				var solarSystem SolarSystem
				if err = yaml.Unmarshal(content, &solarSystem); err != nil {
					if err != nil {
						fmt.Fprintf(os.Stderr, "unmarshal error in %s: %s\n", name, err.Error())
						os.Exit(1)
					}
				}
				solarSystems[solarSystemName] = solarSystem
			}
		}
	}
	return solarSystems
}
