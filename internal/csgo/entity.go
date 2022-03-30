package csgo

import (
	"gosource/internal/configs"
	"gosource/internal/hackFunctions/vector"
	"gosource/internal/memory"
	"math"
)

var LocalEntity uintptr
var MaxPlayers int

type PlayerEntity struct {
}

func UpdatePlayerVars() error {
	var err error
	LocalEntity, err = GetLocalPlayer()

	if err != nil {
		return err
	}

	MaxPlayers, err = GetMaxPlayers()

	if err != nil {
		return err
	}

	return nil
}

func IsDefusing(entity uintptr) bool {

	isDef, _ := memory.GameProcess.ReadInt(entity + configs.Offsets.Netvars.MBIsDefusing)
	return isDef != 0

}

func GetPlayerHealth(entity uintptr) float32 {

	health, _ := memory.GameProcess.ReadInt(entity + 0x100)
	return float32(health)

}

func GetLocalPlayer() (uintptr, error) {
	player, err := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwLocalPlayer)

	if err != nil {
		return 0, err
	}

	return player, nil
}

func GetMaxPlayers() (int, error) {
	moduleBase, err := memory.GameProcess.ReadIntPtr(memory.GameEngine + configs.Offsets.Signatures.DwClientState)

	if err != nil {
		return 0, err
	}

	return memory.GameProcess.ReadInt(moduleBase + configs.Offsets.Signatures.DwClientStateMaxPlayer)
}

func GetPlayerByIndex(index int) (uintptr, error) {
	player, err := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwEntityList + uintptr(index)*0x10)

	if err != nil {
		return 0, err
	}

	return player, nil
}

func PlayerIsEnemy(entity uintptr) (bool, error) {

	if LocalEntity == 0 {
		return false, nil
	}

	localEntityTeamID, err := memory.GameProcess.ReadInt(LocalEntity + 0xF4)

	if err != nil {
		return false, err
	}

	playerEntityTeamID, err := memory.GameProcess.ReadInt(entity + 0xF4)

	if err != nil {
		return false, err
	}

	if playerEntityTeamID == 0 || localEntityTeamID == 0 {
		return false, nil
	}

	return localEntityTeamID != playerEntityTeamID, nil
}

func PlayerIsValid(entity uintptr) (bool, error) {
	entityLife, err := memory.GameProcess.ReadInt(entity + 0x100)

	if err != nil {
		return false, err
	}

	return entityLife > 0, nil
}

func IsEntityImmune(entity uintptr) (bool, error) {
	isImmune, _ := memory.GameProcess.ReadInt(entity + configs.Offsets.Netvars.MBGunGameImmunity)
	return isImmune == 1, nil
}

func getDistance(entity uintptr, other vector.Vector3) float64 {
	entPosition, _ := memory.GameProcess.ReadVec3(entity + configs.Offsets.Netvars.MVecOrigin)
	delta := other.CalcVector3WithOtherVector3(*entPosition, "-")

	return math.Sqrt(delta.X*delta.X + delta.Y*delta.Y + delta.Z*delta.Z)
}

func GetClosestPlayer() (uintptr, error) {

	if MaxPlayers <= 0 || LocalEntity == 0 {
		return 0, nil
	}

	closestDistance := 1000000.0
	closestEnt := uintptr(0)

	for i := 0; i < MaxPlayers; i++ {
		curEnt, _ := GetPlayerByIndex(i)
		isEnemy, _ := PlayerIsEnemy(curEnt)
		valid, _ := PlayerIsValid(curEnt)

		if curEnt == 0 || !isEnemy || !valid {
			continue
		}

		currentEntityVecOrigin, _ := memory.GameProcess.ReadVec3(curEnt + configs.Offsets.Netvars.MVecOrigin)

		curDistance := getDistance(LocalEntity, *currentEntityVecOrigin)

		if curDistance < closestDistance {
			closestDistance = curDistance
			closestEnt = curEnt
		}

	}

	return closestEnt, nil
}

func GetBonePos(entity uintptr, bone int) (vector.Vector3, error) {
	bonePointer, err := memory.GameProcess.ReadIntPtr(entity + configs.Offsets.Netvars.MDwBoneMatrix)

	if err != nil {
		return vector.Vector3{}, err
	}

	x, err := memory.GameProcess.ReadFloat64(bonePointer + 0x30*uintptr(bone) + 0x0C)

	if err != nil {
		return vector.Vector3{}, err
	}

	y, err := memory.GameProcess.ReadFloat64(bonePointer + 0x30*uintptr(bone) + 0x1C)

	if err != nil {
		return vector.Vector3{}, err
	}

	z, err := memory.GameProcess.ReadFloat64(bonePointer + 0x30*uintptr(bone) + 0x2C)

	if err != nil {
		return vector.Vector3{}, err
	}

	return vector.Vector3{X: x, Y: y, Z: z}, nil
}
