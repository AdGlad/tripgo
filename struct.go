// structs.go
package main

type City struct {
	Name    string `firestore:"city"`
	Country string `firestore:"country"`
}

type CitiesDocument struct {
	Top50 []City `firestore:"top50"`
}

// User represents the structure of user data in Firestore.
type User struct {
	UserID    string                   `json:"userId"` // Renaming ID to UserID
	Badges    []map[string]interface{} `json:"badges"`
	Countries map[string]Country       `json:"countries"`
	// Optional: Include any other fields relevant to your application.
}

// type User struct {
// 	Badges    []map[string]interface{} `json:"badges"`
// 	Countries map[string]Country       `json:"countries"`
// 	//   PlaceHistory map[string]PlaceHistory  `json:"placeHistory"`
// }

// Country struct represents a country with regions.
type Country struct {
	CountryCode string            `json:"countryCode"`
	Regions     map[string]Region `json:"regions"`
}

// Region struct represents a region with place histories.
type Region struct {
	RegionCode   string                  `json:"regionCode"`
	PlaceHistory map[string]PlaceHistory `json:"placehistory"`
}

// PlaceHistory struct represents the detailed information of a visit.
type PlaceHistory struct {
	ID            string                 `firestore:"id"`
	Name          string                 `firestore:"name"`
	Location      string                 `firestore:"location"`
	Latitude      float64                `firestore:"latitude"`
	Longitude     float64                `firestore:"longitude"`
	Distance      int                    `firestore:"distance"`
	StreetAddress string                 `firestore:"streetAddress"`
	City          string                 `firestore:"city"`
	CountryName   string                 `firestore:"countryName"`
	CountryCode   string                 `firestore:"countryCode"`
	Postal        string                 `firestore:"postal"`
	Region        string                 `firestore:"region"`
	RegionCode    string                 `firestore:"regionCode"`
	ApiregionCode string                 `firestore:"apiregionCode"`
	Timezone      string                 `firestore:"timezone"`
	Elevation     int                    `firestore:"elevation"`
	Timestamp     int                    `firestore:"timestamp"`
	Arrivaldate   int                    `firestore:"arrivaldate"`
	FirstVisit    bool                   `firestore:"firstVisit"`
	UserId        string                 `firestore:"userId"`
	Description   string                 `firestore:"description"`
	Diary         string                 `firestore:"diary"`
	Rating        string                 `firestore:"rating"`
	Poi           string                 `firestore:"poi"`
	Category      []string               `firestore:"category"`
	PoiId         int                    `firestore:"poiId"`
	PoiName       string                 `firestore:"poiName"`
	PoiGroupIds   []string               `firestore:"poiGroupIds"`
	LocationRaw   map[string]interface{} `firestore:"locationRaw"`
	ImagePaths    []string               `firestore:"imagePaths"`
	ImageSize     int                    `firestore:"imageSize"`
}

// Badge represents the information about a badge awarded to a user.
type Badge struct {
	Name        string `firestore:"name"`
	AwardedOn   string `firestore:"awardedOn"` // Consider using time.Time for actual timestamps
	Description string `firestore:"description"`
}
