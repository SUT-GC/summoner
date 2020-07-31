package biz

import (
	"summoner/repo"
	"summoner/utils"
)

func BlogPO2BlogSummaryDTO(po repo.BlogInfoPO) BlogSummaryDTO {
	return BlogSummaryDTO{
		Id:           po.Id,
		Title:        utils.DecodeWithBase64(po.Title),
		Summary:      utils.DecodeWithBase64(po.Summary),
		ViewCount:    po.ViewCount,
		CommentCount: po.CommentCount,
		CreatedAt:    po.CreatedAt,
		UpdatedAt:    po.UpdatedAt,
	}
}

func BlogPO2BlogDTO(po repo.BlogInfoPO) BlogDTO {
	return BlogDTO{
		Id:           po.Id,
		UserID:       po.UserID,
		Title:        utils.DecodeWithBase64(po.Title),
		Summary:      utils.DecodeWithBase64(po.Summary),
		Content:      utils.DecodeWithBase64(po.Content),
		ViewCount:    po.ViewCount,
		CommentCount: po.CommentCount,
		Status:       po.Status,
		CreatedAt:    po.CreatedAt,
		UpdatedAt:    po.UpdatedAt,
		DeletedAt:    po.DeletedAt,
	}
}

func BlogDTO2BlogPO(dto BlogDTO) repo.BlogInfoPO {
	return repo.BlogInfoPO{
		Id:           dto.Id,
		UserID:       dto.UserID,
		Title:        utils.EncodeWithBase64(dto.Title),
		Summary:      utils.EncodeWithBase64(dto.Summary),
		Content:      utils.EncodeWithBase64(dto.Content),
		ViewCount:    dto.ViewCount,
		CommentCount: dto.CommentCount,
		Status:       dto.Status,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
		DeletedAt:    dto.DeletedAt,
	}
}
