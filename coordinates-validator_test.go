package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"regexp"
	"strconv"
	"strings"
)

var _ = Describe("Valid coordinates", func() {

	validCoordinates := []string{
		"-23, 25",
		"4, -3",
		"24.53525235, 23.45235",
		"04, -23.234235",
		"43.91343345, 143"}

	FIt("should validate the coordinates", func() {
		for _, coordinates := range validCoordinates {
			Expect(IsValidCoordinates(coordinates)).To(Equal(true))
		}
	})
})

var _ = Describe("Invalid coordinates", func() {

	invalidCoordinates := []string{
		"23.234, - 23.4234",
		"2342.43536, 34.324236",
		"N23.43345, E32.6457",
		"99.234, 12.324",
		"6.325624, 43.34345.345",
		"0, 1,2",
		"0.342q0832, 1.2324",
		"23.245, 1e1"}

	It("should invalidate the coordinates", func() {
		for _, coordinates := range invalidCoordinates {
			Expect(IsValidCoordinates(coordinates)).To(Equal(false))
		}
	})
})

func IsValidCoordinates(coordinates string) bool {
	var re = regexp.MustCompile(`^-?(\d+\.?\d*),\s-?(\d+\.?\d*)$`)
	var match = re.MatchString(coordinates)
	if match {
		matchs := strings.Split(coordinates, ",")
		lat, _ := strconv.ParseFloat(matchs[0], 32)
		lon, _ := strconv.ParseFloat(matchs[1], 32)
		return lat <= 90 && lon <= 180
	}
	return false
}
