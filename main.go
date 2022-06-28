package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 480
	fps          = 60
)

var (
	running  = true
	bkgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite  rl.Texture2D
	hillSprite   rl.Texture2D
	fenceSprite  rl.Texture2D
	houseSprite  rl.Texture2D
	waterSprite  rl.Texture2D
	tilledSprite rl.Texture2D
	tex          rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc                                     rl.Rectangle
	playerDest                                    rl.Rectangle
	playerMoving                                  bool
	playerDir                                     int
	playerUp, playerDown, playerRight, playerLeft bool
	playerSpeed                                   float32 = 3.0
	playerFrame                                   int

	frameCount int

	tileDest  rl.Rectangle
	tileSrc   rl.Rectangle
	tileMap   []int
	srcMap    []string
	mapWidth  int
	mapHeight int

	musicPaused bool
	music       rl.Music

	cam rl.Camera2D
)

func drawScene() {
	// rl.DrawTexture(grassSprite, 100, 50, rl.White)
	for i := 0; i < len(tileMap); i++ {
		if tileMap[i] != 0 {
			tileDest.X = tileDest.Width * float32(i%mapWidth)
			tileDest.Y = tileDest.Height * float32(i/mapWidth)

			switch srcMap[i] {
			case "g":
				tex = grassSprite
			case "l":
				tex = hillSprite
			case "f":
				tex = fenceSprite
			case "h":
				tex = houseSprite
			case "w":
				tex = waterSprite
			case "t":
				tex = tilledSprite
			}
			// you need to draw grass below a fence or house part first
			if srcMap[i] == "f" || srcMap[i] == "h" {
				tileSrc.X = 0
				tileSrc.Y = 0
				rl.DrawTexturePro(grassSprite, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White)
			}

			tileSrc.X = tileSrc.Width * float32((tileMap[i]-1)%int(tex.Width/int32(tileSrc.Width)))
			tileSrc.Y = tileSrc.Height * float32((tileMap[i]-1)/int(tex.Width/int32(tileSrc.Width)))

			rl.DrawTexturePro(tex, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White)
		}
	}

	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerDir = 1
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDir = 0
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDir = 2
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDir = 3
		playerRight = true
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func update() {
	running = !rl.WindowShouldClose()

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed
		}
		if playerDown {
			playerDest.Y += playerSpeed
		}
		if playerLeft {
			playerDest.X -= playerSpeed
		}
		if playerRight {
			playerDest.X += playerSpeed
		}
		if frameCount%8 == 1 {
			playerFrame++
		}
	} else if frameCount%45 == 1 {
		playerFrame++
	}

	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}
	if !playerMoving && playerFrame > 1 {
		playerFrame = 0
	}

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	playerSrc.Y = playerSrc.Height * float32(playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)))

	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false

}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)

	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func loadMap(mapFile string) {
	file, err := ioutil.ReadFile(mapFile)
	if err != nil {
		fmt.Errorf("Map File did not load: %v", err)
		os.Exit(1)
	}
	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")
	mapWidth = -1
	mapHeight = -1
	for i, me := range sliced {
		s, _ := strconv.ParseInt(me, 10, 64)
		m := int(s)
		if mapWidth == -1 {
			mapWidth = m
		} else if mapHeight == -1 {
			mapHeight = m
		} else if i < mapHeight*mapWidth+2 {
			tileMap = append(tileMap, m)
		} else {
			srcMap = append(srcMap, sliced[i])
		}
	}
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sproutlings!")
	rl.SetExitKey(0)
	rl.SetTargetFPS(fps)

	grassSprite = rl.LoadTexture("res/SproutLands/Tilesets/Grass.png")
	hillSprite = rl.LoadTexture("res/SproutLands/Tilesets/Hills.png")
	fenceSprite = rl.LoadTexture("res/SproutLands/Tilesets/Building Parts/Fences.png")
	houseSprite = rl.LoadTexture("res/SproutLands/Tilesets/Building Parts/Wooden House.png")
	waterSprite = rl.LoadTexture("res/SproutLands/Tilesets/Water.png")
	tilledSprite = rl.LoadTexture("res/SproutLands/Tilesets/Tilled Dirt.png")
	tileDest = rl.NewRectangle(0, 0, 16, 16)
	tileSrc = rl.NewRectangle(0, 0, 16, 16)

	playerSprite = rl.LoadTexture("res/SproutLands/characters/Basic Charakter Spritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/sproutlings_loopable.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))),
		0, 1.5)

	loadMap("one.map")
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}

	quit()
}
