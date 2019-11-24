package ginterface

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

const (
	// Resolucao Janela Principal
	ResJanPrincX = 500
	ResJanPrincY = 700

	// Posicao Botao Adicionar Diretorio
	PosBotAddDirX = 5
	PosBotAddDirY = 5

	// Posicao Botao Deletar Diretorio
	PosBotDelDirX = 5
	PosBotDelDirY = 40

	// Posicao Lista Diretorios Selecionados
	PosListaDirsX = 130
	PosListaDirsY = 5

	// Posicao Campo Pesquisa
	PosCampoPesqX = 5
	PosCampoPesqY = 160

	// Posicao Botao Pesquisar
	PosBotPesquiX = 235
	PosBotPesquiY = 160

	// Posicao Botao Criar B Tree
	PosBotBTreeX = 360
	PosBotBTreeY = 160

	// Posicao Barra de Progresso
	PosBarraProgX = 5
	PosBarraProgY = 195

	// Posicao Janela de Log
	PosJanelaLogX = 5
	PosJanelaLogY = 235

	// Posicao Tabela de Resultados
	PosTabResultX = 5
	PosTabResultY = 425
)

// Executar inicia a execucao da Janela Principal
func Executar() {
	jp := new(JanelaPrincipal)
	jp.dirs = new(ModelListaDirs)
	jp.resultados = new(ModelResultadoPesquisa)
	jp.linhasLog = new(ModelLinhaLog)

	if err := (declarative.MainWindow{
		AssignTo: &jp.MainWindow,
		Title:    "Pesquisador",
		Size:     declarative.Size{Width: ResJanPrincX, Height: ResJanPrincY},
		MinSize:  declarative.Size{Width: ResJanPrincX, Height: ResJanPrincY},
	}).Create(); err != nil {
		log.Fatal(err)
	}

	bad, err := NewBotaoAdicionarDir(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		bad.SetX(PosBotAddDirX)
		bad.SetY(PosBotAddDirY)
	}

	bdd, err := NewBotaoDeletarDir(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		bdd.SetX(PosBotDelDirX)
		bdd.SetY(PosBotDelDirY)
	}

	lds, err := NewListaDirsSelecionados(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		lds.SetX(PosListaDirsX)
		lds.SetY(PosListaDirsY)
	}

	cp, err := NewCampoPesquisa(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		cp.SetX(PosCampoPesqX)
		cp.SetY(PosCampoPesqY)
	}

	bp, err := NewBotaoPesquisar(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		bp.SetX(PosBotPesquiX)
		bp.SetY(PosBotPesquiY)
	}

	bcbt, err := NewBotaoCriarBTree(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		bcbt.SetX(PosBotBTreeX)
		bcbt.SetY(PosBotBTreeY)
	}

	jp.barraProgresso, err = NewBarraProgresso(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		jp.barraProgresso.SetX(PosBarraProgX)
		jp.barraProgresso.SetY(PosBarraProgY)
	}

	jl, err := NewJanelaLog(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		jl.SetX(PosJanelaLogX)
		jl.SetY(PosJanelaLogY)
	}

	tr, err := NewTabelaResultados(jp)
	if err != nil {
		log.Fatal(err)
	} else {
		tr.SetX(PosTabResultX)
		tr.SetY(PosTabResultY)
	}

	jp.Run()
}

// JanelaPrincipal contem a base da interface grafica e informacoes que precisam ser acessadas por outros elementos graficos
type JanelaPrincipal struct {
	*walk.MainWindow
	dirs           *ModelListaDirs
	resultados     *ModelResultadoPesquisa
	barraProgresso *BarraProgresso
	linhasLog      *ModelLinhaLog
	indiceDirSelec int
	textoPesquisa  string
	textoProgresso string
}

// AdicionarDirSelecionado adiciona um novo diretorio a lista dos diretorios onde sera realizada pesquisa
func (jp *JanelaPrincipal) AdicionarDirSelecionado(pathDir string) {
	if jp.dirs.itens == nil {
		jp.dirs.itens = make([]*ItemListaDirs, 1)
		jp.dirs.itens[0] = &ItemListaDirs{pathDir: pathDir}

		jp.dirs.PublishItemsInserted(0, 0)
		return
	} else if !jp.VerificarDirSelecUnico(pathDir) {
		return
	}

	qtdDirs := len(jp.dirs.itens)
	dirsSelecAux := make([]*ItemListaDirs, qtdDirs+1)

	copy(dirsSelecAux[:qtdDirs], jp.dirs.itens[:qtdDirs])

	dirsSelecAux[qtdDirs] = &ItemListaDirs{pathDir: pathDir}

	jp.dirs.itens = make([]*ItemListaDirs, qtdDirs+1)
	jp.dirs.PublishItemsRemoved(0, qtdDirs)

	copy(jp.dirs.itens, dirsSelecAux)
	jp.dirs.PublishItemsInserted(0, qtdDirs)
}

// VerificarDirSelecUnico verifica se um diretorio ja existe entre os diretorios selecionados para pesquisa
func (jp *JanelaPrincipal) VerificarDirSelecUnico(pathDir string) bool {
	for _, dirSelec := range jp.dirs.itens {
		if pathDir == dirSelec.pathDir {
			return false
		}
	}
	return true
}

// DeletarDirSelecionado remove um diretorio da lista de diretorios selecionados para pesquisa
func (jp *JanelaPrincipal) DeletarDirSelecionado() {
	qtdDirs := len(jp.dirs.itens)
	dirsSelecAux := make([]*ItemListaDirs, qtdDirs)

	copy(dirsSelecAux, jp.dirs.itens)
	jp.dirs.itens = nil
	jp.dirs.PublishItemsRemoved(0, qtdDirs)

	for i, item := range dirsSelecAux {
		if i != jp.indiceDirSelec {
			jp.AdicionarDirSelecionado(item.pathDir)
		}
	}
}

// AdicionarResultado adiciona as informacoes de path do arquivo e localizacao da informacao busca na tabela de resultados
func (jp *JanelaPrincipal) AdicionarResultado(pathArquivo, localizacao string) {
	if jp.resultados.itens == nil {
		jp.resultados.itens = make([]*ItemResultadoPesquisa, 1)

		jp.resultados.itens[0] = &ItemResultadoPesquisa{
			pathArquivo: pathArquivo,
			localizacao: localizacao,
		}
		return
	}

	qtdResult := len(jp.resultados.itens)
	resultadosAux := make([]*ItemResultadoPesquisa, qtdResult+1)

	copy(resultadosAux[:qtdResult], jp.resultados.itens[:qtdResult])

	resultadosAux[qtdResult] = &ItemResultadoPesquisa{
		pathArquivo: pathArquivo,
		localizacao: localizacao,
	}

	jp.resultados.itens = make([]*ItemResultadoPesquisa, qtdResult+1)
	copy(jp.resultados.itens, resultadosAux)

	jp.resultados.PublishRowsInserted(0, qtdResult)
}

// RegistrarQtdArquivos define o valor de referencia de quantos arquivos serao pesquisados para a barra de progresso
func (jp *JanelaPrincipal) RegistrarQtdArquivos(qtdArquivos int) {
	jp.barraProgresso.SetRange(0, qtdArquivos)
	jp.barraProgresso.SetValue(0)
	jp.barraProgresso.SetToolTipText(fmt.Sprintf("Pesquisando [0 / %d]", qtdArquivos))
}

// AtualizarProgresso informa a barra de progresso que um arquivo foi concluido
func (jp *JanelaPrincipal) AtualizarProgresso() {
	if jp.barraProgresso.Value() < jp.barraProgresso.MaxValue() {
		jp.barraProgresso.SetValue(jp.barraProgresso.Value() + 1)
		jp.barraProgresso.SetToolTipText(fmt.Sprintf("Pesquisando [%d / %d]", jp.barraProgresso.Value(), jp.barraProgresso.MaxValue()))
	}
}

// EscreverLog registra uma mensagem na caixa de logs
func (jp *JanelaPrincipal) EscreverLog(mensagem string) {
	if jp.linhasLog.linhas == nil {
		jp.linhasLog.linhas = make([]*LinhaLog, 1)
		jp.linhasLog.linhas[0] = &LinhaLog{mensagem: mensagem}

		jp.linhasLog.PublishItemsInserted(0, 0)
		return
	}
	qtdLinhas := len(jp.linhasLog.linhas)
	linhasLogAux := make([]*LinhaLog, qtdLinhas+1)

	copy(linhasLogAux[:qtdLinhas], jp.linhasLog.linhas[:qtdLinhas])

	linhasLogAux[qtdLinhas] = &LinhaLog{mensagem: mensagem}

	jp.linhasLog.linhas = make([]*LinhaLog, qtdLinhas+1)
	jp.linhasLog.PublishItemsRemoved(0, qtdLinhas)

	copy(jp.linhasLog.linhas, linhasLogAux)
	jp.linhasLog.PublishItemsInserted(0, qtdLinhas)
}

// ResetarLog limpa todas as mensagens na caixa de log
func (jp *JanelaPrincipal) ResetarLog() {
	qtdDirs := len(jp.linhasLog.linhas)
	jp.linhasLog.linhas = nil
	jp.dirs.PublishItemsRemoved(0, qtdDirs)
}
