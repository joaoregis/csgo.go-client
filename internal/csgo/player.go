package csgo

import (
	"gosource/internal/csgo/sdk"
	"gosource/internal/global/configs"
	"gosource/internal/hackFunctions/vector"
	"gosource/internal/memory"
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

func GetPlayerByIndex(index int) (uintptr, error) {
	player, err := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwEntityList + uintptr(index)*0x10)

	if err != nil {
		return 0, err
	}

	return player, nil
}

func PlayerIsLocalEntity(entity uintptr) bool {
	dwLocalEntity, _ := GetLocalPlayer()
	return entity == dwLocalEntity
}

func GetPlayerTeam(entity uintptr) sdk.CSTeam {

	playerEntityTeamID, _ := memory.GameProcess.ReadInt(entity + 0xF4)
	return sdk.CSTeam(playerEntityTeamID)

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
	entityLife, err := memory.GameProcess.ReadInt(entity + configs.Offsets.Netvars.MIHealth)

	if err != nil {
		return false, err
	}

	return entityLife > 0, nil
}

func PlayerIsDormant(entity uintptr) (bool, error) {
	bDormant, err := memory.GameProcess.ReadInt(entity + configs.Offsets.Signatures.MBDormant)

	if err != nil {
		return false, err
	}

	return bDormant == 1, nil
}

func GetEntity2DPos(entity uintptr) *vector.Vector2 {

	entityLocation, _ := memory.GameProcess.ReadVec3(entity + configs.Offsets.Netvars.MVecOrigin)
	var screenCoordinates vector.Vector2
	if !WorldToScreen(entityLocation, &screenCoordinates) {
		// if couldn't convert world to screen, ignores
		return nil
	}

	return &screenCoordinates
}

func IsEntityImmune(entity uintptr) (bool, error) {
	isImmune, _ := memory.GameProcess.ReadInt(entity + configs.Offsets.Netvars.MBGunGameImmunity)
	return isImmune == 1, nil
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

		curDistance := GetDistance(LocalEntity, *currentEntityVecOrigin)

		if curDistance < closestDistance {
			closestDistance = curDistance
			closestEnt = curEnt
		}

	}

	return closestEnt, nil
}

func GetEntityVecOrigin(entity uintptr) *vector.Vector3 {

	currentEntityVecOrigin, _ := memory.GameProcess.ReadVec3(entity + configs.Offsets.Netvars.MVecOrigin)
	return currentEntityVecOrigin

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

var UserInfoTableItems uintptr

func LoadUserInfoTable() {

	clientState, _ := memory.GameProcess.ReadIntPtr(memory.GameEngine + configs.Offsets.Signatures.DwClientState)
	userInfoTable, _ := memory.GameProcess.ReadIntPtr(clientState + configs.Offsets.Signatures.DwClientStatePlayerInfo)
	userInfoTable, _ = memory.GameProcess.ReadIntPtr(userInfoTable + 0x40)
	UserInfoTableItems, _ = memory.GameProcess.ReadIntPtr(userInfoTable + 0xC)

}

func GetPlayerInfo(entity uintptr) *sdk.PlayerInfo_T {

	entityIndex, _ := memory.GameProcess.ReadInt(entity + 0x64)
	dwPlayerInfo, _ := memory.GameProcess.ReadIntPtr((UserInfoTableItems + 0x28) + (uintptr((entityIndex - 1) * 0x34)))
	pInfo, _ := memory.Read[sdk.PlayerInfo_T](&memory.GameProcess, dwPlayerInfo)
	return pInfo
}

func GetPlayerName(entity uintptr) string {
	pInfo := GetPlayerInfo(entity)
	return string(pInfo.SzPlayerName[:])
}

func GetRadarPlayer(entityIndex uintptr) *sdk.RadarPlayer_T {

	dwRadarBase, _ := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwRadarBase)
	dwRadarBase_CSGOHudRadar, _ := memory.GameProcess.ReadIntPtr(dwRadarBase + 0x78)
	radarPlayer, _ := memory.Read[sdk.RadarPlayer_T](&memory.GameProcess, dwRadarBase_CSGOHudRadar+(0x174*(entityIndex+2))-0x18)
	return radarPlayer

}

func GetEntityGlow(entity uintptr) *sdk.EntityGlowStruct {

	dwGlowObjectManager, _ := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwGlowObjectManager)
	iGlowIndex, _ := memory.GameProcess.ReadIntPtr(entity + configs.Offsets.Netvars.MIGlowIndex)
	v, _ := memory.Read[sdk.EntityGlowStruct](&memory.GameProcess, dwGlowObjectManager+(iGlowIndex*0x38)+0x8)
	return v

}

func SetEntityGlow(v *sdk.EntityGlowStruct, entity uintptr, value int) {

	dwGlowObjectManager, _ := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwGlowObjectManager)
	iGlowIndex, _ := memory.GameProcess.ReadIntPtr(entity + configs.Offsets.Netvars.MIGlowIndex)
	memory.Write(&memory.GameProcess, dwGlowObjectManager+(iGlowIndex*0x38)+0x8, v)
	memory.GameProcess.WriteInt(dwGlowObjectManager+iGlowIndex*0x38+0x27, value) // Enables Glow at Entity
	memory.GameProcess.WriteInt(dwGlowObjectManager+iGlowIndex*0x38+0x28, value) // Enables Glow at Entity

}
