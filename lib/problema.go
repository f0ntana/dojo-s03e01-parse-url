package lib

import (
	"strings"
)

//Problema estrutura que representa o problema
type Problema struct {
	Protocol   string
	Host       string
	Domain     string
	Path       string
	Parameters string
}

//NewProblema cria um novo problema
func NewProblema() *Problema {
	return &Problema{}
}

//Resolver devolve o problema resolvido
func (problema *Problema) Resolver(url string) *Problema {
	return problema.resolverEscondido(url)
}

func (problema *Problema) extractProtocol(url string) {
	if strings.Contains(url, "://") {
		problema.Protocol = strings.Split(url, "://")[0]
	} else {
		problema.Protocol = ""
	}
}

func (problema *Problema) extractHost(url string) {
	if strings.Contains(url, "://") {
		protocol := strings.Split(url, "://")
		url = protocol[1]
	}

	parts := strings.Split(url, ".")
	problema.Host = parts[0]
}

func (problema *Problema) extractDomain(url string) {

	url = strings.Replace(url, problema.Protocol, "", -1)
	url = strings.Replace(url, "://", "", -1)

	urlSplited := strings.SplitN(url, ".", 2)
	// split no ponto pedindo 2 valores
	// pegar o ultimo valor e splitar ate o primeiro slash (/)
	if len(urlSplited) > 0 {
		problema.Domain = urlSplited[1]
	} else {
		problema.Domain = urlSplited[0]
	}

}

//resolverEscondido metodo privado que resolve o problema
func (problema *Problema) resolverEscondido(url string) *Problema {
	problema.extractProtocol(url)
	problema.extractHost(url)
	problema.extractDomain(url)

	return problema
}
