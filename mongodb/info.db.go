package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

const infoDB string = "info-db"

/*
Insert : Function description of the function, input, output, and notes if any
*/
func Insert(ctx context.Context, info *InfoToDB) error {
	if info != nil {
		/*
			Kiểm tra tính hợp lệ của dữ liệu
		*/
		if err := info.Validate(); err != nil {
			return err
		}
		/*
			Gắn thời điểm hiện tại cho trường CreateAt
		*/
		info.CreateAt = time.Now()
		/*
			Insert dữ liệu vào database
		*/
		ir, err := db.Collection(infoDB).InsertOne(ctx, info)
		if err != nil {
			return err
		}
		/*
			Gán id đã insert cho info
		*/
		info.ID = ir.InsertedID.(primitive.ObjectID)
		/*
			Không xảy ra lỗi -> thành công
		*/
		return nil
	}

	return errors.New("info is invalid")
}

/*
Update : Function description of the function, input, output, and notes if any
*/
func Update(ctx context.Context, info *InfoToDB) error {
	if info != nil {
		/*
			Kiểm tra tính hợp lệ của trường filter
		*/
		if info.ID.IsZero() {
			return errors.New("ID field is invalid")
		}

		/*
			Gắn thời điểm hiện tại cho trường CreateAt
		*/
		info.CreateAt = time.Now()

		var (
			filter = bson.M{"_id": info.ID}
			update = bson.M{"$set": info}
		)
		/*
			Cập nhật dữ liệu vào database
		*/
		ur, err := db.Collection(infoDB).UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
		/*
			Nếu matched với nhiều hơn 0 thì thành công
		*/
		if ur.MatchedCount > 0 {
			return nil
		}
		return mongo.ErrNoDocuments
	}

	return errors.New("info is invalid")
}

/*
GetOne : Function description of the function, input, output, and notes if any
*/
func GetOne(ctx context.Context, id primitive.ObjectID) (*InfoToDB, error) {
	/*
		Kiểm tra tính hợp lệ ID
	*/
	if id.IsZero() {
		return nil, errors.New("ID field is invalid")
	}

	var (
		filter = bson.M{"_id": id}
	)
	/*
		Tìm kiếm dữ liệu khớp với filter
	*/
	fr := db.Collection(infoDB).FindOne(ctx, filter)
	/*
		Nếu lỗi là không tìm thấy -> trả về mismatched
	*/
	if fr.Err() == mongo.ErrNoDocuments {
		return nil, mongo.ErrNoDocuments
	}
	if fr.Err() != nil {
		return nil, fr.Err()
	}

	/*
		Tạo một biến tạm chữa dữ liệu decoding
	*/
	var item = &InfoToDB{}

	/*
		Giải mã dữ liệu
	*/
	if err := fr.Decode(item); err != nil {
		return nil, err
	}
	return item, nil

}
