package dto

// 通用对象
type CommonDTO struct {
	ID uint `json:"id" formData:"id" uri:"id"`
}


// 分页相关的 DTO
type Paginate struct {
  Page  int `json:"page,omitempty" form:"page"`
  Limit int `json:"limit,omitempty" form:"limit"`
}

// 获取分页参数，对参数合理性做一些校验，如果不合理使用默认值
func (m *Paginate) GetPage() int {
  if m.Page <= 0 {
    m.Page = 1
  }
  return m.Page
}

func (m *Paginate) GetLimit() int {
  if m.Limit <= 0 {
    m.Limit = 1
  }
  return m.Limit
}
