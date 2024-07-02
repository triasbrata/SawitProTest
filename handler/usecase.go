package handler

import (
	"sort"

	"github.com/SawitProRecruitment/UserService/constants"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/repository"
)

func calculateDronePlan(estate repository.Estate, trees []repository.Tree, max_distance int) model.DronePlanResponseSuccess {
	res := model.DronePlanResponseSuccess{}
	md := []int{}
	mapTree := make(map[int]map[int]repository.Tree) // map by x then y then h
	firstTreeHight := 0
	for _, t := range trees {
		if t.X == 1 && t.Y == 1 {
			firstTreeHight = int(t.Height)
		}
		xPos, existX := mapTree[int(t.X)]
		if !existX {
			xPos = make(map[int]repository.Tree)
		}
		pos, exists := xPos[int(t.Y)]
		if !exists {
			pos = t
		}
		xPos[int(t.Y)] = pos
		mapTree[int(t.X)] = xPos
	}
	if max_distance > 0 {
		md = append(md, max_distance)
	}
	drone := model.InitDrone(md...)
	droneDirection := "right"
	drone.UpDown(model.SafeInt(firstTreeHight + 1))

	for {
		if drone.X == int(estate.Length) {
			droneDirection = "left"
			checkTreeY(drone, mapTree)
			drone.Forward()
		}
		if drone.X == 1 && drone.Y > 1 {
			droneDirection = "right"
			checkTreeY(drone, mapTree)
			drone.Forward()

		}
		if droneDirection == "left" {
			checkTreeX(drone, mapTree, -1)
			drone.Left()

		}
		if droneDirection == "right" {
			checkTreeX(drone, mapTree, 1)
			drone.Right()

		}
		//drone battery is low
		if drone.IsRest() {
			break
		}
		//end of estate
		if drone.Y >= int(estate.Width) && drone.X >= int(estate.Length) {
			break
		}
		//out of bound
		if drone.Y > constants.MAX_ESTATE_WIDTH_LENGTH || drone.X > constants.MAX_ESTATE_WIDTH_LENGTH || drone.X < constants.MIN_ESTATE_WIDTH_LENGTH || drone.Y < constants.MIN_ESTATE_WIDTH_LENGTH {
			break
		}

	}
	drone.UpDown(0)
	res.Distance = int64(drone.Distance())
	if drone.IsRest() {
		res.Rest = &model.DronePlanResponseSuccessRes{
			X: int64(drone.X),
			Y: int64(drone.Y),
		}
	}
	return res
}

func checkTreeY(drone *model.Drone, mapTree map[int]map[int]repository.Tree) {

	var (
		foundTree bool
		tree      repository.Tree
	)
	treeX, safe := mapTree[drone.X]
	if safe {
		tree, foundTree = treeX[drone.Y+1]
		if safe {
			drone.UpDown(model.SafeInt(tree.Height + 1))
		}
	}
	if !foundTree {
		drone.UpDown(1)
	}
}
func checkTreeX(drone *model.Drone, mapTree map[int]map[int]repository.Tree, mod int) {
	nx := drone.X + mod
	if nx < 1 {
		nx = 1
	}
	var (
		foundTree bool
		tree      repository.Tree
	)
	treeX, safe := mapTree[nx]
	if safe {
		tree, foundTree = treeX[drone.Y]
		if foundTree {
			drone.UpDown(model.SafeInt(tree.Height + 1))
		}
	}
	if !foundTree {
		drone.UpDown(1)
	}
}

// calculate min, max, and median
func calculateStats(res repository.GetEstateTreeResponse) model.EstateStatsResponse {
	count := len(res.Data)
	max := int64(0)
	min := int64(0)
	median := float64(0)
	heights := make([]int64, 0, count)
	for _, t := range res.Data {
		if t.Height > max {
			max = t.Height
		}
		if t.Height < min || min == 0 {
			min = t.Height
		}
		heights = append(heights, t.Height)
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})

	if count%2 == 1 {
		median = float64(heights[(count-1)/2])
	} else {
		median = float64(heights[count/2-1]+heights[count/2]) / 2.0
	}
	calc := model.EstateStatsResponse{
		Count:  count,
		Max:    max,
		Min:    min,
		Median: median,
	}
	return calc
}
