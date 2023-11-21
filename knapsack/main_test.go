package main

import "testing"

// createTestPackages creates a fixed set of packages for testing
func createTestPackages() []Package {
	var packages []Package
	idCounter := 0

	// Creating a few packages of each type
	for i := 0; i < 5; i++ {
		packages = append(packages, Package{ID: idCounter, Type: "A", Size: SizeA, Priority: 1})
		idCounter++
		packages = append(packages, Package{ID: idCounter, Type: "B", Size: SizeB, Priority: 2})
		idCounter++
		packages = append(packages, Package{ID: idCounter, Type: "C", Size: SizeC, Priority: 3})
		idCounter++
	}

	return packages
}

func TestMaxPackages(t *testing.T) {
	testPackages := createTestPackages()

	// Define test cases
	tests := []struct {
		name        string
		packages    []Package
		truckVol    int
		expectedMax int
	}{
		{
			name:        "Sample Test Case",
			packages:    testPackages,
			truckVol:    2500000,
			expectedMax: 49,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualMax := len(MaxPackages(tc.packages, tc.truckVol))
			if actualMax != tc.expectedMax {
				t.Errorf("Test %v failed. Expected %v, got %v", tc.name, tc.expectedMax, actualMax)
			}
		})
	}
}
