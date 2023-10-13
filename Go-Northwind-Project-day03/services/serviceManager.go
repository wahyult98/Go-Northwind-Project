package services

import "codeid.northwind/repositories"

type ServiceManager struct {
	CategoryService
	//ProductService
}

// constructor
func NewServiceManager(repoMgr *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		CategoryService: *NewCategoryService(repoMgr),
		//ProductService: *NewProductService(&repoMgr.ProductCategory),
	}
}
