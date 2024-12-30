package server

func Checker(gameMap [][]string) string {

	if gameMap[0][0] == gameMap[0][1] && gameMap[0][1] == gameMap[0][2] { // верхняя строка
		return gameMap[0][0]
	}

	if gameMap[1][0] == gameMap[1][1] && gameMap[1][1] == gameMap[1][2] { // средняя строка
		return gameMap[1][0]
	}

	if gameMap[2][0] == gameMap[2][1] && gameMap[2][1] == gameMap[2][2] { // нижняя строка
		return gameMap[2][0]
	}

	if gameMap[0][0] == gameMap[1][1] && gameMap[1][1] == gameMap[2][2] { // диагоняль лево верх - право низ
		return gameMap[0][0]
	}

	if gameMap[0][2] == gameMap[1][1] && gameMap[2][0] == gameMap[1][1] { // диагональ право верх - лево низ
		return gameMap[0][2]
	}

	if gameMap[0][0] == gameMap[1][0] && gameMap[1][0] == gameMap[2][0] { // вертикаль левая
		return gameMap[0][0]
	}

	if gameMap[0][1] == gameMap[1][1] && gameMap[2][1] == gameMap[1][1] { // вертикаль средняя
		return gameMap[1][1]
	}

	if gameMap[0][2] == gameMap[1][2] && gameMap[2][2] == gameMap[0][2] { // вертикаль правая
		return gameMap[0][2]
	}

	return ""
}

func PlaceCreater() [][]string {
	return [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
}
