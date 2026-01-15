package main

func bubbleSort(game *Game) {
	if game.i < len(game.numbers) {
		if game.j < len(game.numbers)-1 {
			if game.numbers[game.j] > game.numbers[game.j+1] {
				game.numbers[game.j], game.numbers[game.j+1] = game.numbers[game.j+1], game.numbers[game.j]
			}
			game.j++
		} else {
			game.j = 0
			game.i++
		}
	} else {
		game.sortingFinished = true
		game.controlButton.Text = "Start"
	}
}
