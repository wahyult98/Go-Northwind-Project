package controllers

import "codeid.northwind/services"

type ControllerManager struct {
	CategoryController
}

// constructor
func NewControllerManager(serviceMgr *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewCategoryController(&serviceMgr.CategoryService),
	}
}
