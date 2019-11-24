package ginterface

import "github.com/lxn/walk"

const (
	// Resolucao da lista de diretorios
	ResListaDirsX = 350
	ResListaDirsY = 145
)

// ListaDirsSelecionados elemento grafico que indica os diretorios selecionados para pesquisa
type ListaDirsSelecionados struct {
	*walk.ListBox
	jp *JanelaPrincipal
}

// ItemListaDirs struct contendo a informacao de um diretorio alvo de pesquisa
type ItemListaDirs struct {
	pathDir string
}

// ModelListaDirs representa a informacao da lista de diretorios selecionados para pesquisa
type ModelListaDirs struct {
	walk.ListModelBase
	itens []*ItemListaDirs
}

// NewListaDirsSelecionados cria uma nova ListaDirSelecionados
func NewListaDirsSelecionados(jp *JanelaPrincipal) (*ListaDirsSelecionados, error) {
	lb, err := walk.NewListBox(jp.MainWindow)
	if err != nil {
		return nil, err
	}

	lds := &ListaDirsSelecionados{lb, jp}
	if err := walk.InitWrapperWindow(lds); err != nil {
		return nil, err
	}

	lds.SetSize(walk.Size{Width: ResListaDirsX, Height: ResListaDirsY})
	lds.SetModel(jp.dirs)
	lds.CurrentIndexChanged().Attach(lds.atualizarIndiceDirSelecionado)

	return lds, nil
}

// atualizarIndiceDirSelecionado altera a variavel da janela principal do indice do diretorio selecionado na lista para poder deletar
func (lds *ListaDirsSelecionados) atualizarIndiceDirSelecionado() {
	lds.jp.indiceDirSelec = lds.CurrentIndex()
}

// ItemCount funcao requerida pela biblioteca "github.com/lxn/walk" para exibicao de elementos visuais
func (mld *ModelListaDirs) ItemCount() int {
	return len(mld.itens)
}

// Value funcao requerida pela biblioteca "github.com/lxn/walk" para exibicao de elementos visuais
func (mld *ModelListaDirs) Value(index int) interface{} {
	return mld.itens[index].pathDir
}
