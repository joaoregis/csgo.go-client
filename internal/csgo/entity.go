package csgo

import (
	"gosource/internal/configs"
	"gosource/internal/csgo/weapons"
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

func GetActiveWeapon() (int, error) {

	dwLocalPlayer, _ := GetLocalPlayer()
	hActiveWeapon, _ := memory.GameProcess.ReadIntPtr(dwLocalPlayer + configs.Offsets.Netvars.MHActiveWeapon)
	dwWeaponEntity, _ := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwEntityList + ((hActiveWeapon&0xFFF)-1)*0x10)
	nWeaponIndex, _ := memory.GameProcess.ReadInt(dwWeaponEntity + configs.Offsets.Netvars.MIItemDefinitionIndex)
	return nWeaponIndex & 0xFFF, nil

}

func IsC4Item(wIdx int) bool {

	return wIdx == weapons.C4

}

func IsGranade(wIdx int) bool {

	var items []int = []int{
		weapons.Flashbang,
		weapons.HeGrenade,
		weapons.SmokeGrenade,
		weapons.Molotov,
		weapons.Decoy,
		weapons.IncGrenade,
	}

	return contains(items, wIdx)

}

func IsAimbotable(wIdx int) bool {

	var items []int = []int{
		weapons.Usp_s,
		weapons.Glock,
		weapons.Fiveseven,
		weapons.Cz75a,
		weapons.P250,
		weapons.Hkp2000,
		weapons.Tec9,
		weapons.Deagle,
		weapons.Revolver,
		weapons.Awp,
		weapons.Ssg08,
		weapons.Mac10,
		weapons.P90,
		weapons.Mp5sd,
		weapons.Ump45,
		weapons.Bizon,
		weapons.Mp7,
		weapons.Mp9,
		weapons.Ak47,
		weapons.M4a1_s,
		weapons.M4A4,
		weapons.Famas,
		weapons.Aug,
		weapons.GalilAr,
		weapons.Sg553,
		weapons.M249,
		weapons.Xm1014,
		weapons.Mag7,
		weapons.Negev,
		weapons.Sawedoff,
		weapons.Nova,
	}

	return contains(items, wIdx)

}

func IsPistol(wIdx int) bool {

	var items []int = []int{
		weapons.Usp_s,
		weapons.Glock,
		weapons.Fiveseven,
		weapons.Cz75a,
		weapons.P250,
		weapons.Hkp2000,
		weapons.Tec9,
	}

	return contains(items, wIdx)

}

func IsBoltSniper(wIdx int) bool {

	var items []int = []int{
		weapons.Awp,
		weapons.Ssg08,
	}

	return contains(items, wIdx)

}

func IsSubmachineGun(wIdx int) bool {

	var items []int = []int{
		weapons.Mac10,
		weapons.P90,
		weapons.Mp5sd,
		weapons.Ump45,
		weapons.Bizon,
		weapons.Mp7,
		weapons.Mp9,
	}

	return contains(items, wIdx)

}

func IsAutoSniper(wIdx int) bool {

	var items []int = []int{
		weapons.Scar20,
		weapons.G3SG1,
	}

	return contains(items, wIdx)

}

func IsSniper(wIdx int) bool {

	var items []int = []int{
		weapons.Awp,
		weapons.Ssg08,
		weapons.Scar20,
		weapons.G3SG1,
	}

	return contains(items, wIdx)

}

func IsHeavyPistol(wIdx int) bool {

	var items []int = []int{
		weapons.Deagle,
		weapons.Revolver,
	}

	return contains(items, wIdx)

}

func IsDeagle(wIdx int) bool {

	return wIdx == weapons.Deagle

}

func IsRevolver(wIdx int) bool {

	return wIdx == weapons.Revolver

}

func IsTaser(wIdx int) bool {

	return wIdx == weapons.Taser

}

func IsKnife(wIdx int) bool {

	isInKnifeModelRange := wIdx >= weapons.KnifeBayonet && wIdx <= weapons.KnifeSkeleton

	var items []int = []int{
		weapons.KnifeCT,
		weapons.KnifeT,
		weapons.GhostKnife,
		weapons.GoldenKnife,
	}

	return contains(items, wIdx) || isInKnifeModelRange

}

func IsRifle(wIdx int) bool {

	var items []int = []int{
		weapons.Ak47,
		weapons.M4a1_s,
		weapons.M4A4,
		weapons.Famas,
		weapons.Aug,
		weapons.GalilAr,
		weapons.Sg553,
	}

	return contains(items, wIdx)

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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
