package constants

type ConditionType string

const (
	Condition_InDivision      ConditionType = "in_division"
	Condition_DivisionVisitor ConditionType = "division_visitor"
	Condition_HomeFacility    ConditionType = "home"
	Condition_VisitFacility   ConditionType = "visit"
	Condition_HomeOrVisit     ConditionType = "home_or_visit"
	Condition_Rating          ConditionType = "rating"
	Condition_Role            ConditionType = "role"
	Condition_All             ConditionType = "all" // Will assign to everyone with a linked discord account
	Condition_DivisionStaff   ConditionType = "division_staff"
	Condition_FacilityStaff   ConditionType = "facility_staff"
	Condition_FacilityRole    ConditionType = "facility_role"
)
