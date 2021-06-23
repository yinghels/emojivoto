package emoji

//go:generate generateEmojiCodeMap -pkg emojivoto

type Emoji struct {
	Unicode   string `json:"unicode"`
	Shortcode string `json:"shortcode"`
}

type AllEmoji interface {
	WithShortcode(shortcode string) *Emoji
	List() []*Emoji
}

type inMemoryAllEmoji struct {
	emojiList []*Emoji
}

var top10Emoji = []string{
	":joy:",
	":sunglasses:",
	":doughnut:",
	":stuck_out_tongue_winking_eye:",
	":money_mouth_face:",
	":flushed:",
	":mask:",
	":nerd_face:",
	":ghost:",
	":skull_and_crossbones:",
}

func (allEmoji *inMemoryAllEmoji) List() []*Emoji {
	return allEmoji.emojiList
}

func (allEmoji *inMemoryAllEmoji) WithShortcode(shortcode string) *Emoji {
	for _, emoji := range allEmoji.List() {
		if emoji.Shortcode == shortcode {
			return emoji
		}
	}
	return nil
}

func NewAllEmoji() AllEmoji {
	emojiList := make([]*Emoji, 0)

	for _, name := range top10Emoji {
		e := &Emoji{
			Unicode:   emojiCodeMap[name],
			Shortcode: name,
		}
		emojiList = append(emojiList, e)
	}

	return &inMemoryAllEmoji{
		emojiList,
	}
}
