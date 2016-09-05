package controllers

import (
	"github.com/lfq7413/tomato/errs"
	"github.com/lfq7413/tomato/hooks"
	"github.com/lfq7413/tomato/types"
	"github.com/lfq7413/tomato/utils"
)

// HooksController ...
type HooksController struct {
	ClassesController
}

// HandleGetAllFunctions ...
// @router /functions [get]
func (h *HooksController) HandleGetAllFunctions() {
	if h.Auth.IsMaster == false {
		h.Data["json"] = errs.ErrorMessageToMap(errs.OperationForbidden, "Need master key!")
		h.ServeJSON()
		return
	}

	results, err := hooks.GetFunctions()
	if err != nil {
		h.Data["json"] = errs.ErrorToMap(err)
		h.ServeJSON()
		return
	}
	if results == nil {
		results = types.S{}
	}
	h.Data["json"] = results
	h.ServeJSON()
}

// HandleGetFunction ...
// @router /functions/:functionName [get]
func (h *HooksController) HandleGetFunction() {
	if h.Auth.IsMaster == false {
		h.Data["json"] = errs.ErrorMessageToMap(errs.OperationForbidden, "Need master key!")
		h.ServeJSON()
		return
	}

	functionName := h.Ctx.Input.Param(":functionName")
	result, err := hooks.GetFunction(functionName)
	if err != nil {
		h.Data["json"] = errs.ErrorToMap(err)
		h.ServeJSON()
		return
	}
	if result == nil {
		h.Data["json"] = errs.ErrorMessageToMap(errs.WebhookError, "no function named: "+functionName+" is defined")
		h.ServeJSON()
		return
	}
	h.Data["json"] = result
	h.ServeJSON()
}

// HandleCreateFunction ...
// @router /functions [post]
func (h *HooksController) HandleCreateFunction() {
	if h.Auth.IsMaster == false {
		h.Data["json"] = errs.ErrorMessageToMap(errs.OperationForbidden, "Need master key!")
		h.ServeJSON()
		return
	}

	result, err := hooks.CreateHook(h.JSONBody)
	if err != nil {
		h.Data["json"] = errs.ErrorToMap(err)
		h.ServeJSON()
		return
	}
	h.Data["json"] = result
	h.ServeJSON()
}

// HandleUpdateFunction ...
// @router /functions/:functionName [put]
func (h *HooksController) HandleUpdateFunction() {
	if h.Auth.IsMaster == false {
		h.Data["json"] = errs.ErrorMessageToMap(errs.OperationForbidden, "Need master key!")
		h.ServeJSON()
		return
	}

	functionName := h.Ctx.Input.Param(":functionName")
	var err error
	var result = types.M{}
	if utils.S(h.JSONBody["__op"]) == "Delete" {
		// delete
		err = hooks.DeleteFunction(functionName)
	} else {
		// update
		if h.JSONBody["url"] == nil {
			h.Data["json"] = errs.ErrorMessageToMap(errs.WebhookError, "invalid hook declaration")
			h.ServeJSON()
			return
		}
		hook := types.M{
			"functionName": functionName,
			"url":          h.JSONBody["url"],
		}
		result, err = hooks.UpdateHook(hook)
	}
	if err != nil {
		h.Data["json"] = errs.ErrorToMap(err)
		h.ServeJSON()
		return
	}
	h.Data["json"] = result
	h.ServeJSON()
}

// HandleGetAllTriggers ...
// @router /triggers [get]
func (h *HooksController) HandleGetAllTriggers() {
	if h.Auth.IsMaster == false {
		h.Data["json"] = errs.ErrorMessageToMap(errs.OperationForbidden, "Need master key!")
		h.ServeJSON()
		return
	}

	results, err := hooks.GetTriggers()
	if err != nil {
		h.Data["json"] = errs.ErrorToMap(err)
		h.ServeJSON()
		return
	}
	if results == nil {
		results = types.S{}
	}
	h.Data["json"] = results
	h.ServeJSON()
}

// HandleGetTrigger ...
// @router /triggers/:className/:triggerName [get]
func (h *HooksController) HandleGetTrigger() {
	if h.Auth.IsMaster == false {
		h.Data["json"] = errs.ErrorMessageToMap(errs.OperationForbidden, "Need master key!")
		h.ServeJSON()
		return
	}

	className := h.Ctx.Input.Param(":className")
	triggerName := h.Ctx.Input.Param(":triggerName")
	result, err := hooks.GetTrigger(className, triggerName)
	if err != nil {
		h.Data["json"] = errs.ErrorToMap(err)
		h.ServeJSON()
		return
	}
	if result == nil {
		h.Data["json"] = errs.ErrorMessageToMap(errs.WebhookError, "class "+className+" does not exist")
		h.ServeJSON()
		return
	}
	h.Data["json"] = result
	h.ServeJSON()
}

// HandleCreateTrigger ...
// @router /triggers [post]
func (h *HooksController) HandleCreateTrigger() {
	if h.Auth.IsMaster == false {
		h.Data["json"] = errs.ErrorMessageToMap(errs.OperationForbidden, "Need master key!")
		h.ServeJSON()
		return
	}

	result, err := hooks.CreateHook(h.JSONBody)
	if err != nil {
		h.Data["json"] = errs.ErrorToMap(err)
		h.ServeJSON()
		return
	}
	h.Data["json"] = result
	h.ServeJSON()
}

// HandleUpdateTrigger ...
// @router /triggers/:className/:triggerName [put]
func (h *HooksController) HandleUpdateTrigger() {
	if h.Auth.IsMaster == false {
		h.Data["json"] = errs.ErrorMessageToMap(errs.OperationForbidden, "Need master key!")
		h.ServeJSON()
		return
	}

	className := h.Ctx.Input.Param(":className")
	triggerName := h.Ctx.Input.Param(":triggerName")
	var err error
	var result = types.M{}
	if utils.S(h.JSONBody["__op"]) == "Delete" {
		// delete
		err = hooks.DeleteTrigger(className, triggerName)
	} else {
		// update
		if h.JSONBody["url"] == nil {
			h.Data["json"] = errs.ErrorMessageToMap(errs.WebhookError, "invalid hook declaration")
			h.ServeJSON()
			return
		}
		hook := types.M{
			"className":   className,
			"triggerName": triggerName,
			"url":         h.JSONBody["url"],
		}
		result, err = hooks.UpdateHook(hook)
	}
	if err != nil {
		h.Data["json"] = errs.ErrorToMap(err)
		h.ServeJSON()
		return
	}
	h.Data["json"] = result
	h.ServeJSON()
}

// Get ...
// @router / [get]
func (h *HooksController) Get() {
	h.ClassesController.Get()
}

// Post ...
// @router / [post]
func (h *HooksController) Post() {
	h.ClassesController.Post()
}

// Delete ...
// @router / [delete]
func (h *HooksController) Delete() {
	h.ClassesController.Delete()
}

// Put ...
// @router / [put]
func (h *HooksController) Put() {
	h.ClassesController.Put()
}
