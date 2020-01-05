package leetcode



type Trie struct {
	Prefix string
	IsWord bool
	Index []*Trie
}

var dict = make(map[string]int)


/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.Prefix = ""
	t.IsWord = false
	return t
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	if word == ""{
		this.IsWord = true
		return
	}
	if this.Index == nil{
		this.Index = make([]*Trie,26)
	}
	iTrie := this.Index[dict[string(word[0])]]
	if iTrie == nil {//如果word第一个字母的子树还未建立
		iTrie = &Trie{}
		iTrie.Prefix = word
		iTrie.IsWord = true
		this.Index[dict[string(word[0])]] = iTrie
		return
	}
	if iTrie.Prefix == word {
		if !iTrie.IsWord {
			iTrie.IsWord = true
		}
		return
	}
	lmi := FindLastMatchIndex(iTrie.Prefix,word)
	if lmi+1 == len(iTrie.Prefix){
		iTrie.Insert(word[lmi+1:])
		return
	}
	if lmi+1 == len(word){
		newTrie := &Trie{}
		newTrie.Prefix = word
		newTrie.IsWord = true
		this.Index[dict[string(word[0])]] = newTrie
		newTrie.Index = make([]*Trie,26)
		iTrie.Prefix = iTrie.Prefix[lmi+1:]
		newTrie.Index[dict[string(iTrie.Prefix[0])]] = iTrie
		return

	}
	newTrie := &Trie{}
	newTrie.Prefix = word[0:lmi+1]
	newTrie.IsWord = false
	newTrie.Index = make([]*Trie,26)
	this.Index[dict[string(word[0])]] = newTrie


	wordTrie := &Trie{}
	wordTrie.Prefix = word[lmi+1:]
	wordTrie.IsWord = true
	newTrie.Index[dict[string(wordTrie.Prefix[0])]] = wordTrie


	iTrie.Prefix = iTrie.Prefix[lmi+1:]
	newTrie.Index[dict[string(iTrie.Prefix[0])]] = iTrie

	return

}

func FindLastMatchIndex(wordA,wordB string)int{
	lenA := len(wordA)
	lenB := len(wordB)
	if lenA == 0 || lenB == 0 {
		return -1
	}
	for i:=0;;i++{
		if i==lenA || i== lenB{
			return i-1
		}
		if string(wordA[i]) != string(wordB[i]) {
			return i-1
		}
	}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this.Prefix == word{
		if this.IsWord {
			return true
		}
		return false
	}
	lmi := FindLastMatchIndex(this.Prefix,word)
	if lmi+1 == len(this.Prefix){
		if this.Index != nil && this.Index[dict[string(word[lmi+1])]] != nil {
			return this.Index[dict[string(word[lmi+1])]].Search(word[lmi+1:])
		}
	}
	return false
}




/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	lmi := FindLastMatchIndex(this.Prefix,prefix)
	if lmi+1 == len(prefix){
		return true
	}
	if lmi+1 == len(this.Prefix){
		if this.Index != nil && this.Index[dict[string(prefix[lmi+1])]] != nil{
			return this.Index[dict[string(prefix[lmi+1])]].StartsWith(prefix[lmi+1:])
		}
	}
	return false
}


func init(){
	dict["a"]=0
	dict["b"]=1
	dict["c"]=2
	dict["d"]=3
	dict["e"]=4
	dict["f"]=5
	dict["g"]=6
	dict["h"]=7
	dict["i"]=8
	dict["j"]=9
	dict["k"]=10
	dict["l"]=11
	dict["m"]=12
	dict["n"]=13
	dict["o"]=14
	dict["p"]=15
	dict["q"]=16
	dict["r"]=17
	dict["s"]=18
	dict["t"]=19
	dict["u"]=20
	dict["v"]=21
	dict["w"]=22
	dict["x"]=23
	dict["y"]=24
	dict["z"]=25
}