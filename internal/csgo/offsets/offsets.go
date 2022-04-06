package offsets

import (
	"encoding/json"
	"fmt"
	"gosource/internal/global/logs"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const uri string = "https://raw.githubusercontent.com/frk1/hazedumper/master/csgo.min.json"

func InitOffsets(o *GameOffset) {

	logs.Info("getting updated offsets from source ...")
	resp, err := http.Get(uri)
	if err != nil {
		logs.Fatal(errors.Wrap(err, "cannot get updated offsets. aborting ...").Error())
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Fatal(errors.Wrap(err, "cannot read offsets. aborting ...").Error())
	}

	err = json.Unmarshal(b, o)
	if err != nil {
		logs.Fatal(errors.Wrap(err, "could not decode game offsets. aborting ...").Error())
		fmt.Println(err)
		panic("")
	}

}

const (
	// MODULES
	Game   = "csgo.exe"
	Client = "client.dll"
	Engine = "engine.dll"
)

type GameOffset struct {
	Timestamp  int `json:"timestamp"`
	Signatures struct {
		AnimOverlays                   uintptr `json:"anim_overlays"`
		ClientstateChokedCommands      uintptr `json:"clientstate_choked_commands"`
		ClientstateDeltaTicks          uintptr `json:"clientstate_delta_ticks"`
		ClientstateLastOutgoingCommand uintptr `json:"clientstate_last_outgoing_command"`
		ClientstateNetChannel          uintptr `json:"clientstate_net_channel"`
		ConvarNameHashTable            uintptr `json:"convar_name_hash_table"`
		DwClientState                  uintptr `json:"dwClientState"`
		DwClientStateGetLocalPlayer    uintptr `json:"dwClientState_GetLocalPlayer"`
		DwClientStateIsHLTV            uintptr `json:"dwClientState_IsHLTV"`
		DwClientStateMap               uintptr `json:"dwClientState_Map"`
		DwClientStateMapDirectory      uintptr `json:"dwClientState_MapDirectory"`
		DwClientStateMaxPlayer         uintptr `json:"dwClientState_MaxPlayer"`
		DwClientStatePlayerInfo        uintptr `json:"dwClientState_PlayerInfo"`
		DwClientStateState             uintptr `json:"dwClientState_State"`
		DwClientStateViewAngles        uintptr `json:"dwClientState_ViewAngles"`
		DwEntityList                   uintptr `json:"dwEntityList"`
		DwForceAttack                  uintptr `json:"dwForceAttack"`
		DwForceAttack2                 uintptr `json:"dwForceAttack2"`
		DwForceBackward                uintptr `json:"dwForceBackward"`
		DwForceForward                 uintptr `json:"dwForceForward"`
		DwForceJump                    uintptr `json:"dwForceJump"`
		DwForceLeft                    uintptr `json:"dwForceLeft"`
		DwForceRight                   uintptr `json:"dwForceRight"`
		DwGameDir                      uintptr `json:"dwGameDir"`
		DwGameRulesProxy               uintptr `json:"dwGameRulesProxy"`
		DwGetAllClasses                uintptr `json:"dwGetAllClasses"`
		DwGlobalVars                   uintptr `json:"dwGlobalVars"`
		DwGlowObjectManager            uintptr `json:"dwGlowObjectManager"`
		DwInput                        uintptr `json:"dwInput"`
		DwInterfaceLinkList            uintptr `json:"dwInterfaceLinkList"`
		DwLocalPlayer                  uintptr `json:"dwLocalPlayer"`
		DwMouseEnable                  uintptr `json:"dwMouseEnable"`
		DwMouseEnablePtr               uintptr `json:"dwMouseEnablePtr"`
		DwPlayerResource               uintptr `json:"dwPlayerResource"`
		DwRadarBase                    uintptr `json:"dwRadarBase"`
		DwSensitivity                  uintptr `json:"dwSensitivity"`
		DwSensitivityPtr               uintptr `json:"dwSensitivityPtr"`
		DwSetClanTag                   uintptr `json:"dwSetClanTag"`
		DwViewMatrix                   uintptr `json:"dwViewMatrix"`
		DwWeaponTable                  uintptr `json:"dwWeaponTable"`
		DwWeaponTableIndex             uintptr `json:"dwWeaponTableIndex"`
		DwYawPtr                       uintptr `json:"dwYawPtr"`
		DwZoomSensitivityRatioPtr      uintptr `json:"dwZoomSensitivityRatioPtr"`
		DwbSendPackets                 uintptr `json:"dwbSendPackets"`
		DwppDirect3DDevice9            uintptr `json:"dwppDirect3DDevice9"`
		FindHudElement                 uintptr `json:"find_hud_element"`
		ForceUpdateSpectatorGlow       uintptr `json:"force_update_spectator_glow"`
		InterfaceEngineCvar            uintptr `json:"interface_engine_cvar"`
		IsC4Owner                      uintptr `json:"is_c4_owner"`
		MBDormant                      uintptr `json:"m_bDormant"`
		MFlSpawnTime                   uintptr `json:"m_flSpawnTime"`
		MPStudioHdr                    uintptr `json:"m_pStudioHdr"`
		MPitchClassPtr                 uintptr `json:"m_pitchClassPtr"`
		MYawClassPtr                   uintptr `json:"m_yawClassPtr"`
		ModelAmbientMin                uintptr `json:"model_ambient_min"`
		SetAbsAngles                   uintptr `json:"set_abs_angles"`
		SetAbsOrigin                   uintptr `json:"set_abs_origin"`
	} `json:"signatures"`
	Netvars struct {
		CsGamerulesData                uintptr `json:"cs_gamerules_data"`
		MArmorValue                    uintptr `json:"m_ArmorValue"`
		MCollision                     uintptr `json:"m_Collision"`
		MCollisionGroup                uintptr `json:"m_CollisionGroup"`
		MLocal                         uintptr `json:"m_Local"`
		MMoveType                      uintptr `json:"m_MoveType"`
		MOriginalOwnerXuidHigh         uintptr `json:"m_OriginalOwnerXuidHigh"`
		MOriginalOwnerXuidLow          uintptr `json:"m_OriginalOwnerXuidLow"`
		MSurvivalGameRuleDecisionTypes uintptr `json:"m_SurvivalGameRuleDecisionTypes"`
		MSurvivalRules                 uintptr `json:"m_SurvivalRules"`
		MAimPunchAngle                 uintptr `json:"m_aimPunchAngle"`
		MAimPunchAngleVel              uintptr `json:"m_aimPunchAngleVel"`
		MAngEyeAnglesX                 uintptr `json:"m_angEyeAnglesX"`
		MAngEyeAnglesY                 uintptr `json:"m_angEyeAnglesY"`
		MBBombDefused                  uintptr `json:"m_bBombDefused"`
		MBBombPlanted                  uintptr `json:"m_bBombPlanted"`
		MBBombTicking                  uintptr `json:"m_bBombTicking"`
		MBFreezePeriod                 uintptr `json:"m_bFreezePeriod"`
		MBGunGameImmunity              uintptr `json:"m_bGunGameImmunity"`
		MBHasDefuser                   uintptr `json:"m_bHasDefuser"`
		MBHasHelmet                    uintptr `json:"m_bHasHelmet"`
		MBInReload                     uintptr `json:"m_bInReload"`
		MBIsDefusing                   uintptr `json:"m_bIsDefusing"`
		MBIsQueuedMatchmaking          uintptr `json:"m_bIsQueuedMatchmaking"`
		MBIsScoped                     uintptr `json:"m_bIsScoped"`
		MBIsValveDS                    uintptr `json:"m_bIsValveDS"`
		MBSpotted                      uintptr `json:"m_bSpotted"`
		MBSpottedByMask                uintptr `json:"m_bSpottedByMask"`
		MBStartedArming                uintptr `json:"m_bStartedArming"`
		MBUseCustomAutoExposureMax     uintptr `json:"m_bUseCustomAutoExposureMax"`
		MBUseCustomAutoExposureMin     uintptr `json:"m_bUseCustomAutoExposureMin"`
		MBUseCustomBloomScale          uintptr `json:"m_bUseCustomBloomScale"`
		MClrRender                     uintptr `json:"m_clrRender"`
		MDwBoneMatrix                  uintptr `json:"m_dwBoneMatrix"`
		MFAccuracyPenalty              uintptr `json:"m_fAccuracyPenalty"`
		MFFlags                        uintptr `json:"m_fFlags"`
		MFlC4Blow                      uintptr `json:"m_flC4Blow"`
		MFlCustomAutoExposureMax       uintptr `json:"m_flCustomAutoExposureMax"`
		MFlCustomAutoExposureMin       uintptr `json:"m_flCustomAutoExposureMin"`
		MFlCustomBloomScale            uintptr `json:"m_flCustomBloomScale"`
		MFlDefuseCountDown             uintptr `json:"m_flDefuseCountDown"`
		MFlDefuseLength                uintptr `json:"m_flDefuseLength"`
		MFlFallbackWear                uintptr `json:"m_flFallbackWear"`
		MFlFlashDuration               uintptr `json:"m_flFlashDuration"`
		MFlFlashMaxAlpha               uintptr `json:"m_flFlashMaxAlpha"`
		MFlLastBoneSetupTime           uintptr `json:"m_flLastBoneSetupTime"`
		MFlLowerBodyYawTarget          uintptr `json:"m_flLowerBodyYawTarget"`
		MFlNextAttack                  uintptr `json:"m_flNextAttack"`
		MFlNextPrimaryAttack           uintptr `json:"m_flNextPrimaryAttack"`
		MFlSimulationTime              uintptr `json:"m_flSimulationTime"`
		MFlTimerLength                 uintptr `json:"m_flTimerLength"`
		MHActiveWeapon                 uintptr `json:"m_hActiveWeapon"`
		MHBombDefuser                  uintptr `json:"m_hBombDefuser"`
		MHMyWeapons                    uintptr `json:"m_hMyWeapons"`
		MHObserverTarget               uintptr `json:"m_hObserverTarget"`
		MHOwner                        uintptr `json:"m_hOwner"`
		MHOwnerEntity                  uintptr `json:"m_hOwnerEntity"`
		MHViewModel                    uintptr `json:"m_hViewModel"`
		MIAccountID                    uintptr `json:"m_iAccountID"`
		MIClip1                        uintptr `json:"m_iClip1"`
		MICompetitiveRanking           uintptr `json:"m_iCompetitiveRanking"`
		MICompetitiveWins              uintptr `json:"m_iCompetitiveWins"`
		MICrosshairID                  uintptr `json:"m_iCrosshairId"`
		MIDefaultFOV                   uintptr `json:"m_iDefaultFOV"`
		MIEntityQuality                uintptr `json:"m_iEntityQuality"`
		MIFOV                          uintptr `json:"m_iFOV"`
		MIFOVStart                     uintptr `json:"m_iFOVStart"`
		MIGlowIndex                    uintptr `json:"m_iGlowIndex"`
		MIHealth                       uintptr `json:"m_iHealth"`
		MIItemDefinitionIndex          uintptr `json:"m_iItemDefinitionIndex"`
		MIItemIDHigh                   uintptr `json:"m_iItemIDHigh"`
		MIMostRecentModelBoneCounter   uintptr `json:"m_iMostRecentModelBoneCounter"`
		MIObserverMode                 uintptr `json:"m_iObserverMode"`
		MIShotsFired                   uintptr `json:"m_iShotsFired"`
		MIState                        uintptr `json:"m_iState"`
		MITeamNum                      uintptr `json:"m_iTeamNum"`
		MLifeState                     uintptr `json:"m_lifeState"`
		MNBombSite                     uintptr `json:"m_nBombSite"`
		MNFallbackPaintKit             uintptr `json:"m_nFallbackPaintKit"`
		MNFallbackSeed                 uintptr `json:"m_nFallbackSeed"`
		MNFallbackStatTrak             uintptr `json:"m_nFallbackStatTrak"`
		MNForceBone                    uintptr `json:"m_nForceBone"`
		MNTickBase                     uintptr `json:"m_nTickBase"`
		MNViewModelIndex               uintptr `json:"m_nViewModelIndex"`
		MRgflCoordinateFrame           uintptr `json:"m_rgflCoordinateFrame"`
		MSzCustomName                  uintptr `json:"m_szCustomName"`
		MSzLastPlaceName               uintptr `json:"m_szLastPlaceName"`
		MThirdPersonViewAngles         uintptr `json:"m_thirdPersonViewAngles"`
		MVecOrigin                     uintptr `json:"m_vecOrigin"`
		MVecVelocity                   uintptr `json:"m_vecVelocity"`
		MVecViewOffset                 uintptr `json:"m_vecViewOffset"`
		MViewPunchAngle                uintptr `json:"m_viewPunchAngle"`
		MZoomLevel                     uintptr `json:"m_zoomLevel"`
	} `json:"netvars"`
}
