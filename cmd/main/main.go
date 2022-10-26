package main

import (
	grblfeeder "plotter/pkg/grbl_feeder"
	"plotter/pkg/util"
)



func main() {
	var grbl *grblfeeder.GrblFeeder = grblfeeder.NewGrblFeeder("/dev/ttyUSB0")
	gcode := util.ReadFile("gcode.gcode")

	grbl.SendGcodes(gcode)
}
