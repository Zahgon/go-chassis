package schemas

import (
	"context"

	"github.com/go-chassis/go-chassis/v2/examples/schemas/employ"
)

// EmployServer is a struct
type EmployServer struct{}

// AddEmploy 这里实现服务端接口中的方法。
func (s *EmployServer) AddEmploy(ctx context.Context, in *employ.EmployRequest) (*employ.EmployResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EditEmploy is a method used to edit employ
func (s *EmployServer) EditEmploy(ctx context.Context, in *employ.EmployRequest) (*employ.EmployResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEmploys is a method used to get employs
func (s *EmployServer) GetEmploys(ctx context.Context, in *employ.EmployRequest) (*employ.EmployResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteEmploys is a method used to delete employ
func (s *EmployServer) DeleteEmploys(ctx context.Context, in *employ.EmployRequest) (*employ.EmployResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
