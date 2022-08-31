package hw_leetcode

import "strings"

// 11. 字典树  代表题目：820 触类旁通：648.208🉑

//	208. 实现 Trie (前缀树)	https://leetcode.cn/problems/implement-trie-prefix-tree/
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

//	820. 单词的压缩编码	https://leetcode.cn/problems/short-encoding-of-words/
//输入：words = ["time", "me", "bell"]
//输出：10
func minimumLengthEncoding(words []string) int {
	ans := 0
	m := map[string]bool{}
	for _, w := range words {
		m[w] = true
	}
	for w, _ := range m {
		for i := 1; i < len(w); i++ {
			delete(m, w[i:])
		}
	}
	for w, _ := range m {
		ans += len(w) + 1
	}
	return ans
}

//648. 单词替换	https://leetcode.cn/problems/replace-words/
func replaceWords(dictionary []string, sentence string) string {
	t := Constructor()
	for i := range dictionary {
		t.Insert(dictionary[i])
	}
	sList := strings.Split(sentence, " ")
	for i := 0; i < len(sList); i++ {
		if sub, ok := t.Search(sList[i]); ok {
			sList[i] = sub
		}
	}
	return strings.Join(sList, " ")
}

type Trie struct {
	Child map[byte]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	for i, _ := range word {
		v := byte(word[i])
		if t.Child == nil {
			t.Child = make(map[byte]*Trie)
		}
		if _, ok := t.Child[v]; !ok {
			t.Child[v] = &Trie{}
		}
		t = t.Child[v]
	}
	t.IsEnd = true
}

func (t *Trie) Search(word string) (string, bool) {
	sub := ""
	for i, _ := range word {
		v := byte(word[i])
		if t.Child == nil {
			return "", false
		}
		if _, ok := t.Child[v]; !ok {
			return "", false
		}
		t = t.Child[v]
		sub += string(v)
		if t.IsEnd {
			return sub, true
		}
	}
	return "", false
}
