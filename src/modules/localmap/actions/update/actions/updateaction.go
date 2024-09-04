package actions

//Actions for form update .
//You can  bind actions to  router by using below code :
//import   updateactions "modules/localmap/actions/update/actions"
//
//	Router.POST("/update").
//		Handle(updateactions.ActionUpdate)

import (
	"net/http"

	"modules/localmap"
	"modules/localmap/actions/update/forms"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/ui/render"
	"github.com/herb-go/herb/ui/validator/formdata"
)

// ActionUpdate action that verifyupdate form in json format.
var ActionUpdate = action.New(func(w http.ResponseWriter, r *http.Request) {
	form := forms.NewUpdateForm()
	if formdata.MustValidateJSONRequest(r, form) {
		err := form.Exec()
		if err != nil {
			panic(err)
		}
		d := localmap.NewRawData(form.ID, form.Name, form.Data)
		localmap.DefaultManager.Import(d)
		localmap.MustSave()
		render.MustJSON(w, form, 200)
	} else {
		formdata.MustRenderErrorsJSON(w, form)
	}
})
