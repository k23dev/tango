package templates

func (t *Templates) ViewsFullMenu() string {

	t.setReplacements()

	template := `
package menus

templ $PC$(){

    <div class=" relative ">
        <a href={ templ.URL("new") } class="flex-shrink-0 px-4 py-2 text-base font-semibold text-white bg-purple-600 rounded-lg shadow-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 focus:ring-offset-purple-200" type="submit">
            Nuevo
        </a>
    </div>
}
	`

	return t.Replacements.Replace(template)

}

func (t *Templates) ViewsFullForms() string {

	t.setReplacements()

	template := `
package forms

templ $SC$(action string,item models.$SC$){
    <form class="flex flex-col justify-center w-3/4 max-w-sm space-y-3 md:flex-row md:w-full md:space-x-3 md:space-y-0" action={ templ.URL(action) } method="post">
        <div class=" relative ">
            <input type="text" id="&quot;form-subscribe-Subscribe" class=" rounded-lg border-transparent flex-1 appearance-none border border-gray-300 w-full py-2 px-4 bg-white text-gray-700 placeholder-gray-400 shadow-sm text-base focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" 
            name="name" placeholder="Nombre..." value={item.Name}/>
        </div>
        <button class="flex-shrink-0 px-4 py-2 text-base font-semibold text-white bg-purple-600 rounded-lg shadow-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 focus:ring-offset-purple-200" type="submit">
            Guardar
        </button>
        <a href="/$PL$/" class="flex-shrink-0 px-4 py-2 text-base font-semibold text-white bg-red-600 rounded-lg shadow-md hover:bg-red-500 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 focus:ring-offset-red-200 text-center">
            Cancelar
        </a>
    </form>
}
	`

	return t.Replacements.Replace(template)
}

func (t *Templates) ViewsFullFormsWithSelector() string {

	t.setReplacements()

	template := `
package forms

import(
    "github.com/k23dev/tango/app/views/components"
    "github.com/k23dev/tango/app/models"
)

templ $SC$(action string,selectorList []models.$SC$){
    <form class="flex flex-col justify-center w-3/4 max-w-sm space-y-3 md:flex-row md:w-full md:space-x-3 md:space-y-0" action={ templ.URL(action) } method="post">
        <div class=" relative ">
            <input type="text" id="&quot;form-subscribe-Subscribe" class=" rounded-lg border-transparent flex-1 appearance-none border border-gray-300 w-full py-2 px-4 bg-white text-gray-700 placeholder-gray-400 shadow-sm text-base focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" 
            name="name" placeholder="Nombre..." value=""/>
        </div>
        // Selector component $PC$
        @components.Selector$SC$(selectorList)
        <button class="flex-shrink-0 px-4 py-2 text-base font-semibold text-white bg-purple-600 rounded-lg shadow-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 focus:ring-offset-purple-200" type="submit">
            Guardar
        </button>
        <a href="/$PL$/" class="flex-shrink-0 px-4 py-2 text-base font-semibold text-white bg-red-600 rounded-lg shadow-md hover:bg-red-500 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 focus:ring-offset-red-200 text-center">
            Cancelar
        </a>
    </form>
}
	`

	return t.Replacements.Replace(template)
}

func (t *Templates) ViewsFullFormsWithSelectorUpdate() string {

	t.setReplacements()

	template := `
package forms

import(
    "github.com/k23dev/tango/app/views/components"
    "github.com/k23dev/tango/app/models"
)

templ $SC$Update(action string,item models.$SC$,selectorList []models.$SC$){
    <form class="flex flex-col justify-center w-3/4 max-w-sm space-y-3 md:flex-row md:w-full md:space-x-3 md:space-y-0" action={ templ.URL(action) } method="post">
        <div class=" relative ">
            <input type="text" id="&quot;form-subscribe-Subscribe" class=" rounded-lg border-transparent flex-1 appearance-none border border-gray-300 w-full py-2 px-4 bg-white text-gray-700 placeholder-gray-400 shadow-sm text-base focus:outline-none focus:ring-2 focus:ring-purple-600 focus:border-transparent" 
            name="name" placeholder="Nombre..." value={item.Name}/>
        </div>
        // Selector component $PC$
        @components.Selector$SC$_update(selectorList,item.ID)
        <button class="flex-shrink-0 px-4 py-2 text-base font-semibold text-white bg-purple-600 rounded-lg shadow-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 focus:ring-offset-purple-200" type="submit">
            Guardar
        </button>
        <a href="/$PL$/" class="flex-shrink-0 px-4 py-2 text-base font-semibold text-white bg-red-600 rounded-lg shadow-md hover:bg-red-500 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 focus:ring-offset-red-200 text-center">
            Cancelar
        </a>
    </form>
}
	`

	return t.Replacements.Replace(template)
}

func (t *Templates) ViewsFullTable() string {

	t.setReplacements()

	template := `
package tables

import "github.com/k23dev/tango/app/models"
import "github.com/k23dev/tango/app/views/menus"
import "github.com/k23dev/tango/pkg/pagination"

templ $PC$($PL$ []models.$SC$,pagination pagination.Pagination){

<div class="">
    <div class="py-8">
        <div class="flex flex-row justify-between w-full mb-1 sm:mb-0">
            <h2 class="text-2xl leading-tight">
                $PC$
            </h2>
            <div class="text-end">
                @menus.$PC$()
            </div>
        </div>
            <div class="px-4 py-4 -mx-4 overflow-x-auto sm:-mx-8 sm:px-8">
                <div class="inline-block min-w-full overflow-hidden rounded-lg shadow">
                    <table class="min-w-full leading-normal">
                        <thead>
                            <tr>
                                <th scope="col" class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-200">
                                    Nombre
                                </th>
                                <th scope="col" class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-200">
                                    Created at
                                </th>
                                <th scope="col" class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-200">
                                    status
                                </th>
                                <th scope="col" class="px-5 py-3 text-sm font-normal text-left text-gray-800 uppercase bg-white border-b border-gray-200">
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            for _,$FL$:=range($PL$){
                            <tr>
                                <td class="px-5 py-5 text-sm bg-white border-b border-gray-200">
                                    <p class="text-gray-900 whitespace-no-wrap">
                                        {$FL$.Name}
                                    </p>
                                </td>
                                <td class="px-5 py-5 text-sm bg-white border-b border-gray-200">
                                    <p class="text-gray-900 whitespace-no-wrap">
                                        {$FL$.CreatedAt.String()}
                                    </p>
                                </td>
                                <td class="px-5 py-5 text-sm bg-white border-b border-gray-200">
                                    <span class="relative inline-block px-3 py-1 font-semibold leading-tight text-green-900">
                                        <span aria-hidden="true" class="absolute inset-0 bg-green-200 rounded-full opacity-50">
                                        </span>
                                        <span class="relative">
                                            {$FL$.DeletedAt.Time.String()}
                                        </span>
                                    </span>
                                </td>
                                <td class="px-5 py-5 text-sm bg-white border-b border-gray-200">
                                    <a href={templ.URL("/$PL$/edit/"+$FL$.GetIDAsString())} class="text-indigo-600 hover:text-indigo-900">
                                        Edit
                                    </a>
                                    |
                                    <a href={templ.URL("/$PL$/delete/"+$FL$.GetIDAsString())} class="text-red-600 hover:text-red-900">
                                        Delete
                                    </a>
                                </td>
                            </tr>
                            }
                        </tbody>
                    </table>
                    <div class="flex flex-col items-center px-5 py-5 bg-white xs:flex-row xs:justify-between">
                        <div class="flex items-center">
                            <a href={ templ.URL("?page="+pagination.ToString("prev")) } class="w-full p-4 text-base text-gray-600 bg-white border rounded-l-xl hover:bg-gray-100">
                                <svg width="9" fill="currentColor" height="8" class="" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M1427 301l-531 531 531 531q19 19 19 45t-19 45l-166 166q-19 19-45 19t-45-19l-742-742q-19-19-19-45t19-45l742-742q19-19 45-19t45 19l166 166q19 19 19 45t-19 45z">
                                    </path>
                                </svg>
                            </a>
                            <a href={ templ.URL("?page="+pagination.ToString("start")) } class="w-full px-4 py-2 text-base text-indigo-500 bg-white border-t border-b hover:bg-gray-100 ">
                                { pagination.ToString("start") }
                            </a>
                            <a href={ templ.URL("?page="+pagination.ToString("current")) } class="w-full px-4 py-2 text-base text-gray-600 bg-white border hover:bg-gray-100">
                                { pagination.ToString("current") }
                            </a>
                            <a href={ templ.URL("?page="+pagination.ToString("end")) } class="w-full px-4 py-2 text-base text-gray-600 bg-white border hover:bg-gray-100">
                                { pagination.ToString("end") }
                            </a>
                            <a href={ templ.URL("?page="+pagination.ToString("next")) } class="w-full p-4 text-base text-gray-600 bg-white border-t border-b border-r rounded-r-xl hover:bg-gray-100">
                                <svg width="9" fill="currentColor" height="8" class="" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M1363 877l-742 742q-19 19-45 19t-45-19l-166-166q-19-19-19-45t19-45l531-531-531-531q-19-19-19-45t19-45l166-166q19-19 45-19t45 19l742 742q19 19 19 45t-19 45z">
                                    </path>
                                </svg>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
</div>
}

templ $PC$Empty(){
<div class="">
    <div class="py-8">
        <div class="flex flex-row justify-between w-full mb-1 sm:mb-0">
            <h2 class="text-2xl leading-tight">
                $PC$
            </h2>
            <div class="text-end">
                @menus.$PC$()
            </div>
        </div>
    	<div>
            <h1>No items</h1>
    	</div>
    </div>
</div>

}
	`

	return t.Replacements.Replace(template)

}
