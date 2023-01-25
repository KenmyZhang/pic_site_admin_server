package pic_service

import (
	"fmt"
	"strings"
	"time"

	"github.com/KenmyZhang/pic_site_admin_server/app/models"
	"github.com/KenmyZhang/pic_site_admin_server/app/models/vo"
	productDto "github.com/KenmyZhang/pic_site_admin_server/app/service/pic_service/dto"
	"github.com/KenmyZhang/pic_site_admin_server/pkg/logging"
	"github.com/jinzhu/copier"
)

type Product struct {
	Id   int64
	Name string

	Enabled int
	CatId   int64

	PageNum  int
	PageSize int

	//M *models.YshopStoreProductRule

	Ids []int64

	Dto productDto.StorePic

	JsonObj map[string]interface{}

	Order int

	News       string
	PriceOrder string
	SalesOrder string
	Sid        string

	Uid int64

	Unique string

	Type string
}

func (d *Product) AddOrSaveProduct() (err error) {
	var (
		model models.YshopStorePic
	)
	m := d.Dto
	copier.Copy(&model, &m)

	images := strings.Join(m.ListPicUrl, ",")
	model.ListPicUrl = images

	if model.PicType == 0 {
		model.PicType = 1
	}
	if m.Id > 0 {
		err = models.UpdateByProduct(m.Id, model)
	} else {
		model.CreateTime = time.Now()
		model.UpdateTime = model.CreateTime
		err = models.AddProduct(&model)
	}

	if err != nil {
		return err
	}

	return err
}

func (d *Product) GetProductInfo() map[string]interface{} {
	var (
		mapData    = make(map[string]interface{})
		picProduct models.YshopStorePic
		picDto     productDto.StorePicInfo
	)

	if d.Id == 0 {
		return mapData
	}

	picProduct = models.GetProduct(d.Id)
	ee := copier.Copy(&picDto, picProduct)
	if ee != nil {
		logging.Error(ee)
	}
	fmt.Printf("picProduct:%+v\n", picProduct)
	fmt.Printf("picProduct.ListPicUrl:%+v\n", picDto.ListPicUrl)
	picDto.ListPicUrl = strings.Split(picProduct.ListPicUrl, ",")
	mapData["productInfo"] = picDto

	return mapData
}

func (d *Product) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}
	if d.Enabled > 0 {
		maps["is_show"] = 1
	} else {
		maps["is_show"] = -1
	}

	if d.CatId > 0 {
		maps["category_id"] = d.CatId
	}

	total, list := models.GetAllProduct(d.PageNum, d.PageSize, maps)
	for idx, val := range list {
		picList := strings.Split(val.ListPicUrl, ",")
		if len(picList) > 0 {
			list[idx].PrimaryPicUrl = picList[0]
		}
	}
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *Product) Del() error {
	return models.DelByProduct(d.Ids)
}
