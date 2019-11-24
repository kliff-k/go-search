package ginterface

import "github.com/lxn/walk"

const (
	// Resolucao da tabela de resultados
	ResTabResultX = 475
	ResTabResultY = 230

	//Largura das colunas
	LargPathArquivo = 325
	LargLocalizacao = 150
	LargTotalInical = LargPathArquivo + LargLocalizacao
)

// TabelaResultados exibe uma tabela de duas colunas informando os resultados da pesquisa
type TabelaResultados struct {
	*walk.TableView
	jp *JanelaPrincipal
}

// ItemResultadoPesquisa contem a informacao do path de um arquivo e a posicao nele onde o texto foi encontrado
type ItemResultadoPesquisa struct {
	pathArquivo string
	localizacao string
}

// ModelResultadoPesquisa contem a lista de resultados da pesquisa
type ModelResultadoPesquisa struct {
	walk.TableModelBase
	itens []*ItemResultadoPesquisa
}

// NewTabelaResultados cria uma nova TabelaResultados
func NewTabelaResultados(jp *JanelaPrincipal) (*TabelaResultados, error) {
	tv, err := walk.NewTableView(jp.MainWindow)

	if err != nil {
		return nil, err
	}

	tr := &TabelaResultados{tv, jp}

	if err := walk.InitWrapperWindow(tr); err != nil {
		return nil, err
	}

	tvcPathArquivo := walk.NewTableViewColumn()
	tvcPathArquivo.SetTitle("Path Arquivo")
	tvcPathArquivo.SetWidth(LargPathArquivo)

	tvcLocalizacao := walk.NewTableViewColumn()
	tvcLocalizacao.SetTitle("Localizacão")
	tvcLocalizacao.SetWidth(LargLocalizacao)
	tvcLocalizacao.SetAlignment(walk.AlignCenter)

	tr.SetSize(walk.Size{Width: ResTabResultX, Height: ResTabResultY})
	tr.SetModel(jp.resultados)
	tr.SetWidth(LargTotalInical)
	tr.SetAlternatingRowBG(true)
	tr.SetColumnsOrderable(false)

	tvcl := tr.Columns()
	tvcl.Add(tvcPathArquivo)
	tvcl.Add(tvcLocalizacao)

	return tr, nil
}

// RowCount funcao requerida pela biblioteca "github.com/lxn/walk" para exibicao de elementos visuais
func (mrp *ModelResultadoPesquisa) RowCount() int {
	return len(mrp.itens)
}

// Value funcao requerida pela biblioteca "github.com/lxn/walk" para exibicao de elementos visuais
func (mrp *ModelResultadoPesquisa) Value(linha, coluna int) interface{} {
	item := mrp.itens[linha]

	switch coluna {
	case 0:
		return item.pathArquivo

	case 1:
		return item.localizacao
	}

	panic("unexpected col")
}

// ResetarTabela limpa os dados da tabela de resultados
func (mrp *ModelResultadoPesquisa) ResetarTabela() {
	mrp.itens = nil
	mrp.PublishRowsReset()
}
