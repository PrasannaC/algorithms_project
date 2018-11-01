package DisjointSets

import (
	"graphs/DisjointSets/Set"
)

/*
	Alias Sets for the type Map<object,Set>
*/
type Sets map[interface{}]*Set.Set

type UnionByType int32

/*
	Enum for UnionBy options
*/
const (
	BY_RANK = iota
	BY_SIZE = iota
)

/*
	Structure for the disjoint set data structure instance
*/
type DisjointSet struct {
	setMap  Sets
	usePC   bool
	unionBy UnionByType
}

/*
	Factory method to create instance of disjoint set with given params
*/
func CreateDisjointSet(usePathCompression bool, unionBy UnionByType) *DisjointSet {
	return &DisjointSet{usePC: usePathCompression, unionBy: unionBy}
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
			if d.usePC {
				v.Parent = root
			}
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
		if (d.unionBy == BY_RANK && root1.Rank < root2.Rank) ||
			(d.unionBy == BY_SIZE && root1.Size < root2.Size) {
			root1.Parent = root2
			root2.Size += root1.Size
			if root1.Rank == root2.Rank {
				root2.Rank++
			}
		} else {
			root2.Parent = root1
			root1.Size += root2.Size
			if root2.Rank == root1.Rank {
				root1.Rank++
			}
		}
	}
}
