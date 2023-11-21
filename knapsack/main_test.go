// main_test.go
package main

import "testing"

func TestMaxPackages(t *testing.T) {
	var testPackages []Package
	for i := 0; i < 20; i++ { // Adjust the number of each package type as needed
		testPackages = append(testPackages, Package{Width: 25, Height: 30, Depth: 10, Weight: 500, Priority: i%3 + 1})
		testPackages = append(testPackages, Package{Width: 10, Height: 40, Depth: 65, Weight: 10_000, Priority: i%3 + 1})
		testPackages = append(testPackages, Package{Width: 20, Height: 40, Depth: 65, Weight: 15_000, Priority: i%3 + 1})
		testPackages = append(testPackages, Package{Width: 40, Height: 40, Depth: 64, Weight: 20_000, Priority: i%3 + 1})
	}

	// Define test cases
	tests := []struct {
		name        string
		packages    []Package
		truckVol    int
		expectedMax int
	}{
		{
			name:        "Test Case 1",
			packages:    testPackages,
			truckVol:    2000000, // 2 cubic meters in cubic cm
			expectedMax: 62,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualMax := MaxPackages(tc.packages, tc.truckVol)
			if actualMax != tc.expectedMax {
				t.Errorf("Test %v failed. Expected %v, got %v", tc.name, tc.expectedMax, actualMax)
			}
		})
	}
}
