package gcodemanipulator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	xValueRegex = regexp.MustCompile(`X(-*\d+(?:\.\d+|\d*))`)	
	yValueRegex = regexp.MustCompile(`Y(-*\d+(?:\.\d+|\d*))`)
)

func MoveToTopLeft(gcode []string, machineMaxX float64, machineMaxY float64) []string {
	minX, maxX, minY, maxY := getGcodeExtremes(gcode)

	fmt.Println(minX, maxX, minY, maxY)

	offset := machineMaxY - maxY


	var modifiedGcode = []string{}

	for _, line := range gcode {
		originalValue := yValueRegex.FindStringSubmatch(line)
		if (len(originalValue) > 0) {
			yValue, _ := strconv.ParseFloat(originalValue[1], 64)
			newLine := strings.Replace(line, originalValue[0], fmt.Sprintf("Y%.4f", yValue + offset), 1)
			// fmt.Println(line + "\t\t" + newLine)
			modifiedGcode = append(modifiedGcode, newLine)
		} else {
			modifiedGcode = append(modifiedGcode, line)
		}
	}
	return modifiedGcode
}

func getGcodeExtremes(gcode []string) (float64, float64, float64, float64) {
	var minX float64 = 0
	var maxX float64 = 0
	var minY float64 = 0
	var maxY float64 = 0
	for _, line := range gcode {
		xStringParts := xValueRegex.FindStringSubmatch(line)
		if (len(xStringParts) > 0) {
			value, _ := strconv.ParseFloat(xStringParts[1], 64)
			if (minX > value) {
				minX = value
			}
			if (maxX < value) {
				maxX = value
			}
		}

		yStringParts := yValueRegex.FindStringSubmatch(line)
		if (len(yStringParts) > 0) {
			fmt.Println(yStringParts[1])
			value, _ := strconv.ParseFloat(yStringParts[1], 64)
			if (minY > value) {
				minY = value
			}
			if (maxY < value) {
				maxY = value
			}
		}


	}
	return minX, maxX, minY, maxY
}


func getGcodeMaxY(gcode []string) float32 {
	return 0
}