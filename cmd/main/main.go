package main

import (
	grblfeeder "plotter/pkg/grbl_feeder"
	// "fmt"
	gcodemanipulator "plotter/pkg/gcode_manipulator"
	"plotter/pkg/util"
)



func main() {
	var grbl *grblfeeder.GrblFeeder = grblfeeder.NewGrblFeeder("/dev/ttyUSB0")
	gcode := util.ReadFile("lilien.gcode")
	modifiedGcode := gcodemanipulator.MoveToTopLeft(gcode, 305, 200)
	// fmt.Println(modifiedGcode)

	util.WriteFile("m_lilien.gcode", modifiedGcode)

	grbl.SendGcodes(modifiedGcode)
}
