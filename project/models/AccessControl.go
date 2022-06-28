package models

type AccessControl map[string][]string

var(
	CreateUser AccessControl=map[string][]string{
		"admin":{"admin"},
		"manager":{"admin"},
		"teacher":{"admin","manager"},
		"student":{"admin","manager","teacher"},
	}

	GetUser  []string=[]string{
		"admin",
		"teacher",
		"manager",
	}
	UpdateUserRole	AccessControl=map[string][]string{
		"admin":{"admin"},
		"manager":{"admin"},
		"teacher":{"admin"},
		"student":{"admin"},
	}

	UpdateUserData	AccessControl=map[string][]string{
		"admin":{"admin"},
		"manager":{"admin"},
		"teacher":{"admin","manager"},
		"student":{"admin","manager","teacher"},
	}

	UpdateUserMobile	AccessControl=map[string][]string{
		"admin":{"admin"},
		"manager":{"admin"},
		"teacher":{"admin","manager"},
		"student":{"admin","manager","teacher"},
	}

	DeleteUser	AccessControl=map[string][]string{
		"admin":{"admin"},
		"manager":{"admin"},
		"teacher":{"admin","manager"},
		"student":{"admin","manager"},
	}
)
var(
	CreateClassroom  []string=[]string{
		"admin",
		"manager",
	}
	GetClassroom  []string=[]string{
		"admin",
		"manager",
		"teacher",
	}
	UpdateClassroom []string=[]string{
		"admin",
		"manager",
	}
	DeleteClassroom []string=[]string{
		"admin",
		"manager",
	}
	AssignClassStudent []string=[]string{
		"admin",
		"manager",
	}
	DeleteClassStudent []string=[]string{
		"admin",
		"manager",
	}
	AssignClassTeacher []string=[]string{
		"admin",
		"manager",
	}
	UpdateClassTeacher []string=[]string{
		"admin",
		"manager",
	}
	DeleteClassTeacher []string=[]string{
		"admin",
		"manager",
	}
)


// admin:=CreateAccessControl.Role