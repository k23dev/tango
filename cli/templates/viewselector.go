package templates

func (t *Templates) ViewSelector() string {
	t.setReplacements()

	template := `
package views

import (
    "github.com/k23dev/tango/app/views/layouts"
    "github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/app/views/tables"
	"github.com/k23dev/tango/app/views/forms"
    "github.com/k23dev/tango/pkg/pagination"
)

templ $PC$ShowList(appTitle string,$PL$ []models.$SC$,pagination pagination.Pagination){
    @layouts.Default(appTitle){
        @tables.$PC$($PL$,pagination)
    }

}

templ $PC$ShowListEmpty(appTitle string){
    @layouts.Default(appTitle){
        @tables.$PC$Empty()
    }

}

templ $PC$ShowOne(appTitle string,$FL$ models.$SC$){
    @layouts.Default(appTitle){
        <h1>$SC$</h1>
        <h2>{$FL$.Name}</h2>
    }
}

templ $PC$FormCreate(appTitle string,list *[]models.$SC$){
    @layouts.Default(appTitle){
        @forms.$SC$("/$PL$/create",*list)
    }
}

templ $PC$FormUpdate(appTitle string,$FL$ *models.$SC$,list *[]models.$SC$){
    @layouts.Default(appTitle){
        @forms.$SC$Update("/$PL$/update/"+$FL$.GetIDAsString(),*$FL$,*list)
    }
}

templ $PC$Delete(appTitle string,$FL$ *models.$SC$){
    @layouts.Default(appTitle){
        <h1>Borrar $SL$</h1>
    }
}
	`
	return t.Replacements.Replace(template)
}
