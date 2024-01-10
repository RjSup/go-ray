package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	//* constand screen width and heights
	screenWidth  = 1000
	screenHeight = 480
)

var (
	//* boolean defining if the game is running or not
	running = true

	//* defining a new colour
	bkgColor = rl.NewColor(147, 211, 196, 255)

	//* grass tiles
	grassSprite rl.Texture2D
	//* player sprite
	playerSprite rl.Texture2D

	//* the actual player image in a rectangle to show only one frame at a time
	playerSrc  rl.Rectangle
	playerDest rl.Rectangle
	//* for player animation
	playerMoving                                  bool
	playerDirection                               int
	playerUp, playerDown, playerLeft, playerRight bool
	//* referencing which tile in the sprite sheet its on
	playerFrame int

	//*
	frameCount int

	//* speed the player is moving
	playerSpeed float32 = 3

	//* initialise music features
	musicPaused bool
	music       rl.Music

	//* camera
	cam rl.Camera2D
)

func drawScene() {
	//* draw the grass texture to the window -
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	//* draw the player sprite to the window -
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

//* controls for player movement
func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerDirection = 1
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDirection = 2
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDirection = 0
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDirection = 3
		playerRight = true
	}
	//* start and stop music
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func update() {
	//* running = true
	running = !rl.WindowShouldClose()

	//* if the key is pressd and the player is moving is true then move the player
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
		//* move along the frames of the sprite sheet
		if frameCount%8 == 1 {
			playerFrame++
		}
	}
	//* if player is moving update every frame and every player frame to move accross the sprite sheet
	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}

	playerSrc.X = playerSrc.Width * float32(playerFrame)
	playerSrc.Y = playerSrc.Height * float32(playerDirection)

	//* update music
	rl.UpdateMusicStream(music)

	//* pause music
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	//* camera follows the player
	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)))

	playerMoving = false
	playerUp, playerDown, playerLeft, playerRight = false, false, false, false
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)

	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()

	rl.EndDrawing()
}

func init() {
	//* initialise the window for game
	rl.InitWindow(screenWidth, screenHeight, "Animal Crossing Clone")
	//* sets esc not to close screen
	rl.SetExitKey(0)
	//* traget fps
	rl.SetTargetFPS(60)

	//* initialise timemaps into the window
	grassSprite = rl.LoadTexture("resource/Tilesets/Grass.png")
	//* initialise player sprite into the window
	playerSprite = rl.LoadTexture("resource/Characters/Char.png")

	// size of image
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	//* scaled up
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	//* initialise music player
	rl.InitAudioDevice()
	//* load the music file
	music = rl.LoadMusicStream("resource/music/Childhood_Memories.mp3")
	//* music running
	musicPaused = false
	//* play the file
	rl.PlayMusicStream(music)

	//* initialise the camera
	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2),
		float32(screenHeight/2)),
		rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)),
			float32(playerDest.Y-(playerDest.Height/2))), 0.0, 1.5)
}

func quit() {
	//* used to remove the memory allocations on quit
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
func main() {

	//* while running = true
	for running {
		//* allow the player to move
		input()
		update()
		//* update the window with textures and what not
		render()
	}

	quit()
}
