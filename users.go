package trakt

var (
	UserURL            = Hyperlink("users/{username}")
	UsersAllWatchedURL = Hyperlink("users/{username}/watched/shows")
)

// Create a ShowsService with the base url.URL
func (c *Client) Users() (shows *ShowsService) {
	shows = &ShowsService{client: c}
	return
}

type UsersService struct {
	client *Client
}

// One returns a single user identified by a Trakt username . It also returns a Result
// object to inspect the returned response of the server.
func (r *UsersService) One(username string) (user *User, result *Result) {
	url, _ := UserURL.Expand(M{"username": username})
	result = r.client.get(url, &user)
	return
}

func (r *ShowsService) AllWatched(username string) (shows []WatchedShow, result *Result) {
	url, _ := UsersAllWatchedURL.Expand(M{"username": username})
	result = r.client.get(url, &shows)
	return
}

type WatchedShow struct {
	Show          Show   `json:"show"`
	LastWatchedAt string `json:"last_watched_at"`
}

// Show struct for the Trakt v2 API
type User struct {
	Username string `json:"username"`
	Private  bool   `json:"private"`
	Name     string `json:"name"`
	VIP      bool   `json:"vip"`
	VIP_EP   bool   `json:"vip_ep"`
}
