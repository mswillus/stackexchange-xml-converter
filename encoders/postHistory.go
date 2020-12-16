package encoders

import "html"

// PostHistory entity
type PostHistory struct {
	ID                string `xml:"Id,attr" json:"Id"`
	PostID            string `xml:"PostId,attr" json:"PostId"`
	UserID            string `xml:"UserId,attr" json:"UserId,omitempty"`
	PostHistoryTypeID string `xml:"PostHistoryTypeId,attr" json:"PostHistoryTypeId"`
	UserDisplayName   string `xml:"UserDisplayName,attr" json:"UserDisplayName,omitempty"`
	ContentLicense    string `xml:"ContentLicense,attr" json:"ContentLicense,omitempty"`
	RevisionGUID      string `xml:"RevisionGUID,attr" json:"RevisionGUID,omitempty"`
	Text              string `xml:"Text,attr" json:"Text,omitempty"`
	Comment           string `xml:"Comment,attr" json:"Comment,omitempty"`
	CreationDate      string `xml:"CreationDate,attr" json:"CreationDate"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (ph PostHistory) GetCSVHeaderRow() []string {
	return []string{"Id", "PostId", "UserId",
		"PostHistoryTypeId", "UserDisplayName",
		"ContentLicense", "RevisionGUID",
		"Text", "Comment", "CreationDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (ph *PostHistory) GETCSVRow(skipHTMLDecoding bool) []string {

	text := ph.Text
	if skipHTMLDecoding {
		text = html.EscapeString(text)
	}

	return []string{ph.ID, ph.PostID, ph.UserID,
		ph.PostHistoryTypeID, ph.UserDisplayName,
		ph.ContentLicense, ph.RevisionGUID,
		text, ph.Comment, ph.CreationDate}
}
