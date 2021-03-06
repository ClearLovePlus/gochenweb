package model

type Tag struct {
	Model
	Name       string `json:"name"`
	CreateBy   string `json:"createBy"`
	ModifiedBy string `json:"modifiedBy"`
	State      int    `json:"state"`
}

//获取标签值
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

//获取标签总数
func GetTotalTags(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
