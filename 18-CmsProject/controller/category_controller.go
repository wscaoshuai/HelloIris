package controller

import (
	"HelloIris/18-CmsProject/model"
	"HelloIris/18-CmsProject/service"
	"HelloIris/18-CmsProject/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

/**
 * 食品种类控制器
 */
type CategoryController struct {
	Ctx     iris.Context
	Service service.CategoryService
}

/**
 * 添加食品种类实体
 */
type CategoryEntity struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	RestaurantId string `json:"restaurant_id"`
}

func (cc *CategoryController) BeforeActivation(a mvc.BeforeActivation) {

	//通过商铺Id获取对应的食品种类
	a.Handle("GET", "/getcategory/{shopId}", "GetCategoryByShopId")

	//获取全部的食品种类
	a.Handle("GET", "/v2/restaurant/category", "GetAllCategory")

}

/**
 * 即将添加的食品记录实体
 */
type AddFoodEntity struct {
	Name         string   `json:"name"`          //食品名称
	Description  string   `json:"description"`   //食品描述
	ImagePath    string   `json:"image_path"`    //食品图片地址
	Activity     string   `json:"activity"`      //食品活动
	Attributes   []string `json:"attributes"`    //食品特点
	Specs        []Specs  `json:"specs"`         //食品规格
	CategoryId   int      `json:"category_id"`   //食品种类  种类id
	RestaurantId string   `json:"restaurant_id"` //哪个店铺的食品 店铺id
}

// 食品规格
type Specs struct {
	Specs      string `json:"specs"`
	PackingFee int    `json:"packing_fee"`
	Price      int    `json:"price"`
}

/**
 * url: /shopping/addfood
 * type：post
 * descs：添加食品数据功能
 */
func (cc *CategoryController) PostaAddfood() mvc.Result {

	var foodEntity AddFoodEntity
	err := cc.Ctx.ReadJSON(&foodEntity)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_FOODADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_FOODADD),
			},
		}
	}

	var food model.Food
	food.Name = foodEntity.Name
	food.Description = foodEntity.Description
	food.ImagePath = foodEntity.ImagePath
	food.Activity = foodEntity.Activity
	food.CategoryId = int64(foodEntity.CategoryId)
	//food.Restaurant = foodEntity.RestaurantId
	food.DelFlag = 0
	food.Rating = 0 //初始评分为零

	return mvc.Response{}
}

/**
 * url：/shopping/getcategory/1
 * type：get
 * desc：根据商铺Id获取对应的商铺的食品种类列表信息
 */
func (cc *CategoryController) GetCategoryByShopId() mvc.Result {

	shopIdStr := cc.Ctx.Params().Get("shopId")
	if shopIdStr == "" {
		iris.New().Logger().Info(shopIdStr)
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORIES),
			},
		}
	}

	shopId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		iris.New().Logger().Info(shopId)
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORIES),
			},
		}
	}

	//调用服务实体功能类查询商铺对应的食品种类列表
	categories, err := cc.Service.GetCategoryByShopId(int64(shopId))
	if err != nil {
		iris.New().Logger().Info(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORIES),
			},
		}
	}

	//返回对应店铺的食品种类类型
	return mvc.Response{
		Object: map[string]interface{}{
			"status":        utils.RECODE_OK,
			"category_list": &categories,
		},
	}
}

/**
 * url：/shopping/addcategory
 * type：post
 * desc：添加食品种类记录
 */
func (cc *CategoryController) PostAddcategory() mvc.Result {

	var categoryEntity *CategoryEntity
	cc.Ctx.ReadJSON(&categoryEntity)

	if categoryEntity.Name == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORYADD),
			},
		}
	}

	iris.New().Logger().Info(categoryEntity)
	restaurant_id, _ := strconv.Atoi(categoryEntity.RestaurantId)
	//构造要添加的数据记录
	foodCategory := &model.FoodCategory{
		CategoryName:     categoryEntity.Name,
		CategoryDesc:     categoryEntity.Description,
		Level:            1,
		ParentCategoryId: 0,
		RestaurantId:     int64(restaurant_id),
	}

	saveSucc := cc.Service.AddCategory(foodCategory)
	if !saveSucc {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORYADD),
			},
		}
	}

	//成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_CATEGORYADD),
		},
	}
}

/**
 * url：/shopping/v2/restaurant/category
 * type：get
 * desc：获取所有食品种类供添加商铺时进行添加
 */
func (cc *CategoryController) GetAllCategory() mvc.Result {

	cc.Service.GetAllCategory()

	categories, err := cc.Service.GetAllCategory()
	if err != nil {
		iris.New().Logger().Info(err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_CATEGORIES),
			},
		}
	}

	//返回所有的食品类型列表
	return mvc.Response{
		Object: &categories,
	}
}
