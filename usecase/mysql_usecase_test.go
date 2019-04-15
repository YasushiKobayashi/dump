package usecase

import (
	"testing"

	"github.com/YasushiKobayashi/dump/models"
	"github.com/YasushiKobayashi/dump/usecase/mock_usecase"
	"github.com/golang/mock/gomock"
)

func Test_MysqlInterActor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mysqlMock := mock_usecase.NewMockMysqlRepository(ctrl)
	uploadMysqlRepository := mock_usecase.NewMockMysqlRepository(ctrl)
	localMock := mock_usecase.NewMockLocalRepository(ctrl)
	awsS3Mock := mock_usecase.NewMockAwsS3Repository(ctrl)

	mysql := &models.MysqlModel{}

	i := &MysqlInterActor{
		MysqlRepository:     mysqlMock,
		SyncMysqlRepository: uploadMysqlRepository,
		LocalRepository:     localMock,
		AwsS3Repository:     awsS3Mock,
	}

	t.Run("type dump", func(t *testing.T) {
		gomock.InOrder(
			mysqlMock.EXPECT().GetDump().Return(mysql, nil).Times(1),
		)

		param := &models.Param{
			Type: models.Dump,
		}
		err := i.Upload(param)
		if err != nil {
			t.Fatalf("faile test %#v", err)
		}
	})

	t.Run("type sync", func(t *testing.T) {
		gomock.InOrder(
			mysqlMock.EXPECT().GetSync().Return(mysql, nil).Times(1),
		)

		param := &models.Param{
			Type: models.Sync,
		}
		err := i.Upload(param)
		if err != nil {
			t.Fatalf("faile test %#v", err)
		}
	})

	t.Run("output File", func(t *testing.T) {
		path := models.Path{}
		param := &models.Param{
			Type:   models.Dump,
			Output: models.File,
			Path:   path,
		}

		gomock.InOrder(
			mysqlMock.EXPECT().GetDump().Return(mysql, nil).Times(1),
			localMock.EXPECT().Upload(mysql.GetQuery(param.Type), path).Return(nil).Times(1),
		)

		err := i.Upload(param)
		if err != nil {
			t.Fatalf("faile test %#v", err)
		}
	})

	t.Run("output S3", func(t *testing.T) {
		path := models.Path{}
		awsS3 := models.AwsS3Param{}
		param := &models.Param{
			Type:   models.Dump,
			Output: models.AwsS3,
			Path:   path,
			AwsS3:  awsS3,
		}

		gomock.InOrder(
			mysqlMock.EXPECT().GetDump().Return(mysql, nil).Times(1),
			awsS3Mock.EXPECT().Upload(mysql.GetQuery(param.Type), path, awsS3).Return(nil).Times(1),
		)

		err := i.Upload(param)
		if err != nil {
			t.Fatalf("faile test %#v", err)
		}
	})

	t.Run("output other database", func(t *testing.T) {
		param := &models.Param{
			Type:   models.Sync,
			Output: models.OtherDatabase,
		}

		gomock.InOrder(
			mysqlMock.EXPECT().GetSync().Return(mysql, nil).Times(1),
			uploadMysqlRepository.EXPECT().UploadSync(mysql).Return(nil).Times(1),
		)

		err := i.Upload(param)
		if err != nil {
			t.Fatalf("faile test %#v", err)
		}
	})
}
