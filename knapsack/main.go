package main

import (
	"fmt"
	"math"
)

type Package struct {
	Width, Height, Depth, Weight, Priority int
}

func (p Package) Volume() int {
	return p.Width * p.Height * p.Depth
}

func main() {
	fmt.Println("Knapsack Problem")

	var testPackages []Package
	for i := 0; i < 150; i++ {
		testPackages = append(testPackages, Package{Width: 25, Height: 30, Depth: 15, Priority: i%3 + 1})
		testPackages = append(testPackages, Package{Width: 35, Height: 40, Depth: 25, Priority: i%3 + 1})
		testPackages = append(testPackages, Package{Width: 45, Height: 50, Depth: 35, Priority: i%3 + 1})
	}

	truckVolume := 2000000 // 2 cubic meters in cubic cm

	maxPackagesCount := MaxPackages(testPackages, truckVolume)

	fmt.Printf("Maximum number of packages that can be loaded: %d\n", maxPackagesCount)
}

func MaxPackages(packages []Package, truckVol int) int {
	// Initialize DP table
	dp := make([][2]int, truckVol+1)
	for i := range dp {
		// Use a large number for initial priority sum
		dp[i][0] = math.MaxInt32
	}

	// Use a large number for initial priority sum
	dp[0] = [2]int{0, 0}

	for _, pkg := range packages {
		pkgVol := pkg.Volume()

		for j := truckVol; j >= pkgVol; j-- {
			// We check if adding this package improves the number of packages without excessively worsening the
			// priority sum. We also allow some degradation in priority sum.
			if dp[j-pkgVol][1]+1 > dp[j][1] &&
				dp[j-pkgVol][0]+pkg.Priority < dp[j][0]+10 {
				dp[j][0] = dp[j-pkgVol][0] + pkg.Priority // Priority sum
				dp[j][1] = dp[j-pkgVol][1] + 1            // Package count
			}
		}
	}

	return dp[truckVol][1]
}
