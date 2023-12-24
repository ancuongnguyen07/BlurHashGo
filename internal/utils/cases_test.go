package utils_test

var linearTosRGBTestCases = []struct {
	in  float64
	out int
}{
	{-3.25, 0},
	{3.25, 255},
	{0.5, 188},
}

var sRGBToLinearTestCases = []struct {
	in  int
	out float64
}{
	{0, 0.0},
	{150, 0.304987},
	{255, 1.0},
}

var signPowTestCases = []struct {
	in  float64
	exp float64
	out float64
}{
	{-2.0, 4, -16.0},
	{2.0, 4, 16.0},
}
