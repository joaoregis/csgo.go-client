package csgo

import (
	"gosource/internal/global/configs"
	"gosource/internal/global/utils"
	"gosource/internal/memory"
)

func GetActiveWeapon() (int, error) {

	dwLocalPlayer, _ := GetLocalPlayer()
	hActiveWeapon, _ := memory.GameProcess.ReadIntPtr(dwLocalPlayer + configs.Offsets.Netvars.MHActiveWeapon)
	dwWeaponEntity, _ := memory.GameProcess.ReadIntPtr(memory.GameClient + configs.Offsets.Signatures.DwEntityList + ((hActiveWeapon&0xFFF)-1)*0x10)
	nWeaponIndex, _ := memory.GameProcess.ReadInt(dwWeaponEntity + configs.Offsets.Netvars.MIItemDefinitionIndex)
	return nWeaponIndex & 0xFFF, nil

}

func IsC4Item(wIdx int) bool {

	return wIdx == C4

}

func IsGranade(wIdx int) bool {

	var items []int = []int{
		Flashbang,
		HeGrenade,
		SmokeGrenade,
		Molotov,
		Decoy,
		IncGrenade,
	}

	return utils.Contains(items, wIdx)

}

func IsAimbotable(wIdx int) bool {

	var items []int = []int{
		Usp_s,
		Glock,
		Fiveseven,
		Cz75a,
		P250,
		Hkp2000,
		Tec9,
		Deagle,
		Revolver,
		Awp,
		Ssg08,
		Mac10,
		P90,
		Mp5sd,
		Ump45,
		Bizon,
		Mp7,
		Mp9,
		Ak47,
		M4a1_s,
		M4A4,
		Famas,
		Aug,
		GalilAr,
		Sg553,
		M249,
		Xm1014,
		Mag7,
		Negev,
		Sawedoff,
		Nova,
		Scar20,
		G3SG1,
	}

	return utils.Contains(items, wIdx)

}

func IsPistol(wIdx int) bool {

	var items []int = []int{
		Usp_s,
		Glock,
		Fiveseven,
		Cz75a,
		P250,
		Hkp2000,
		Tec9,
	}

	return utils.Contains(items, wIdx)

}

func IsBoltSniper(wIdx int) bool {

	var items []int = []int{
		Awp,
		Ssg08,
	}

	return utils.Contains(items, wIdx)

}

func IsSubmachineGun(wIdx int) bool {

	var items []int = []int{
		Mac10,
		P90,
		Mp5sd,
		Ump45,
		Bizon,
		Mp7,
		Mp9,
	}

	return utils.Contains(items, wIdx)

}

func IsAutoSniper(wIdx int) bool {

	var items []int = []int{
		Scar20,
		G3SG1,
	}

	return utils.Contains(items, wIdx)

}

func IsSniper(wIdx int) bool {

	var items []int = []int{
		Awp,
		Ssg08,
		Scar20,
		G3SG1,
	}

	return utils.Contains(items, wIdx)

}

func IsHeavyPistol(wIdx int) bool {

	var items []int = []int{
		Deagle,
		Revolver,
	}

	return utils.Contains(items, wIdx)

}

func IsDeagle(wIdx int) bool {

	return wIdx == Deagle

}

func IsRevolver(wIdx int) bool {

	return wIdx == Revolver

}

func IsTaser(wIdx int) bool {

	return wIdx == Taser

}

func IsKnife(wIdx int) bool {

	isInKnifeModelRange := wIdx >= KnifeBayonet && wIdx <= KnifeSkeleton

	var items []int = []int{
		KnifeCT,
		KnifeT,
		GhostKnife,
		GoldenKnife,
	}

	return utils.Contains(items, wIdx) || isInKnifeModelRange

}

func IsRifle(wIdx int) bool {

	var items []int = []int{
		Ak47,
		M4a1_s,
		M4A4,
		Famas,
		Aug,
		GalilAr,
		Sg553,
	}

	return utils.Contains(items, wIdx)

}
