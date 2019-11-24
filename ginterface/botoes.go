package ginterface

import (
	"fmt"
	"log"
	"time"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

const (
	// Resolucao dos botoes
	ResBotoesX = 120
	ResBotoesY = 25
)

// BotaoAdicionarDir abre uma janela de selecao de diretorio e o adiciona a lista de diretorios a serem pesquisados
type BotaoAdicionarDir struct {
	*walk.PushButton
	jp *JanelaPrincipal
}

// NewBotaoAdicionarDir cria um novo BotaoAdicionarDir
func NewBotaoAdicionarDir(jp *JanelaPrincipal) (*BotaoAdicionarDir, error) {
	pb, err := walk.NewPushButton(jp.MainWindow)

	if err != nil {
		return nil, err
	}

	bad := &BotaoAdicionarDir{pb, jp}

	if err := walk.InitWrapperWindow(bad); err != nil {
		return nil, err
	}

	bad.SetText("Adicionar Diretório")
	bad.SetSize(walk.Size{Width: ResBotoesX, Height: ResBotoesY})

	return bad, nil
}

// WndProc controla a interacao ao clicar no botao "Adicionar Diretorio"
func (bad *BotaoAdicionarDir) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		dialogSelecDir := new(walk.FileDialog)

		dialogSelecDir.ShowBrowseFolder(bad.jp)
		bad.jp.AdicionarDirSelecionado(dialogSelecDir.FilePath)
	}

	return bad.PushButton.WndProc(hwnd, msg, wParam, lParam)
}

// BotaoDeletarDir remove um diretorio da lista de selecionados
type BotaoDeletarDir struct {
	*walk.PushButton
	jp *JanelaPrincipal
}

// NewBotaoDeletarDir cria um novo BotaoDeletarDir
func NewBotaoDeletarDir(jp *JanelaPrincipal) (*BotaoDeletarDir, error) {
	pb, err := walk.NewPushButton(jp.MainWindow)

	if err != nil {
		return nil, err
	}

	bdd := &BotaoDeletarDir{pb, jp}

	if err := walk.InitWrapperWindow(bdd); err != nil {
		return nil, err
	}

	bdd.SetText("Deletar Diretório")
	bdd.SetSize(walk.Size{Width: ResBotoesX, Height: ResBotoesY})

	return bdd, nil
}

// WndProc controla a interacao ao clicar no botao "Deletar Diretorio"
func (bdd *BotaoDeletarDir) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		// Chamada de funcao que deleta um item da lista de diretorios
		bdd.jp.DeletarDirSelecionado()
	}

	return bdd.PushButton.WndProc(hwnd, msg, wParam, lParam)
}

// BotaoPesquisar inicia a pesquisa do texto nos arquivos dos diretorios selecionados
type BotaoPesquisar struct {
	*walk.PushButton
	jp *JanelaPrincipal
}

// NewBotaoPesquisar cria um novo BotaoPesquisar
func NewBotaoPesquisar(jp *JanelaPrincipal) (*BotaoPesquisar, error) {
	pb, err := walk.NewPushButton(jp.MainWindow)
	if err != nil {
		return nil, err
	}

	bp := &BotaoPesquisar{pb, jp}
	if err := walk.InitWrapperWindow(bp); err != nil {
		return nil, err
	}

	bp.SetText("Pesquisar")
	bp.SetSize(walk.Size{Width: ResBotoesX, Height: ResBotoesY})

	return bp, nil
}

// WndProc controla a interacao ao clicar no botao "Pesquisar"
func (bp *BotaoPesquisar) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		// TODO: Substituir a alimentacão por dados reais. O texto de busca e acessado por "bp.jp.textoPesquisa"
		qtdArquivos := 30 // Alterar para quantidade real de arquivos a pesquisar

		bp.jp.RegistrarQtdArquivos(qtdArquivos) // Registro para referência na barra de progresso
		bp.jp.ResetarLog()                      // Limpeza do texto na janela de log
		bp.jp.resultados.ResetarTabela()        // Limpeza dos dados anteriores na tabela de resultados

		for i := 1; i <= qtdArquivos; i++ {
			bp.jp.AdicionarResultado("C:\\Algum\\Dir\\arquivo.txt", fmt.Sprintf("Linha %d", i)) // Adiciona um resultado na tabela
			bp.jp.AtualizarProgresso()                                                          // Atualiza a barra de progresso, usar a cara nova insercão
			bp.jp.EscreverLog(fmt.Sprintf("C:\\Algum\\Dir\\arquivo%d.txt", i))                  // Registra mensagem na janela de log
			time.Sleep(500 * time.Millisecond)                                                  // Sleep apenas para simulacão, remover ao usar dados reais
		}
	}

	return bp.PushButton.WndProc(hwnd, msg, wParam, lParam)
}

// BotaoCriarBTree cria uma B-Tree dos arquivos nos diretorios selecionados
type BotaoCriarBTree struct {
	*walk.PushButton
	jp *JanelaPrincipal
}

// NewBotaoCriarBTree cria um novo BotaoCriarBTree
func NewBotaoCriarBTree(jp *JanelaPrincipal) (*BotaoCriarBTree, error) {
	pb, err := walk.NewPushButton(jp.MainWindow)

	if err != nil {
		return nil, err
	}

	bcbt := &BotaoCriarBTree{pb, jp}

	if err := walk.InitWrapperWindow(bcbt); err != nil {
		return nil, err
	}

	bcbt.SetText("Criar B-Tree")
	bcbt.SetSize(walk.Size{Width: ResBotoesX, Height: ResBotoesY})

	return bcbt, nil
}

// WndProc controla a interacao ao clicar no botao "Criar B-Tree"
func (bcbt *BotaoCriarBTree) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		log.Printf("%s: WM_LBUTTONDOWN", bcbt.Text())
		// TODO: Chamada de funcao de criar B-Tree
	}

	return bcbt.PushButton.WndProc(hwnd, msg, wParam, lParam)
}
