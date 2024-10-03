package mentordm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type MentorRecruitmentID string

func NewMentorRecruitmentID() MentorRecruitmentID {
	return MentorRecruitmentID(uuid.New().String())
}

func NewMentorRecruitmentIDByVal(val string) (MentorRecruitmentID, error) {
	if val == "" {
		return MentorRecruitmentID(""), xerrors.New("mentor recruitment id must not be empty")
	}
	return MentorRecruitmentID(val), nil
}

func (id MentorRecruitmentID) String() string {
	return string(id)
}

func (id MentorRecruitmentID) Equal(mentorID2 MentorRecruitmentID) bool {
	return string(id) == string(mentorID2)
}
