package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed testdata/input.txt
var input []byte

type Point struct {
	x, y, z int
}

func (p *Point) Subtract(other Point) Point {
	return Point{
		x: p.x - other.x,
		y: p.y - other.y,
		z: p.z - other.z,
	}
}

func (p Point) Add(other Point) Point {
	return Point{
		x: p.x + other.x,
		y: p.y + other.y,
		z: p.z + other.z,
	}
}

type Rotation func(p Point) Point

func IdentityRotation(p Point) Point { return p }

var rotations = []Rotation{
	func(p Point) Point { return Point{x: p.x, y: p.y, z: p.z} },
	func(p Point) Point { return Point{x: p.z, y: p.y, z: -p.x} },
	func(p Point) Point { return Point{x: -p.x, y: p.y, z: -p.z} },
	func(p Point) Point { return Point{x: p.z, y: p.y, z: -p.x} },

	func(p Point) Point { return Point{x: p.x, y: -p.z, z: -p.y} },
	func(p Point) Point { return Point{x: p.y, y: -p.z, z: -p.x} },
	func(p Point) Point { return Point{x: -p.x, y: -p.z, z: -p.y} },
	func(p Point) Point { return Point{x: -p.y, y: -p.z, z: p.x} },

	func(p Point) Point { return Point{x: p.x, y: p.z, z: -p.y} },
	func(p Point) Point { return Point{x: -p.y, y: p.z, z: -p.x} },
	func(p Point) Point { return Point{x: -p.x, y: p.z, z: p.y} },
	func(p Point) Point { return Point{x: p.y, y: p.z, z: p.x} },

	func(p Point) Point { return Point{x: -p.x, y: -p.y, z: p.z} },
	func(p Point) Point { return Point{x: p.z, y: -p.y, z: p.x} },
	func(p Point) Point { return Point{x: p.x, y: -p.y, z: -p.z} },
	func(p Point) Point { return Point{x: -p.z, y: -p.y, z: -p.x} },

	func(p Point) Point { return Point{x: p.y, y: -p.x, z: p.z} },
	func(p Point) Point { return Point{x: p.z, y: -p.x, z: -p.y} },
	func(p Point) Point { return Point{x: -p.y, y: -p.x, z: -p.z} },
	func(p Point) Point { return Point{x: -p.z, y: -p.x, z: p.y} },

	func(p Point) Point { return Point{x: -p.y, y: p.x, z: p.z} },
	func(p Point) Point { return Point{x: p.z, y: p.x, z: p.y} },
	func(p Point) Point { return Point{x: p.y, y: p.x, z: -p.z} },
	func(p Point) Point { return Point{x: -p.z, y: p.x, z: -p.y} },
}

func (p *Point) Equal(other Point) bool {
	return p.x == other.x && p.y == other.y && p.z == other.z
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.x, p.y, p.z)
}

type Scanner struct {
	name    string
	beacons []Point
}

type ScannerWithScanners struct {
	*Scanner
	scanners []Point
}

var zero = Point{}

func main() {
	input, err := readInput(bytes.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	scanner := solve(input)

	fmt.Println(len(scanner.beacons))
	fmt.Println(manhattanDistance(scanner))
}

func compare(refRelativeBeacons, compRelativeBeacons map[Point]bool) ([]Point, []Point) {
	same := []Point{}
	distinct := []Point{}

	for compPoint, _ := range compRelativeBeacons {
		if _, ok := refRelativeBeacons[compPoint]; ok {
			same = append(same, compPoint)
		} else {
			distinct = append(distinct, compPoint)
		}
	}

	return same, distinct
}

func readInput(r io.Reader) ([]*Scanner, error) {
	scanner := bufio.NewScanner(r)

	scanners := []*Scanner{}
	for scanner.Scan() {
		s := scanner.Text()
		if s[:3] != "---" {
			break
		}

		scan, err := readScanner(s, scanner)
		if err != nil {
			return nil, err
		}

		scanners = append(scanners, scan)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return scanners, nil
}

func readScanner(name string, s *bufio.Scanner) (*Scanner, error) {
	points := []Point{}

	for s.Scan() {
		s := s.Text()
		if len(s) == 0 {
			break
		}

		coordinates := []int{}
		for _, part := range strings.Split(s, ",") {
			coordinate, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("parsing coordinate: %w", err)
			}
			coordinates = append(coordinates, coordinate)
		}

		points = append(points, Point{
			x: coordinates[0],
			y: coordinates[1],
			z: coordinates[2],
		})
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return &Scanner{
		name:    name,
		beacons: points,
	}, nil
}

func solve(input []*Scanner) *ScannerWithScanners {
	unsolved := map[string]*ScannerWithScanners{}
	for i := 0; i < len(input); i++ {
		unsolved[input[i].name] = &ScannerWithScanners{Scanner: input[i]}
	}

	for len(unsolved) > 1 {
		for _, refScanner := range unsolved {
		start:
			for _, refBeacon := range refScanner.beacons {
				refRelative := buildBeaconRelativeCoordinates(zero.Subtract(refBeacon), refScanner.beacons, IdentityRotation)

				for _, compScanner := range unsolved {
					if refScanner == compScanner {
						continue
					}

					for _, compBeacon := range compScanner.beacons {
						for _, rotation := range rotations {
							compBeacon := rotation(compBeacon)

							diff := zero.Subtract(compBeacon)

							compRelativeBeacons := buildBeaconRelativeCoordinates(diff, compScanner.beacons, rotation)
							same, distinct := compare(refRelative, compRelativeBeacons)
							if len(same) >= 12 {
								for _, newBeacon := range distinct {
									refScanner.beacons = append(refScanner.beacons, newBeacon.Add(refBeacon))
								}

								for _, newScanner := range compScanner.scanners {
									newScanner = rotation(newScanner)
									refScanner.scanners = append(refScanner.scanners, newScanner.Add(diff).Add(refBeacon))
								}

								refScanner.scanners = append(refScanner.scanners, zero.Add(diff).Add(refBeacon))

								delete(unsolved, compScanner.name)
								break start
							}
						}
					}
				}
			}
		}
	}

	keys := make([]string, 0, len(unsolved))
	for k := range unsolved {
		keys = append(keys, k)
	}

	return unsolved[keys[0]]
}

func buildBeaconRelativeCoordinates(diff Point, beacons []Point, rotation Rotation) map[Point]bool {
	result := map[Point]bool{}

	for _, comparison := range beacons {
		result[rotation(comparison).Add(diff)] = true
	}

	return result
}

func manhattanDistance(scanner *ScannerWithScanners) int {
	maxDistance := float64(0)
	allPoints := append(scanner.scanners, Point{x: 0, y: 0, z: 0})

	for _, p1 := range allPoints {
		for _, p2 := range allPoints {
			diff := p1.Subtract(p2)

			distance := math.Abs(float64(diff.x)) + math.Abs(float64(diff.y)) + math.Abs(float64(diff.z))
			maxDistance = math.Max(float64(distance), float64(maxDistance))
		}
	}

	return int(maxDistance)
}
