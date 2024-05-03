package templates

func (t *Templates) ViewsFullSelector() string {

	t.setReplacements()

	template := `
package components

import(
    "github.com/k23dev/tango/app/models"
)

templ Selector$SC$(list []models.$SC$){

<label class="text-gray-700" for="item_id">
    $PL$
    <select name="$SL$_id" id="$SL$_id" class="block px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm w-52 focus:outline-none focus:ring-primary-500 focus:border-primary-500">
    for _, item :=range(list){
        <option value={item.GetIDAsString()}>{item.Name}</option>
    }
    </select>
</label>
}
	`

	return t.Replacements.Replace(template)

}

func (t *Templates) ViewsFullSelectorUpdate() string {

	t.setReplacements()

	template := `
package components

import(
    "github.com/k23dev/tango/app/models"
)

templ Selector$SC$_update(list []models.$SC$,currentSelectedID uint){

<label class="text-gray-700" for="item_id">
    $PL$
    <select name="$SL$_id" id="$SL$_id" class="block px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm w-52 focus:outline-none focus:ring-primary-500 focus:border-primary-500">
    for _, item :=range(list){
        if currentSelectedID == item.ID{
            <option value={item.GetIDAsString()} selected>{item.Name}</option>
        }else{
            <option value={item.GetIDAsString()}>{item.Name}</option>
        }
    }
    </select>
</label>
}
	`

	return t.Replacements.Replace(template)

}
