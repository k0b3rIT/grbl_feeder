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
	modifiedGcode := gcodemanipulator.ScaleTo(gcode, 50)
	modifiedGcode = gcodemanipulator.MoveToTopLeft(modifiedGcode, 305, 200)
	grbl.SendGcodes(modifiedGcode)

	modifiedGcode = gcodemanipulator.OffsetX(modifiedGcode, 60)
	modifiedGcode = gcodemanipulator.OffsetY(modifiedGcode, -60)
	grbl.SendGcodes(modifiedGcode)


	// fmt.Println(modifiedGcode)

	// util.WriteFile("m_lilien.gcode", modifiedGcode)

	
}
