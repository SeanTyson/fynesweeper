package game

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/AnkushJadhav/fynesweeper/components"
)

// Game is a game damnit
type Game struct {
	Board       *fyne.Container
	Tiles       [][]*components.Tile
	Smiley      *components.SmileyMan
	MineCounter *components.MineCounter
	TimeCounter *components.TimeCounter

	OpenCount int
	WinCount  int
	IsRunning bool

	Win fyne.Window
}

// NewGame creates a new game
func NewGame() *Game {
	return &Game{}
}

// SeedGame creates a new game damnit
func (g *Game) SeedGame(rows, cols, mineCount int) {
	plan := make([][]int, 0)

	// generate a blank plan
	for itrRow := 0; itrRow < rows; itrRow++ {
		row := make([]int, 0)
		for itrCol := 0; itrCol < cols; itrCol++ {
			row = append(row, 0)
		}
		plan = append(plan, row)
	}

	// generate mines with surrounding info
	plan = generatePlan(plan, mineCount)

	// generate game tiles based on plan
	tiles := make([][]*components.Tile, 0)
	for itrRow := 0; itrRow < rows; itrRow++ {
		row := make([]*components.Tile, 0)
		for itrCol := 0; itrCol < cols; itrCol++ {
			var tile *components.Tile
			switch plan[itrRow][itrCol] {
			case -1:
				tile = components.NewTile(components.TileTypeMine, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 0:
				tile = components.NewTile(components.TileType0, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 1:
				tile = components.NewTile(components.TileType1, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 2:
				tile = components.NewTile(components.TileType2, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 3:
				tile = components.NewTile(components.TileType3, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 4:
				tile = components.NewTile(components.TileType4, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 5:
				tile = components.NewTile(components.TileType5, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 6:
				tile = components.NewTile(components.TileType6, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 7:
				tile = components.NewTile(components.TileType7, itrRow, itrCol, g.openTile, g.markTile)
				break
			case 8:
				tile = components.NewTile(components.TileType8, itrRow, itrCol, g.openTile, g.markTile)
				break
			default:
			}

			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	board := components.NewBoard(tiles)
	sm := components.NewSmileyMan(components.GameStateOngoing, g.resetHandler)
	mc := components.NewMineCounter(mineCount)
	tc := components.NewTimeCounter(0)

	g.Board = board
	g.Tiles = tiles
	g.Smiley = sm
	g.MineCounter = mc
	g.TimeCounter = tc
	g.OpenCount = 0
	g.IsRunning = true
	g.WinCount = (rows * cols) - mineCount
}

func generatePlan(plan [][]int, mineCount int) [][]int {
	rand.Seed(time.Now().UnixNano())
	maxRow := len(plan) - 1
	maxCol := len(plan[0]) - 1
	itr := 0
	for {
		if itr == mineCount {
			break
		}

		r := rand.Intn(maxRow + 1)
		c := rand.Intn(maxCol + 1)
		if isMine(plan, r, c) {
			continue
		}
		plan[r][c] = -1
		if isInBounds(r-1, c-1, 0, 0, maxRow, maxCol) && !isMine(plan, r-1, c-1) {
			plan[r-1][c-1]++
		}
		if isInBounds(r-1, c, 0, 0, maxRow, maxCol) && !isMine(plan, r-1, c) {
			plan[r-1][c]++
		}
		if isInBounds(r-1, c+1, 0, 0, maxRow, maxCol) && !isMine(plan, r-1, c+1) {
			plan[r-1][c+1]++
		}
		if isInBounds(r, c-1, 0, 0, maxRow, maxCol) && !isMine(plan, r, c-1) {
			plan[r][c-1]++
		}
		if isInBounds(r, c+1, 0, 0, maxRow, maxCol) && !isMine(plan, r, c+1) {
			plan[r][c+1]++
		}
		if isInBounds(r+1, c-1, 0, 0, maxRow, maxCol) && !isMine(plan, r+1, c-1) {
			plan[r+1][c-1]++
		}
		if isInBounds(r+1, c, 0, 0, maxRow, maxCol) && !isMine(plan, r+1, c) {
			plan[r+1][c]++
		}
		if isInBounds(r+1, c+1, 0, 0, maxRow, maxCol) && !isMine(plan, r+1, c+1) {
			plan[r+1][c+1]++
		}

		itr++
	}

	return plan
}

// Render the game
func (g *Game) Render() {
	t := container.NewVBox(container.NewBorder(nil, nil, g.MineCounter.Container, g.TimeCounter.Container, g.Smiley), g.Board)

	g.Win.SetContent(t)
}

func (g *Game) resetHandler() {
	g.SeedGame(20, 20, 1)
	g.Render()
}

func isInBounds(row, col, minRow, minCol, maxRow, maxCol int) bool {
	return row >= minRow && row <= maxRow && col >= minCol && col <= maxCol
}

func isMine(plan [][]int, row, col int) bool {
	return plan[row][col] == -1
}
