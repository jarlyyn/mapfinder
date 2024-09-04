package forms

import (
	"net/http"
	"strings"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
	"github.com/herb-go/util/form/commonform"
)

// UpdateFormFieldLabels form field labels map.
var UpdateFormFieldLabels = map[string]string{
	"ID":   "ID",
	"Name": "名称",
	"Data": "地图",
}

// UpdateForm form struct forupdate
type UpdateForm struct {
	formdata.Form
	ID   string
	Name string
	Data string
}

// UpdateFormID form id of formupdate
const UpdateFormID = "formlocalmap.actions.update"

// NewUpdateForm create newupdate form
func NewUpdateForm() *UpdateForm {
	form := &UpdateForm{}
	return form
}

// ComponentID return form component id.
func (f *UpdateForm) ComponentID() string {
	return UpdateFormID
}

// Validate Validate form and return any error if raised.
func (f *UpdateForm) Validate() error {
	commonform.ValidateRequiredString(f, f.ID, f.GetFieldLabel("ID"))
	commonform.ValidateRequiredString(f, f.Name, f.GetFieldLabel("Name"))
	commonform.ValidateRequiredString(f, f.Data, f.GetFieldLabel("Data"))
	if !f.HasError() {
		if strings.Contains(f.ID, " ") {
			f.AddError("ID", "ID不能包含空格")
		}
	}
	if !f.HasError() {
		if strings.Contains(f.Name, " ") {
			f.AddError("Name", "Name不能包含空格")
		}
	}
	return nil
}

// Exec execwhen form validated.
func (f *UpdateForm) Exec() error {
	return nil
}

// InitWithRequest init update form  with http request.
func (f *UpdateForm) InitWithRequest(r *http.Request) error {

	//Put your request form code here.
	//such as get current user id or ip address.

	//Set form labels with translated messages
	f.SetComponentLabels(ui.GetMessages(f.Lang(), "app").Collection(UpdateFormFieldLabels))
	return nil
}
