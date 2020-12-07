package encoders

func getVoteColumnMap() map[string]int {
	return map[string]int{
		"Id":           0,
		"UserId":       1,
		"PostId":       2,
		"VoteTypeId":   3,
		"BountyAmount": 4,
		"CreationDate": 5}
}

// // Vote entity
// type Vote struct {
// 	ID           string `xml:"Id,attr"`
// 	UserID       string `xml:"UserId,attr"`
// 	PostID       string `xml:"PostId,attr"`
// 	VoteTypeID   string `xml:"VoteTypeId,attr"`
// 	BountyAmount string `xml:"BountyAmount,attr"`
// 	CreationDate string `xml:"CreationDate,attr"`
// }

// func (v Vote) GetCSVHeaderRow() []string {
// 	return []string{"Id", "UserId", "PostId",
// 		"VoteTypeId", "BountyAmount", "CreationDate"}
// }

// func (v *Vote) GETCSVRow() []string {
// 	return []string{v.ID, v.UserID, v.PostID,
// 		v.VoteTypeID, v.BountyAmount, v.CreationDate}
// }
