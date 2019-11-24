package ginterface

import "github.com/lxn/walk"

const (
	// Resolucao do formulario de pesquisa
	ResCampoPesqX = 225
	ResCampoPesqY = 25
)

// CampoPesquisa campo de entrada para o texto a ser pesquisados nos arquivos dos diretorios selecionados
type CampoPesquisa struct {
	*walk.TextEdit
	jp *JanelaPrincipal
}

// NewCampoPesquisa cria um novo CampoPesquisa
func NewCampoPesquisa(jp *JanelaPrincipal) (*CampoPesquisa, error) {
	te, err := walk.NewTextEdit(jp.MainWindow)

	if err != nil {
		return nil, err
	}

	cp := &CampoPesquisa{te, jp}

	if err := walk.InitWrapperWindow(cp); err != nil {
		return nil, err
	}

	cp.SetSize(walk.Size{Width: ResCampoPesqX, Height: ResCampoPesqY})
	cp.TextChanged().Attach(cp.atualizarTextoPesquisa)

	return cp, nil
}

// atualizarTextoPesquisa altera a variavel da janela principal do texto de busca quando ocorre alteracao no campo de pesquisa
func (cp *CampoPesquisa) atualizarTextoPesquisa() {
	cp.jp.textoPesquisa = cp.Text()
}
