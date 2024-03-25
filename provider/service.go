package provider

type IProvider interface {
	BLockNumber() (uint64, error)
}
