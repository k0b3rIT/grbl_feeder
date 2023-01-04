package gcodemanipulator

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

	
type gcode struct {
    symbol string
    pattern *regexp.Regexp
}

var (
	GCODE_X = gcode{"X", regexp.MustCompile(`X(-*\d+(?:\.\d+|\d*))`)}
	GCODE_Y = gcode{"Y", regexp.MustCompile(`Y(-*\d+(?:\.\d+|\d*))`)}
)

func MoveToTopLeft(gcode []string, machineMaxX float64, machineMaxY float64) []string {
	minX, maxX, minY, maxY := getGcodeExtremes(gcode)

	fmt.Println(minX, maxX, minY, maxY)

	offset := machineMaxY - maxY

	var modifiedGcode = []string{}

	for _, line := range gcode {
		yValue, err := getValueFromGcodeLine(line, GCODE_Y)
		if (err == nil) {
			newLine := replaceValueInGcodeLine(line, GCODE_Y, yValue + offset)
			modifiedGcode = append(modifiedGcode, newLine)
		} else {
			modifiedGcode = append(modifiedGcode, line)
		}
	}
	return modifiedGcode
}

func ScaleTo(gcode []string, maxWidthAndHight float64) []string {
	minX, maxX, minY, maxY := getGcodeExtremes(gcode)
	width := maxX - minX
	height := maxY - minY

	if (height > maxWidthAndHight || width > maxWidthAndHight) {
		largestDimension := math.Max(height, width)
		scaleFactor := largestDimension / maxWidthAndHight

		var modifiedGcode = []string{}

		for _, line := range gcode {
			lineTmp := line
			yValue, err := getValueFromGcodeLine(lineTmp, GCODE_Y)
			if (err == nil) {
				lineTmp = replaceValueInGcodeLine(lineTmp, GCODE_Y, yValue / scaleFactor)
			}

			xValue, err := getValueFromGcodeLine(lineTmp, GCODE_X)
			if (err == nil) {
				lineTmp = replaceValueInGcodeLine(lineTmp, GCODE_X, xValue / scaleFactor)
			}

			modifiedGcode = append(modifiedGcode, lineTmp)
		}

		return modifiedGcode
	}

	return gcode
}

func OffsetX(gcode []string, offset float64) []string {
	return offsetOnGcode(GCODE_X, gcode, offset)
}

func OffsetY(gcode []string, offset float64) []string {
	return offsetOnGcode(GCODE_Y, gcode, offset)
}

func offsetOnGcode(gcodeType gcode, gcode []string, offset float64) []string {
	var modifiedGcode = []string{}
	for _, line := range gcode {
		value, err := getValueFromGcodeLine(line, gcodeType)
		if (err == nil) {
			newLine := replaceValueInGcodeLine(line, gcodeType, value + offset)
			modifiedGcode = append(modifiedGcode, newLine)
		} else {
			modifiedGcode = append(modifiedGcode, line)
		}
	}
	return modifiedGcode
}

func getValueFromGcodeLine(line string, gcodeType gcode) (float64, error) {
	values := gcodeType.pattern.FindStringSubmatch(line)
	if (len(values) > 0) {
		value, _ := strconv.ParseFloat(values[1], 64)
		return value, nil
	} else {
		return 0, fmt.Errorf("no [%s] value found", gcodeType.symbol)
	}
}

func replaceValueInGcodeLine(line string, gcodeType gcode, newValue float64) string {
	originalValue := gcodeType.pattern.FindStringSubmatch(line)
	if (len(originalValue) > 0) {
		newLine := strings.Replace(line, originalValue[0], fmt.Sprintf("%s%.4f", gcodeType.symbol, newValue), 1)
		return newLine
	} else {
		panic(fmt.Sprintf("Gcode [%s] not found!", gcodeType.symbol))
	}
}

func getGcodeExtremes(gcode []string) (float64, float64, float64, float64) {
	var minX float64 = 0
	var maxX float64 = 0
	var minY float64 = 0
	var maxY float64 = 0
	for _, line := range gcode {
		xValue, err := getValueFromGcodeLine(line, GCODE_X)
		if (err == nil) {
			if (minX > xValue) {
				minX = xValue
			}
			if (maxX < xValue) {
				maxX = xValue
			}
		}

		yValue, err := getValueFromGcodeLine(line, GCODE_Y)
		if (err == nil) {
			if (minY > yValue) {
				minY = yValue
			}
			if (maxY < yValue) {
				maxY = yValue
			}
		}
	}
	return minX, maxX, minY, maxY
}