package hashtable


type HashTable struct {
	Items map[string] int
	Path map[string]map[string]int
}

func (h *HashTable) Add(key string, val int) {
	if h.Items == nil {
		h.Items = map[string]int{}
	}
	h.Items[key] = val
}

func (h *HashTable) Update(key string, val int)  {
	if h.Path == nil {
		h.Path = map[string]map[string]int{
		}
	}

	if h.Path[key] == nil {
		h.Path[key] = map[string]int{}
	}

	h.Path[key]["man"] = 100


}