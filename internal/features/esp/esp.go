package esp

import (
	"gosource/internal/csgo"
	"gosource/internal/global/configs"
	"gosource/internal/hackFunctions/color"
	"gosource/internal/hackFunctions/vector"
	"math"

	math2 "github.com/google/gxui/math"
)

func renderSnapLines(entity uintptr) {

	if entity2DPosition := csgo.GetEntity2DPos(entity); entity2DPosition != nil {
		DrawLine(lineOrigin, *entity2DPosition, 1, &colorA)
	}

}

type BOUNDING_BOX_LAYOUT int

const (
	L_2D        BOUNDING_BOX_LAYOUT = 0
	L_3D        BOUNDING_BOX_LAYOUT = 1
	L_2D_CORNER BOUNDING_BOX_LAYOUT = 2
)

func renderBox2d(bottom vector.Vector2, top vector.Vector2, cfg configs.ConfigDataESPBoundingBox) {

	if BOUNDING_BOX_LAYOUT(cfg.Layout) == L_2D {

		H := math.Abs(top.Y - bottom.Y)
		var topLeft, topRight, bottomLeft, bottomRight vector.Vector2

		topLeft.X = top.X - H/4
		topRight.X = top.X + H/4
		topLeft.Y = top.Y
		topRight.Y = top.Y
		bottomLeft.X = top.X - H/4
		bottomRight.X = top.X + H/4
		bottomLeft.Y = bottom.Y
		bottomRight.Y = bottom.Y

		if cfg.FullfillBox {
			cFullfillBox := color.HexToRGBA(color.Hex(cfg.FullfillBoxColor), &cfg.FullfillBoxColorAlpha)
			DrawFilledRect(topLeft, bottomRight, cFullfillBox)
		}

		cBoxLines := color.HexToRGBA(color.Hex(cfg.Color), &cfg.ColorAlpha)
		DrawLine(topLeft, topRight, cfg.Thickness, cBoxLines)
		DrawLine(bottomLeft, bottomRight, cfg.Thickness, cBoxLines)
		DrawLine(topLeft, bottomLeft, cfg.Thickness, cBoxLines)
		DrawLine(topRight, bottomRight, cfg.Thickness, cBoxLines)

		if cfg.Outline {

			topLeft.X -= .002
			topRight.X += .002
			topRight.Y += .005
			topLeft.Y += .005

			bottomLeft.X -= .002
			bottomRight.X += .002
			bottomRight.Y -= .005
			bottomLeft.Y -= .005

			// draw outline for 2d box
			cBoxLines := color.HexToRGBA(color.Hex(cfg.OutlineColor), &cfg.ColorAlpha)
			DrawLine(topLeft, topRight, cfg.Thickness*.5, cBoxLines)
			DrawLine(bottomLeft, bottomRight, cfg.Thickness*.5, cBoxLines)
			DrawLine(topLeft, bottomLeft, cfg.Thickness*.5, cBoxLines)
			DrawLine(topRight, bottomRight, cfg.Thickness*.5, cBoxLines)

		}

	}

}

func renderHealth(bottom vector.Vector2, top vector.Vector2, entity uintptr, alpha float32) {

	height := float64(math.Abs(top.Y - bottom.Y))

	playerHealth := csgo.GetPlayerHealth(entity)
	healthPerc := float64(playerHealth / 100.0)

	var topLeft, bottomRight, topLeftHealth vector.Vector2

	healthHeight := height * healthPerc

	topLeft.X = (top.X - height/4) - 0.02
	topLeftHealth.X = topLeft.X

	topLeft.Y = top.Y - height
	topLeftHealth.Y = topLeft.Y + healthHeight

	bottomRight.X = topLeft.X + 0.01
	bottomRight.Y = bottom.Y

	c := float32((math2.Lerpf(0, 1, playerHealth/100)))
	rgba := color.NewRGBA(1-c, c, 0, alpha)

	// DrawFilledRect(topLeft, bottomRight, color.NewRGBA(0, 0, 0, .5))
	DrawFilledRect(topLeftHealth, bottomRight, rgba)

}

func renderName(bottom vector.Vector2, top vector.Vector2, entity uintptr, entityIndex int, cfg configs.ConfigDataESPBoundingBox) {

	if csgo.PlayerIsLocalEntity(entity) {
		return
	}

	dwLocalPlayer, _ := csgo.GetLocalPlayer()
	distanceToTargetEntity := csgo.GetDistance(dwLocalPlayer, *csgo.GetEntityVecOrigin(entity))

	var offsetName float64

	if distanceToTargetEntity < 400 {
		offsetName = distanceToTargetEntity / 100000
	} else if distanceToTargetEntity < 600 {
		offsetName = distanceToTargetEntity / 10000
	} else {
		offsetName = distanceToTargetEntity / 100000
	}

	espColor := color.HexToRGBA(color.Hex(cfg.Color), nil)
	entityName := csgo.GetPlayerName(entity)
	pNamePos := vector.Vector2{X: top.X, Y: top.Y}

	pNamePos.X -= offsetName
	pNamePos.Y += +0.065

	DrawStringf(pNamePos, espColor, entityName)

}

func Esp(entity uintptr, entityIndex int) {

	if !ValidatePlayerESP(entity) {
		return
	}

	entityHead3D, _ := csgo.GetBonePos(entity, 8)
	entityHead3D.Z += 8 // fix box height
	var entityHead2D vector.Vector2
	if !csgo.WorldToScreen(&entityHead3D, &entityHead2D) {
		return
	}

	var entityPosition2D *vector.Vector2
	if entityPosition2D = csgo.GetEntity2DPos(entity); entityPosition2D == nil {
		return
	}

	/*
	* [[ --
	*
	* ESP FEATURES
	*
	* -- ]]
	 */

	/***** SNAPLINES *****/
	if configs.G.D.ESP.DrawSnapLines {
		renderSnapLines(entity)
	}

	if isEnemy, _ := csgo.PlayerIsEnemy(entity); isEnemy && configs.G.D.ESP.EnemyBoundingBox.Enabled {

		/***** ENEMY BOUNDING BOXES *****/
		if configs.G.D.ESP.EnemyBoundingBox.DrawBox {
			renderBox2d(*entityPosition2D, entityHead2D, configs.G.D.ESP.EnemyBoundingBox)
		}

		/***** ENEMY HEALTHBAR *****/
		if configs.G.D.ESP.EnemyBoundingBox.DrawHealth {
			renderHealth(*entityPosition2D, entityHead2D, entity, configs.G.D.ESP.EnemyBoundingBox.ColorAlpha)
		}

		/***** ENEMY NAME *****/
		if configs.G.D.ESP.EnemyBoundingBox.DrawName {
			renderName(*entityPosition2D, entityHead2D, entity, entityIndex, configs.G.D.ESP.EnemyBoundingBox)
		}

	} else if configs.G.D.ESP.AllyBoundingBox.Enabled {

		/***** ALLY BOXES *****/
		if configs.G.D.ESP.AllyBoundingBox.DrawBox {
			renderBox2d(*entityPosition2D, entityHead2D, configs.G.D.ESP.AllyBoundingBox)
		}

		/***** ALLY HEALTHBAR *****/
		if configs.G.D.ESP.AllyBoundingBox.DrawHealth {
			renderHealth(*entityPosition2D, entityHead2D, entity, configs.G.D.ESP.AllyBoundingBox.ColorAlpha)
		}

		/***** ALLY NAME *****/
		if configs.G.D.ESP.AllyBoundingBox.DrawName {
			renderName(*entityPosition2D, entityHead2D, entity, entityIndex, configs.G.D.ESP.EnemyBoundingBox)
		}

	}

}
