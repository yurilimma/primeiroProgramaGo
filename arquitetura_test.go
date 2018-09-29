package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T){
	//executa os testes em paralelo
	t.Parallel()
	//verifica se a arquitetura atual é amd64
	if runtime.GOARCH == "amd64"{
		t.Skip("Não funciona em arquitetura amd64")
	}
	t.Fail()
}