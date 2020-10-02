package main

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-begin-training/mongodb-driver.first-step/mongodb"

	"github.com/gin-gonic/gin"
)

func setInfo(c *gin.Context) {
	var (
		request = &struct {
			Name string `json:"name"`
		}{}

		r = c.Request
	)
	/*
		Đọc dữ liệu được gửi từ client
	*/
	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
		return
	}
	/*
		Kiểm tra tính hợp lệ của dữ liệu
	*/
	if len(request.Name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "name field is invalid"})
		return
	}
	/*
		Tạo một biến mới tương thích với hàm Insert
	*/
	var item = &mongodb.InfoToDB{
		Name: request.Name,
	}
	/*
		Insert vào database
	*/
	if err := mongodb.Insert(r.Context(), item); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"code": 503, "message": err.Error()})
		return
	}
	/*
		Không xảy ra lỗi trả về thành công
	*/
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "successfully", "info": item})
	return
}

func updateInfo(c *gin.Context) {
	var (
		request = &struct {
			ID   primitive.ObjectID `json:"id"`
			Name string             `json:"name"`
		}{}

		r = c.Request
	)
	/*
		Đọc dữ liệu được gửi từ client
	*/
	if err := c.BindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
		return
	}
	/*
		Kiểm tra tính hợp lệ của dữ liệu
	*/
	switch true {
	case request.ID.IsZero():
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "id field is invalid"})
		return
	case len(request.Name) == 0:
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "name field is invalid"})
		return
	}
	/*
		Tạo một biến mới tương thích với hàm Update
	*/
	var item = &mongodb.InfoToDB{
		ID:   request.ID,
		Name: request.Name,
	}

	/*
		Cập nhật database
	*/
	if err := mongodb.Update(r.Context(), item); err != nil {
		switch true {
		case err == mongo.ErrNoDocuments:
			c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
			return
		default:
			c.JSON(http.StatusServiceUnavailable, gin.H{"code": 503, "message": err.Error()})
			return
		}
	}
	/*
		Không xảy ra lỗi trả về thành công
	*/
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "successfully", "info": item})
	return
}

func getInfo(c *gin.Context) {
	var (
		r = c.Request
	)
	/*
		Đọc dữ liệu được gửi từ client
	*/
	objectID, err := primitive.ObjectIDFromHex(c.DefaultQuery("id", primitive.NilObjectID.Hex()))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
		return
	}
	/*
		Query dữ liệu từ database
	*/
	item, err := mongodb.GetOne(r.Context(), objectID)
	switch true {
	case err == mongo.ErrNoDocuments:
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
		return
	case err != nil:
		c.JSON(http.StatusServiceUnavailable, gin.H{"code": 503, "message": err.Error()})
		return
	}
	/*
		Không xảy ra lỗi trả về thành công
	*/
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "successfully", "info": item})
	return
}
