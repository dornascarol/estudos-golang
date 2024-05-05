// Rock Paper Scissors!
// Let's play! You have to return which player won! In case of a draw return Draw!.

// Lógica:
// 01) Caso de empate vai ser quando o p1 = p2 => return "Draw!"
// 02) Caso do p1 ganhar => pedra ganha da tesoura | tesoura ganha do papel | papel ganha da pedra
// 03) Caso do p2 ganhar => pedra ganha da tesoura | tesoura ganha do papel | papel ganha da pedra
// 04) O que "sobrar" é o ganhador p2.

package kata

func Rps(p1, p2 string) string {

	if p1 == p2 {
		return "Draw!"
	}

	if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	}

	if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	}

	if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	} else {
		return "Player 2 won!"
	}
}
