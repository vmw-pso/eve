package main

import (
	"fmt"
	"os"

	"github.com/vmw-pso/eve/sde"
)

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
	// 		Planets:         len(system.Planets)
	// 		Moons:           systemMoonCount(system.Planets),
	// 		AsteroidBelts:   systemAsteroidBeltCount(system.Planets),
	// 		Stargates:       systemStargateCount(system.Stargates),
	// 	}
	// 	solarSystemDump = append(solarSystemDump, systemData)
	// }

	// jsonData, _ := json.Marshal(solarSystemDump)
	// os.WriteFile("jove.json", jsonData, os.ModePerm)

	// var kspace []SolarSystemData
	// file, _ := os.Open("solarsystems.json")
	// data, _ := io.ReadAll(file)
	// json.Unmarshal(data, &kspace)

	// highsec := highsecSystems(kspace)
	// lowsec := lowsecSystems(kspace)
	// nullsec := nullsecSystems(kspace)

	// var wormholes []SolarSystemData
	// file, _ = os.Open("wormholes.json")
	// data, _ = io.ReadAll(file)
	// json.Unmarshal(data, &wormholes)

	// var jove []SolarSystemData
	// file, _ = os.Open("jove.json")
	// data, _ = io.ReadAll(file)
	// json.Unmarshal(data, &jove)

	// fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Highsec", len(highsec), totalPlanets(highsec), totalMoons(highsec), totalAsteroidBelts(highsec), totalStargates(highsec))
	// fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Lowsec", len(lowsec), totalPlanets(lowsec), totalMoons(lowsec), totalAsteroidBelts(lowsec), totalStargates(lowsec))
	// fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Nullsec", len(nullsec), totalPlanets(nullsec), totalMoons(nullsec), totalAsteroidBelts(nullsec), totalStargates(nullsec))
	// fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Jove", len(jove), totalPlanets(jove), totalMoons(jove), totalAsteroidBelts(jove), totalStargates(jove))
	// fmt.Printf("%-10s%4d, Planets%8d, Moons%10d, Asteroid Belts%10d, Stargates%10d\n", "Wormholes", len(wormholes), totalPlanets(wormholes), totalMoons(wormholes), totalAsteroidBelts(wormholes), 0)

	solarsystems, err := sde.NewFromJSON("solarsystems.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	kspace := solarsystems.KSpaceSystems()

	highsec := solarsystems.HighsecSystems()
	lowsec := solarsystems.LowsecSystems()
	nullsec := solarsystems.NullsecSystems()
	wormholes := solarsystems.JSpaceSystems()

	fmt.Println("K Space", kspace.SystemCount(), ", Regions", kspace.RegionCount())
	fmt.Println("Highsec", highsec.SystemCount(), ", Regions", highsec.RegionCount())
	fmt.Println("Lowsec", lowsec.SystemCount(), ", Regions", lowsec.RegionCount())
	fmt.Println("Nullsec", nullsec.SystemCount(), ", Regions", nullsec.RegionCount())
	fmt.Println("Wormholes", wormholes.SystemCount(), ", Regions", wormholes.RegionCount())
}

// func systemStargateCount(stargates map[int]StarGate) int {
// 	return len(stargates)
// }

// func systemPlanetCount(planets map[int]Planet) int {
// 	return len(planets)
// }

// func systemMoonCount(planets map[int]Planet) int {
// 	moons := 0
// 	for _, planet := range planets {
// 		moons += len(planet.Moons)
// 	}
// 	return moons
// }

// func totalPlanets(systems []SolarSystemData) int {
// 	planets := 0
// 	for _, system := range systems {
// 		planets += system.Planets
// 	}
// 	return planets
// }

// func totalMoons(systems []SolarSystemData) int {
// 	moons := 0
// 	for _, system := range systems {
// 		moons += system.Moons
// 	}
// 	return moons
// }

// func totalAsteroidBelts(systems []SolarSystemData) int {
// 	belts := 0
// 	for _, system := range systems {
// 		belts += system.AsteroidBelts
// 	}
// 	return belts
// }

// func totalStargates(systems []SolarSystemData) int {
// 	gates := 0
// 	for _, system := range systems {
// 		gates += system.Stargates
// 	}
// 	return gates
// }

// func systemAsteroidBeltCount(planets map[int]Planet) int {
// 	asteroidBelts := 0
// 	for _, planet := range planets {
// 		asteroidBelts += len(planet.AsteroidBelts)
// 	}
// 	return asteroidBelts
// }

// func Add(regions []string, region string) []string {
// 	for _, name := range regions {
// 		if name == region {
// 			return regions
// 		}
// 	}
// 	regions = append(regions, region)
// 	return regions
// }

// func readUniverse(zf *zip.ReadCloser) {
// 	space := []string{}
// 	regions := make(map[string][]string)
// 	for _, file := range zf.File {
// 		name := file.Name
// 		if strings.Contains(name, "sde/fsd/universe") {
// 			spaceSubstring := name[17:]
// 			index := strings.Index(spaceSubstring, "/")
// 			space = Add(space, spaceSubstring[:index])
// 		}
// 		for _, s := range space {
// 			if strings.Contains(name, s) {
// 				index := 18 + len(s)
// 				region := name[index:]
// 				indexSlash := strings.Index(region, "/")
// 				regions[s] = Add(regions[s], region[:indexSlash])
// 			}
// 		}
// 	}
// 	// fmt.Println(space)
// 	// fmt.Println(regions)

// 	for k, v := range regions {
// 		fmt.Println(k, len(v))
// 	}
// 	fmt.Println()
// 	fmt.Println(regions["eve"])
// }

// func dumpFilenames(archive string, saveFilename string) error {
// 	zf, err := zip.OpenReader(archive)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
// 	}
// 	defer zf.Close()

// 	file, err := os.OpenFile(saveFilename, os.O_CREATE|os.O_RDWR, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)

// 	for _, document := range zf.File {
// 		writer.WriteString(document.Name + "\n")
// 	}
// 	return nil
// }

// func getName(filename string) string {
// 	lastIndex := strings.LastIndex(filename, "/")
// 	directories := filename[0:lastIndex]
// 	lastIndex = strings.LastIndex(directories, "/")
// 	return directories[lastIndex+1:]
// }

// func isJove(filename string) bool {
// 	jove := []string{"A821-A", "J7HZ-F", "UUA-F4"}
// 	for _, region := range jove {
// 		if strings.Contains(filename, region) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func kSpaceStaticData(zipFilename string, kspace bool) map[string]SolarSystem {
// 	solarSystems := make(map[string]SolarSystem)
// 	zf, err := zip.OpenReader(zipFilename)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
// 		os.Exit(1)
// 	}
// 	defer zf.Close()

// 	for _, file := range zf.File {
// 		name := file.Name
// 		solarSystemName := getName(name)
// 		if strings.Contains(name, "sde/fsd/universe/eve") {
// 			if kspace && isJove(name) {
// 				continue
// 			}
// 			if strings.Contains(name, "solarsystem.staticdata") {
// 				data, err := file.Open()
// 				if err != nil {
// 					fmt.Fprintf(os.Stderr, "%s: %s\n", name, err.Error())
// 					os.Exit(1)
// 				}

// 				content, err := io.ReadAll(data)
// 				if err != nil {
// 					fmt.Fprintf(os.Stderr, "%s read error: %s\n", name, err.Error())
// 					os.Exit(1)
// 				}

// 				var solarSystem SolarSystem
// 				if err = yaml.Unmarshal(content, &solarSystem); err != nil {
// 					if err != nil {
// 						fmt.Fprintf(os.Stderr, "unmarshal error in %s: %s\n", name, err.Error())
// 						os.Exit(1)
// 					}
// 				}
// 				solarSystems[solarSystemName] = solarSystem
// 			}
// 		}
// 	}
// 	return solarSystems
// }

// func wormholeStaticData(zipFilename string) map[string]SolarSystem {
// 	solarSystems := make(map[string]SolarSystem)
// 	zf, err := zip.OpenReader(zipFilename)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
// 		os.Exit(1)
// 	}
// 	defer zf.Close()

// 	for _, file := range zf.File {
// 		name := file.Name
// 		solarSystemName := getName(name)
// 		if strings.Contains(name, "sde/fsd/universe/wormhole") {
// 			if strings.Contains(name, "solarsystem.staticdata") {
// 				data, err := file.Open()
// 				if err != nil {
// 					fmt.Fprintf(os.Stderr, "%s: %s\n", name, err.Error())
// 					os.Exit(1)
// 				}

// 				content, err := io.ReadAll(data)
// 				if err != nil {
// 					fmt.Fprintf(os.Stderr, "%s read error: %s\n", name, err.Error())
// 					os.Exit(1)
// 				}

// 				var solarSystem SolarSystem
// 				if err = yaml.Unmarshal(content, &solarSystem); err != nil {
// 					if err != nil {
// 						fmt.Fprintf(os.Stderr, "unmarshal error in %s: %s\n", name, err.Error())
// 						os.Exit(1)
// 					}
// 				}
// 				solarSystems[solarSystemName] = solarSystem
// 			}
// 		}
// 	}
// 	return solarSystems
// }

// func joveStaticData(zipFilename string) map[string]SolarSystem {
// 	solarSystems := make(map[string]SolarSystem)
// 	zf, err := zip.OpenReader(zipFilename)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
// 		os.Exit(1)
// 	}
// 	defer zf.Close()

// 	for _, file := range zf.File {
// 		name := file.Name
// 		solarSystemName := getName(name)
// 		if strings.Contains(name, "sde/fsd/universe/eve") {
// 			if !isJove(name) {
// 				continue
// 			}
// 			if strings.Contains(name, "solarsystem.staticdata") {
// 				data, err := file.Open()
// 				if err != nil {
// 					fmt.Fprintf(os.Stderr, "%s: %s\n", name, err.Error())
// 					os.Exit(1)
// 				}

// 				content, err := io.ReadAll(data)
// 				if err != nil {
// 					fmt.Fprintf(os.Stderr, "%s read error: %s\n", name, err.Error())
// 					os.Exit(1)
// 				}

// 				var solarSystem SolarSystem
// 				if err = yaml.Unmarshal(content, &solarSystem); err != nil {
// 					if err != nil {
// 						fmt.Fprintf(os.Stderr, "unmarshal error in %s: %s\n", name, err.Error())
// 						os.Exit(1)
// 					}
// 				}
// 				solarSystems[solarSystemName] = solarSystem
// 			}
// 		}
// 	}
// 	return solarSystems
// }
