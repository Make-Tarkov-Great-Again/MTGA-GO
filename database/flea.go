package database

var flea = Flea{}

func GetFlea() *Flea {
	return &flea
}

func setFlea() {

}

type Flea struct {
	Offers           []interface{}
	OffersCount      int
	SelectedCategory string
	Categories       map[string]int
}
type MemberCategory int

/* const (
	defaultCategory                  MemberCategory = 0
	developerCategory                MemberCategory = 1
	uniqueIDCategory                 MemberCategory = 2
	traderCategory                   MemberCategory = 4
	groupCategory                    MemberCategory = 8
	systemCategory                   MemberCategory = 16
	chatModeratorCategory            MemberCategory = 32
	chatModeratorWithPermBanCategory MemberCategory = 64
	unitTestCategory                 MemberCategory = 128
	sherpaCategory                   MemberCategory = 256
	emissaryCategory                 MemberCategory = 512
) */
