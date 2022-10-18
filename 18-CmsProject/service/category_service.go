package service

import (
	"HelloIris/18-CmsProject/model"
	"github.com/kataras/iris"
	"xorm.io/xorm"
)

/**
 * 食品种类服务接口
 */
type CategoryService interface {
	AddCategory(model *model.FoodCategory) bool
	GetCategoryByShopId(shopId int64) ([]model.FoodCategory, error)
	GetAllCategory() ([]model.FoodCategory, error)
}

/**
 * 种类服务实现结构体
 */
type categoryService struct {
	Engine *xorm.Engine
}

/**
 * 实例化种类服务:服务器
 */
func NewCategoryService(engine *xorm.Engine) CategoryService {
	return &categoryService{
		Engine: engine,
	}
}

/**
 * 通过商铺Id获取食品类型
 */
func (cs *categoryService) GetCategoryByShopId(shopId int64) ([]model.FoodCategory, error) {
	categories := make([]model.FoodCategory, 0)
	err := cs.Engine.Where(" restaurant_id = ? ", shopId).Find(&categories)
	return categories, err
}

/**
 * 添加食品种类记录
 */
func (cs *categoryService) AddCategory(category *model.FoodCategory) bool {
	_, err := cs.Engine.Insert(category)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 获取所有的种类信息
 */
func (cs *categoryService) GetAllCategory() ([]model.FoodCategory, error) {
	categories := make([]model.FoodCategory, 0)
	err := cs.Engine.Where(" parent_category_id = ? ", 0).Find(&categories)
	return categories, err
}
