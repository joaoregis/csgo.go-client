package features

import (
	"gosource/internal/configs"
	"gosource/internal/csgo"
	kb "gosource/internal/hackFunctions/keyboard"
	"gosource/internal/hackFunctions/vector"
	"gosource/internal/memory"
	"math"
)

// Calculating distance
func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)*1.0)
}

func getSmoothedValue(target float64, origin float64) float64 {
	differenceBetween2 := target - origin

	return origin + differenceBetween2/configs.G.Aimbot.Smooth
}

func AimbotAt(closestEnemy uintptr) {

	// 8 is head position
	// TODO: look at nearest hitbox or make selectable hitbox
	target, err := csgo.GetBonePos(closestEnemy, 8)
	if err != nil {
		return
	}

	localPlayer, err := csgo.GetLocalPlayer()
	if err != nil || localPlayer == 0 {
		return
	}

	dwClientStatePointer, err := memory.GameProcess.ReadIntPtr(memory.GameEngine + configs.Offsets.Signatures.DwClientState)
	if err != nil {
		return
	}

	viewAngles, err := memory.GameProcess.ReadVec3(dwClientStatePointer + configs.Offsets.Signatures.DwClientStateViewAngles)
	if err != nil {
		return
	}

	origin, err := memory.GameProcess.ReadVec3(localPlayer + configs.Offsets.Netvars.MVecOrigin)
	if err != nil {
		return
	}

	viewOffset, err := memory.GameProcess.ReadVec3(localPlayer + configs.Offsets.Netvars.MVecViewOffset)
	if err != nil {
		return
	}

	myPosition := origin.CalcVector3WithOtherVector3(*viewOffset, "+")
	deltaVec := target.CalcVector3WithOtherVector3(*myPosition, "-")
	deltaVecLength := math.Sqrt(deltaVec.X*deltaVec.X + deltaVec.Y*deltaVec.Y + deltaVec.Z*deltaVec.Z)
	pitch := -math.Asin(deltaVec.Z/deltaVecLength) * (180 / math.Pi)
	yaw := math.Atan2(deltaVec.Y, deltaVec.X) * (180 / math.Pi)

	smoothedPitch := getSmoothedValue(pitch, viewAngles.X)
	smoothedYaw := getSmoothedValue(yaw, viewAngles.Y)

	if distance(viewAngles.X, viewAngles.Y, pitch, yaw) > configs.G.Aimbot.Fov {
		// prevent to aim when target is out of fov circle
		return
	}

	// check the limits of csgo. if you go through this limits you'll get vacced
	if smoothedPitch >= -89 && smoothedPitch <= 89 && smoothedYaw >= -180 && smoothedYaw <= 180 {

		// if configs.G.Aimbot.NonStickyAim {

		// // TODO:
		// // get mouse delta
		// // calculate distance between aim delta and mouse delta
		// // determines mouse direction and aim delta direction (derivative)
		// // if converges, should perform aim, if not, should cancel aim

		// }

		// just execute aimbot without prevent sticky aims
		memory.GameProcess.WriteVec3(dwClientStatePointer+configs.Offsets.Signatures.DwClientStateViewAngles, vector.Vector3{X: smoothedPitch, Y: smoothedYaw, Z: viewAngles.Z})

	}
}

func Aimbot() {

	if !configs.G.Aimbot.Enabled {
		return
	}

	hWId, _ := csgo.GetActiveWeapon()
	if !csgo.IsAimbotable(hWId) {
		return
	}

	if !kb.GetAsyncKeyState(kb.GetKey(configs.G.Aimbot.Key)) {
		return
	}

	closestEnemy, err := csgo.GetClosestPlayer()
	if err != nil {
		return
	}

	if closestEnemy != 0 {
		AimbotAt(closestEnemy)
	}
}
