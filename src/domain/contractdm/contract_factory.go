package contractdm

func GenContractIfCreate(id, planID, userID, contractApprovalID, message string, status uint8) (*Contract, error) {
	contractID := ContractID(id)
	return newContract(
		contractID,
		planID,
		userID,
		contractApprovalID,
		message,
		status,
	)
}
