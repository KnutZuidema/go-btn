package btn

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KnutZuidema/go-btn/pkg/jsonrpc"
	"github.com/KnutZuidema/go-btn/pkg/model"
)

type Client struct {
	client *jsonrpc.Client
	apiKey string
}

func NewClient(client *http.Client, apiKey string) *Client {
	return &Client{
		client: jsonrpc.NewClient(client, "http://api.broadcasthe.net/"),
		apiKey: apiKey,
	}
}

func (c Client) GetUser() (*model.User, error) {
	var user model.User
	if err := c.client.GetInto(&user, "userInfo", c.apiKey); err != nil {
		return nil, err
	}
	return &user, nil
}

func (c Client) GetNews() ([]model.NewsItem, error) {
	var news model.News
	if err := c.client.GetInto(&news, "getNews", c.apiKey); err != nil {
		return nil, err
	}
	return news, nil
}

func (c Client) GetNewsByID(id int) (*model.NewsItem, error) {
	var news model.NewsItem
	if err := c.client.GetInto(&news, "getNewsById", c.apiKey, id); err != nil {
		return nil, err
	}
	return &news, nil
}

func (c Client) GetBlog() ([]model.NewsItem, error) {
	var news model.News
	if err := c.client.GetInto(&news, "getBlog", c.apiKey); err != nil {
		return nil, err
	}
	return news, nil
}

func (c Client) GetBlogByID(id int) (*model.NewsItem, error) {
	var news model.NewsItem
	if err := c.client.GetInto(&news, "getBlogById", c.apiKey, id); err != nil {
		return nil, err
	}
	return &news, nil
}

func (c Client) GetTVNews() ([]model.NewsItem, error) {
	var news model.News
	if err := c.client.GetInto(&news, "getTVNews", c.apiKey); err != nil {
		return nil, err
	}
	return news, nil
}

func (c Client) GetTVNewsByID(id int) (*model.NewsItem, error) {
	var news model.NewsItem
	if err := c.client.GetInto(&news, "getTVNewsById", c.apiKey, id); err != nil {
		return nil, err
	}
	return &news, nil
}

func (c Client) GetForumPage(id, page int) (*model.ForumPage, error) {
	var res model.ForumPage
	err := c.client.GetInto(&res, "getForumsPage", c.apiKey, id, page)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c Client) GetChangelog() ([]model.ChangelogEntry, error) {
	var log []model.ChangelogEntry
	if err := c.client.GetInto(&log, "getChangelog", c.apiKey); err != nil {
		return nil, err
	}
	return log, nil
}

func (c Client) GetInbox(page int) (*model.Inbox, error) {
	var inbox model.Inbox
	if err := c.client.GetInto(&inbox, "getInbox", c.apiKey, page); err != nil {
		return nil, err
	}
	return &inbox, nil
}

func (c Client) GetInboxConversation(id int) (*model.Conversation, error) {
	var conv model.Conversation
	if err := c.client.GetInto(&conv, "getInboxConversation", c.apiKey, id); err != nil {
		return nil, err
	}
	return &conv, nil
}

func (c Client) SendInboxConversation(id int, content string) error {
	var resp string
	if err := c.client.GetInto(&resp, "sendInboxConversation", c.apiKey, id, content); err != nil {
		return err
	}
	if resp != "success" {
		return fmt.Errorf("send conversation was not succesfull")
	}
	return nil
}

func (c Client) SearchTorrents(options model.SearchTorrentOptions, count, offset int) ([]model.Torrent, error) {
	var res struct {
		Results  string `json:"Results"`
		Torrents map[string]model.Torrent
	}
	if err := c.client.GetInto(&res, "getTorrents", c.apiKey, options, count, offset); err != nil {
		return nil, err
	}
	var torrents []model.Torrent
	for _, torrent := range res.Torrents {
		torrents = append(torrents, torrent)
	}
	return torrents, nil
}

func (c Client) GetTorrentDownloadURL(id int) (string, error) {
	var url string
	if err := c.client.GetInto(&url, "getTorrentsUrl", c.apiKey, id); err != nil {
		return "", err
	}
	return url, nil
}

func (c Client) GetForumsIndex(lastPost int) (map[string][]model.Forum, error) {
	var index map[string][]model.Forum
	if err := c.client.GetInto(&index, "getForumsIndex", c.apiKey, lastPost); err != nil {
		return nil, err
	}
	return index, nil
}

func (c Client) GetTorrentByID(id int) (*model.Torrent, error) {
	var torrent model.Torrent
	if err := c.client.GetInto(&torrent, "getTorrentById", c.apiKey, id); err != nil {
		return nil, err
	}
	return &torrent, nil
}

func (c Client) GetUserSubscriptions() ([]model.Subscription, error) {
	var subs model.Subscriptions
	if err := c.client.GetInto(&subs, "getUserSubscriptions", c.apiKey); err != nil {
		return nil, err
	}
	return subs, nil
}

func (c Client) GetUserSnatchlist(count, offset int) ([]model.Snatch, error) {
	var resp struct {
		Results  string                  `json:"results"`
		Torrents map[string]model.Snatch `json:"torrents"`
	}
	if err := c.client.GetInto(&resp, "getUserSnatchlist", c.apiKey, count, offset); err != nil {
		return nil, err
	}
	var snatches []model.Snatch
	for idStr, snatch := range resp.Torrents {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		snatch.ID = id
		snatches = append(snatches, snatch)
	}
	return snatches, nil
}
