package model

type AccountStatus string

const ( 
	Activated	AccountStatus = "Activated"
	Deactivated	AccountStatus = "Deactivated"
	Blocked		AccountStatus = "Blocked"
)

func (s AccountStatus) isValid() bool {
	switch s {
		case Activated, Deactivated, Blocked:
			return true
	}
	return false;	
}

func (s AccountStatus) isActivated() bool{
	return s == Activated
}

func (s AccountStatus) isDeactivated() bool{
	return s == Deactivated
}

func (s AccountStatus) isBlocked() bool{
	return s == Blocked
}
