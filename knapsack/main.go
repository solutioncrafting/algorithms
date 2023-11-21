package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type PackageSize struct {
	Width, Height, Depth int
}

func (ps PackageSize) Volume() int {
	return ps.Width * ps.Height * ps.Depth
}

type Package struct {
	ID       int
	Size     PackageSize
	Weight   int
	Priority int
	Type     string
}

type PackageInfo struct {
	ID     int
	Type   string
	Volume int
}

// Predefined package sizes
var (
	SizeA = PackageSize{Width: 25, Height: 30, Depth: 15}
	SizeB = PackageSize{Width: 35, Height: 40, Depth: 25}
	SizeC = PackageSize{Width: 45, Height: 50, Depth: 35}
)

func main() {
	fmt.Println("Knapsack Problem")

	testPackages := sampleList()
	truckVolume := 2500000 // 2.5 cubic meters in cubic cm

	selectedPackages := MaxPackages(testPackages, truckVolume)

	Summarize(selectedPackages)
}

func MaxPackages(packages []Package, truckVol int) []PackageInfo {
	dp := make([][2]int, truckVol+1)
	lastAdded := make([]int, truckVol+1)

	for i := range dp {
		dp[i][0] = math.MaxInt32
		lastAdded[i] = -1
	}
	dp[0] = [2]int{0, 0}

	for i, pkg := range packages {
		pkgVol := pkg.Size.Volume()

		for j := truckVol; j >= pkgVol; j-- {
			if dp[j-pkgVol][1]+1 > dp[j][1] && dp[j-pkgVol][0]+pkg.Priority < dp[j][0]+10 {
				dp[j][0] = dp[j-pkgVol][0] + pkg.Priority
				dp[j][1] = dp[j-pkgVol][1] + 1
				lastAdded[j] = i
			}
		}
	}

	var selected []PackageInfo
	for vol := truckVol; vol > 0; {
		if lastAdded[vol] == -1 {
			break
		}

		pkgIndex := lastAdded[vol]
		pkg := packages[pkgIndex]

		selected = append([]PackageInfo{{
			ID:     pkg.ID,
			Type:   pkg.Type,
			Volume: pkg.Size.Volume(),
		}}, selected...)

		vol -= packages[pkgIndex].Size.Volume()
	}

	return selected
}

func Summarize(packages []PackageInfo) {
	totalCount := len(packages)
	typeCount := make(map[string]int)
	totalVolume := 0

	for _, pkg := range packages {
		typeCount[pkg.Type]++
		totalVolume += pkg.Volume
	}

	fmt.Printf("Type A: %d\n", typeCount["A"])
	fmt.Printf("Type B: %d\n", typeCount["B"])
	fmt.Printf("Type C: %d\n", typeCount["C"])
	fmt.Printf("Total Qty : %d\n", totalCount)
	fmt.Printf("Toal Volume: %d\n", totalVolume)
}

func sampleList() []Package {
	var packages []Package
	idCounter := 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		priorityA := rand.Intn(3) + 1
		priorityB := rand.Intn(3) + 1
		priorityC := rand.Intn(3) + 1

		packages = append(packages, Package{ID: idCounter, Type: "A", Size: SizeA, Priority: priorityA})
		idCounter++
		packages = append(packages, Package{ID: idCounter, Type: "B", Size: SizeB, Priority: priorityB})
		idCounter++
		packages = append(packages, Package{ID: idCounter, Type: "C", Size: SizeC, Priority: priorityC})
		idCounter++
	}

	return packages
}
