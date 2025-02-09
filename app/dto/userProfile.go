package dto

type UpdateProfilePayload struct {
	UserID    string  `json:"userId"`
	FullName  *string `json:"fullName,omitempty"`
	Avatar    string  `json:"avatar,omitempty"`
	Cover     string  `json:"cover,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
	Gender    *string `json:"gender,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Location  *string `json:"location,omitempty"`
	AboutMe   *string `json:"aboutMe,omitempty"`
	Website   *string `json:"website,omitempty"`

	// Social Media Links
	Facebook  *string `json:"facebook,omitempty"`
	Twitter   *string `json:"twitter,omitempty"`
	LinkedIn  *string `json:"linkedin,omitempty"`
	Instagram *string `json:"instagram,omitempty"`
	GitHub    *string `json:"github,omitempty"`
	YouTube   *string `json:"youtube,omitempty"`
	TikTok    *string `json:"tiktok,omitempty"`
}
