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

func AimbotAt(target vector.Vector3) {
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
		return
	}

	if smoothedPitch >= -89 && smoothedPitch <= 89 && smoothedYaw >= -180 && smoothedYaw <= 180 {
		memory.GameProcess.WriteVec3(dwClientStatePointer+configs.Offsets.Signatures.DwClientStateViewAngles, vector.Vector3{X: smoothedPitch, Y: smoothedYaw, Z: viewAngles.Z})
	}
}

func Aimbot() {

	if !configs.G.Aimbot.Enabled || !kb.GetAsyncKeyState(kb.GetKey(configs.G.Aimbot.Key)) {
		return
	}

	closestEnemy, err := csgo.GetClosestPlayer()

	if err != nil {
		return
	}

	if closestEnemy != 0 {
		headPos, _ := csgo.GetBonePos(closestEnemy, 8)

		AimbotAt(headPos)
	}
}
