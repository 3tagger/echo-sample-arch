package usecase_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/3tagger/echo-sample-arch/internal/user"
	"github.com/3tagger/echo-sample-arch/internal/user/usecase"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_GetAll(t *testing.T) {
	type fields struct {
		userRepositoryMock func(t testing.TB) user.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*user.User
		wantErr bool
	}{
		{
			name: "repository returns error",
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("GetAll", mock.Anything).Return(nil, errors.New("test"))

					return repoMock
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no error, return a user",
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					u := &user.User{
						Id: 1,
					}
					users := []*user.User{u}
					repoMock.On("GetAll", mock.Anything).Return(users, nil)

					return repoMock
				},
			},
			want: []*user.User{
				&user.User{Id: 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewUserUsecase(tt.fields.userRepositoryMock(t))
			got, err := u.GetAll(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_GetOneById(t *testing.T) {
	type fields struct {
		userRepositoryMock func(t testing.TB) user.Repository
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		want    *user.User
		wantErr bool
	}{
		{
			name: "repository returns error",
			args: args{
				userId: 1,
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("GetOneById", mock.Anything, int64(1)).Return(nil, errors.New("test"))

					return repoMock
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "repository does not return a user",
			args: args{
				userId: 1,
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("GetOneById", mock.Anything, int64(1)).Return(nil, nil)

					return repoMock
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no error, return a user",
			args: args{
				userId: 1,
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					u := &user.User{
						Id: 1,
					}
					repoMock.On("GetOneById", mock.Anything, int64(1)).Return(u, nil)

					return repoMock
				},
			},
			want:    &user.User{Id: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewUserUsecase(tt.fields.userRepositoryMock(t))
			got, err := u.GetOneById(context.Background(), tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.GetOneById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.GetOneById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_InsertOne(t *testing.T) {
	type fields struct {
		userRepositoryMock func(t testing.TB) user.Repository
	}
	type args struct {
		user user.User
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		want    *user.User
		wantErr bool
	}{
		{
			name: "repository returns error",
			args: args{
				user: user.User{Name: "test"},
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("InsertOne", mock.Anything, user.User{Name: "test"}).Return(nil, errors.New("test"))

					return repoMock
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no error, return a user",
			args: args{
				user: user.User{Name: "test"},
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					u := &user.User{
						Id:   1,
						Name: "test",
					}
					repoMock.On("InsertOne", mock.Anything, user.User{Name: "test"}).Return(u, nil)

					return repoMock
				},
			},
			want:    &user.User{Id: 1, Name: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewUserUsecase(tt.fields.userRepositoryMock(t))
			got, err := u.InsertOne(context.Background(), tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.InsertOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.InsertOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_InsertMany(t *testing.T) {
	type fields struct {
		userRepositoryMock func(t testing.TB) user.Repository
	}
	type args struct {
		users []user.User
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		want    *user.User
		wantErr bool
	}{
		{
			name: "repository returns error",
			args: args{
				users: []user.User{{Name: "test"}},
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("InsertMany", mock.Anything, []user.User{{Name: "test"}}).Return(errors.New("test"))

					return repoMock
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no error, return a user",
			args: args{
				users: []user.User{{Name: "test"}},
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("InsertMany", mock.Anything, []user.User{{Name: "test"}}).Return(nil)

					return repoMock
				},
			},
			want:    &user.User{Id: 1, Name: "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewUserUsecase(tt.fields.userRepositoryMock(t))
			err := u.InsertMany(context.Background(), tt.args.users)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.InsertMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserUsecase_DeleteOne(t *testing.T) {
	type fields struct {
		userRepositoryMock func(t testing.TB) user.Repository
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		wantErr bool
	}{
		{
			name: "repository returns error",
			args: args{
				userId: 1,
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("DeleteOne", mock.Anything, int64(1)).Return(errors.New("test"))

					return repoMock
				},
			},
			wantErr: true,
		},
		{
			name: "no error, delete a user",
			args: args{
				userId: 1,
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("DeleteOne", mock.Anything, int64(1)).Return(nil)

					return repoMock
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewUserUsecase(tt.fields.userRepositoryMock(t))
			err := u.DeleteOne(context.Background(), tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.DeleteOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserUsecase_UpdateOneById(t *testing.T) {
	type fields struct {
		userRepositoryMock func(t testing.TB) user.Repository
	}
	type args struct {
		userId int64
		user   user.User
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		wantErr bool
	}{
		{
			name: "repository returns error",
			args: args{
				userId: 1,
				user:   user.User{Name: "test"},
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("UpdateOneById", mock.Anything, int64(1), user.User{Name: "test"}).Return(errors.New("test"))

					return repoMock
				},
			},
			wantErr: true,
		},
		{
			name: "no error",
			args: args{
				userId: 1,
				user:   user.User{Name: "test"},
			},
			fields: fields{
				userRepositoryMock: func(t testing.TB) user.Repository {
					repoMock := user.NewMockRepository(t)
					repoMock.On("UpdateOneById", mock.Anything, int64(1), user.User{Name: "test"}).Return(nil)

					return repoMock
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usecase.NewUserUsecase(tt.fields.userRepositoryMock(t))
			err := u.UpdateOneById(context.Background(), tt.args.userId, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUsecase.UpdateOneById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
