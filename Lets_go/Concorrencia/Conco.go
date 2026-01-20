package concorrencia

// type VerificadorWebsite func(string) bool

// func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
// 	resultados := make(map[string]bool)

// 	for _, url := range urls {
// 		resultados[url] = vw(url)
// 	}

// 	return resultados
// }
//----------
// import "time"

// type VerificadorWebsite func(string) bool

// func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
// 	resultados := make(map[string]bool)

// 	for _, url := range urls {
// 		go func(u string) {
// 			resultados[u] = vw(u)
// 		}(url)
// 	}

// 	time.Sleep(2 * time.Second)

// 	return resultados
// }
//=============

type VerificadorWebsite func(string) bool
type resultado struct {
	string
	bool
}

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)
	canalResultado := make(chan resultado)

	for _, url := range urls {
		go func(u string) {
			canalResultado <- resultado{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resultado := <-canalResultado
		resultados[resultado.string] = resultado.bool
	}

	return resultados
}
