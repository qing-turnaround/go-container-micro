package handler

import (
	"context"

	"github.com/xing-you-ji/go-container-micro/category/domain/model"
	"github.com/xing-you-ji/go-container-micro/category/domain/service"
	category "github.com/xing-you-ji/go-container-micro/category/proto/category"
	"github.com/xing-you-ji/go-container-micro/common"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

func (h *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest,
	response *category.CreateCategoryResponse) error {
	category := new(model.Category)
	// json Tag赋值
	if err := common.SwapTo(request, category); err != nil {
		return err
	}
	categoryID, err := h.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryID
	return nil
}

func (h *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest,
	response *category.UpdateCategoryResponse) error {
	category := new(model.Category)
	if err := common.SwapTo(request, category); err != nil {
		return err
	}
	if err := h.CategoryDataService.UpdateCategory(category); err != nil {
		return err
	}
	response.Message = "分类更新成功"
	return nil
}

func (h *Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest,
	response *category.DeleteCategoryResponse) error {
	if err := h.CategoryDataService.DeleteCategory(request.CategoryId); err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}

func (h *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest,
	response *category.CategoryResponse) error {
	category, err := h.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return nil
	}

	return common.SwapTo(category, response)
}

func (h *Category) FindCategoryByID(ctx context.Context, request *category.FindByIdRequest,
	response *category.CategoryResponse) error {
	category, err := h.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

func (h *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest,
	response *category.FindAllResponse) error {
	categorySlice, err := h.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	return common.SwapCategoryTo(categorySlice, response)
}
func (h *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest,
	response *category.FindAllResponse) error {
	categorySlice, err := h.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	return common.SwapCategoryTo(categorySlice, response)
}
func (h *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest,
	response *category.FindAllResponse) error {
	categorySlice, err := h.CategoryDataService.FindAll()
	if err != nil {
		return err
	}
	return common.SwapCategoryTo(categorySlice, response)
}
