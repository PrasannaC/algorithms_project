package DisjointSets

import (
	"disjoint_sets/DisjointSets/Set"
)

/*
	Alias Sets for the type Map<object,Set>
*/
type Sets map[interface{}]*Set.Set

/*
	Structure for the disjoint set data structure instance
*/
type DisjointSet struct {
	setMap Sets
}

/*
	MakeSet operation.
	Accepts value and adds to the map of sets in the instance of disjoint set
*/
func (d *DisjointSet) MakeSet(data interface{}) {
	set := new(Set.Set)
	set.Parent = set
	set.Data = data
	if d.setMap == nil {
		d.setMap = make(map[interface{}]*Set.Set, 0)
	}
	//store the set as value against they key of data in the map
	//setMap will contain pointers to set of the data we send.
	d.setMap[data] = set
}

/*
	Find operation.
	Accepts data/set as parameter and find the root if exists.
*/
func (d *DisjointSet) Find(data interface{}) *Set.Set {
	value, ok := data.(*Set.Set)
	if ok {
		data = value.Data
	}

	if d.setMap == nil {
		return nil
	} else if v, ok := d.setMap[data]; !ok {
		return nil
	} else {
		if v == nil {
			return nil
		}
		if v == v.Parent {
			return v // root found
		} else {
			//recursively call find till you find root
			root := d.Find(v.Parent.Data)
			return root
		}
	}
}

/*
	Union operation.
	Accepts two values/sets, and performs union operation on their respective roots
	Order of union: v1 <- v2
*/
func (d *DisjointSet) Union(v1, v2 interface{}) {
	// check if user has passed set references of values

	_, ok := v1.(*Set.Set)
	if !ok {
		v1 := d.setMap[v1]
		v1 = v1 //have to do this to avoid go compiler to think this variable is unused (-_-)"
	}

	_, ok = v2.(*Set.Set)
	if !ok {
		v2 := d.setMap[v2]
		v2 = v2 //have to do this to avoid go compiler to think this variable is unused (-_-)"
	}

	if v1 == nil || v2 == nil {
		panic("Values are not in the right format")
	}

	root1 := d.Find(v1)
	root2 := d.Find(v2)

	if root1 == root2 {
		return
	} else {
		root2.Parent = root1
		root1.Size++
		if root1.Rank <= 1+root2.Rank {
			root1.Rank = 1 + root2.Rank
		}
	}
}
