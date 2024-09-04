package routers

import (
	"modules/localmap/actions/actiongbkupdate"
	"modules/localmap/actions/actionlistmap"
	"modules/localmap/actions/actionutf8update"
	"modules/localmap/actions/actionviewmap"
	updateaction "modules/localmap/actions/update/actions"
	"modules/middlewares"
	"modules/search/actions/actiongbksearch"
	"modules/search/actions/actionsearch"
	"modules/search/actions/actionutf8search"

	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/herb/middleware/errorpage"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/middleware/router/httprouter"
)

// APIMiddlewares middlewares that should used in api requests
var APIMiddlewares = func() middleware.Middlewares {
	return middleware.Middlewares{
		middlewares.MiddlewareCsrfVerifyHeader,
		errorpage.MiddlewareDisable,
	}
}

// RouterAPIFactory api router factory.
var RouterAPIFactory = router.NewFactory(func() router.Router {
	var Router = httprouter.New()
	//Put your router configure code here
	Router.POST("/utf8update/:id").HandleFunc(actionutf8update.Action)
	Router.POST("/gbkupdate/:id").HandleFunc(actiongbkupdate.Action)
	Router.POST("/utf8search").HandleFunc(actionutf8search.Action)
	Router.POST("/gbksearch").HandleFunc(actiongbksearch.Action)
	Router.POST("/search").HandleFunc(actionsearch.Action)
	Router.GET("/map/list").HandleFunc(actionlistmap.Action)
	Router.GET("/map/view/:id").HandleFunc(actionviewmap.Action)
	Router.POST("/map/update").Handle(updateaction.ActionUpdate)
	Router.POST("/map/delete/:id").Handle(updateaction.ActionDelete)
	return Router
})
