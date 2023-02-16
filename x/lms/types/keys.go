package types

const (

	ModuleName = "lms"

	StoreKey = ModuleName

	// RouterKey = ModuleName
	// QuerierRoute = ModuleName
)

var (
	AdminKey  = []byte{0x01}
	StudentKey  = []byte{0x02}
	LeavesKey = []byte{0x03}
)

func AdminStoreKey(classID string) []byte {
	key := make([]byte, len(AdminKey)+len(classID))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], classID)
	return key
}

func StudentStoreKey(classID string) []byte {
	key := make([]byte, len(StudentKey)+len(classID))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], classID)
	return key
}

func LeavesStoreKey(classID string) []byte {
	key := make([]byte, len(LeavesKey)+len(classID))
	copy(key, LeavesKey)
	copy(key[len(LeavesKey):], classID)
	return key
}