package mentordm

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

type MentorRecruitmentID uuid.UUID

func NewMentorRecruitmentID() MentorRecruitmentID {
	return MentorRecruitmentID(uuid.New())
}

func (id MentorRecruitmentID) String() string {
	return uuid.UUID(id).String()
}

func (id *MentorRecruitmentID) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		u, err := uuid.ParseBytes(v)
		if err != nil {
			return err
		}
		*id = MentorRecruitmentID(u)
		return nil
	case string:
		u, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		*id = MentorRecruitmentID(u)
		return nil
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type MentorRecruitmentID", v)
	}
}

func (id MentorRecruitmentID) Value() (driver.Value, error) {
	return uuid.UUID(id).String(), nil
}
