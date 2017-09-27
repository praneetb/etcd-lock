package etcdlock

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type LockerSuite struct {
	suite.Suite
	ctx    context.Context
	locker *Locker
}

func (s *LockerSuite) SetupSuite() {
	s.ctx = context.Background()
	locker, err := NewLocker(LockerOptions{
		Address:        "127.0.0.1:2379",
		DialOptions:    []grpc.DialOption{grpc.WithInsecure()},
		DefaultTimeout: 3 * time.Second,
	})
	s.Nil(err)
	s.locker = locker
}

func (s *LockerSuite) TestNewLocker() {
	locker, err := NewLocker(LockerOptions{
		Address:        "127.0.0.1:2379",
		DialOptions:    []grpc.DialOption{grpc.WithInsecure()},
		DefaultTimeout: 3 * time.Second,
	})

	s.Nil(err)
	s.Equal(locker.etcdKeyPrefix, defaultEtcdKeyPrefix)
}

func (s *LockerSuite) TestLockNewKey() {
	lock, err := s.locker.Lock(s.ctx, "test_lock_new_key")

	s.Nil(err)
	s.Contains(string(lock.keyName), "test_lock_new_key")
}

func TestLocker(t *testing.T) {
	suite.Run(t, new(LockerSuite))
}