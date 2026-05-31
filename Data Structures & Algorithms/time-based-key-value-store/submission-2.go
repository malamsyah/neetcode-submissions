type TimeMap struct {
	t map[string][]int
	kv map[string]map[int]string
}

func Constructor() TimeMap {
	t := make(map[string][]int)
	kv := make(map[string]map[int]string)
	return TimeMap{
		t: t,
		kv: kv,
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.t[key] == nil {
		this.t[key] = []int{timestamp}
		this.kv[key] = make(map[int]string)
		this.kv[key][timestamp] = value 
	} else {
		this.t[key] = append(this.t[key], timestamp)
		this.kv[key][timestamp] = value
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if this.t[key] == nil {
		return ""
	}

	vals, _ := this.t[key]

	fmt.Println(vals)
	fmt.Println("target: ", timestamp)

	n := len(vals) -1
	if timestamp >= vals[n] {
		return this.kv[key][vals[n]]
	}

	if timestamp == vals[0] {
		return this.kv[key][vals[0]]
	}

	if timestamp < vals[0] {
		return ""
	}

	l, r := 0, n
	for l < r {
		m := l + (r - l) / 2
		if vals[m] == timestamp {
			l = m
			break
		}
		if vals[m] > timestamp {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if vals[l] <= timestamp {
		return this.kv[key][vals[l]]
	}

	if vals[l] > timestamp {
		return this.kv[key][vals[l-1]]
	}

	return ""
}
