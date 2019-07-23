package Robot

import "encoding/json"

/*
curl 'https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxx' \
   -H 'Content-Type: application/json' \
   -d '{"msgtype": "text",
        "text": {
             "content": "我就是我, 是不一样的烟火"
        }
      }'



text:

{
    "msgtype": "text",
    "text": {
        "content": "我就是我, 是不一样的烟火@156xxxx8827"
    },
    "at": {
        "atMobiles": [
            "156xxxx8827",
            "189xxxx8325"
        ],
        "isAtAll": false
    }
}


link

{
    "msgtype": "link",
    "link": {
        "text": "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。
而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？",
        "title": "时代的火车向前开",
        "picUrl": "",
        "messageUrl": "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
    }
}


markdown:

{
     "msgtype": "markdown",
     "markdown": {
         "title":"杭州天气",
         "text": "#### 杭州天气 @156xxxx8827\n" +
                 "> 9度，西北风1级，空气良89，相对温度73%\n\n" +
                 "> ![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)\n"  +
                 "> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n"
     },
    "at": {
        "atMobiles": [
            "156xxxx8827",
            "189xxxx8325"
        ],
        "isAtAll": false
    }
 }




{
    "actionCard": {
        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
        "text": "![screenshot](serverapi2/@lADOpwk3K80C0M0FoA)
 ### 乔布斯 20 年前想打造的苹果咖啡厅
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
        "hideAvatar": "0",
        "btnOrientation": "0",
        "singleTitle" : "阅读全文",
        "singleURL" : "https://www.dingtalk.com/"
    },
    "msgtype": "actionCard"
}


{
    "actionCard": {
        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
        "text": "![screenshot](serverapi2/@lADOpwk3K80C0M0FoA)
 ### 乔布斯 20 年前想打造的苹果咖啡厅
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
        "hideAvatar": "0",
        "btnOrientation": "0",
        "btns": [
            {
                "title": "内容不错",
                "actionURL": "https://www.dingtalk.com/"
            },
            {
                "title": "不感兴趣",
                "actionURL": "https://www.dingtalk.com/"
            }
        ]
    },
    "msgtype": "actionCard"
}


{
    "feedCard": {
        "links": [
            {
                "title": "时代的火车向前开",
                "messageURL": "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
                "picURL": "https://www.dingtalk.com/"
            },
            {
                "title": "时代的火车向前开2",
                "messageURL": "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI",
                "picURL": "https://www.dingtalk.com/"
            }
        ]
    },
    "msgtype": "feedCard"
}


*/

const (
	MSGTYPE_NULL = iota
	MSGTYPE_TEXT
	MSGTYPE_LINK
	MSGTYPE_MARKDOWN
	MSGTYPE_ACTIONCARD
	MSGTYPE_FEEDCARD
)

type Msg interface {
	JsonEncode() (string, error)
}

type MsgText struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func (m *MsgText) JsonEncode() (string, error) {
	by, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(by), nil
}

type MsgLink struct {
	Msgtype string `json:"msgtype"`
	Link    struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicURL     string `json:"picUrl"`
		MessageURL string `json:"messageUrl"`
	} `json:"link"`
}

func (m *MsgLink) JsonEncode() (string, error) {
	by, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(by), nil
}

type MsgMarkDown struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func (m *MsgMarkDown) JsonEncode() (string, error) {
	by, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(by), nil
}

type MsgActionCard struct {
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		HideAvatar     string `json:"hideAvatar"`
		BtnOrientation string `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
		SingleURL      string `json:"singleURL"`
	} `json:"actionCard"`
	Msgtype string `json:"msgtype"`
}

func (m *MsgActionCard) JsonEncode() (string, error) {
	by, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(by), nil
}

type MsgFeedCard struct {
	FeedCard struct {
		Links []struct {
			Title      string `json:"title"`
			MessageURL string `json:"messageURL"`
			PicURL     string `json:"picURL"`
		} `json:"links"`
	} `json:"feedCard"`
	Msgtype string `json:"msgtype"`
}

func (m *MsgFeedCard) JsonEncode() (string, error) {
	by, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(by), nil
}
