package model

import "fmt"

type SafeInt int

func (a SafeInt) Abs() int {
	if a < 0 {
		return int(a) * -1
	}
	return int(a)
}
func (a SafeInt) Int() int {

	return int(a)
}

type Drone struct {
	X           int
	Y           int
	z           int
	distance    int
	maxDistance int
	isRunning   bool
	directionX  string
	directionZ  string
	directionY  string
}

func (d *Drone) Forward() {
	d.updateDir("", "forward", "")
	d.calcY(1)
	d.directionY = ""
	d.State()
}
func (d *Drone) Right() {
	d.updateDir("right", "", "")

	d.calcX(1)
	d.State()
}
func (d *Drone) Left() {
	d.updateDir("left", "", "")
	d.calcX(-1)
	d.State()
}
func (d *Drone) UpDown(distance ...SafeInt) {
	nd := 1
	if len(distance) == 1 {
		nd = distance[0].Int()
	}
	targZ := nd - d.z
	if targZ == 0 {
		return
	}
	tdz := "up"
	if targZ < 0 {
		tdz = "down"
	}
	d.directionZ = tdz
	d.calcZ(SafeInt(targZ))
	d.updateDir("", "", tdz)
	d.State()
}
func (d *Drone) updateDir(x, y, z string) {
	d.directionX = x
	d.directionY = y
	d.directionZ = z
}
func (d *Drone) calcZ(distance SafeInt) {
	if d.isRunning {
		d.z += distance.Int()
		d.distance += distance.Abs()
		d.calcDistanceUsage()

	}
}
func (d *Drone) calcX(distance SafeInt) {
	if d.isRunning {
		d.X += distance.Int()
		d.distance += distance.Abs() * 10
		d.calcDistanceUsage()

	}
}
func (d *Drone) calcY(distance SafeInt) {
	if d.isRunning {
		d.Y += distance.Int()
		d.distance += distance.Abs() * 10
		d.calcDistanceUsage()
	}
}
func (d *Drone) Distance() int {
	return d.distance
}

func (d *Drone) calcDistanceUsage() {
	if d.maxDistance == -1 {
		return
	}
	if d.distance >= d.maxDistance {
		//reset distance to max distance because using backup battery
		d.distance = d.maxDistance
		d.isRunning = false
	}
}
func (d *Drone) IsRest() bool {
	return !d.isRunning
}
func (d *Drone) State() {
	fmt.Printf(`
	====
	drone state :
		direction X: %s
		direction Y: %s 
		direction Z: %s
		Pos(x,y,z): (%v,%v,%v) 
		IsRest: %v 
		maxDistance: %v  
		distance: %v
	
	`, d.directionX, d.directionY, d.directionZ, d.X, d.Y, d.z, d.IsRest(), d.maxDistance, d.distance)
}

func InitDrone(distance ...int) *Drone {
	d := -1
	if len(distance) > 0 && distance[0] > 0 {
		d = distance[0]
	}
	return &Drone{X: 1, Y: 1, z: 0, distance: 0, maxDistance: d, isRunning: true, directionX: "right"}
}
