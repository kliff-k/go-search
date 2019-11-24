package ginterface

import (
	"github.com/lxn/walk"
)

const (
	// Resolucao da Barra de Progresso
	ResBarraProgX = 475
	ResBarraProgY = 30
	// Resolucao da Janela de Log
	ResJanelaLogX = 475
	ResJanelaLogY = 180
)

// BarraProgresso indica visualmente o progresso de conclusao das pesquisas
type BarraProgresso struct {
	*walk.ProgressBar
	jp *JanelaPrincipal
}

// NewBarraProgresso cria uma nova BarraProgresso
func NewBarraProgresso(jp *JanelaPrincipal) (*BarraProgresso, error) {
	pb, err := walk.NewProgressBar(jp.MainWindow)

	if err != nil {
		return nil, err
	}

	bpr := &BarraProgresso{pb, jp}

	if err := walk.InitWrapperWindow(bpr); err != nil {
		return nil, err
	}

	bpr.SetSize(walk.Size{Width: ResBarraProgX, Height: ResBarraProgY})

	return bpr, nil
}

// JanelaLog exibe mensagens sobre o progresso das pesquisas
type JanelaLog struct {
	*walk.ListBox
	jp *JanelaPrincipal
}

// LinhaLog armazena a informa de uma mensagem de log
type LinhaLog struct {
	mensagem string
}

// ModelLinhaLog armazena o conjunto de mensagens de log
type ModelLinhaLog struct {
	walk.ListModelBase
	linhas []*LinhaLog
}

// NewJanelaLog cria uma nova JanelaLog
func NewJanelaLog(jp *JanelaPrincipal) (*JanelaLog, error) {
	lb, err := walk.NewListBox(jp.MainWindow)
	if err != nil {
		return nil, err
	}

	lds := &JanelaLog{lb, jp}
	if err := walk.InitWrapperWindow(lds); err != nil {
		return nil, err
	}

	lds.SetSize(walk.Size{Width: ResJanelaLogX, Height: ResJanelaLogY})
	lds.SetModel(jp.linhasLog)

	return lds, nil
}

// ItemCount funcao requerida pela biblioteca "github.com/lxn/walk" para exibicao de elementos visuais
func (mld *ModelLinhaLog) ItemCount() int {
	return len(mld.linhas)
}

// Value funcao requerida pela biblioteca "github.com/lxn/walk" para exibicao de elementos visuais
func (mld *ModelLinhaLog) Value(index int) interface{} {
	return mld.linhas[index].mensagem
}
