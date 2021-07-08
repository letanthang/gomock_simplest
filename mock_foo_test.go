package foo

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/letanthang/gomock_simplest/mocks"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	// assert that Bar() is invoked
	defer ctrl.Finish()
	m := mocks.NewMockFoo(ctrl)
	m.EXPECT().Do(123).Return(101)
	// m.EXPECT().Do(gomock.Any()).Return(101)
	m.EXPECT().Do2(gomock.Nil())

	Bar(m)
}
