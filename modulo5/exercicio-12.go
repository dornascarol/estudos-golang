// Well of Ideas - Easy Version
// For every good kata idea there seem to be quite a few bad ones!
// In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. If there are one or two good ideas, return 'Publish!', if there are more than 2 return 'I smell a series!'. If there are no good ideas, as is often the case, return 'Fail!'.

// Lógica:
// 01) Receber uma lista e percorrer com for range os elementos de boas ideias.
// 02) Se só tiver "bad" => o retorno será "Fail!"
// 03) Se tiver 1 ou 2 "good" => o retorno será "Publish!"
// 04) Se tiver 3 ou mais "good" => o retorno será "I smell a series!"
// 05) Variável para guardar a quantidade de "good" que aparecem na lista.

package kata

func Well(x []string) string {
	var numberGood int

	for _, elementGood := range x {
		if elementGood == "good" {
			numberGood = numberGood + 1
		}
	}

	if numberGood == 0 {
		return "Fail!"
	}

	if numberGood <= 2 {
		return "Publish!"
	} else {
		return "I smell a series!"
	}

}
