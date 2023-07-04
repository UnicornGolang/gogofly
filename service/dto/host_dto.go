package dto

type ShowdownDTO struct {
	HostIP string `json:"host_ip" binding:"required" message:"host required"`
}
